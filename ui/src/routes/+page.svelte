<script>
import "$lib/styles.css"
import LoginForm from "./components/login_form.svelte"

let searching = $state(false)
</script>

<div class="bg-stone-800 h-dvh flex justify-center items-center">
	{#if !searching}
		<LoginForm bind:searching={searching} />
	{:else }
		carregando
	{/if}
</div>
