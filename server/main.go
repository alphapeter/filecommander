package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alphapeter/filecommander/server/cfg"
	"github.com/alphapeter/filecommander/server/commands"
	"github.com/alphapeter/filecommander/server/static"
)

func main() {
	settings := cfg.GetSettings()

	staticHandler := static.CreateHandler(static.Data, "index.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			staticHandler.ServeHTTP(w, r)
			break
		default:
			w.WriteHeader(400)
			w.Write([]byte("bad request"))
		}
	})

	http.HandleFunc("/static/", staticHandler.ServeHTTP)

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
		response := Response{Id: command.Id, RpcVersion: "2.0"}

		var err error
		switch command.Method {
		case "mv":
			err = command.Move()
			break
		case "cp":
			err = command.Copy()
			break
		case "rm":
			err = command.Delete()
			break
		case "ls":
			response.Result, err = command.Ls()
			break
		case "df":
			response.Result, err = command.Df()
			break
		case "mkdir":
			err = command.Mkdir()
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
		w.Header().Set("Content-Type", "application/json")
		w.Write(a)

	})

	err := http.ListenAndServe(settings.Binding, nil)
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
