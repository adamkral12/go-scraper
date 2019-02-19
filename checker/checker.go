package checker

import (
	"github.com/gocolly/colly"
)


func GetSoldOutTobaccos(url string) ([]string, []string) {
	var soldOutUrls []string
	var notSoldOutUrls []string
	c := colly.NewCollector()
	c.OnHTML(".product-image-container", func(e *colly.HTMLElement) {
		tobaccoUrl, exists := e.DOM.Find(".product_img_link").Attr("href")
		if exists {

			_, soldOutExists := e.DOM.Find(".soldout-box").Attr("href")

			if soldOutExists {
				soldOutUrls = append(soldOutUrls, tobaccoUrl)
			} else {
				notSoldOutUrls = append(notSoldOutUrls, tobaccoUrl)
			}
		}
	})

	err := c.Visit(url)

	if err != nil {
		panic(err)
	}

	return soldOutUrls, notSoldOutUrls
}