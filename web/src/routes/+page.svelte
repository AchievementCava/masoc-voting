<script lang="ts">
	import Panel from "$lib/panel.svelte";
	import { currentPoll, polls } from "../store";
	import { goto } from "$app/navigation";
	const images = [
		"original",
		"christmas",
		"ball",
		"pride",
		"bbq",
		"graduation",
		"old-joe",
		"halloween",
	];
	const image = images[Math.floor(Math.random() * images.length)];

	$: upcomingPolls = $polls?.filter((e) => !e.isActive && !e.isConcluded) ?? [];

	$: if ($currentPoll && !$currentPoll.hasVoted) {
		goto(`/vote`);
	} else if (upcomingPolls.length > 0) {
		let poll = upcomingPolls[0];
		goto(`/${poll.pollType.name.toLowerCase()}/${poll.id}`);
	}
</script>

<svelte:head>
	<title>MaSoc Election Platform</title>
</svelte:head>

<Panel title="There are no upcoming elections">
	<p>Check this space later for updates.</p>
</Panel>

<style>
	img {
		margin-top: 16px;
	}
</style>
