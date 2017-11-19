<template>
    <div>
        <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
            {{ this.errorMessage }}
        </v-alert>
        <h1 class="mb-4 display-1">Your Challenges
            <v-btn class="create-challenge"
                   color="error"
                   fab
                   dark
                   to="/company/challenge/create"
            >
                <v-icon>add</v-icon>
            </v-btn>
        </h1>
        <p class="mb-4" v-html="randomText.text_out"></p>
        <v-alert v-if="companyChallenges.length === 0" color="info" value="true">
            You don't have any yet. Go create some!! Hush hush!!
        </v-alert>
        <template v-for="challenge in companyChallenges">
            <challenge-item :challenge="challenge" class="challenge-item"></challenge-item>
        </template>
        <div class="create-challenge">
        </div>
    </div>
</template>
<script>
	import {getChallengeForTechfugee, getRandomText} from '../../api/api';
	import ChallengeItem from '../../components/challenge-item/ChallengeItem.vue';
	import {mapActions, mapState} from 'vuex';

	export default {
		name: 'YourChallengeList',
		data() {
			return {
				randomText: String,
				errorMessage: '',
			};
		},
		computed: {
			...mapState({
				donator: state => state.donator,
				companyChallenges: state => state.companyChallenges,
			}),
		},
		mounted() {
			this.getChallengesForDonator({id: JSON.parse(window.localStorage.getItem('companyId'))}).then(response => {
				if (response.status !== 200) {
					this.errorMessage = 'Oooops, something did not work. Please try again.';
				}
			});

			getRandomText('30-50').then(response => {
				this.randomText = response.data;
			});
		},
		methods: {
			...mapActions([
				'getChallengesForDonator',
			])
		},
		components: {
			ChallengeItem,
		},
	};
</script>

<style lang="scss" type="text/scss">
    .create-challenge {
        position: absolute;
        bottom: 1rem;
        right: 1rem;
    }
</style>
