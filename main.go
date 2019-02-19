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
	//checkUrls()
}

func checkUrls() {
	for index, url := range getUrls() {
		soldOut, notSoldOut := checker.GetSoldOutTobaccos(url)
		message := "%d %s tobaccos available out of %d tobaccos \n Sold out: \n ```%s``` \n Available: ```%s```"
		notifier.SendMessage(fmt.Sprintf(message, len(notSoldOut), index, len(notSoldOut) + len(soldOut), strings.Join(soldOut, "\n"), strings.Join(notSoldOut, "\n")))
	}
}

func getUrls() map[string]string {
	m := make(map[string]string)
	m["Fumari"] = "https://www.smoking-shisha.de/542-fumari"
	m["Zomo"] = "https://www.smoking-shisha.de/suche?search_query=zomo%2C+200g&orderby=position&orderway=desc&search_query=zomo%2C+200g&submit_search=&n=48"
	return m
}