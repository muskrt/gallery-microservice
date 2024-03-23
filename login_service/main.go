package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()

	fmt.Println("server started at 30300")
	server.Run(":30300")

}
