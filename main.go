package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


type Recipe struct {
	Name 			string 		`json:"name"`
	Tags 			[]string 	`json:"tags"`
	Ingredients 	[]string 	`json:"ingredients"`
	Instructions 	[]string 	`json:"instructions"`
	PublishedAt 	time.Time 	`json:"publishedAt"`
	}





func main() {
	fmt.Println("Starting Server on localhost:8080 ... ")
	router := gin.Default()
	router.Run(":8080")
}