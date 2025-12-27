package services

func CallLLM(prompt string) map[string]interface{} {

	return map[string]interface{}{
		"variants": []map[string]string{
			{
				"headline": "Land More Interviews with AI",
				"body":     "Build an ATS-friendly resume in minutes using AI.",
				"cta":      "Try Free",
			},
			{
				"headline": "Your Resume, Optimized by AI",
				"body":     "Stand out to recruiters with AI-powered resumes.",
				"cta":      "Get Started",
			},
			{
				"headline": "Smart Resumes for Smart Careers",
				"body":     "AI-crafted resumes designed to convert.",
				"cta":      "Build Now",
			},
		},
	}
}
