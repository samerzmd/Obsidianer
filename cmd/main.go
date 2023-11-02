package main

import (
	"fmt"
	"log"
	"obsidianer/pkg/git_ops"
	"obsidianer/pkg/watcher"
)

func main() {
	var path string
	fmt.Println("Enter the path of the directory to watch:")
	fmt.Scanln(&path)

	// Create an instance of RealGitRepo
	realGitRepo := git_ops.RealGitRepo{}

	if err := git_ops.InitGitRepo(realGitRepo, path); err != nil {
		log.Fatal("Could not initialize Git repository:", err)
	}

	if err := watcher.WatchFiles(path); err != nil {
		log.Fatal("Could not watch files:", err)
	}
}
