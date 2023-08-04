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

const defaultFake = "https://docs.monk.io"

const defaultRickroll = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

const defaultRoute = "/docs.monk.io"

func main() {
	if err := server.Start(defaultHost, defaultRoute, defaultRickroll, defaultFake, defaultPreviews); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
