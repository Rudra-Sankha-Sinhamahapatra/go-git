package repo

import (
	"fmt"
)

/*
  Walks commit history starting from HEAD
  @LogCommits
*/

func LogCommits(repoPath string, path string) error {
	ref, err := ResolveHEAD(repoPath)
	if err != nil {
		return err
	}

	current, exists, err := ReadCurrentCommit(repoPath, ref)
	if err != nil {
		return err
	}
	if !exists {
		fmt.Println("no commits yet")
		return nil
	}

	for {
		commit, err := ReadCommit(path, current)
		if err != nil {
			return err
		}

		fmt.Println("commit", commit.Hash)
		fmt.Println(commit.Message)
		fmt.Println()

		if commit.Parent == nil {
			break
		}
		current = *commit.Parent
	}

	return nil
}
