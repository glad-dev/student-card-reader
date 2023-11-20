<script lang="ts">
	import logo from './assets/images/IKOM.svg'
	import {AddStudent} from '../wailsjs/go/main/App.js'
	import Loading from "./Loading.svelte";
	import Input from "./Input.svelte";
	import Result from "./Result.svelte";

	let firstCall: boolean = true
	let promise: Promise<void> = new Promise(resolve => setTimeout(resolve, 1));

	function checkID(event): void {
		firstCall = false

		const uni = event.detail.uni
		const id = event.detail.id

		promise = AddStudent(id, uni)
	}
</script>

<main>
	<div>
		<img alt="Wails logo" id="logo" src="{logo}">
		<h2>Ausgabe</h2>
	</div>

	<div>
		{#await promise}
			<Loading/>
		{:then _}
			{#if !firstCall}
				<Result error="" />
			{/if}

			<Input on:message={checkID}/>
		{:catch err}
			<Result error={err} />

			<Input on:message={checkID}/>
		{/await}
	</div>
</main>

<style>
	#logo {
		display: block;
		width: 50%;
		height: 50%;
		margin: auto;
		padding: 10% 0 0;
		background-position: center;
		background-repeat: no-repeat;
		background-size: 100% 100%;
		background-origin: content-box;
	}

	.result {
		height: 20px;
		line-height: 20px;
		margin: 1.5rem auto;
	}
</style>
