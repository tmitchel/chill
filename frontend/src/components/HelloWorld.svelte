<script>
	import { getContext } from 'svelte';
	
	const skipBreak = () => {
		window.backend.Timer.SkipBreak().then(() => console.log('break skipped'))
	};
	const startBreak = () => {
		window.backend.Timer.StartBreak().then(() => console.log('break started'))
	};

	const getQuote = async () => {
		return await window.backend.Quotes.RandomQuote();
	}
	const gotQuote = getQuote()

	const getTasks = async () => {
		return await window.backend.Tasks.Tasks();
	}
	const gotTasks = getTasks();
</script>

<main>
	<h1>Chill</h1>
	{#await gotQuote}
		<p>wait</p>
	{:then quote}
		<blockquote>{quote.Content}</blockquote>
		<p>- {quote.Author}</p>
	{/await}

	<h2>Tasks</h2>
	{#await gotTasks}
		<p>wait</p>
	{:then tasks}
	<ul>
		{#each tasks as task}
			<li>{task.Content}</li>
		{/each}
	</ul>
	{/await}

	<h2>Actions</h2>
	<p><button type="button" class="button is-primary" on:click={skipBreak}>Skip Break</button></p>
	<p><button type="button" class="button is-primary" on:click={startBreak}>Start Break</button></p>
</main>

<style></style>