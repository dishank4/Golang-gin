package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Make sure the fields of your struct are exported (start with an uppercase letter). If a field is unexported (starts with a lowercase letter), it won't be accessible for JSON serialization.
// if you want to send isError instead of IsError then you should use JSON Tags, for example  `json: "isError"`
type Response struct {
	IsError bool   `json:"isError"`
	Msg     string `json:"msg"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Not able to load env")
	}

	port := os.Getenv("PORT")
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		res := Response{
			Msg:     "This is our first route",
			IsError: false,
		}
		c.JSON(http.StatusOK, res) // you can give response in json like this, first argument http status code and 2nd is struct
	}) //you should be do " " only ' consider as rune in go

	router.GET("/gitDefaultMap", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "using default map[string]interface{}", "isError": false}) // gin.H is a shorthand type for a map[string]interface{}
	})
	router.Run(port)
}
