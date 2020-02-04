package router

import (
	"eirevpn/api/config"
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"eirevpn/api/server"
	"eirevpn/api/settings"
	"eirevpn/api/util/jwt"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"eirevpn/api/plan"
	"eirevpn/api/user"
	"eirevpn/api/userplan"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const secretkey = "verysecretkey1995"

func Init(logging bool) *gin.Engine {

	conf := config.GetConfig()

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
	protected.GET("/user/:id", user.User)
	protected.PUT("/user/update/:id", user.UpdateUser)
	protected.DELETE("/users/delete/:id", user.DeleteUser)
	protected.GET("/users", user.AllUsers)
	public.POST("/user/webhook", user.Webhook)
	private.GET("/user/updatepayment", user.StripeUpdatePaymentSession)
	private.GET("/user/session/:planid", user.StripeSession)
	private.GET("/user/cancel", user.CancelSubscription)
	public.GET("/user/logout", user.Logout) //public so this router can skip auth middleware

	protected.GET("/plans/:id", plan.Plan)
	protected.POST("/plans/create", plan.CreatePlan)
	protected.PUT("/plans/update/:id", plan.UpdatePlan)
	protected.DELETE("/plans/delete/:id", plan.DeletePlan)
	protected.GET("/plans", plan.AllPlans)
	public.GET("/plans", plan.AllPlansPublic)

	protected.GET("/userplans/:id", userplan.UserPlan)
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

	router.Static("/assets", "./assets")
	return router
}

func auth(secret string, protected bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.GetConfig()
		if conf.App.EnableAuth {
			var usersession models.UserAppSession

			// Fetch refresh token
			authToken, err := c.Request.Cookie(conf.App.AuthCookieName)
			if err != nil {
				logger.Log(logger.Fields{
					Loc:  "router.go - auth()",
					Code: errors.AuthCookieMissing.Code,
					Err:  err.Error(),
				})
				c.AbortWithStatusJSON(errors.AuthCookieMissing.Status, errors.AuthCookieMissing)
				return
			}

			// Check auth token is valid
			authClaims, err := jwt.ValidateToken(authToken.Value)
			if err != nil {
				logger.Log(logger.Fields{
					Loc:  "router.go - auth()",
					Code: errors.TokenInvalid.Code,
					Err:  err.Error(),
				})
				c.SetCookie(conf.App.AuthCookieName, "", -1, "/", conf.App.Domain, false, true)
				c.AbortWithStatusJSON(errors.TokenInvalid.Status, errors.TokenInvalid)
				return
			}

			usersession = models.UserAppSession{
				UserID:     authClaims.UserID,
				Identifier: authClaims.SessionIdentifier,
			}

			if err := usersession.Find(); err != nil {
				logger.Log(logger.Fields{
					Loc:   "router.go - auth()",
					Code:  errors.InvalidIdentifier.Code,
					Extra: map[string]interface{}{"Identifier": usersession.Identifier},
					Err:   err.Error(),
				})
				c.SetCookie(conf.App.AuthCookieName, "", -1, "/", conf.App.Domain, false, true)
				c.AbortWithStatusJSON(errors.InvalidIdentifier.Status, errors.InvalidIdentifier)
				return
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
			newAuthToken, newCsrfToken, err := jwt.Tokens(newUserSession)
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
			c.SetCookie(conf.App.AuthCookieName, newAuthToken, authCookieMaxAge, "/", conf.App.Domain, false, true)
			c.Header("X-CSRF-Token", newCsrfToken)
		}
	}
}
