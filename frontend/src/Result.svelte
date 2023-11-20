<script lang="ts">
	import {onMount} from "svelte";

	export let error: string

	let show: boolean = true

	onMount(() => {
		const timeout = setTimeout(() => {
			show = false
		}, 4000)

		return () => clearTimeout(timeout)
	})
</script>

{#if show}
	<div>
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
		{/if}

		<br>
	</div>
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