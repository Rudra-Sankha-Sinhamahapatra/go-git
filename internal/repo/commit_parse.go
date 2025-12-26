package repo

import "strings"

type Commit struct {
	Hash    string
	Tree    string
	Parent  *string
	Message string
}

/*
Parses commit object content
@ParseCommitObject
*/

func ParseCommitObject(content []byte) (*Commit, error) {
	lines := strings.Split(string(content), "\n")

	var tree string
	var parent *string
	var messageLines []string

	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]

		if line == "" {
			i++
			break
		}

		if strings.HasPrefix(line, "tree") {
			tree = strings.TrimPrefix(line, "tree ")
		} else if strings.HasPrefix(line, "parent ") {
			p := strings.TrimPrefix(line, "parent ")
			parent = &p
		}
	}

	for ; i < len(lines); i++ {
		messageLines = append(messageLines, lines[i])
	}

	return &Commit{
		Tree:    tree,
		Parent:  parent,
		Message: strings.Join(messageLines, "\n"),
	}, nil
}

/*
  Reads and parses a commit object by hash
  @ReadCommit
*/

func ReadCommit(repoPath string, hash string) (*Commit, error) {
	compressed, err := ReadCompressedGoGitObject(hash, repoPath)
	if err != nil {
		return nil, err
	}

	raw, err := DecompressesGoGitObject(compressed)
	if err != nil {
		return nil, err
	}

	_, _, content, err := ParseGoGitObjectHeader(raw)
	if err != nil {
		return nil, err
	}

	commit, err := ParseCommitObject(content)
	if err != nil {
		return nil, err
	}

	commit.Hash = hash
	return commit, nil
}
