package prompt

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"go-commitizen/model"
	"log"
	"os"
)

func Ask(questions []model.Question) map[string]string {
	answers := make(map[string]string)
	for _, question := range questions {
		answers[question.Name] = prompt(question)
	}
	return answers
}


func prompt(question model.Question) string {
	switch question.Type {
	case "list":
		return promptList(question)
	case "input":
		return promptInput(question)
	case "confirm":
		return promptConfirm(question)
	case "multiline":
		return promptMultiline(question)

	}
	fmt.Println("Unknown question type")
	os.Exit(1)
	return ""
}


func promptList(question model.Question) string {
	prompt := &survey.Select{
		Message: question.Message,
		VimMode: true,
		PageSize: 8,
	}
	for _,option := range question.Options {
		prompt.Options = append(prompt.Options, option.Name + ":" + option.Value)
	}

	var answer survey.OptionAnswer
	err := survey.AskOne(prompt, &answer)
	processError(err)
	return question.Options[answer.Index].Value
}

func promptInput(question model.Question) (answer string) {
	prompt := &survey.Input{
		Message: question.Message,
	}
	err := survey.AskOne(prompt, &answer)
	processError(err)
	return
}

func promptMultiline(question model.Question) (answer string) {
	prompt := &survey.Multiline{
		Message: question.Message,
	}
	err := survey.AskOne(prompt, &answer)
	processError(err)
	return
}

func promptConfirm(question model.Question) (answer string) {
	result := false
	prompt := &survey.Confirm{
		Message: question.Message,
		Default: true,
	}
	err := survey.AskOne(prompt, &result)
	processError(err)
	if result {
		return "y"
	} else {
		return "n"
	}

	return
}

func processError(err error) {
	if err == nil {
		return
	}

	if err == terminal.InterruptErr {
		os.Exit(1)
		return
	}

	log.Fatal(err)
}