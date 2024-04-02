package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userNotes struct {
	Note string `json:"note"`
	Todo string `json:"todo"`
	Date string `json:"date"`
}

func notesGetController(ctx *gin.Context) {
	username, _ := ctx.Params.Get("username")

	data := []userNotes{
		userNotes{
			Note: string(username + "test"), Todo: "yes", Date: "today"},
		userNotes{
			Note: string(username + "quiz"), Todo: "yes", Date: "tomorrow"}}
	usernotes, _ := json.Marshal(data)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": string(usernotes), // cast it to string before showing
	})

}

func newServer() *gin.Engine {
	server := gin.New()
	server.GET("/notes/:username", notesGetController)

	return server

}

func main() {

	server := newServer()

	fmt.Println("landing service starting on 30300")

	server.Run(":30300")

}
