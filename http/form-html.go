package http

func formHtml() string {
	return `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Cero Macro Form</title>
	<style type="text/css">
		.directions {float: left;}
		.commands {float: left; padding-left: 10px;}
		.directionButton {width: 40px; height: 40px;}
		.commandButton {width: 80px; height: 40px;}
	</style>
	<script type="text/javascript">
function send(obj, player) {
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		// Do something.
	}
	xhr.open('POST', "./", true);
	xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
	xhr.send(player + "=" + obj.value + " 2");
	return false;
}

function submitForm(form) {
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		// Do something.
	}
	xhr.open('POST', "./", true);
	xhr.setRequestHeader( 'Content-Type', 'application/x-www-form-urlencoded' );
	xhr.send("player1=" + form['player1'].value + "&" + "player2=" + form['player2'].value);
	return false;
}
	</script>
</head>

<body>
<form id="macroForm" name="macroForm" action="." method="POST" onsubmit="return submitForm(this);">
	<b>Player1</b><br/>
	<textarea id="player1" name="player1" rows="10" cols="40">{{.Player1}}</textarea><br/>

	<b>Player2</b><br/>
	<textarea id="player2" name="player2" rows="10" cols="40">{{.Player2}}</textarea><br/>
	<input type="submit" id="submit" name="submit" value="submit"/>
</form>
<br/>
<br/>

<b>Player1</b>
<div>
	<div class="directions">
		<div>
			<button id="d7" value="7" class="directionButton" onclick="return send(this, 'player1');">7</button>
			<button id="d8" value="8" class="directionButton" onclick="return send(this, 'player1');">8</button>
			<button id="d9" value="9" class="directionButton" onclick="return send(this, 'player1');">9</button>
		</div>
		<div>
			<button id="d4" value="4" class="directionButton" onclick="return send(this, 'player1');">4</button>
			<button id="d5" value="5" class="directionButton" onclick="return send(this, 'player1');">5</button>
			<button id="d6" value="6" class="directionButton" onclick="return send(this, 'player1');">6</button>
		</div>
		<div>
			<button id="d1" value="1" class="directionButton" onclick="return send(this, 'player1');">1</button>
			<button id="d2" value="2" class="directionButton" onclick="return send(this, 'player1');">2</button>
			<button id="d3" value="3" class="directionButton" onclick="return send(this, 'player1');">3</button>
		</div>
	</div>
	<div class="commands">
		<div>
			<button id="lp" value="lp" class="commandButton" onclick="return send(this, 'player1');">LP</button>
			<button id="mp" value="mp" class="commandButton" onclick="return send(this, 'player1');">MP</button>
			<button id="hp" value="hp" class="commandButton" onclick="return send(this, 'player1');">HP</button>
		</div>
		<div>
			<button id="lk" value="lk" class="commandButton" onclick="return send(this, 'player1');">LK</button>
			<button id="mk" value="mk" class="commandButton" onclick="return send(this, 'player1');">MK</button>
			<button id="hk" value="hk" class="commandButton" onclick="return send(this, 'player1');">HK</button>
		</div>
		<div>
			<button id="pause" value="pause" class="commandButton" onclick="return send(this, 'player1');">Pause</button>
			<button id="save" value="save" class="commandButton" onclick="return send(this, 'player1');">Save</button>
			<button id="reload" value="reload" class="commandButton" onclick="return send(this, 'player1');">Reload</button>
		</div>
	</div>
	<div style="clear: both;"></div>
</div>
<br/>
<b>Player2</b>
<div>
	<div class="directions">
		<div>
			<button id="d7" value="7" class="directionButton" onclick="return send(this, 'player2');">7</button>
			<button id="d8" value="8" class="directionButton" onclick="return send(this, 'player2');">8</button>
			<button id="d9" value="9" class="directionButton" onclick="return send(this, 'player2');">9</button>
		</div>
		<div>
			<button id="d4" value="4" class="directionButton" onclick="return send(this, 'player2');">4</button>
			<button id="d5" value="5" class="directionButton" onclick="return send(this, 'player2');">5</button>
			<button id="d6" value="6" class="directionButton" onclick="return send(this, 'player2');">6</button>
		</div>
		<div>
			<button id="d1" value="1" class="directionButton" onclick="return send(this, 'player2');">1</button>
			<button id="d2" value="2" class="directionButton" onclick="return send(this, 'player2');">2</button>
			<button id="d3" value="3" class="directionButton" onclick="return send(this, 'player2');">3</button>
		</div>
	</div>
	<div class="commands">
		<div>
			<button id="lp" value="lp" class="commandButton" onclick="return send(this, 'player2');">LP</button>
			<button id="mp" value="mp" class="commandButton" onclick="return send(this, 'player2');">MP</button>
			<button id="hp" value="hp" class="commandButton" onclick="return send(this, 'player2');">HP</button>
		</div>
		<div>
			<button id="lk" value="lk" class="commandButton" onclick="return send(this, 'player2');">LK</button>
			<button id="mk" value="mk" class="commandButton" onclick="return send(this, 'player2');">MK</button>
			<button id="hk" value="hk" class="commandButton" onclick="return send(this, 'player2');">HK</button>
		</div>
		<div>
			<button id="pause" value="pause" class="commandButton" onclick="return send(this, 'player2');">Pause</button>
			<button id="save" value="save" class="commandButton" onclick="return send(this, 'player2');">Save</button>
			<button id="reload" value="reload" class="commandButton" onclick="return send(this, 'player2');">Reload</button>
		</div>
	</div>
	<div style="clear: both;"></div>
</div>
</body>
</html>
`
}
