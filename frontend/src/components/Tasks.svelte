<script>
    import * as Wails from '@wailsapp/runtime';

    export let taskContent;

    // fetch tasks from backend
    const getTasks = async () => {
        return await window.backend.Tasks.Tasks();
    };
    let gotTasks = getTasks();

    // handle events and send to the backend
    const toggleStatus = async (i) => {
        await window.backend.Tasks.ToggleStatus(i);
        gotTasks = getTasks();
    };

    const createTask = async () => {
        await window.backend.Tasks.Create(taskContent);
        gotTasks = getTasks();
        taskContent = '';
    };
</script>

<style>
    .new-task {
        width: 80%;
        margin: auto;
    }
</style>

<div class="tile is-child is-8">
    <h2 class="subtitle is-3">Tasks</h2>
    <div class="new-task">
        <div class="field has-addons">
            <div class="control is-expanded"><input class="input" type="text" placeholder="Task" bind:value={taskContent} /></div>
            <div class="control"><button class="button is-info" on:click={createTask}>Add</button></div>
        </div>
    </div>
    {#await gotTasks}
        <p>wait</p>
    {:then tasks}
        {#each tasks as task, i}
            {#if task.Completed}
                <div class="new-task">
                    <div class="field has-addons">
                        <div class="control"><button class="button is-success" on:click={() => toggleStatus(i)}>Done</button></div>
                        <div class="control is-expanded"><input class="input" type="text" placeholder={task.Content} readonly /></div>
                    </div>
                </div>
            {:else}
                <div class="new-task">
                    <div class="field has-addons">
                        <div class="control"><button class="button is-danger" on:click={() => toggleStatus(i)}>Todo</button></div>
                        <div class="control is-expanded"><input class="input" type="text" placeholder={task.Content} readonly /></div>
                    </div>
                </div>
            {/if}
        {/each}
    {/await}
</div>
