package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginGetController(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", nil)
}

func loginPostController(ctx *gin.Context) {

	username := ctx.PostForm("username")
	fmt.Println("username is -------" + username)
	//resp, _ := http.Get("login_service/"+username+"&"+password)
	// fmt.Println(sucess.Close)
	ctx.HTML(http.StatusOK, "success.html", nil)
}
func main() {

	server := gin.New()

	server.LoadHTMLGlob("templates/*.html")
	server.GET("/login", loginGetController)
	server.POST("/login", loginPostController)

	fmt.Println("landing service starting on 8080")

	server.Run(":8080")

}
