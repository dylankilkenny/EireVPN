package server

import (
	"eirevpn/api/errors"
	"eirevpn/api/logger"
	"eirevpn/api/models"
	"net/http"
	"strconv"

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
			"connection": server,
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
		IP   string `json:"ip" binding:"required"`
		Port int    `json:"port" binding:"required"`
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
	if err := server.Save(); err != nil {
		logger.Log(logger.Fields{
			Loc:   "/server/update/:id - UpdateServer()",
			Code:  errors.InternalServerError.Code,
			Extra: map[string]interface{}{"ServerID": server.ID},
			Err:   err.Error(),
		})
		c.AbortWithStatusJSON(errors.InternalServerError.Status, errors.InternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
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

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"servers": servers,
		},
	})
}
