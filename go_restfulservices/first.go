package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"username"`
	Age  int    `json:"userage"`
}

func main() {
	router := gin.Default()

	user := User{
		Name: "John",
		Age:  20,
	}
	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, user)
	})
	var port string = ":9999"
	fmt.Println("Server listening on ", port)
	router.Run(port)
}
