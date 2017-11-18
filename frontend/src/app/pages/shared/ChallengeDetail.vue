<template>
    <div>
        <h1>Chalenge Detail {{ idChallenge }}
        </h1>
    </div>
</template>
<script>
    export default {
    	name: 'ChallengeDetail',
        data() {
    		return {
    		    idChallenge: this.$route.params.id,
            }
        },
        mounted() {

        }
    }
</script>