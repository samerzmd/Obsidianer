package git_ops

import (
	"fmt"
	"time"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type GitRepo interface {
	PlainOpen(string) (*git.Repository, error)
	PlainClone(string, bool, *git.CloneOptions) (*git.Repository, error)
}

type RealGitRepo struct{}

func (r RealGitRepo) PlainOpen(path string) (*git.Repository, error) {
	return git.PlainOpen(path)
}

func (r RealGitRepo) PlainClone(path string, isBare bool, options *git.CloneOptions) (*git.Repository, error) {
	return git.PlainClone(path, isBare, options)
}

func InitGitRepo(gitRepo GitRepo, path string) error {
	_, err := git.PlainOpen(path)
	if err == nil {
		return nil
	}

	fmt.Println("No Git repository found. Please enter an empty repository URL to clone:")
	var repoURL string
	fmt.Scanln(&repoURL)

	_, err = git.PlainClone(path, false, &git.CloneOptions{URL: repoURL})
	return err
}

func CommitAndPush(path string, fileName string, gitUserName string, gitUserEmail string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	_, err = w.Add(".")
	if err != nil {
		return err
	}

	commitMessage := fmt.Sprintf("Modified: %s", fileName)

	_, err = w.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Name:  gitUserName,
			Email: gitUserEmail,
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
	})
	if err != nil {
		return err
	}
	return nil
}
