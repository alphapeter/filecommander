package commands

import "errors"

type Command struct {
	RpcVersion string   `josn:"jsonrpc"`
	Method     string   `json:"method"`
	Params     []string `json:"params"`
	Id         string   `json:"id"`
}

func (command Command) validateBinaryParameters() error{
	switch len(command.Params) {
	case 0:
		return errors.New("missing params ['source', 'destination']")
	case 1:
		return errors.New("missing param [..., 'destination'] ")
	case 2:
		return nil
	default:
		return errors.New("too many params, should only be ['source', 'destination'] ")
	}
}

func (command Command) validateUnaryParameters() error{
	switch len(command.Params) {
	case 0:
		return errors.New("missing params ['path/(file)'] ")
	case 1:
		return nil
	default:
		return errors.New("too many params, should only be ['path/(file)'] ")
	}
}
