<template>
    <div>
        <h1 class="mb-2">Challenge {{ idChallenge }}</h1>
        <h3 class="mb-2">powered by Trivago</h3>

        <p class="mb-4">Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet,</p>

        <v-card class="mb-4">
            <v-card-title>
                <h3 class="title mb-2">What you will get</h3>
                <div>
                    <v-chip color="orange" text-color="white">
                        <v-icon left>laptop</v-icon>
                        Laptop
                    </v-chip>
                    <v-chip color="indigo" text-color="white">
                        <v-icon left>code</v-icon>
                        Udemy
                    </v-chip>
                </div>
            </v-card-title>
        </v-card>

        <v-btn class="mb-4" :to="{ path: '/refugee/register' }" color="primary">Sign up</v-btn>

        <v-card>
            <v-card-title>
                <h3 class="headline mb-2">Trivago</h3>
                <v-layout>
                    <v-flex xs2>
                        <v-avatar
                                :tile="tile"
                                :size="avatarSize"
                                class="grey lighten-4"
                        >
                            <img src="/static/apple-touch-icon-180x180.png" alt="avatar">
                        </v-avatar>
                    </v-flex>
                    <v-flex xs9 offset-xs1>
                        <p class="mb-0">Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt</p>
                    </v-flex>
                </v-layout>
            </v-card-title>
            <v-card-title primary-title class="pt-0">
                <h3>Düsseldorf, Germany</h3>
                <a href="trivago.de">Website</a>
            </v-card-title>
        </v-card>
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