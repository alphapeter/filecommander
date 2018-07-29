package commands

import (
	"github.com/alphapeter/filecommander/server/cfg"
	"errors"
)

func (command Command) Df() ([]string, error){
	if len(command.Params) > 0{
		return nil, errors.New("invalid number of params, the command does not require parameters")
	}
	return df()
}

func df() ([]string, error) {
	settings := cfg.GetSettings()
	rootNames := make([]string, 0)
	for i := 0; i < len(settings.Roots); i++ {
		rootNames = append(rootNames, settings.Roots[i].Name)
	}
	return rootNames, nil
}
