# 🧾 Website Brochure Generator (Go + LLM)

Generate professional or humorous brochures for any company website using a combination of web scraping and LLM-based content generation — all in Go.

---

## 🚀 Features

- 🔍 **Scrapes Website Content:** Extracts meaningful content (title, text, links) from landing and relevant pages.
- 🤖 **LLM-Powered Analysis:** Uses OpenAI API to select relevant pages and generate brochures.
- 🧠 **Custom Prompting:** Fine-tuned prompts for both serious and humorous brochure generation.
- 📄 **Outputs:** Brochures in either Markdown, styled HTML, or PDF (coming soon).
- 📁 **Clean Project Structure:** Modular, extensible, and easy to build on.

---

## 🔧 Installation

```bash
git clone https://github.com/yourusername/website-brochure-generator.git
cd website-brochure-generator
go mod tidy
```

---

## ⚙️ Environment Setup

```
OPENAI_API_KEY=your-openai-key
OPENAI_API_URL=https://api.openai.com/v1/chat/completions
OPENAI_MODEL=gpt-4o-mini
```

Make sure .env is listed in .gitignore.

---

## ▶️ Running the Generator

1. Update main.go (or build your own CLI) to use the core functions:

```
website, _ := core.ScrapeWebsite("https://example.com")
links, _ := core.GetRelevantLinks(website)
brochureHTML, _ := core.CreateBrochure("Example Co", website.URL, links.Links)
core.ExportBrochureAsHTML(brochureHTML, "brochure.html")
```

2. Run the program:

```
go run main.go
```

3. ✅ Check your generated brochure.html.


---

## 🧪 Customization

### Prompt Tuning

- Edit prompt.go to customize tone, format, length, and structure.
- Switch between GetBrochureSystemPrompt() and GetBrochureHumorousSystemPrompt().

### Output Styling

- Modify the CSS in ExportBrochureAsHTML() (or via prompt) for different styles.
- Add logos, company colors, etc., by analyzing website assets.

---

_Project maintained by [Ovijeet](https://github.com/ovijeet26) | [Report Issue](https://github.com/ovijeet26/brochure-generator-llm/issues)_
