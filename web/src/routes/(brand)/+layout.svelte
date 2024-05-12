<script lang="ts">
	import Icon from '@iconify/svelte';

	let navOpen = false;
</script>

<div>
	<div class="bg-base border-base-content sticky top-0 border-b-[1px] border-opacity-15 px-8">
		<div class="m-auto flex max-w-screen-xl flex-row items-center justify-between py-4">
			<figure class="font-header text-3xl"><a href="/">deltaform</a></figure>
			<nav class="hidden md:block">
				<ul class="flex flex-row gap-8">
					<li><a href="/">Guide</a></li>
					<li><a href="/">Licence</a></li>
					<li><a href="/">Community</a></li>
					<li>
						<a href="/">
							<Icon icon="ri:external-link-line" class="inline text-sm" />
							GitHub
						</a>
					</li>
				</ul>
			</nav>

			<a
				href="/login"
				class="bg-accent text-base-content hidden rounded-md p-2 px-6 text-base text-xl font-semibold md:block"
			>
				Login
			</a>
			<button on:click={() => (navOpen = !navOpen)}>
				{#if navOpen}
					<Icon class="text-xl md:hidden" icon="ri:close-large-line" />
				{:else}
					<Icon class="text-xl md:hidden" icon="ri:menu-line" />
				{/if}
			</button>
		</div>
		{#if navOpen}
			<div
				class="border-base-content bg-base absolute right-0 z-50 w-full max-w-48 border border-opacity-10 p-4 md:hidden"
			>
				<a
					href="/login"
					class="bg-accent text-base-content mb-4 block rounded-md p-2 px-6 text-base"
				>
					Login
				</a>
				<nav class="md:hidden">
					<ul class="flex flex-col gap-4">
						<li><a href="/">Guide</a></li>
						<li><a href="/"> Community </a></li>
						<li><a href="/">Licence</a></li>
						<li>
							<a href="/">
								<Icon icon="ri:external-link-line" class="inline text-sm" />
								GitHub
							</a>
						</li>
					</ul>
				</nav>
			</div>
		{/if}
	</div>

	<slot></slot>
</div>
