package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var api_server = os.Getenv("API_SERVER")

func loginGetController(ctx *gin.Context) {
	fmt.Println("api server is ", api_server)
	ctx.HTML(http.StatusOK, "login.html", nil)
}

type loginResponse struct {
	Message string `json:"message"`
}

func notesGetController(ctx *gin.Context) {
	fmt.Println("got the request")
	username, _ := ctx.Params.Get("username")

	url := strings.TrimSpace("http://" + api_server + ":30200/notes/" + username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Body)
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
	fmt.Println(res.Message)
	ctx.JSON(http.StatusOK, gin.H{"message": res.Message})

}
func loginPostController(ctx *gin.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println("username is -------" + username)
	fmt.Println("password is ------" + password)

	url := strings.TrimSpace("http://" + api_server + ":30200/user/check/" + username + "/" + password)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Body)
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
		ctx.HTML(http.StatusOK, "success.html", gin.H{"username": username})
	} else if res.Message == "error" {
		ctx.HTML(http.StatusOK, "error.html", nil)
	}

}

func newServer() *gin.Engine {
	server := gin.New()
	server.LoadHTMLGlob("templates/*.html")
	server.GET("/login", loginGetController)
	server.POST("/login", loginPostController)
	server.GET("/notes/:username", notesGetController)
	return server

}

func main() {

	server := newServer()

	fmt.Println("landing service starting on 8080")

	server.Run(":8080")

}
