package commitizen

import (
	"fmt"
	"go-commitizen/model"
)

type ConventionalCommit struct {

}

func (cz ConventionalCommit)Questions() []model.Question {
	return []model.Question{
		{
			Type: "list",
			Name: "prefix",
			Message: "Select the type of change you are committing",
			Options: []model.Option{
				{
					"feat: A new feature",
					"feat",
				},
				{
					"fix: A bug fix",
					"fix",
				},
			},
		},
		{
			Type: "input",
			Name: "scope",
			Message: "Scope. Specifying place of the commit change(users, db, poll):",
		},
		{
			Type: "input",
			Name: "subject",
			Message: "Subject. Write a SHORT, IMPERATIVE tense description of the change:",
		},
		{
			Type: "multiline",
			Name: "body",
			Message: "Body. Provide a LONGER description of the change(optional):",
		},
		{
			Type: "input",
			Name: "footer",
			Message: "Footer. List any Breaking Changes or reference issues that this commit closes:",
		},
	}
}

func (cz ConventionalCommit)Message(answers map[string]string) string {
	prefix := answers["prefix"]
	scope := answers["scope"]
	subject := answers["subject"]
	body := answers["body"]
	footer := answers["footer"]
	message := ""
	if prefix != "" {
		message := prefix
		if scope != "" {
			message += fmt.Sprintf("(%s)", scope)
		}
		message += ":"
	}
	message += subject
	if body != "" {
		message += "\n\n" + body
	}
	if footer != "" {
		 message += "\n\n" + footer
	}
	return message
}

func (cz ConventionalCommit)Schema() string {
	return "<type>(<scope>): <subject>\n" +
		   "<BLANK LINE>\n" +
		   "<body>\n" +
		   "<BLANK LINE>\n" +
		   "<footer>"
}