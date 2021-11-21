package main

import "go-commitizen/commitizen"

func main() {
	cz := new(commitizen.ConventionalCommit)
	commitizen.Run(cz)
}
