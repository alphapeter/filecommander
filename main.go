package main

import (
	"./commands"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write(html)
			break
		default:
			w.WriteHeader(400)
			w.Write([]byte("bad request"))
		}
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

var html = []byte(`
	<html>
		<head>
		<title>
			Filecommander
		</title>
		</head>
		<body>
		<div id="directory1" class="filebrowser left"></div>
		<div id="directory2" class="filebrowser right"></div>

		<div class="commands">
			<button id="copy" class="icon-clone">copy</button>
			<button id="move" class="icon-exchange">move</button>
			<button id="mkdir" class="icon-folder-empty-1">mkdir</button>
			<button id="delete" class="icon-trash-empty">delete</button>
			<button id="rename" class="icon-pencil-squared">rename</button>
		</div>
		</body>
	</html>

	<script type="text/javascript">

	</script>

	<style>
		.filebrowser { border: 1px solid black; width:calc(50% - 4px); height: calc(100% - 40px); display: inline-block;}
		.left { left: 0px; }
		.right { right: 0px; }
		.commands { width: 100%; text-align: center; }
	</style>
	<style>
		@font-face {
  font-family: 'fontello';
  src: url('../font/fontello.eot?39893882');
  src: url('../font/fontello.eot?39893882#iefix') format('embedded-opentype'),
       url('../font/fontello.svg?39893882#fontello') format('svg');
  font-weight: normal;
  font-style: normal;
}
@font-face {
  font-family: 'fontello';
  src: url('data:application/octet-stream;base64,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') format('woff'),
       url('data:application/octet-stream;base64,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sEsssgABQSstsEwssgEAQSstsE0ssgEBQSstsE4ssgAAQistsE8ssgABQistsFAssgEAQistsFEssgEBQistsFIssgAAPistsFMssgABPistsFQssgEAPistsFUssgEBPistsFYssgAAQCstsFcssgABQCstsFgssgEAQCstsFkssgEBQCstsFossgAAQystsFsssgABQystsFwssgEAQystsF0ssgEBQystsF4ssgAAPystsF8ssgABPystsGAssgEAPystsGEssgEBPystsGIssDcrLrErARQrLbBjLLA3K7A7Ky2wZCywNyuwPCstsGUssAAWsDcrsD0rLbBmLLA4Ky6xKwEUKy2wZyywOCuwOystsGgssDgrsDwrLbBpLLA4K7A9Ky2waiywOSsusSsBFCstsGsssDkrsDsrLbBsLLA5K7A8Ky2wbSywOSuwPSstsG4ssDorLrErARQrLbBvLLA6K7A7Ky2wcCywOiuwPCstsHEssDorsD0rLbByLLMJBAIDRVghGyMhWUIrsAhlsAMkUHiwARUwLQBLuADIUlixAQGOWbABuQgACABjcLEABUKyAAEAKrEABUKzCgIBCCqxAAVCsw4AAQgqsQAGQroCwAABAAkqsQAHQroAQAABAAkqsQMARLEkAYhRWLBAiFixA2REsSYBiFFYugiAAAEEQIhjVFixAwBEWVlZWbMMAgEMKrgB/4WwBI2xAgBEAAA=') format('truetype');
}
/* Chrome hack: SVG is rendered more smooth in Windozze. 100% magic, uncomment if you need it. */
/* Note, that will break hinting! In other OS-es font will be not as sharp as it could be */
/*
@media screen and (-webkit-min-device-pixel-ratio:0) {
  @font-face {
    font-family: 'fontello';
    src: url('../font/fontello.svg?39893882#fontello') format('svg');
  }
}
*/

 [class^="icon-"]:before, [class*=" icon-"]:before {
  font-family: "fontello";
  font-style: normal;
  font-weight: normal;
  speak: none;

  display: inline-block;
  text-decoration: inherit;
  width: 1em;
  margin-right: .2em;
  text-align: center;
  /* opacity: .8; */

  /* For safety - reset parent styles, that can break glyph codes*/
  font-variant: normal;
  text-transform: none;

  /* fix buttons height, for twitter bootstrap */
  line-height: 1em;

  /* Animation center compensation - margins should be symmetric */
  /* remove if not needed */
  margin-left: .2em;

  /* you can be more comfortable with increased icons size */
  /* font-size: 120%; */

  /* Uncomment for 3D effect */
  /* text-shadow: 1px 1px 1px rgba(127, 127, 127, 0.3); */
}
.icon-trash-empty:before { content: '\e800'; } /* '' */
.icon-doc-1:before { content: '\e801'; } /* '' */
.icon-edit:before { content: '\e802'; } /* '' */
.icon-folder-1:before { content: '\e807'; } /* '' */
.icon-folder-open:before { content: '\e808'; } /* '' */
.icon-spin5:before { content: '\e838'; } /* '' */
.icon-exchange:before { content: '\f0ec'; } /* '' */
.icon-folder-empty-1:before { content: '\f114'; } /* '' */
.icon-folder-open-empty:before { content: '\f115'; } /* '' */
.icon-pencil-squared:before { content: '\f14b'; } /* '' */
.icon-doc-inv-1:before { content: '\f15b'; } /* '' */
.icon-trash:before { content: '\f1f8'; } /* '' */
.icon-clone:before { content: '\f24d'; } /* '' */
.icon-question-circle-o:before { content: '\f29c'; } /* '' */
	</style>
`)
