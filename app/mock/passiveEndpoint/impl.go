package passiveEndpoint

import (
	"awesomeProject/app/mock/reflection"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ServeFromReflection(c *gin.Context) {
	var request RegisterEndpointRequest
	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		reflect := reflection.ReflectAndDescribe(request.DataPattern)

		data := reflection.GenerateReflectionData(reflect, 5)

		fmt.Println("NEW")
		fmt.Println(data)
		fmt.Println("NEW")

		res := RegisterEndpointResponse{
			EndpointName: request.EndpointName,
			Description:  request.Description,
			DataPattern:  reflect,
		}

		c.JSON(http.StatusCreated, res)
	}
}
