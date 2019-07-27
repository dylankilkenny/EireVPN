package router

import (
	"io/ioutil"
	"net/http"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"eirevpn/api/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const secretkey = "verysecretkey1995"

func SetupRouter(db *gorm.DB, logging bool) *gin.Engine {
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

	router.Use(apiMiddleware(db))

	public := router.Group("/api")
	// private := router.Group("/api/private")
	// secret := router.Group("/api/secret")

	public.POST("/signup", user.SignUpUser)
	// public.POST("/validate", user.ValidateToken)
	public.POST("/login", user.LoginUser)
	// private.Use(auth(secretkey))
	// private.GET("/address", user.GetSubscribedAddresses)
	// private.POST("/address", user.SubscribeToAddress)
	// private.DELETE("/remove", user.RemoveSubscribedAddress)
	// secret.POST("/users", user.SubscribedUsers)

	return router
}

func apiMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid token",
			})
		}
	}
}
