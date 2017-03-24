package gui

var Html = []byte(`
	<html>
		<head>
		<title>
			Filecommander
		</title>
		<link rel="stylesheet" href="styles.css">
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
	<script src="bundle.js"></script>
`)
