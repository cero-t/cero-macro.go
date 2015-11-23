package http

func formHtml() string {
	return `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Cero Macro Form</title>
	</head>

	<body>
		<form id="macroForm" name="macroForm" action="." method="POST">
			<b>Player1</b><br/>
			<textarea id="player1" name="player1" rows="10" cols="40">{{.Player1}}</textarea><br/>
			<b>Player2</b><br/>
			<textarea id="player2" name="player2" rows="10" cols="40">{{.Player2}}</textarea><br/>
			<input id="submit" name="submit" type="submit" value="submit" />
		</form>
	</body>

	</html>`
}
