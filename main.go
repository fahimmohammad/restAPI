package main

import (
	"fmt"
	"time"

	"github.com/fahimsGit/restAPI/article"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func main() {
	router := gin.New()
	router.Use(Logger())
	v1 := router.Group("api/v1")
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))
	fmt.Println("Startder")
	router.Run(":8080")
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect")
		return
	}
	fmt.Println("DB server started")

	initializeServices(v1, session)
	router.Run(":8080")
}

func initializeServices(router *gin.RouterGroup, session *mgo.Session) {

	article.Init(router, session)

}
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("======================================>")
		fmt.Println("Url Hit : " + c.Request.URL.String() + " Method : " + c.Request.Method)
		c.Next()
		since := time.Since(t)
		fmt.Println("Time Took : " + since.String())
	}
}
