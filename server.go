package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Make sure the fields of your struct are exported (start with an uppercase letter). If a field is unexported (starts with a lowercase letter), it won't be accessible for JSON serialization.
// if you want to send isError instead of IsError then you should use JSON Tags, for example  `json: "isError"`
type Response struct {
	IsError bool   `json:"isError"`
	Msg     string `json:"msg"`
}

type APIConfig struct {
	BaseURL string `mapstructure:"base_url"`
	Timeout int    `mapstructure:"timeout"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Not able to load env")
	}

	//read env and load it
	readEnvs() // remember suppose if you declate same variable in config and env file then it will ignore the config and return env file. so it means env file has higher precedents.
	readConfigs("./configs", "app.config", "json", true)
	readConfigs("./configs", "app.log", "json", false)

	//get value form loaded envs
	port := viper.GetString("PORT")
	dbConnectionString := viper.GetString("DATABASE_CONNECTION")
	fmt.Println("db Connection string from envs ", dbConnectionString)

	//get value from the configs
	fmt.Println("appname from config : ", viper.GetString("app.name"))
	fmt.Println("appname from config : ", viper.GetString("logging.level"))

	// read whole object and bind to variable
	var apiConfig APIConfig

	err = viper.UnmarshalKey("api", &apiConfig)
	if err != nil {
		fmt.Printf("Error unmarshaling API configuration: %s\n", err)
		return
	}

	fmt.Printf("API Base URL: %v", apiConfig)

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

func readEnvs() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Access environment variables using Viper
	viper.AutomaticEnv()
}

func readConfigs(path, filename, extension string, isFirstconfig bool) {
	// Set up Viper to read configuration from a file
	viper.SetConfigName(filename)
	viper.SetConfigType(extension)
	viper.AddConfigPath(path)

	// Read configuration from the file
	var err error
	if isFirstconfig {
		err = viper.ReadInConfig()
	} else {
		err = viper.MergeInConfig()
	}
	if err != nil {
		fmt.Printf("Error reading configuration file: %s\n", err)
		return
	}

}
