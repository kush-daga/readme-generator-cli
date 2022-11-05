package prompts

import (
	"github.com/AlecAivazis/survey/v2"
)

func createCustomValidatorOpt(validators ...survey.Validator) survey.AskOpt {
	customValidator := survey.WithValidator(survey.ComposeValidators(validators...))
	return customValidator
}

func StringPrompt(
	promptQuestion string,
	multiLine bool,
	validators ...survey.Validator,
) string {
	promptAnswer := ""

	if multiLine {
		prompt := &survey.Multiline{
			Message: promptQuestion,
		}
		survey.AskOne(prompt, &promptAnswer, createCustomValidatorOpt(validators...))
	} else {
		prompt := &survey.Input{
			Message: promptQuestion,
		}
		survey.AskOne(prompt, &promptAnswer, createCustomValidatorOpt(validators...))
	}

	return promptAnswer
}

func SelectOnePrompt(promptLabel string, promptQs []string, validators ...survey.Validator) string {
	res := ""

	prompt := &survey.Select{
		Message: promptLabel,
		Options: promptQs,
	}
	survey.AskOne(prompt, &res, createCustomValidatorOpt(validators...))

	return res
}

func MultiSelectPrompt(promptLabel string, promptQs []string, validators ...survey.Validator) []string {
	res := []string{}

	prompt := &survey.MultiSelect{
		Message: promptLabel,
		Options: promptQs,
	}

	survey.AskOne(prompt, &res, createCustomValidatorOpt(validators...))

	return res
}
