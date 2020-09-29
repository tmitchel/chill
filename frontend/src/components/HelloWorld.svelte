<script>
	import { getContext } from 'svelte';
	import * as Wails from '@wailsapp/runtime';

	export let seconds;
	export let taskContent;
	export let chilling;

	Wails.Events.On('tick', s => seconds = s);
	Wails.Events.On('start-break', () => chilling = true);
	Wails.Events.On('end-break', () => chilling = false);

	const skipBreak = () => {
		window.backend.Timer.SkipBreak().then(() => console.log('break skipped'))
	};
	const startBreak = () => {
		window.backend.Timer.StartBreak().then(() => console.log('break started'))
	};
	const endBreak = () => {
		window.backend.Timer.EndBreak().then(() => console.log('break ended'))
	};

	const createTask = async () => {
		await window.backend.Tasks.Create(taskContent);
		gotTasks = getTasks();
		taskContent = "";
	}

	const toggleStatus = async(i) => {
		await window.backend.Tasks.ToggleStatus(i);
		gotTasks = getTasks();
	}

	const getQuote = async () => {
		return await window.backend.Quotes.RandomQuote();
	}
	const gotQuote = getQuote()

	const getTasks = async () => {
		return await window.backend.Tasks.Tasks();
	}
	let gotTasks = getTasks();
</script>

<main>
	<div class="main">
		<h1 class="title is-1">Chill out, my guy</h1>
		{#if seconds > 0}
			<h1 class="subtitle is-3">Chillin' - {seconds}s</h1>
		{:else}
			<h1 class="subtitle is-3">Working</h1>
		{/if}

		{#await gotQuote}
			<p>wait</p>
		{:then quote}
			<article class="message is-dark quote">
				<div class="message-header">
				Quote of the Break
				</div>
				<div class="message-body">
					<p>{quote.Content}</p>
					<p class="is-italic">- {quote.Author}</p>
				</div>
			</article>
		{/await}

		<div class="tile is-ancestor">
			<div class="tile is-child is-8">
				<h2 class="subtitle is-3">Tasks</h2>
				<div class="new-task">
					<div class="field has-addons">
						<div class="control is-expanded">
							<input class="input" type="text" placeholder="Task" bind:value={taskContent}>
						</div>
						<div class="control">
							<button class="button is-info" on:click={createTask}>Add</button>
						</div>
					</div>
				</div>
				{#await gotTasks}
					<p>wait</p>
				{:then tasks}
				{#each tasks as task, i}
					{#if task.Completed}
						<div class="new-task">
							<div class="field has-addons">
								<div class="control">
									<button class="button is-success" on:click={() => toggleStatus(i)}>Done</button>
								</div>
								<div class="control is-expanded">
									<input class="input" type="text" placeholder={task.Content} readonly>
								</div>
							</div>
						</div>
					{:else}
						<div class="new-task">
							<div class="field has-addons">
								<div class="control">
									<button class="button is-danger" on:click={() => toggleStatus(i)}>Todo</button>
								</div>
								<div class="control is-expanded">
									<input class="input" type="text" placeholder={task.Content} readonly>
								</div>
							</div>
						</div>
					{/if}
				{/each}
				{/await}
			</div>

			<div class="tile is-parent is-vertical is-4">
				<div class="tile is-child">
					<table class="table quote">
						<thead>
						<tr>
							<th>Daily Stats</th>
							<td></td>
						</tr>
						</thead>
						<tbody>
							<tr>
								<td>Time Working</td>
								<td>0</td>
							</tr>
							<tr>
								<td>Time Chillin</td>
								<td>0</td>
							</tr>
							<tr>
								<td>Water Drank</td>
								<td>0</td>
							</tr>
							<tr>
								<td>Tasks Complete</td>
								<td>0</td>
							</tr>
						</tbody>
					</table>
				</div>
				<div class="tile is-child">
					<h2 class="subtitle is-3 mt-4">Actions</h2>
					{#if chilling}
						<button type="button" class="button is-info" on:click={skipBreak}>Skip Break</button>
						{#if seconds > 15}
							<button type="button" class="button is-info" on:click={endBreak}>End Break</button>
						{:else}
							<button disabled type="button" class="button is-info" on:click={endBreak}>End Break</button>
						{/if}
					{:else}
						<button type="button" class="button is-info" on:click={startBreak}>Start Break</button>
					{/if}
				</div>
			</div>
		</div>
	</div>
</main>

<style>
.main{
	margin-top: 5vh;	
}
.quote{
	width: 60%;
	margin: auto;
}

.new-task{
	width: 80%;
	margin: auto;
}
</style>