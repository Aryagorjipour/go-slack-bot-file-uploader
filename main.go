package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	os.Setenv("SLACK-BOT-TOKEN", "xoxb-3815292028437-3826628254999-6M46PT05ILclKCT3AVxhgAO4")
	os.Setenv("CHANNEL_ID", "C03Q25DE15H")
	api := slack.New(os.Getenv("SLACK-BOT-TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"resume.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
