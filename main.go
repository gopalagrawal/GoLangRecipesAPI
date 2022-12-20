
package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Server on localhost:8080 ... ")
	router := gin.Default()
	router.Run(":8080")
}