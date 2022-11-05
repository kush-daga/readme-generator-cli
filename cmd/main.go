package cmd

import (
	"fmt"
	"os"
	"text/template"

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

type data struct {
	Title       string
	Description string
	Components  []string
}

func CreateReadme() (bool, error) {
	title := prompts.StringPrompt("Please enter the title of your project", false)
	description := prompts.StringPrompt("Please enter the description of your project", true)

	components := GetReadmeComponents()

	d := data{title, description, components}

	paths := []string{
		"./cmd/md.go.tpl",
	}

	t := template.Must(template.New("md-tmpl").ParseFiles(paths...))
	f, err := os.Create("./README.md")

	if err != nil {
		return false, err
	}

	fmt.Print(t.DefinedTemplates())
	err = t.ExecuteTemplate(f, "md.go.tpl", d)

	if err != nil {
		return false, err
	}

	return true, nil
}
