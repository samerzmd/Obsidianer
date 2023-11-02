package watcher

import (
	"fmt"
	"log"
	"obsidianer/pkg/config"
	"obsidianer/pkg/git_ops"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func AddWatcher(watcher *fsnotify.Watcher, path string) error {
	return filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !strings.Contains(walkPath, ".git") {
			return watcher.Add(walkPath)
		}
		return nil
	})
}

func WatchFiles(path string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := AddWatcher(watcher, path); err != nil {
		return err
	}
	// Create an instance of RealExecCommander
    realExecCommander := config.RealExecCommander{}
    

	gitUserName, err := config.GetGitConfig(realExecCommander, "user.name", "Please enter your Git username:")
    if err != nil {
        return err
    }

	gitUserEmail, err := config.GetGitConfig(realExecCommander, "user.email", "Please enter your Git email:")
    if err != nil {
        return err
    }

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("File modified:", event.Name)
					relativePath, _ := filepath.Rel(path, event.Name)
					if err := git_ops.CommitAndPush(path, relativePath, gitUserName, gitUserEmail); err != nil {
						log.Println("Could not commit and push:", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	<-done
	return nil
}
