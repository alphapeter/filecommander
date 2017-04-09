package commands

import (
	"io"
	"os"
)


func (command Command) Copy() error{
	if err := command.validateBinaryParameters(); err != nil{
		return err
	}
	return copy(command.Params[0], command.Params[1])
}

func copy(source string, destination string) error {
	sourcePath, err := getPath(source)

	if err != nil {
		return err
	}

	destinationPath, err := getPath(destination)

	if err != nil {
		return err
	}

	return copyFileContents(sourcePath, destinationPath)
}

func copyFileContents(src string, dst string) (err error) {
	sourceFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer sourceFile.Close()
	destinationFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := destinationFile.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(destinationFile, sourceFile); err != nil {
		return
	}
	err = destinationFile.Sync()
	return
}
