package router

import (
	"eirevpn/api/config"
	"eirevpn/api/errors"
	"eirevpn/api/handlers/message"
	"eirevpn/api/handlers/plan"
	"eirevpn/api/handlers/server"
	"eirevpn/api/handlers/settings"
	"eirevpn/api/handlers/user"
	"eirevpn/api/handlers/userplan"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const secretkey = "verysecretkey1995"

func Init(logging bool) *gin.Engine {

	conf := config.Load()

	var router *gin.Engine

	if logging {
		var file, err = os.OpenFile(logger.LogFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
		router = gin.New()
		router.Use(gin.Recovery())
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = conf.App.AllowedOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowBrowserExtensions = true
	corsConfig.ExposeHeaders = []string{"X-CSRF-Token", "X-Auth-Token"}
	corsConfig.AddAllowHeaders("Origin", "Content-Length", "Content-Type", "Authorization", "X-CSRF-Token", "X-Auth-Token")
	router.Use(cors.New(corsConfig))

	public := router.Group("/api")
	private := router.Group("/api/private")
	protected := router.Group("/api/protected")
	private.Use(auth(secretkey, false))
	protected.Use(auth(secretkey, true))

	public.POST("/user/signup", user.SignUpUser)
	public.POST("/user/login", user.LoginUser)
	private.GET("/user/get/:id", user.User)
	private.PUT("/user/changepassword", user.ChangePassword)
	private.PUT("/user/update/:id", user.UpdateUser)
	protected.DELETE("/users/delete/:id", user.DeleteUser)
	protected.GET("/users", user.AllUsers)
	public.POST("/user/webhook", user.Webhook)
	private.GET("/user/updatepayment", user.StripeUpdatePaymentSession)
	private.GET("/user/session/:planid", user.StripeSession)
	private.GET("/user/cancel", user.CancelSubscription)
	public.GET("/user/logout", user.Logout) //public so this router can skip auth middleware

	public.GET("/user/confirm_email/:token", user.ConfirmEmail)

	protected.GET("/plans/:id", plan.Plan)
	protected.POST("/plans/create", plan.CreatePlan)
	protected.PUT("/plans/update/:id", plan.UpdatePlan)
	protected.DELETE("/plans/delete/:id", plan.DeletePlan)
	protected.GET("/plans", plan.AllPlans)
	public.GET("/plans", plan.AllPlansPublic)

	private.GET("/userplans/:userid", userplan.UserPlan)
	protected.POST("/userplans/create", userplan.CreateUserPlan)
	protected.PUT("/userplans/update/:id", userplan.UpdateUserPlan)
	protected.DELETE("/userplans/delete/:id", userplan.DeleteUserPlan)
	protected.GET("/userplans", userplan.AllUserPlans)

	protected.GET("/servers/:id", server.Server)
	protected.POST("/servers/create", server.CreateServer)
	protected.PUT("/servers/update/:id", server.UpdateServer)
	protected.DELETE("/servers/delete/:id", server.DeleteServer)
	private.GET("/servers/connect/:id", server.Connect)
	private.GET("/servers", server.AllServers)

	protected.GET("/settings", settings.Settings)
	protected.PUT("/settings/update", settings.UpdateSettings)

	public.POST("/message", message.Message)

	router.Static("/assets", "./assets")
	return router
}

func auth(secret string, protected bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.Load()
		if conf.App.EnableAuth {
			var usersession models.UserAppSession

			// Fetch auth token
			authToken, err := c.Request.Cookie(conf.App.AuthCookieName)
			if err != nil || authToken.Value == "" {
				var errMsg string
				if err != nil {
					errMsg = err.Error()
				} else {
					errMsg = "Auth cookie missing"
				}
				logger.Log(logger.Fields{
					Loc:  "router.go - auth()",
					Code: errors.AuthCookieMissing.Code,
					Err:  errMsg,
				})
				c.AbortWithStatusJSON(errors.AuthCookieMissing.Status, errors.AuthCookieMissing)
				return
			}

			// Check auth token is valid
			authClaims, err := jwt.ValidateToken(authToken.Value)
			if err != nil {

				// Fetch refresh token
				refreshToken, err := c.Request.Cookie(conf.App.RefreshCookieName)
				if err != nil || refreshToken.Value == "" {
					var errMsg string
					if err != nil {
						errMsg = err.Error()
					} else {
						errMsg = "Refresh cookie missing"
					}
					logger.Log(logger.Fields{
						Loc:  "router.go - auth()",
						Code: errors.RefresCookieMissing.Code,
						Err:  errMsg,
					})
					clearCookies(c)
					c.AbortWithStatusJSON(errors.RefresCookieMissing.Status, errors.RefresCookieMissing)
					return
				}

				refreshClaims, err := jwt.ValidateToken(refreshToken.Value)
				if err != nil {
					logger.Log(logger.Fields{
						Loc:  "router.go - auth()",
						Code: errors.RefresCookieMissing.Code,
						Err:  err.Error(),
					})
					clearCookies(c)
					c.AbortWithStatusJSON(errors.TokenInvalid.Status, errors.TokenInvalid)
					return
				}

				usersession = models.UserAppSession{
					UserID:     refreshClaims.UserID,
					Identifier: refreshClaims.SessionIdentifier,
				}

				if err := usersession.Find(); err != nil {
					logger.Log(logger.Fields{
						Loc:   "router.go - auth()",
						Code:  errors.InvalidIdentifier.Code,
						Extra: map[string]interface{}{"Identifier": usersession.Identifier},
						Err:   err.Error(),
					})
					clearCookies(c)
					c.AbortWithStatusJSON(errors.InvalidIdentifier.Status, errors.InvalidIdentifier)
					return
				}

			}

			// Check CSRF token
			if conf.App.EnableCSRF {
				if authClaims.CSRF != c.GetHeader("X-CSRF-Token") {
					var reason string
					authCSRF := ""
					if c.GetHeader("X-CSRF-Token") == "" {
						reason = "CSRF token is missing from header"
					}
					if authClaims.CSRF == "" {
						reason = "CSRF token is missing from auth claims"
					} else {
						authCSRF = authClaims.CSRF
					}
					logger.Log(logger.Fields{
						Loc:   "router.go - auth()",
						Code:  errors.CSRFTokenInvalid.Code,
						Extra: map[string]interface{}{"auth-CSRF": authCSRF, "head-CSRF": c.GetHeader("X-CSRF-Token")},
						Err:   reason,
					})
					c.AbortWithStatusJSON(errors.CSRFTokenInvalid.Status, errors.CSRFTokenInvalid)
					return
				}
			}

			if protected {
				var user models.User
				user.ID = authClaims.UserID
				if err := user.Find(); err != nil {
					logger.Log(logger.Fields{
						Loc: "router.go - auth()",
						Extra: map[string]interface{}{
							"UserID": authClaims.UserID,
							"Detail": "User Not found when checking user type",
						},
						Err: err.Error(),
					})
					c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
					return
				}

				if user.Type != models.UserTypeAdmin {
					logger.Log(logger.Fields{
						Loc:  "router.go - auth()",
						Code: errors.ProtectedRouted.Code,
						Extra: map[string]interface{}{
							"UserID": authClaims.UserID,
						},
						Err: "User does not have permission to access route",
					})
					c.AbortWithStatusJSON(errors.ProtectedRouted.Status, errors.ProtectedRouted)
					return
				}
			}

			// create a new user session
			var newUserSession models.UserAppSession
			if err := newUserSession.New(authClaims.UserID); err != nil {
				logger.Log(logger.Fields{
					Loc:   "/login - LoginUser() - Create session",
					Code:  errors.InternalServerError.Code,
					Extra: map[string]interface{}{"UserID": authClaims.UserID},
					Err:   err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}

			// If all auth checks pass create fresh tokens
			newAuthToken, newRefreshToken, newCsrfToken, err := jwt.Tokens(newUserSession)
			if err != nil {
				logger.Log(logger.Fields{
					Loc:   "router.go - auth()",
					Code:  errors.InternalServerError.Code,
					Extra: map[string]interface{}{"UserID": authClaims.UserID},
					Err:   err.Error(),
				})
				c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
				return
			}

			// Add user id to the context for use within the routes
			c.Set("UserID", newUserSession.UserID)

			// TODO: Change the domain name and add correct maxAge time
			authCookieMaxAge := 24 * 60 * conf.App.AuthCookieAge
			refreshCookieMaxAge := 24 * 60 * conf.App.RefreshCookieAge
			c.SetCookie(conf.App.AuthCookieName, newAuthToken, authCookieMaxAge, "/", conf.App.Domain, false, true)
			c.SetCookie(conf.App.RefreshCookieName, newRefreshToken, refreshCookieMaxAge, "/", conf.App.Domain, false, true)
			c.SetCookie("uid", strconv.FormatUint(uint64(newUserSession.UserID), 10), authCookieMaxAge, "/", conf.App.Domain, false, false)
			c.Header("X-CSRF-Token", newCsrfToken)
		}
	}
}

func clearCookies(c *gin.Context) {
	conf := config.Load()
	c.SetCookie(conf.App.AuthCookieName, "", -1, "/", conf.App.Domain, false, true)
	c.SetCookie(conf.App.RefreshCookieName, "", -1, "/", conf.App.Domain, false, true)
	c.SetCookie("uid", "", -1, "/", conf.App.Domain, false, false)
}
