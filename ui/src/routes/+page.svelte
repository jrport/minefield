<script>
	let userName = $state("");
	let buttonDisabled = $derived(userName.length === 0);
</script>

<div id="login_box">
	<span class="material-symbols-outlined"> bomb </span>
	<h1>Minado Campo</h1>
	<form action="api/get_match" method="POST">
		<input
			type="text"
			bind:value={userName}
			placeholder="Apelido..."
		/>
		<button
			disabled={buttonDisabled}
			aria-label="Button to submit the login form"
			type="submit">Buscar partida</button
		>
	</form>
</div>

<style>
	@import url("https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@100..800&display=swap");
	@import url("https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200&icon_names=bomb");

	:global(body) {
		font-family: "JetBrains Mono", monospace;
		height: 100dvh;
		margin: 0;
		background-color: #1c1c1c;
		color: white;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	#login_box {
		box-shadow: 0.3rem 1rem 1rem rgba(33, 31, 39, 0.6);
		border-radius: 1rem;
		padding: 2rem;
		display: flex;
		flex-direction: column;
		align-items: center;
		background-color: #0c0c0c;
		width: 300px;
		gap: 1rem;
	}

	#login_box > span {
		margin-bottom: 1rem;
		font-size: 48pt;
		color: #ffcc00;
	}

	#login_box > h1 {
		margin: 0;
		font-size: 1.5rem;
	}

	form {
		width: 100%;
		display: flex;
		flex-direction: column;
		gap: 0.8rem;
	}

	input {
		width: 100%;
		padding: 0.5rem 0rem 0.5rem 0.5rem;
		font-size: 1rem;
		border: none;
		border-radius: 0.5rem;
		background: #262626;
		color: white;
		outline: none;
		box-sizing: border-box;
		box-shadow: inset 0 0 5px rgba(255, 255, 255, 0.1);
		transition: all 0.2s ease-in-out;
	}

	input:focus {
		box-shadow: 0 0 5px rgba(255, 255, 255, 0.3);
	}

	button {
		padding: 0.75rem;
		border: none;
		border-radius: 0.5rem;
		background: linear-gradient(135deg, #ffcc00, #ff8800);
		color: black;
		font-size: 1rem;
		font-weight: bold;
		cursor: pointer;
		transition: all 0.2s ease-in-out;
	}

	button:hover {
		background: linear-gradient(135deg, #ffdd44, #ff9933);
		transform: scale(1.05);
	}

	button:disabled {
		background: #444;
		cursor: not-allowed;
		opacity: 0.6;
		transform: none;
	}
</style>
