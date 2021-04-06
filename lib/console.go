package lib

import (
	"github.com/AlecAivazis/survey/v2"
)

type Response struct {
	Value string
}

func PrintHeader(message string) {
	println("=========", message, "=========")
	println()
}

func Input(message string) (string, error) {
	var qs = []*survey.Question{
		{
			Name:      "Value",
			Prompt:    &survey.Input{Message: message},
			Transform: survey.Title,
		},
	}

	res := Response{}

	err := survey.Ask(qs, &res)

	if err != nil {
		return "", err
	}

	return res.Value, nil
}

func Select(message string, options []string) (string, error) {
	var qs = []*survey.Question{
		{
			Name: "Value",
			Prompt: &survey.Select{
				Message: message,
				Options: options,
				Default: options[0],
			},
		},
	}

	res := Response{}

	err := survey.Ask(qs, &res)

	if err != nil {
		return "", err
	}

	return res.Value, nil
}

func Continue() {
	Input("Press enter to continue")
}
