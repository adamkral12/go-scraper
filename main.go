package main

import (
	"fmt"
	"github.com/adamkral12/go-scraper/checker"
	"github.com/adamkral12/go-scraper/notifier"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

func main() {
	lambda.Start(checkUrls)
}

func checkUrls() {
	for _, url := range getUrls() {
		soldOut, notSoldOut := checker.GetSoldOutTobaccos(url)
		message := "%d Fumari tobaccos available out of %d tobaccos \n Sold out: \n ```%s``` \n Available: ```%s```"
		notifier.SendMessage(fmt.Sprintf(message, len(notSoldOut), len(notSoldOut) + len(soldOut), strings.Join(soldOut, "\n"), strings.Join(notSoldOut, "\n")))
	}
}

func getUrls() []string {
	return []string{
		"https://www.smoking-shisha.de/542-fumari",
	}
}