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
		$.getJSON('/playing', function (data){
			$('#playing_path').html(data['Path']);
			$('#playing_hash').html(data['Hash']);
			$('#playing_rank').html(data['Rank']);
		});
	}
	updatePlaying();
	setInterval(updatePlaying, 2*1000);

	$('#skip').on('click', function(){ $.post('/skip'); updatePlaying(); });
	$('#next').on('click', function(){ $.post('/next'); updatePlaying(); });
	$('#reward').on('click', function(){ $.post('/reward'); updatePlaying(); });
});

</script>
</head><body>
<h1>RFK!</h1>

<div id='playing'>
	<div id='playing_path'></div>
	<div id='playing_hash'></div>
	<div id='playing_rank'></div>
</div>

<div id='actions'>
<button id='skip'>skip (-)</button>
<button id='next'>next</button>
<button id='reward'>reward (+)</button>
</div>

</body></html>
`
