<template>
    <div>
        <h1>Discover Challenges</h1>
        <p>Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet,</p>
        <challenge-item></challenge-item>
        <challenge-item></challenge-item>
    </div>
</template>
<script>
    import axios from 'axios';
    import ChallengeItem from '../../components/challenge-item/ChallengeItem.vue';

    export default {
    	name: 'ChallengeList',
        data() {
            return {
                challenges: [],
            };
        },
        mounted() {
            axios.get('')
                .then(response => {
                this.challenges = response.data
            })
            .catch(e => {
                console.error(e);
            })
        },
        methods: {

        },
        components: {
            ChallengeItem,
        }
    }
</script>