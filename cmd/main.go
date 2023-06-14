//Reference:- https://betterprogramming.pub/build-a-scalable-api-in-go-with-gin-131af7f780c0

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prag-sri/go-gin-api-medium/pkg/books"
	"github.com/prag-sri/go-gin-api-medium/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")

	//Sets the configuration file path to the specified file.

	viper.ReadInConfig()

	//Reads the configuration values from the file into the Viper configuration object.

	port := viper.Get("PORT").(string)

	//Retrieves the value associated with the key "PORT" from the configuration.

	dbUrl := viper.Get("DB_URL").(string)

	//Retrieves the value associated with the key "DB_URL" from the configuration.
	//.Get() returns an interface{}, so type assertions (string) are used to convert the values to the string type.

	r := gin.Default()

	//Creates a new Gin router with default middleware, including a logger and recovery middleware.

	h := db.Init(dbUrl)

	//Initializes the database connection by calling the Init() function from the db package, passing the dbUrl as the connection URL.

	/*This block is removed in modification

	r.GET("/", func(c *gin.Context) {

		//r.GET("/", ...) registers a handler function for a GET request to the root path. The handler function is defined as an anonymous function that takes a *gin.Context object as a parameter.

		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})

		//Inside the handler function, a JSON response is sent back to the client with a status code of 200 (OK). The response body contains a JSON object with two key-value pairs: "port" and "dbUrl", which represent the values retrieved from the configuration settings.
	})

	*/

	books.RegisterRoutes(r, h)

	r.Run(port)

	//r.Run(port) starts the server
}
