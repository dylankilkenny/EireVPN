package server

import (
	"bytes"
	"eirevpn/api/config"
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"encoding/json"
	"fmt"

	"eirevpn/api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Server fetches a server by ID
func Server(c *gin.Context) {
	serverID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var server models.Server
	server.ID = uint(serverID)
	if err := server.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/:id - Server()",
			Code:  errors.ServerNotFound.Code,
			Extra: map[string]interface{}{"ConnID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.ServerNotFound.Status, errors.ServerNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"server": server,
		},
	})

}

// CreateServer creates a new server
func CreateServer(c *gin.Context) {
	var server models.Server

	if err := c.ShouldBind(&server); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/server/create - CreateServer()",
			Code: errors.InvalidForm.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	file, _ := c.FormFile("img")
	if file != nil {

		err := c.SaveUploadedFile(file, "assets/"+file.Filename)
		if err != nil {
			logger.Log(logger.Fields{
				Loc:  "/server/create - CreateServer()",
				Code: errors.InternalServerError.Code,
				Err:  err.Error(),
			})
			c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
			return
		}

		server.ImagePath = "assets/" + file.Filename
	}

	if err := server.Create(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/server/create - CreateServer()",
			Code: errors.InternalServerError.Code,
			Extra: map[string]interface{}{
				"Server": server,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"server": server,
		},
	})
}

// DeleteServer deletes a given server. It will not delete a server fully however,
// it will just set a DeletedAt datetime on the record
func DeleteServer(c *gin.Context) {
	serverID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var server models.Server
	server.ID = uint(serverID)
	if err := server.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/delete/:id - DeleteServer()",
			Code:  errors.ServerNotFound.Code,
			Extra: map[string]interface{}{"ServerID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.ServerNotFound.Status, errors.ServerNotFound)
		return
	}

	if err := server.Delete(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/delete/:id - DeleteServer()",
			Code:  errors.ServerNotFound.Code,
			Extra: map[string]interface{}{"ServerID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}

// UpdateServer updates an existing server
func UpdateServer(c *gin.Context) {
	ServerID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var server models.Server
	server.ID = uint(ServerID)

	type ServerUpdates struct {
		IP       string `json:"ip" binding:"required"`
		Port     int    `json:"port" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	serverUpdates := ServerUpdates{}

	if err := server.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/update/:id - UpdateServer()",
			Code:  errors.ServerNotFound.Code,
			Extra: map[string]interface{}{"ServerID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.ServerNotFound.Status, errors.ServerNotFound)
		return
	}

	if err := c.BindJSON(&serverUpdates); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/update/:id - UpdateServer()",
			Code:  errors.InvalidForm.Code,
			Extra: map[string]interface{}{"ServerID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InvalidForm.Status, errors.InvalidForm)
		return
	}

	server.IP = serverUpdates.IP
	server.Port = serverUpdates.Port
	server.Username = serverUpdates.Username
	server.Password = serverUpdates.Password
	if err := server.Save(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/update/:id - UpdateServer()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"ServerID": server.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	// update proxy server with new config
	if !config.Load().App.TestMode {
		var credentials = map[string]string{
			"username": server.Username,
			"password": server.Password,
		}
		jsonStr, _ := json.Marshal(credentials)
		url := fmt.Sprintf("http://%s:%v/update_creds", server.IP, 3003)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		if _, err := client.Do(req); err != nil {
			logger.Log(logger.Fields{
				Loc:   "/server/update/:id - UpdateServer()",
				Code:  errors.InternalServerError.Code,
				Extra: map[string]interface{}{"ServerID": server.ID, "Detail": "Error posting new credentials to proxy servers"},
				Err:   err.Error(),
			})
			c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}

// Connect returns a username and password for the server if the user
// has a valid subscription
func Connect(c *gin.Context) {
	serverID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var server models.Server
	server.ID = uint(serverID)
	if err := server.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/:id - Server()",
			Code:  errors.ServerNotFound.Code,
			Extra: map[string]interface{}{"ConnID": c.Param("id")},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.ServerNotFound.Status, errors.ServerNotFound)
		return
	}

	userID, exists := c.Get("UserID")
	if !exists {
		logger.Log(logger.Fields{
			Loc: "/server/connect/:id - Connect()",
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": "User ID does not exist in the context",
			},
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	conf := config.Load()
	if conf.App.EnableSubscriptions {
		var userplan models.UserPlan
		userplan.UserID = userID.(uint)
		if err := userplan.Find(); err != nil {
			logger.Log(logger.Fields{
				Loc:  "/server/connect/:id - Connect()",
				Code: errors.InternalServerError.Code,
				Extra: map[string]interface{}{
					"UserID": userplan.UserID,
					"Detail": "Could not find user_plan record",
				},
				Err: err.Error(),
			})
			c.AbortWithStatusJSON(errors.UserPlanNotFound.Status, errors.UserPlanNotFound)
			return
		}
		userPlanExpired := userplan.ExpiryDate.Before(time.Now())
		if !userplan.Active || userPlanExpired {
			logger.Log(logger.Fields{
				Loc:  "/server/connect/:id - Connect()",
				Code: errors.InternalServerError.Code,
				Extra: map[string]interface{}{
					"UserID": userplan.UserID,
				},
				Err: errors.UserPlanExpired.Detail,
			})
			c.AbortWithStatusJSON(errors.UserPlanExpired.Status, errors.UserPlanExpired)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"username": server.Username,
			"password": server.Password,
			"port":     server.Port,
			"ip":       server.IP,
		},
	})

}

// AllServers returns an array of all available servers
func AllServers(c *gin.Context) {
	var servers models.AllServers

	if err := servers.FindAll(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/servers - AllServers()",
			Code: errors.InternalServerError.Code,
			Err:  err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	userID, exists := c.Get("UserID")
	if !exists {
		logger.Log(logger.Fields{
			Loc: "/servers - AllServers()",
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": "User ID does not exist in the context",
			},
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
		return
	}

	var user models.User
	user.ID = userID.(uint)
	if err := user.Find(); err != nil {
		logger.Log(logger.Fields{
			Loc:  "/servers - AllServers()",
			Code: errors.UserNotFound.Code,
			Extra: map[string]interface{}{
				"UserID": userID,
				"Detail": errors.UserNotFound.Detail,
			},
			Err: err.Error(),
		})
		c.AbortWithStatusJSON(errors.UserNotFound.Status, errors.UserNotFound)
		return
	}

	if user.Type != models.UserTypeAdmin {
		// dont send username and passwords
		for i, s := range servers {
			s.Username = ""
			s.Password = ""
			s.IP = ""
			s.Port = 0000
			servers[i] = s
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"servers": servers,
		},
	})
}
