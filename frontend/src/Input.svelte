<script lang="ts">
	import {createEventDispatcher} from 'svelte';

	const dispatch = createEventDispatcher();

	let uni: string = "TUM"
	let id: string = ""

	function onChange(event) {
		let id = event.target.value;
		if (id.endsWith("$")) {
			document.getElementById("submit-button").click();
		}
	}

	function submit() {
		dispatch('message', {
			uni: uni,
			id: id,
		})

		// Reset values
		uni = "TUM"
		id = ""
	}
</script>

<div class="input-box" id="input">
	<fieldset>
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
					bind:group={uni}
					id="other"
					name="uni"
					type="radio"
					value="other"
			/>
			<label for="other">Other</label>
		</div>

		<br>

		<div>
			<label for="id">Student card ID</label>
			<br>
			<input autocomplete="off" autofocus on:input={onChange} bind:value={id} class="input" id="id" type="text"/>
			<br><br>
			<button class="btn" id="submit-button" on:click={submit}>Submit</button>
		</div>
	</fieldset>
</div>

<style>
	input[type="radio"] {
        margin-right: 10px;
    }

	input[type='radio']:after {
		width: 15px;
		height: 15px;
		border-radius: 15px;
		top: -2px;
		left: -1px;
		position: relative;
		background-color: #d1d3d1;
		content: '';
		display: inline-block;
		visibility: visible;
		border: 2px solid white;
	}

	input[type='radio']:checked:after {
		width: 15px;
		height: 15px;
		border-radius: 15px;
		top: -2px;
		left: -1px;
		position: relative;
		background-color: #A2C617;
		content: '';
		display: inline-block;
		visibility: visible;
		border: 2px solid white;
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