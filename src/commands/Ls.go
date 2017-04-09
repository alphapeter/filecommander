package commands

import (
	"io/ioutil"
	"time"
)

func (command Command) Ls() ([]File, error){
	if err := command.validateUnaryParameters(); err != nil{
		return nil, err
	}
	return ls(command.Params[0])
}

func ls(directory string) ([]File, error) {
	path, err := getPath(directory)
	if err != nil {
		return nil, err
	}

	files, _ := ioutil.ReadDir(path)
	response := []File{}

	for _, file := range files {
		f := File{Name: file.Name(), Size: file.Size(), Modified: file.ModTime()}
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
	Type string `json:"type"`
	Name string `json:"name"`
	Modified time.Time `json:"modified"`
	Size int64 `json:"size"`
}
