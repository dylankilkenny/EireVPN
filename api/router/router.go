package router

import (
	"eirevpn/api/db"
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"
	"io/ioutil"

	"eirevpn/api/plan"
	"eirevpn/api/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const secretkey = "verysecretkey1995"

func SetupRouter(logging bool) *gin.Engine {
	var router *gin.Engine

	if logging {
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
		router = gin.New()
		router.Use(gin.Recovery())
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Origin", "Content-Length", "Content-Type", "Authorization")
	router.Use(cors.New(config))

	public := router.Group("/api")

	private := router.Group("/api/private")
	private.Use(auth(secretkey))

	// secret := router.Group("/api/secret")

	public.POST("/signup", user.SignUpUser)
	public.POST("/login", user.LoginUser)
	// public.POST("/validate", user.ValidateToken)
	private.GET("/plan/:id", plan.Plan)
	private.POST("/plan", plan.CreatePlan)
	private.PUT("/plan", plan.UpdatePlan)
	private.DELETE("/plan/:id", plan.DeletePlan)
	private.GET("/plans", plan.AllPlans)
	// private.GET("/address", user.GetSubscribedAddresses)
	// private.POST("/address", user.SubscribeToAddress)
	// private.DELETE("/remove", user.RemoveSubscribedAddress)
	// secret.POST("/users", user.SubscribedUsers)

	return router
}

func auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var usersession models.UserSession

		// Fetch authentification token
		authToken, err := c.Request.Cookie("authToken")
		if err != nil {
			logger.Log(logger.Fields{
				Loc:  "router.go - auth()",
				Code: errors.AuthCookieMissing.Code,
				Err:  err.Error(),
			})
			c.AbortWithStatusJSON(errors.AuthCookieMissing.Status, errors.AuthCookieMissing)
			return
		}

		// Fetch refresh token
		refreshToken, err := c.Request.Cookie("refreshToken")
		if err != nil {
			logger.Log(logger.Fields{
				Loc:  "router.go - auth()",
				Code: errors.RefresCookieMissing.Code,
				Err:  err.Error(),
			})
			c.AbortWithStatusJSON(errors.RefresCookieMissing.Status, errors.RefresCookieMissing)
			return
		}

		// Check auth token is valid
		authClaims, err := jwt.ValidateToken(authToken.Value)
		if err != nil {

			// If auth token is invalid check refresh token is valid
			refreshClaims, err := jwt.ValidateToken(refreshToken.Value)
			if err != nil {
				logger.Log(logger.Fields{
					Loc:  "router.go - auth()",
					Code: errors.TokenInvalid.Code,
					Err:  err.Error(),
				})
				c.AbortWithStatusJSON(errors.TokenInvalid.Status, errors.TokenInvalid)
				return
			}

			usersession = models.UserSession{
				UserID:     refreshClaims.UserID,
				Identifier: refreshClaims.SessionIdentifier,
			}

			db := db.GetDB()
			if err := db.Find(&usersession).Error; err != nil {
				logger.Log(logger.Fields{
					Loc:  "router.go - auth()",
					Code: errors.InvalidIdentifier.Code,
					Err:  err.Error(),
				})
				c.SetCookie("refreshToken", "", -1, "/", "localhost", true, false)
				c.AbortWithStatusJSON(errors.InvalidIdentifier.Status, errors.InvalidIdentifier)
				return
			}
		}

		// If auth token or refresh token is valid check if crsf token matches the one supplied
		// in the header
		if authClaims == nil || authClaims.CSRF != c.GetHeader("X-CSRF-Token") {
			logger.Log(logger.Fields{
				Loc:   "router.go - auth()",
				Code:  errors.CSRFTokenInvalid.Code,
				Extra: map[string]interface{}{"authClaims": authClaims},
				Err:   "CSRF tokens do not match",
			})
			c.AbortWithStatusJSON(errors.CSRFTokenInvalid.Status, errors.CSRFTokenInvalid)
			return
		}

		if usersession == (models.UserSession{}) {
			usersession = models.UserSession{
				UserID: authClaims.UserID,
			}
		}

		// create a new user session
		usersession, err = user.CreateSession(usersession.UserID)
		if err != nil {
			logger.Log(logger.Fields{
				Loc:   "/login - LoginUser() - Create session",
				Code:  errors.InternalServerError.Code,
				Extra: map[string]interface{}{"UserID": usersession.UserID},
				Err:   err.Error(),
			})
			c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
			return
		}

		// If all auth checks pass create fresh tokens
		newAuthToken, newRefreshToken, newCsrfToken, err := jwt.Tokens(usersession)
		if err != nil {
			logger.Log(logger.Fields{
				Loc:   "router.go - auth()",
				Code:  errors.InternalServerError.Code,
				Extra: map[string]interface{}{"UserID": usersession.UserID},
				Err:   err.Error(),
			})
			c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
			return
		}

		// TODO: Change the domain name and add correct maxAge time
		authCookieMaxAge := 15 * 60 // 15 minutes in seconds
		c.SetCookie("authToken", newAuthToken, authCookieMaxAge, "/", "localhost", true, false)

		// TODO: Change the domain name and add correct maxAge time
		refreshCookieMaxAge := 72 * 60 * 60 // 72 hours in seconds
		c.SetCookie("refreshToken", newRefreshToken, refreshCookieMaxAge, "/", "localhost", true, false)
		c.Header("X-CSRF-Token", newCsrfToken)
	}
}
