package core

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetPageContent returns the text content for a given URL (cleaned).
func GetPageContent(url string) (string, error) {
	website, err := ScrapeWebsite(url)
	if err != nil {
		return "", err
	}

	content := fmt.Sprintf("URL: %s\nTitle: %s\n\n%s", website.URL, website.Title, website.Text)
	return content, nil
}

// ScrapeWebsite fetches and parses the content of a URL
func ScrapeWebsite(url string) (*Website, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return scrapeWebsiteContent(url, resp)
}

// scrapeWebsiteContent is a helper function to extract content and links from the HTML response.
func scrapeWebsiteContent(url string, resp *http.Response) (*Website, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	title := strings.TrimSpace(doc.Find("title").First().Text())
	if title == "" {
		title = "No title found"
	}

	doc.Find("script, style, img, input, noscript, iframe, footer, header, nav, aside, .cookie-banner, .modal, .popup").Remove()

	main := doc.Find("main")
	if main.Length() == 0 {
		main = doc.Find("body")
	}
	bodyText := strings.TrimSpace(main.Text())

	var links []string
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && href != "" {
			links = append(links, href)
		}
	})

	return &Website{
		URL:   url,
		Title: title,
		Text:  bodyText,
		Links: links,
	}, nil
}
