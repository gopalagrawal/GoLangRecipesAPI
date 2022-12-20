package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)


type Recipe struct {
	ID 				string 		`json:"id"`
	Name 			string 		`json:"name"`
	Tags 			[]string 	`json:"tags"`
	Ingredients 	[]string 	`json:"ingredients"`
	Instructions 	[]string 	`json:"instructions"`
	PublishedAt 	time.Time 	`json:"publishedAt"`
}

var recipes []Recipe
func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}


// POST: /recipes
// ID and PublishedAt are auto-generated.
func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

// GET: /recipes
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func main() {
	fmt.Println("Starting Server on localhost:8080 ... ")
	router := gin.Default()

	// Create routes
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)


	router.Run(":8080")
}