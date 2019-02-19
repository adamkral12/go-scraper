package checker

import (
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"net/http"
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


func IsSoldOut(url string) bool {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	htmlContent := html.NewTokenizer(resp.Body)

	for {
		tt := htmlContent.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return false
		case tt == html.StartTagToken:
			t := htmlContent.Token()

			isAnchor := t.Data == "div"

			if !isAnchor {
				continue
			}

			foundId := false
			foundClass := false
			for _, val := range t.Attr {
				if val.Key == "class" && val.Val == "product-image-container" {
					//html.ParseFragment(tt, )
				}
			}

			if foundId && foundClass {
				return true
			}
		}
	}
}