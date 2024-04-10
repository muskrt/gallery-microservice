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

type loginResponse struct {
	Message string `json:"message"`
}

type Notes struct {
	Note string `json:"note"`
	Todo string `json:"todo"`
	Date string `json:"date"`
}

var login_server = os.Getenv("LOGIN_SERVER")
var notes_server = os.Getenv("NOTES_SERVER")

func loginController(ctx *gin.Context) {
	username, _ := ctx.Params.Get("username")
	password, _ := ctx.Params.Get("password")
	fmt.Println(username, password)
	if username == "test_user" && password == "toor" {
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})

	} else if username == "test_user" && password == "root" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error"})

	} else {
		fmt.Println("username is -------" + username)
		fmt.Println("login server is ", login_server)
		url := strings.TrimSpace("http://" + login_server + ":30100/check/" + username + "/" + password)
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		res := &loginResponse{}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(body, &res)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": res.Message})

	}

}
func notesController(ctx *gin.Context) {
	username, _ := ctx.Params.Get("username")
	fmt.Println("notes server is ", notes_server)
	url := strings.TrimSpace("http://" + notes_server + ":30300/notes/" + username)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	res := &loginResponse{}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(body, &res)
	usernotes := []Notes{}
	err = json.Unmarshal([]byte(res.Message), &usernotes)
	fmt.Printf("%v", usernotes)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": res.Message})

}
func newServer() *gin.Engine {
	server := gin.New()

	fmt.Println("server started at 30200")
	user_service := server.Group("/user")
	{
		user_service.GET("/check/:username/:password", loginController)
	}
	notes_service := server.Group("/notes")
	{
		notes_service.GET("/:username", notesController)
	}
	// server.GET("/check/:username/:password", loginController)
	return server
}
func main() {

	server := newServer()

	server.Run(":30200")

}
