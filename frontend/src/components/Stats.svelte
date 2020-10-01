<script>
    export let stats;
</script>

<style>
    .center {
        width: 60%;
        margin: auto;
    }
</style>

<table class="table center">
    <thead>
        <tr>
            <th>Daily Stats</th>
            <td />
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Time Working</td>
            <td>{Math.floor(stats.time_working * 1e-9)}</td>
        </tr>
        <tr>
            <td>Time Chillin</td>
            <td>{Math.floor(stats.time_chilling * 1e-9)}</td>
        </tr>
        <tr>
            <td>Water Drank</td>
            <td>{stats.water_drank}</td>
        </tr>
        <tr>
            <td>Tasks Complete</td>
            <td>{stats.tasks_complete}</td>
        </tr>
    </tbody>
</table>
