<script>
    import { getContext } from 'svelte';
    import * as Wails from '@wailsapp/runtime';
    import Tasks from './Tasks.svelte';
    import Stats from './Stats.svelte';
    import Quote from './Quote.svelte';

    export let seconds;
    export let chilling;
    export let endable;

    Wails.Events.On('tick', (s) => (seconds = s));
    Wails.Events.On('chilling', () => {
        gotStats = getStats();
        chilling = true;
        gotQuote = getQuote();
    });
    Wails.Events.On('working', () => {
        gotStats = getStats();
        chilling = false;
        endable = false;
    });
    Wails.Events.On('endable', () => (endable = true));

    const skipBreak = () => {
        window.backend.Timer.SkipBreak().then(() => console.log('break skipped'));
    };
    const startBreak = () => {
        window.backend.Timer.StartBreak().then(() => console.log('break started'));
    };
    const endBreak = () => {
        window.backend.Timer.EndBreak().then(() => console.log('break ended'));
    };

    const getQuote = async () => {
        return await window.backend.Quotes.RandomQuote();
    };
    let gotQuote = getQuote();

    const getStats = async () => {
        return await window.backend.Stats.Get();
    };
    let gotStats = getStats();
</script>

<style>
    .main {
        margin-top: 5vh;
    }
</style>

<main>
    <div class="main">
        <h1 class="title is-1">Chill out, my guy</h1>
        {#if chilling}
            <h1 class="subtitle is-3">Chillin' - {seconds}s</h1>
        {:else}
            <h1 class="subtitle is-3">Working - {seconds}s</h1>
        {/if}

        {#await gotQuote}
            <p>wait</p>
        {:then quote}
            <svelte:component this={Quote} {quote} />
        {/await}

        <div class="tile is-ancestor">
            <svelte:component this={Tasks} />

            <div class="tile is-parent is-vertical is-4">
                <div class="tile is-child">
                    {#await gotStats}
                        <p>wait</p>
                    {:then stats}
                        <svelte:component this={Stats} {stats} />
                    {/await}
                </div>
                <div class="tile is-child">
                    <h2 class="subtitle is-3 mt-4">Actions</h2>
                    {#if chilling}
                        <button type="button" class="button is-info" on:click={skipBreak}> Skip Break </button>
                        {#if endable}
                            <button type="button" class="button is-info" on:click={endBreak}> End Break </button>
                        {:else}<button disabled type="button" class="button is-info" on:click={endBreak}> End Break </button>{/if}
                    {:else}<button type="button" class="button is-info" on:click={startBreak}> Start Break </button>{/if}
                </div>
            </div>
        </div>
    </div>
</main>
