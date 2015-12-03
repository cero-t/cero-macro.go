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
	return post(player + "=" + obj.value + " 2");
}

function sendName(obj, player) {
	return post(player + "=" + obj.name + " 2");
}

function submitForm(form) {
	return post("player1=" + form['player1'].value + "&" + "player2=" + form['player2'].value);
}

function sendStop() {
	return post("stop=stop");
}

function post(message) {
	var xhr = new XMLHttpRequest();
	xhr.onload = function() {
		// Do something?
	}
	xhr.open('POST', "./", true);
	xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
	xhr.send(message);
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
	<input type="submit" id="run" name="run" value="run"/>
	<input type="button" id="stop" name="stop" value="stop" onclick="return sendStop();"/>
</form>
<br/>
(for training mode only)<br/>
<input type="submit" id="save_1" name="save" value="save (1P)" onclick="return sendName(this, 'player1');"/>
<input type="submit" id="reload_1" name="reload" value="reload (1P)" onclick="return sendName(this, 'player1');"/>
&nbsp;
<input type="submit" id="save_2" name="save" value="save (2P)" onclick="return sendName(this, 'player2');"/>
<input type="submit" id="reload_2" name="reload" value="reload (2P)" onclick="return sendName(this, 'player2');"/>
<br/>
<br/>
<br/>

<b>Player1</b>
<div>
	<div class="directions">
		<div>
			<button id="d7_1" value="7" class="directionButton" onclick="return send(this, 'player1');">7</button>
			<button id="d8_1" value="8" class="directionButton" onclick="return send(this, 'player1');">8</button>
			<button id="d9_1" value="9" class="directionButton" onclick="return send(this, 'player1');">9</button>
		</div>
		<div>
			<button id="d4_1" value="4" class="directionButton" onclick="return send(this, 'player1');">4</button>
			<button id="d5_1" value="5" class="directionButton" onclick="return send(this, 'player1');">5</button>
			<button id="d6_1" value="6" class="directionButton" onclick="return send(this, 'player1');">6</button>
		</div>
		<div>
			<button id="d1_1" value="1" class="directionButton" onclick="return send(this, 'player1');">1</button>
			<button id="d2_1" value="2" class="directionButton" onclick="return send(this, 'player1');">2</button>
			<button id="d3_1" value="3" class="directionButton" onclick="return send(this, 'player1');">3</button>
		</div>
	</div>
	<div class="commands">
		<div>
			<button id="lp_1" value="lp" class="commandButton" onclick="return send(this, 'player1');">LP</button>
			<button id="mp_1" value="mp" class="commandButton" onclick="return send(this, 'player1');">MP</button>
			<button id="hp_1" value="hp" class="commandButton" onclick="return send(this, 'player1');">HP</button>
		</div>
		<div>
			<button id="lk_1" value="lk" class="commandButton" onclick="return send(this, 'player1');">LK</button>
			<button id="mk_1" value="mk" class="commandButton" onclick="return send(this, 'player1');">MK</button>
			<button id="hk_1" value="hk" class="commandButton" onclick="return send(this, 'player1');">HK</button>
		</div>
		<div>
			<button id="pause_1" value="pause" class="commandButton" onclick="return send(this, 'player1');">Pause</button>
		</div>
	</div>
	<div style="clear: both;"></div>
</div>
<br/>
<b>Player2</b>
<div>
	<div class="directions">
		<div>
			<button id="d7_2" value="7" class="directionButton" onclick="return send(this, 'player2');">7</button>
			<button id="d8_2" value="8" class="directionButton" onclick="return send(this, 'player2');">8</button>
			<button id="d9_2" value="9" class="directionButton" onclick="return send(this, 'player2');">9</button>
		</div>
		<div>
			<button id="d4_2" value="4" class="directionButton" onclick="return send(this, 'player2');">4</button>
			<button id="d5_2" value="5" class="directionButton" onclick="return send(this, 'player2');">5</button>
			<button id="d6_2" value="6" class="directionButton" onclick="return send(this, 'player2');">6</button>
		</div>
		<div>
			<button id="d1_2" value="1" class="directionButton" onclick="return send(this, 'player2');">1</button>
			<button id="d2_2" value="2" class="directionButton" onclick="return send(this, 'player2');">2</button>
			<button id="d3_2" value="3" class="directionButton" onclick="return send(this, 'player2');">3</button>
		</div>
	</div>
	<div class="commands">
		<div>
			<button id="lp_2" value="lp" class="commandButton" onclick="return send(this, 'player2');">LP</button>
			<button id="mp_2" value="mp" class="commandButton" onclick="return send(this, 'player2');">MP</button>
			<button id="hp_2" value="hp" class="commandButton" onclick="return send(this, 'player2');">HP</button>
		</div>
		<div>
			<button id="lk_2" value="lk" class="commandButton" onclick="return send(this, 'player2');">LK</button>
			<button id="mk_2" value="mk" class="commandButton" onclick="return send(this, 'player2');">MK</button>
			<button id="hk_2" value="hk" class="commandButton" onclick="return send(this, 'player2');">HK</button>
		</div>
		<div>
			<button id="pause_2" value="pause" class="commandButton" onclick="return send(this, 'player2');">Pause</button>
		</div>
	</div>
	<div style="clear: both;"></div>
</div>
</body>
</html>
`
}
