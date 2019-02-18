package checker

import (
	"golang.org/x/net/html"
	"net/http"
)

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
			return true
		case tt == html.StartTagToken:
			t := htmlContent.Token()

			isAnchor := t.Data == "span"

			if !isAnchor {
				continue
			}

			for _, val := range t.Attr {
				if val.Key == "id" && val.Val == "availability_value" {
					if val.Key == "class" && val.Val == "warning_inline" {
						return true
					} else {
						return false
					}
				}
			}
		}
	}
}