<script>
	import Modal from 'svelte-simple-modal';
	import App from './components/App.svelte'
	import "../node_modules/bulma/css/bulma.min.css";
</script>

<main>
	<div class="App">
		<header class="App-header">
			<Modal>
				<App/>
			</Modal>
		</header>
	</div>
</main>

<style>
	:global(body) {
		margin: 0;
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
			"Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
			sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	code {
		font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New", monospace;
	}

	.App {
		text-align: center;
	}
</style>