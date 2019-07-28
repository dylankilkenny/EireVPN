package router

import (
	"io/ioutil"
	"net/http"

	"eirevpn/api/plan"
	"eirevpn/api/user"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
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
	// private := router.Group("/api/private")
	// secret := router.Group("/api/secret")

	public.POST("/signup", user.SignUpUser)
	public.POST("/login", user.LoginUser)
	// public.POST("/validate", user.ValidateToken)
	public.GET("/plan/:id", plan.Plan)
	public.POST("/plan", plan.CreatePlan)
	public.PUT("/plan", plan.UpdatePlan)
	public.DELETE("/plan/:id", plan.DeletePlan)
	public.GET("/plans", plan.AllPlans)
	// private.Use(auth(secretkey))
	// private.GET("/address", user.GetSubscribedAddresses)
	// private.POST("/address", user.SubscribeToAddress)
	// private.DELETE("/remove", user.RemoveSubscribedAddress)
	// secret.POST("/users", user.SubscribedUsers)

	return router
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
