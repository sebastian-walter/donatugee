<template>
    <div>
        <h1 class="display-1 mb-4">Discover Challenges</h1>
        <p class="mb-4" v-html="randomText.text_out"></p>
        <template v-for="challenge in challenges">
            <challenge-item :challenge="challenge"></challenge-item>
        </template>
    </div>
</template>
<script>
    import { getChallenges, getRandomText } from '../../api/api';
    import ChallengeItem from '../../components/challenge-item/ChallengeItem.vue';

    export default {
    	name: 'ChallengeList',
        data() {
            return {
                challenges: [],
                randomText: String
            };
        },
        mounted() {
            getChallenges().then(response => {
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