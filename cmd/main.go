package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/kush-daga/readme-generator/pkg/prompts"
)

func GetReadmeComponents() []string {
	promptQs := []string{
		"Title & Description",
		"Demo Section",
		"Installation",
		"Run locally",
		"Env Variables",
		"FAQ",
		"ScreenShots",
		"Acknowledgments",
	}
	readmeComponents := prompts.MultiSelectPrompt("Please select all options you want to add in the Readme!", promptQs, survey.Required)
	return readmeComponents
}
