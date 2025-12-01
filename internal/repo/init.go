package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
 Creates .gogit repo like git init
 @InitRepo
*/

func InitRepo(path string) error {
	gitPath := filepath.Join(path, ".gogit")

	if _, err := os.Stat(gitPath); err == nil {
		return fmt.Errorf("gogit repository already exists")
	}

	dirs := []string{
		"objects",
		"objects/info",
		"objects/pack",
		"refs",
		"refs/heads",
		"refs/tags",
	}

	for _, d := range dirs {
		err := os.MkdirAll(filepath.Join(gitPath, d), 0755)
		if err != nil {
			return err
		}
	}

	headContent := []byte("ref: refs/heads/main\n")
	err := os.WriteFile(filepath.Join(gitPath, "HEAD"), headContent, 0644)
	if err != nil {
		return err
	}

	config := []byte(
		"[core]\n" +
			"    repositoryformatversion = 0\n" +
			"    filemode = false\n" +
			"    bare = false\n",
	)

	err = os.WriteFile(filepath.Join(gitPath, "config"), config, 0644)
	if err != nil {
		return err
	}

	desc := []byte("Unnamed gogit repository\n")
	err = os.WriteFile(filepath.Join(gitPath, "description"), desc, 0644)
	if err != nil {
		return err
	}

	return nil
}
