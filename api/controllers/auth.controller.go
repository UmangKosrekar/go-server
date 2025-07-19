package controllers

import (
	"first_server/helper"
	"first_server/validations"
	"log"

	"first_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Request Body
	var requestBody = c.MustGet("body").(validations.LoginStruct)

	helper.ResponseHandler(c, http.StatusOK, "Success", requestBody)
}

func Register(c *gin.Context) {
	// Request Body
	var requestBody = c.MustGet("body").(validations.RegisterStruct)

	document, err := model.GetUserCollection().InsertOne(c, requestBody)
	if err != nil {
		helper.ResponseHandler(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	log.Println(document)

	helper.ResponseHandler(c, http.StatusOK, "Success", requestBody)
}
