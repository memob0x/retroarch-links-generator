package utils

import "log"

func ParseCommandLineArgs(args []string) (string, string) {
	var argsCount = len(args)

	var retroArchPath string

	var linksDestPath string = "."

	if argsCount >= 2 {
		retroArchPath = args[1]
	} else {
		log.Fatalf("No retroarch path specified, aborting.\n")
	}

	if argsCount >= 3 {
		linksDestPath = args[2]
	} else {
		log.Fatalf("No destination path specified, aborting.\n")
	}

	return retroArchPath, linksDestPath
}
