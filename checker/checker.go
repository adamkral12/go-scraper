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
			return false
		case tt == html.StartTagToken:
			t := htmlContent.Token()

			isAnchor := t.Data == "span"

			if !isAnchor {
				continue
			}

			foundId := false
			foundClass := false
			for _, val := range t.Attr {
				if val.Key == "id" && val.Val == "availability_value" {
					foundId = true
				}

				if val.Key == "class" && val.Val == "warning_inline" {
					foundClass = true
				}
			}

			if foundId && foundClass {
				return true
			}
		}
	}
}