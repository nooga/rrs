package main

import (
	"fmt"

	"github.com/Lachann/rrs/pkg/server"
)

var defaultPreviews = []string{
	"Slackbot",
	"Discordbot",
}

const defaultHost = ":8080"

const defaultRickroll = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

func main() {
	if err := server.Start(defaultHost, defaultRickroll, defaultPreviews); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
