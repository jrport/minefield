<script>
	let {searching} = $props();
	let name = $state("");

	let deactivateBtn = $derived(name.length == 0)

	function getMatch() {
		let ws = new WebSocket(new URL("http://localhost:3000/get_match"));

		ws.onerror = (event => {
			console.log(event);
			alert(event.error);
		})
		ws.onclose = (event => {
			alert("fecho");
		})
		ws.onopen = (event => {
			alert("foi");
		})
		ws.onmessage = (event => {
			alert("msg");
		})
	}
</script>


<form onsubmit={getMatch} class="bg-stone-900 text-secondary p-4 w-[20%] flex rounded-lg shadow-xl">
	<span class="material-symbols-outlined col-span-1 row-span-2 w-fit" style="font-size: 82px;">
		bomb
	</span>
	<div class="flex-1 flex items-center flex-col">
		<div class="flex-1 font-bold text-2xl text-center mb-[0.3rem]">Campo Minado</div>
		<div class="flex mb-[0.5rem]">
			<input bind:value={name} type="text" class="mr-2 input input-xs input-bordered w-[70%] text-white" placeholder="Apelido..."/>	
			<button disabled={deactivateBtn} type="submit" class="btn btn-xs btn-secondary w-[30%] material-symbols-outlined">check</button>
		</div>
	</div>
</form>
