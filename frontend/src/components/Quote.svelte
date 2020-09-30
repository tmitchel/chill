<script>
    export let quote;

</script>

<article class="message is-dark quote">
    <div class="message-header">
    Quote of the Break
    </div>
    <div class="message-body">
        <p>{quote.Content}</p>
        <p class="is-italic">- {quote.Author}</p>
    </div>
</article>

<style>
.quote{
	width: 60%;
	margin: auto;
}
</style>