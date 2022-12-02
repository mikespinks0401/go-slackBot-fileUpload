package main

import (
	"fmt"
	"log"
	"os"
	"github.com/slack-go/slack"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Token")
	}

	slackbotToken := os.Getenv("SLACKBOTTOKEN")
	channelId := os.Getenv("CHANNELID")

	api := slack.New(slackbotToken)
	channelArr := []string{channelId}
	fileArr := []string{"dailyStoic.pdf"}

	for i := 0; i<len(fileArr);i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err !=nil{
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}