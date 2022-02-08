package passiveEndpoint

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterEndpointHandler() func(*gin.Context) {
	return func(c *gin.Context) {

		var request RegisterEndpointRequest

		decoder := json.NewDecoder(c.Request.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&request)


		Register(&request)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusCreated, "true")
		}
	}
}
