package main

import (
	"./checker"
	"fmt"
)

func main() {
	for _, url := range getUrlsToCheck() {
		if !(checker.IsSoldOut(url)) {
			fmt.Println(url)
		}
	}
}

func getUrlsToCheck() []string {
	return []string{
		"https://www.smoking-shisha.de/fumari/722-fumari-tobacco-original-white-bear-100g.html?search_query=fumari&results=342",
		"https://www.smoking-shisha.de/fumari/713-fumari-tobacco-wild-beach-100g.html?search_query=fumari&results=342",
	}
}