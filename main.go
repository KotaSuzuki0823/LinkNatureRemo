package main

import (
	"./linkslack"
	"./nuture"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

var (
	apitoken     = os.Getenv("SLACKAPI_TOKEN")
	slackChannel = "general"
)
func service(){
	result := nuture.Curltest()
	//nuture.Dummy()
	err := linkslack.SendMessage(string(result), apitoken, slackChannel)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}
