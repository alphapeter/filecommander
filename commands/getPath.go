package commands

import "strings"
import "../cfg"

func getPath(directory string) (string, error) {
	if err := ValidatePath(directory); err != nil{
		return "", err
	}

	paths := strings.SplitN(directory, "/", 2)
	root := paths[0]
	rootPath, err := cfg.GetPathForRoot(root)
	if err != nil {
		return "", err
	}

	var path string
	if len(paths) > 1 {
		path = rootPath + "/" + paths[1]
	} else {
		path = rootPath
	}
	return path, nil
}
