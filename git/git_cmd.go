package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

// RunGit can execute or run git command
func RunGit(command string) (ok bool, stdout, stderr string) {
	cmd := exec.Command("sh", "-c", command)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	ok = true
	err := cmd.Run()
	stdout = outb.String()
	stderr = errb.String()
	if err != nil {
		stderr = fmt.Sprintf("%v", err)
		ok = false
	}
	return
}
