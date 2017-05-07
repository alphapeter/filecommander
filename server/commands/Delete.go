package commands

import (
	"os"
)

func (command Command) Delete() error{
	if err := command.validateUnaryParameters(); err != nil{
		return err
	}
	return delete(command.Params[0])
}

func delete(file string) error {
	path, err := getPath(file)

	if err != nil {
		return err
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}
