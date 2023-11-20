<script lang="ts">
	import logo from './assets/images/IKOM.svg'
	import {AddStudent} from '../wailsjs/go/main/App.js'

	let allowed: boolean = false
	let errorText: string
	let id: string
	let uni: string = "TUM"

	function checkID(): void {
		allowed = false
		errorText = ""

		AddStudent(id, uni).then(
				() => allowed = true
		).catch(
				error => {
					allowed = false
					errorText = error
				}
		)
	}
</script>

<main>
	<div>
		<img alt="Wails logo" id="logo" src="{logo}">
		<h2>Rucksack Ausgabe</h2>
	</div>

	{#if errorText != null && errorText.length > 0}
		<div class="result" id="result">{errorText}</div>
	{/if}
	<div>
		<p>Allowed: {allowed}</p>
		<p>Uni: {uni}</p>
	</div>
	<div class="input-box" id="input">
		<fieldset>
			<legend>Select a Uni:</legend>
			<div>
				<input
						bind:group={uni}
						id="TUM"
						name="uni"
						type="radio"
						value="TUM"/>
				<label for="TUM">TUM/LMU</label>
			</div>

			<div>
				<input
						id="other"
						name="uni"
						type="radio"
						value="other"
						bind:group={uni}
				/>
				<label for="other">Other</label>
			</div>

			<br>

			<div>
				<label for="id">Student card ID</label>
				<br>
				<input autocomplete="off" bind:value={id} class="input" id="id" type="text"/>
				<br><br>
				<button class="btn" on:click={checkID}>Submit</button>
			</div>
		</fieldset>
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

	.input-box .btn {
		width: 60px;
		height: 30px;
		line-height: 30px;
		border-radius: 3px;
		border: none;
		margin: 0 0 0 20px;
		padding: 0 8px;
		cursor: pointer;
	}

	.input-box .btn:hover {
		background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
		color: #333333;
	}

	.input-box .input {
		border: none;
		border-radius: 3px;
		outline: none;
		height: 30px;
		line-height: 30px;
		padding: 0 10px;
		background-color: rgba(240, 240, 240, 1);
		-webkit-font-smoothing: antialiased;
	}

	.input-box .input:hover {
		border: none;
		background-color: rgba(255, 255, 255, 1);
	}

	.input-box .input:focus {
		border: none;
		background-color: rgba(255, 255, 255, 1);
	}
</style>
