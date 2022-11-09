package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type lockedFiles struct {
	filePath  string
	lockOwner string
	fileId    string
}

func main() {
	fmt.Println("Hello World!")

	cmd := exec.Command("git", "-C", "C:/Users/SAM/src/projectkonoha/", "lfs", "locks")
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err.Error())
	}

	lockedFilesArr := parseLockedFiles(string(out))

	for _, file := range lockedFilesArr {
		fmt.Println(file.lockOwner)
	}

}

// parses the output of a git lfs command, adds each row to an array of lockedFiles[] structs for easier doing things.
func parseLockedFiles(gitlfsOutput string) []lockedFiles {

	var lockedFilesArr []lockedFiles

	outputLines := strings.Split(gitlfsOutput, "\n")
	for _, line := range outputLines {
		splitLfsLine := strings.Split(line, "\t")
		if len(splitLfsLine) > 2 {
			lockedFile := lockedFiles{filePath: splitLfsLine[0], lockOwner: splitLfsLine[1], fileId: splitLfsLine[2]}
			lockedFilesArr = append(lockedFilesArr, lockedFile)
		}
	}

	return lockedFilesArr

}
