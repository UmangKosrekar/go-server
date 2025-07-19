package middleware

import (
	"first_server/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validator[localStruct any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody localStruct

		// validate JSON format
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			helper.ResponseHandler(c, http.StatusBadRequest, "Invalid JSON format", nil)
			return
		}

		// Validate struct

		validate := validator.New()

		if err := validate.Struct(requestBody); err != nil {
			fmt.Print(err)
			helper.ResponseHandler(c, http.StatusBadRequest, "Validation Failed", nil)
			return
		}

		c.Set("body", requestBody)
		c.Next()
	}
}
