package command

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
)

func GetImage(args []string) (string, error) {
	dir, err := getDirectory(args)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return "", err
	}
	files := make([]string, 0)
	for _, name := range names {
		fileInfo, err := os.Stat(args[0] + "/" + name)
		if err != nil {
			break
		}

		if fileInfo.IsDir() {
			break
		}
		files = append(files, name)
	}
	if len(files) == 0 {
		return "", fmt.Errorf("No files in directory")
	}
	absDirPath, err := filepath.Abs(dir.Name())
	if err != nil {
		return "", err
	}
	switch len(files) {
	case 1:
		return filepath.Join(absDirPath, names[0]), nil
	default:
		ind := rand.IntN(len(files) - 1)
		return filepath.Join(absDirPath, names[ind]), nil
	}
}

func getDirectory(args []string) (*os.File, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Missing directory of images")
	}
	dirStats, err := os.Stat(args[0])
	if err != nil {
		return nil, fmt.Errorf("Directory does not exist, %s", err)
	}
	if !dirStats.IsDir() {
		return nil, fmt.Errorf("Directory is not a directory")
	}
	dir, err := os.Open(args[0])
	if err != nil {
		return nil, err
	}
	return dir, nil
}
