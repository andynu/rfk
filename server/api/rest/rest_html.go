package rest

var indexHTML string = `
<html><head>
<style type='text/css'>
#skip { background-color: #ffefef; }
#reward { background-color: #efffef; }
</style>

<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
<script type='text/javascript'>

$(function(){
	updatePlaying = function(){
		$.getJSON('/status', function (data){
			$('#request_count').html(data['RequestCount']);
			$('#playpause_state').html(data['PlayPauseState']);
			$('#playing_path').html(data['CurrentSong']['Path']);
			$('#playing_hash').html(data['CurrentSong']['Hash']);
			$('#playing_rank').html(data['CurrentSong']['Rank']);
		});
	}
	updatePlaying();
	setInterval(updatePlaying, 2*1000);

	$('#skip').on('click', function(){ $.post('/skip'); updatePlaying(); });
	$('#next').on('click', function(){ $.post('/next'); updatePlaying(); });
	$('#reward').on('click', function(){ $.post('/reward'); updatePlaying(); });
	$('#playpause').on('click', function(){ $.post('/playpause'); updatePlaying(); });
});

</script>
</head><body>
<h1>RFK!</h1>

<div id='playpause_state'></div>

<div id='playing'>
	<div id='playing_path'></div>
	<div id='playing_hash'></div>
	<div id='playing_rank'></div>
</div>

<div id='actions'>
<button id='skip'>skip (-)</button>
<button id='next'>next</button>
|
<button id='playpause'>play/pause</button>
|
<button id='reward'>reward (+)</button>
</div>

<div id='requests'>
Requests: <span id='request_count'></span>
</div>

</body></html>
`
