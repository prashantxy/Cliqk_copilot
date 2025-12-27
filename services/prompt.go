package services

import (
	"cliqk-copilot/models"
	"fmt"
)

func BuildAdPrompt(req models.AdRequest) string {
	return fmt.Sprintf(`
You are a senior growth marketer.

Create 3 high-converting ad copies.

Product: %s
Target Audience: %s
Platform: %s
Tone: %s

Rules:
- Short & punchy
- Clear value proposition
- Strong CTA
- No emojis
- Return ONLY valid JSON

JSON format:
{
  "variants": [
    {
      "headline": "",
      "body": "",
      "cta": ""
    }
  ]
}
`, req.Product, req.Audience, req.Platform, req.Tone)
}
