package router

import (
	"eirevpn/api/db"
	"eirevpn/api/errors"
	"eirevpn/api/models"
	"eirevpn/api/util/jwt"
	"fmt"
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

		// Fetch authentification token
		authToken, err := c.Request.Cookie("authToken")
		if err != nil || authToken.Value == "" {
			fmt.Println("no authToken")
			c.AbortWithStatusJSON(errors.AuthCookieMissing.Status, errors.AuthCookieMissing)
			return
		}

		// Fetch refresh token
		refreshToken, err := c.Request.Cookie("refreshToken")
		if err != nil || refreshToken.Value == "" {
			fmt.Println("no refreshToken")
			c.AbortWithStatusJSON(errors.RefresCookieMissing.Status, errors.RefresCookieMissing)
			return
		}

		// Check auth token is valid
		authClaims, err := jwt.ValidateAuthToken(authToken.Value)
		if err != nil {

			// If auth token is invalid check refresh token is valid
			refreshClaims, err := jwt.ValidateRefreshToken(refreshToken.Value)
			if err != nil {
				fmt.Println("Token Invalid")
				c.AbortWithStatusJSON(errors.TokenInvalid.Status, errors.TokenInvalid)
				return
			}

			// Check refresh token identifier matches the users session identifier
			UserID, ok := refreshClaims["Id"].(uint)
			if !ok {
				fmt.Println("refreshClaims['Id'].(uint) -> Type assetion error")
			}
			userSessionIdentifier, ok := refreshClaims["Identifier"].(string)
			if !ok {
				fmt.Println("refreshClaims['Identifier'].(uint) -> Type assetion error")
			}

			usersession := models.UserSession{
				UserID:     UserID,
				Identifier: userSessionIdentifier,
			}

			db := db.GetDB()
			if err := db.Find(&usersession).Error; err != nil {
				fmt.Println("Invlaid identifier")

				c.SetCookie("refreshToken", "", -1, "/", "localhost", true, false)
				c.AbortWithStatusJSON(errors.InvalidIdentifier.Status, errors.InvalidIdentifier)
				return
			}
		}

		// If auth token is valid check if crsf token matches the one supplied
		// in the header
		if authClaims["csrf"] != c.GetHeader("X-CSRF-Token") {
			c.AbortWithStatusJSON(errors.CSRFTokenInvalid.Status, errors.CSRFTokenInvalid)
			return
		}
	}
}
