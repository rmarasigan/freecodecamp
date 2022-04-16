package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Commannd Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}

func main() {
	// Setting environment variable
	os.Setenv("SLACK_BOT_TOKEN", "xxxx-your-token-here")
	os.Setenv("SLACK_APP_TOKEN", "xxxx-socket-token-here")

	// Creating a bot
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Go routine that will print the command events
	// Passing a command to a slack bot
	go printCommandEvents(bot.CommandEvents())

	// <year> is the parameter
	bot.Command("My year of birth if <year>", &slacker.CommandDefinition{
		Description: "Year of birth calculator",
		Example:     "My year of birth is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")

			// Convert string to int
			yearOfBirth, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}

			age := time.Now().Year() - yearOfBirth
			reply := fmt.Sprintf("Age is %d", age)
			response.Reply(reply)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
