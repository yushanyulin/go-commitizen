package commitizen

import (
	"fmt"
	"go-commitizen/git"
	"go-commitizen/model"
	"go-commitizen/prompt"
	"io/ioutil"
	"os"
)

// Commitizen 协议
type Commitizen interface {
	Message(map[string]string) string
	Questions() []model.Question
	Schema() string
}



func Run(commitizen Commitizen) {

	if checkEmptyChange() {
		fmt.Println("no changes added to commit(use 'git add')")
		return
	}

	answers := prompt.Ask(commitizen.Questions())
	message := commitizen.Message(answers)
	fmt.Printf("commit is :\n###----------------------------------------------###\n%s\n" +
		"###----------------------------------------------###\n", message)

	questions := []model.Question{
		{
			Type: "confirm",
			Name: "confirm",
			Message: "Are you sure to proceed with the commit above:",
		},
	}
	answers = prompt.Ask(questions)
	answer := answers["confirm"]
	if answer == "y" || answer == "" {
		commit(message)
	} else {
		fmt.Println("Commit failed.")
	}

}

func checkEmptyChange() bool {
	isEmpty, _, _ := git.RunGit("git diff --quiet --cached")
	return isEmpty
}

func commit(message string) {
	tmpfile, err := ioutil.TempFile("", ".commit_message")
	if err != nil {
		fmt.Println("create temporary file: .commit_message failed")
		return
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(message)); err != nil {
		fmt.Println("write temporary file: .commit_message failed")
		return
	}
	if err := tmpfile.Close(); err != nil {
		fmt.Println("close temporary file: .commit_message failed")
		return
	}

	ok, stdout, stderr := git.RunGit("git commit -F " + tmpfile.Name())
	if ok {
		fmt.Println(stdout)
	} else  {
		fmt.Println(stderr)
	}
 }