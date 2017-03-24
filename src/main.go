package main

import (
	"./commands"
	"./gui"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write(gui.Html)
			break
		default:
			w.WriteHeader(400)
			w.Write([]byte("bad request"))
		}
	})
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		w.Write(gui.Css)


	})
	http.HandleFunc("/bundle.js", func(w http.ResponseWriter, r *http.Request) {
		w.Write(gui.Javascript)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var command commands.Command

		if err := decoder.Decode(&command); err != nil {
			er, _ := json.Marshal(ErrorResponse{
				Error: Error{
					Code:    -32700,
					Message: "could not parse JSON: " + err.Error(),
				},
				Id:         command.Id,
				RpcVersion: "2.0",
			})
			w.Write(er)
			return
		}

		defer r.Body.Close()
		fmt.Print(command)
		response := Response{Id: command.Id, RpcVersion: "2.0"}

		var err error
		switch command.Method {
		case "mv":
			err = commands.Move(command.Params[0], command.Params[1])
			break
		case "cp":
			err = commands.Copy(command.Params[0], command.Params[1])
			break
		case "rm":
			err = commands.Delete(command.Params[0])
			break
		case "ls":
			response.Result, err = commands.Ls(command.Params[0])
			break
		case "df":
			response.Result, err = commands.Df()
			break
		case "mkdir":
			err = commands.Mkdir(command.Params[0], command.Params[1])
			break
		default:
			methodNotFound(command, &w)
			return
		}

		if err != nil {
			er, _ := json.Marshal(ErrorResponse{
				Error: Error{
					Code:    -32603,
					Message: err.Error(),
				},
				Id:         command.Id,
				RpcVersion: "2.0",
			})
			w.Write(er)
			return
		}
		a, _ := json.Marshal(response)
		w.Write(a)

	})

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func methodNotFound(command commands.Command, w *http.ResponseWriter) {
	response, _ := json.Marshal(ErrorResponse{
		Id: command.Id,
		Error: Error{
			Code:    -32601,
			Message: "Method not found",
		},
		RpcVersion: "2.0",
	})
	(*w).Write([]byte(response))
}

type Response struct {
	RpcVersion string      `json:"jsonrpc"`
	Result     interface{} `json:"result"`
	Id         string      `json:"id"`
}
type ErrorResponse struct {
	RpcVersion string `json:"jsonrpc"`
	Id         string `json:"id"`
	Error      Error  `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}