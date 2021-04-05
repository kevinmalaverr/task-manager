package console

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
