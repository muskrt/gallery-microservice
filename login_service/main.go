package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginController(ctx *gin.Context) {
	auth := false
	username, _ := ctx.Params.Get("username")
	password, err := ctx.Params.Get("password")
	if !err {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	fmt.Println("username is -----", username)

	fmt.Println("password is -----", password)
	if username == "admin" && password == "toor" {
		auth = true
	}
	if auth {
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error"})
	}
}

func newServer() *gin.Engine {
	server := gin.New()

	fmt.Println("server started at 30300")
	server.GET("/check/:username/:password", loginController)
	return server
}
func main() {

	server := newServer()

	server.Run(":30300")

}
