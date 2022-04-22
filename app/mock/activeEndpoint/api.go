package activeEndpoint

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterEndpointHandler() func(*gin.Context) {
	return func(c *gin.Context) {

		var request RegisterActiveEndpointRequest

		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&request)

		res := Register(&request)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusCreated, res)
		}
	}
}

func RegisteredEndpointsHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		res := GetRegistered()
		c.JSON(http.StatusCreated, res)
	}
}

func RunCmdhandler() func(*gin.Context) {
	return func(c *gin.Context) {

		var request RunActiveEndpointCommandRequest

		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&request)

		RunCmd(&request)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusCreated, "true")
		}
	}
}
