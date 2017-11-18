<template>
    <div>
        <h1 class="mb-4 display-1">Discover Challenges</h1>
        <p class="mb-4" v-html="randomText.text_out"></p>
        <template v-for="challenge in challenges">
            <challenge-item :challenge="challenge" class="challenge-item"></challenge-item>
        </template>
    </div>
</template>
<script>
	import { getChallengeForTechfugee, getRandomText } from '../../api/api';
	import ChallengeItem from '../../components/challenge-item/ChallengeItem.vue';

	export default {
		name: 'YourChallengeList',
		data() {
			return {
				challenges: [],
				randomText: String
			};
		},
		mounted() {
			getChallengeForTechfugee(window.localStorage.getItem('userId')).then(response => {
				this.challenges = response.data;
			});

			getRandomText('30-50').then(response => {
				this.randomText = response.data;
			});
		},
		methods: {

		},
		components: {
			ChallengeItem,
		}
	}
</script>
