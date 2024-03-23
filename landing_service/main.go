package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginGetController(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", nil)
}

type loginResponse struct {
	Message string `json:"message"`
}

func loginPostController(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println("username is -------" + username)
	resp, _ := http.Get("login_service/check/" + username + "/" + password)
	res := &loginResponse{}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}
	if res.Message == "success" {
		ctx.HTML(http.StatusOK, "success.html", nil)
	}
	ctx.HTML(http.StatusOK, "error.html", nil)

}

func main() {

	server := gin.New()

	server.LoadHTMLGlob("templates/*.html")
	server.GET("/login", loginGetController)
	server.POST("/login", loginPostController)

	fmt.Println("landing service starting on 8080")

	server.Run(":8080")

}
