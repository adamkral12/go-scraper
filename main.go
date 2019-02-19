package main

import (
	"fmt"
	"github.com/adamkral12/go-scraper/checker"
	"github.com/adamkral12/go-scraper/notifier"
)

func main() {
	//lambda.Start(checkUrls)
	checkUrls()
}

func checkUrls() {
	for _, url := range getUrlsToCheck() {
		if !(checker.IsSoldOut(url)) {
			notifier.SendMessage(fmt.Sprintf("Tobbaco is no longer sold out! %s.", url))
		}
	}
}

func getUrlsToCheck() []string {
	return []string{
		"https://www.smoking-shisha.de/fumari/722-fumari-tobacco-original-white-bear-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/713-fumari-tobacco-wild-beach-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/715-fumari-tobacco-original-red-bear-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/717-fumari-tobacco-amorosa-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/2144-fumari-tobacco-ceylon-chi-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/720-fumari-tobacco-tango-melange-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/714-fumari-tobacco-golden-dynasty-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/2143-fumari-tobacco-maison-bleu-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/2142-fumari-tobacco-lumin-freeze-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/721-fumari-tobacco-tropical-sunset-100g.html?search_query=fumari&results=342",
	}
}