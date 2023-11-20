<script lang="ts">
	export let error:string
</script>

{#if error == null}
	<h3 class="denied">An internal error occurred: Error is null</h3>
{:else}
	{#if error.length === 0}
		<h3 class="allowed">The student is allowed to get an item</h3>
	{:else}
		<div class="denied">
			<h3>The student is not allowed to get an item</h3>
			<p>{error}</p>
		</div>
	{/if}

	<br>
{/if}


<style>
	.allowed {
		background-color: #4caf50;
		color: #fff;
		padding: 10px 20px;
		border-radius: 10px;
	}

	.denied {
		background-color: #ff0000;
		color: #fff;
		padding: 10px 20px;
		border-radius: 10px;
	}
</style>