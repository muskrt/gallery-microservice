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

var login_server = os.Getenv("LOGIN_SERVER")

func loginController(ctx *gin.Context) {
	username, _ := ctx.Params.Get("username")
	password, _ := ctx.Params.Get("password")
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
	// server.GET("/check/:username/:password", loginController)
	return server
}
func main() {

	server := newServer()

	server.Run(":30200")

}
