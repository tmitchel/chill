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
            <td>{Math.floor(stats.TimeWorking * 1e-9)}</td>
        </tr>
        <tr>
            <td>Time Chillin</td>
            <td>{Math.floor(stats.TimeChilling * 1e-9)}</td>
        </tr>
        <tr>
            <td>Water Drank</td>
            <td>{stats.WaterDrank}</td>
        </tr>
        <tr>
            <td>Tasks Complete</td>
            <td>{stats.TasksComplete}</td>
        </tr>
    </tbody>
</table>
