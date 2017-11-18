<template>
    <div>
        <h1 class="mb-4">Discover Challenges</h1>
        <p class="mb-4">Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet,</p>
        <template v-for="challenge in challenges">
            <challenge-item :challenge="challenge"></challenge-item>
        </template>
    </div>
</template>
<script>
    import { getChallenges } from '../../api/api';
    import ChallengeItem from '../../components/challenge-item/ChallengeItem.vue';

    export default {
    	name: 'ChallengeList',
        data() {
            return {
                challenges: [],
            };
        },
        mounted() {
            getChallenges().then(response => {
                this.challenges = response.data;
            }).catch(e => {
                console.error(e);
            });
        },
        methods: {

        },
        components: {
            ChallengeItem,
        }
    }
</script>