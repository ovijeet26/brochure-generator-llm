package main

import (
	"fmt"
	"log"

	"github.com/ovijeet26/brochure-generator-llm/src/core"
)

func main() {
	// Load environment variables.
	core.LoadEnv()

	// Add the website URL you want to scrape.
	url := "https://ollama.com/"

	// Get the page content.
	website, err := core.ScrapeWebsite(url)
	if err != nil {
		log.Fatal(err)
	}

	// Get relevant links in the website.
	linkResp, err := core.GetRelevantLinks(website)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating brochure...")

	brochure, err := core.CreateBrochure(website.Title, url, linkResp.Links)
	if err != nil {
		log.Fatalf("Failed to create brochure: %v", err)
	}

	fmt.Println("Generated Brochure:\n")
	fmt.Println(brochure)

	err = core.ExportBrochureAsHTML(brochure, "brochure.html")
	if err != nil {
		log.Fatalf("Failed to export HTML: %v", err)
	}
	fmt.Println("✅ HTML brochure saved as brochure.html")

}
