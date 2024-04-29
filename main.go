package main

import (
	"fmt"
)

var help = `

Usage:
	infinit-task <repository> <revision>

	- repository must contain the repository owner and name separated by a slash
	- revison can be a branch or commit SHA

Example:
	infinit-task lodash/lodash  some-branch-name
	infinit-task lodash/lodash  (revision defaults to branch 'main')
	infinit-task lodash/lodash  c7c70a7da5172111b99bb45e45532ed034d7b5b9 
`

func main() {

	repo, err := getArgument(0)
	if err != nil {
		panic("Repository not provided, " + err.Error() + help)
	}

	revision, err := getArgument(1)
	if err != nil {
		revision = "main"
	}

	fmt.Println("Computing stats for repository:", repo, "at revision:", revision)

	url := "https://api.github.com/repos/" + repo + "/zipball/" + revision

	zipFileReader, err := NewZipFileReader(url)
	if err != nil {
		panic("Could not retrieve repository file reader" + err.Error())
	}

	repoFileContent, err := zipFileReader.GetFileContentByExtention([]string{".js", ".ts"})
	if err != nil {
		panic("Could not retrieve repository file content" + err.Error())
	}
	stats := NewCharFrequency(repoFileContent)

	fmt.Println("Result:")
	fmt.Println("======")
	fmt.Println(stats.ToString())
}
