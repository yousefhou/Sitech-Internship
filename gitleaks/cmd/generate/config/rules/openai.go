package rules

import (
	"github.com/zricethezav/gitleaks/v8/cmd/generate/secrets"
	"github.com/zricethezav/gitleaks/v8/config"
)

func OpenAI() *config.Rule {
	// define rule
	r := config.Rule{
		RuleID:      "openai-api-key",
		Description: "OpenAI API Key",
		Regex:       generateUniqueTokenRegex(`sk-[a-zA-Z0-9]{20}T3BlbkFJ[a-zA-Z0-9]{20}`),
		SecretGroup: 1,
		Keywords: []string{
			"T3BlbkFJ",
		},
	}

	// validate
	tps := []string{
		generateSampleSecret("openaiApiKey", "sk-"+secrets.NewSecret(alphaNumeric("20"))+"T3BlbkFJ"+secrets.NewSecret(alphaNumeric("20"))),
	}
	return validate(r, tps, nil)
}
