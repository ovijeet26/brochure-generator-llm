package core

import (
	"fmt"
	"strings"
)

func GetLinkSystemPrompt() string {
	return `You are provided with a list of links found on a webpage. 
You are able to decide which of the links would be most relevant to include in a brochure about the company, 
such as links to an About page, or a Company page, or Careers/Jobs pages.
You should respond respond ONLY with a raw JSON object (no markdown, no explanation, no backticks) as in this example:
{
    "links": [
        {"type": "about page", "url": "https://full.url/goes/here/about"},
        {"type": "careers page", "url": "https://another.full.url/careers"}
    ]
}`
}

func GetLinksUserPrompt(website *Website) string {
	return fmt.Sprintf("Here is the list of links on the website of %s - please decide which of these are relevant web links for a brochure about the company, respond with the full https URL in JSON format. Do not include Terms of Service, Privacy, email links.\nLinks (some might be relative links):\n%s", website.URL, strings.Join(website.Links, "\n"))
}

func GetBrochureSystemPrompt() string {
	return `You are an assistant that analyzes the contents of several relevant pages from a company website and creates a short brochure about the company for prospective customers, investors, and recruits.

Your output must be:
- Written in clean, semantic HTML5.
- Include well-structured headings, sections, and paragraphs.
- Include embedded CSS inside a <style> block in the <head>.
- Use a clean, modern, professional design with legible fonts, spacing, and colors.
- Include company tone and a light touch of personality/humor if present in the content.

Return a single standalone HTML document. Do not use markdown. Do not include backticks or extra formatting. Just return valid HTML.`
}

// Alternate humorous prompt (you can toggle this in code)
func GetBrochureHumorousSystemPrompt() string {
	return `You are a witty assistant tasked with analyzing the contents of a company's website and creating a humorous and entertaining brochure.

Your output should:
- Be written in clean HTML5 with embedded CSS in the <head>.
- Be fun, quirky, and engaging — like a creative agency's landing page.
- Make use of emoji, puns, and jokes where appropriate.
- Include structured sections for company overview, culture, careers, etc.
- Include relevant styling that looks visually engaging.

Return a full standalone HTML document. Do not return markdown or explanation — only valid HTML.`
}

// GetBrochurePrompt builds the user prompt from homepage and selected relevant pages.
func GetBrochureUserPrompt(companyName string, rootURL string, links []LinkSuggestion) (string, error) {
	var contentBuilder strings.Builder

	// Add landing page
	contentBuilder.WriteString("Landing page:\n")
	landingContent, err := GetPageContent(rootURL)
	if err != nil {
		return "", err
	}
	contentBuilder.WriteString(landingContent)

	// Add each relevant page
	for _, link := range links {
		contentBuilder.WriteString(fmt.Sprintf("\n\n%s:\n", link.Type))
		pageContent, err := GetPageContent(link.URL)
		if err != nil {
			continue // skip bad links, but continue
		}
		contentBuilder.WriteString(pageContent)
	}

	// Truncate to max 5000 characters
	fullContent := contentBuilder.String()
	if len(fullContent) > 5000 {
		fullContent = fullContent[:5000]
	}

	// Build user prompt
	prompt := fmt.Sprintf(`You are looking at a company called: %s
Here are the contents of its landing page and other relevant pages; use this information to build a short brochure of the company in markdown.
%s`, companyName, fullContent)

	return prompt, nil
}
