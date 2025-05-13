package core

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// ExportBrochureAsHTML writes the LLM-generated HTML brochure to a file
func ExportBrochureAsHTML(htmlContent string, filename string) error {
	return os.WriteFile(filename, []byte(htmlContent), 0644)
}

func ExportBrochureMDAsHTML(markdownContent, filename, primaryColor string) error {
	// Set markdown parser options
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.HardLineBreak
	parser := parser.NewWithExtensions(extensions)

	// Set HTML renderer options
	htmlOpts := html.RendererOptions{
		Flags: html.CommonFlags | html.UseXHTML,
	}
	renderer := html.NewRenderer(htmlOpts)

	// Convert markdown to HTML
	htmlBody := markdown.ToHTML([]byte(markdownContent), parser, renderer)

	htmlTemplate := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Company Brochure</title>
	<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700&display=swap" rel="stylesheet">
	<style>
		body {
			font-family: 'Inter', sans-serif;
			background-color: #f9f9fb;
			color: #222;
			max-width: 800px;
			margin: 40px auto;
			padding: 40px;
			box-shadow: 0 0 20px rgba(0,0,0,0.05);
			border-radius: 12px;
			background: white;
			white-space: pre-wrap;
		}
		h1, h2, h3 {
			color: %s;
		}
		a {
			color: %s;
		}
		a:hover {
			text-decoration: underline;
		}
	</style>
</head>
<body>
%s
</body>
</html>`, primaryColor, primaryColor, htmlBody)

	return os.WriteFile(filename, []byte(htmlTemplate), 0644)
}

// DetectPrimaryColor tries to extract a primary color from meta tags or CSS
func DetectPrimaryColor(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Try to find <meta name="theme-color" ...>
	color := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); strings.ToLower(name) == "theme-color" {
			if c, ok := s.Attr("content"); ok {
				color = c
			}
		}
	})

	// Fallback if not found
	if color == "" {
		color = "#007acc" // default blue
	}

	return color, nil
}
