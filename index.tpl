<!DOCTYPE html>
<html>
	<head>
		<title>{{.}}</title>
		<style>
			body {
				font-family: monospace;
			}
			#screen {
				border: 2px outset black;
				border-radius: 5px;
			}
			#logo {
				height: 50px;
				width: auto;
				float: right;
				margin-right: 10px;
			}
		</style>
	</head>
	<body>
		<h1>{{.}}<img id="logo" src="/static/nullptr.png" /></h1>
		<img id="screen" src="/img" alt="Simulation Screen"/>
		<script>
			setInterval(refresh,100)
			function refresh() {
				document.getElementById("screen").src =
					"/img?t=" + new Date().getTime();
			}
		</script>
	</body>
</html>
