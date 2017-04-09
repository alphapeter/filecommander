package commands

import (
	"os"
)

func (command Command) Mkdir() error{
	if err := command.validateBinaryParameters(); err != nil{
		return err
	}
	return mkdir(command.Params[0], command.Params[1])
}

func mkdir(directory, directoryName string) error {
	if err := ValidatePath(directoryName); err != nil {
		return err
	}
	path, err := getPath(directory)

	if err != nil {
		return err
	}

	err = os.Mkdir(path + "/" + directoryName, os.ModePerm)

	return err
}
