package rest

var indexHTML string = `
<html><head>
<link href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.6/css/bootstrap.css" rel="stylesheet"></link>
<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.5.0/css/font-awesome.min.css" rel="stylesheet"></link>
<style type='text/css'>
.title {
	margin-top: 0;
	margin-left: -15px;
	margin-right: -15px;
	margin-bottom: 10px;
	padding: 2px;
	background: #222;
	border-color: #080808;
	color: #999;
	font-size: 10px;
}
</style>
</head><body><div class='container-fluid'>
<div class='row'><div class='col-xs-12'>

	<div class='title'>RFK!</div>

	<div class='pull-right'>
		<div class='btn-toolbar'>
			<div class='btn-group'>
				<button id='reward' class='btn btn-success'><i class='fa fa-thumbs-o-up'/></i></button>
				<button id='skip' class='btn btn-danger'><i class='fa fa-thumbs-o-down'/></i></button>
			</div>
		</div>
		<div class='btn-toolbar'>
			<div class='btn-group'>
				<button id='playpause' class='btn btn-default'><i class='fa fa-play'></i></button>
				<button id='next' class='btn btn-default'><i class='fa fa-forward'/></i></button>
			</div>
		</div>
	</div>

	<div id='playing'>
		<b id='playing_title'></b>
		<span id='playing_rank' class='badge'></span>
		<div id='playing_artist'></div>
		<em id='playing_album'></em>
	</div>

	<hr/>

	<div id='requests' class='well'>
		<small>
		Path: <div id='playing_path'></div>
		Hash: <div id='playing_hash' class='text-muted'></div>
		Requests: <span id='request_count'></span>
		</small>
	</div>

</div></div><!-- /col/row -->

<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.6/js/bootstrap.min.js"></script>
<script type='text/javascript'>

$(function(){
	updatePlaying = function(){
		$.getJSON('/status', function (data){
			$('#request_count').html(data['RequestCount']);
			var playpause_icon = data['PlayPauseState']=='playing' ? 'fa fa-stop' : 'fa fa-play';
			$('#playpause i').attr('class', playpause_icon);
			$('#playing_title').html(data['CurrentSongMeta']['Title']);
			$('#playing_artist').html(data['CurrentSongMeta']['Artist']);
			$('#playing_album').html(data['CurrentSongMeta']['Album']);
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
</div></body></html>
`
