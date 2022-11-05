package cmd

import (
	"errors"
	"fmt"

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
		"License",
	}
	readmeComponents := prompts.MultiSelectPrompt("Please select all options you want to add in the Readme!", promptQs, survey.Required)
	return readmeComponents
}

func AppendSectionToReadme(prompt string, options ...string) (ok bool, err error) {
	fmt.Println("ADDING "+prompt+" TO README WITH OPTIONS", options)
	return true, nil
}

func CreateReadme() (ok bool, err error) {
	title := prompts.StringPrompt("Please enter the title of your project", false)
	description := prompts.StringPrompt("Please enter the description of your project", true)

	components := GetReadmeComponents()

	for _, ele := range components {
		if ele == "Title & Description" {
			_, err := AppendSectionToReadme(ele, title, description)

			if err != nil {
				return false, errors.New("Error in writing Title and Desc to Readme File occurred: " + err.Error())
			}
		} else {
			_, err := AppendSectionToReadme(ele)

			if err != nil {
				return false, errors.New("Error in writing " + ele + " to Readme File occurred: " + err.Error())
			}
		}
	}

	return true, nil
}
