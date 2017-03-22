package commands

import (
	"io/ioutil"
)

func Ls(directory string) ([]File, error) {
	path, err := getPath(directory)
	if err != nil {
		return nil, err
	}

	files, _ := ioutil.ReadDir(path)
	response := []File{}

	for _, file := range files {
		f := File{Name: file.Name()}
		if file.IsDir() {
			f.Type = "d"
		} else {
			f.Type = "f"
		}
		response = append(response, f)
	}
	return response, nil
}

type File struct {
	Type string
	Name string
}
