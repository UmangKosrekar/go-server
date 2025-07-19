package controllers

import (
	"first_server/helper"

	"github.com/gin-gonic/gin"

	"net/http"
)

func ListUsers(c *gin.Context) {
	type Response struct {
		List  []interface{} `json:"list"`
		Count int           `json:"count"`
	}

	response := Response{
		List:  []interface{}{},
		Count: 0,
	}

	helper.ResponseHandler(c, http.StatusOK, "Success", response)
}

func GetUser(c *gin.Context) {
	var id = c.Param("id")
	type Response struct {
		id string
	}

	response := Response{id: id}

	helper.ResponseHandler(c, http.StatusOK, "Success", response)
}

func CreateUser(c *gin.Context) {
	type Response struct {
	}

	response := Response{}

	helper.ResponseHandler(c, http.StatusCreated, "Users Created Successfully", response)
}

func UpdateUser(c *gin.Context) {
	type Response struct {
	}

	response := Response{}

	helper.ResponseHandler(c, http.StatusOK, "Users Updated Successfully", response)
}

func DeleteUser(c *gin.Context) {

	helper.ResponseHandler(c, http.StatusOK, "Users Deleted Successfully", nil)
}
