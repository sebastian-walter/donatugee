<template>
    <div>
        <h1 class="headline mb-2">{{ challenge.Name }}</h1>
        <h3 class="mb-2">powered by {{ donator.Name }}</h3>

        <p class="mb-4">{{ challenge.Description }}</p>

        <v-card class="mb-4">
            <v-card-title>
                <h3 class="title mb-2 display-1">What you will get</h3>
            </v-card-title>
            <v-divider></v-divider>
            <v-card-text>
                <div>
                    <v-chip color="orange" text-color="white">
                        <v-icon left>laptop</v-icon>
                        {{ challenge.LaptopType }}
                    </v-chip>
                    <v-chip color="indigo" text-color="white">
                        <v-icon left>code</v-icon>
                        {{ challenge.Amount }} € Udemy Coupon
                    </v-chip>
                </div>
            </v-card-text>
        </v-card>

        <v-btn v-if="!userId" class="mb-4" @click="handleSignUp()" color="primary">Sign up</v-btn>
        <v-btn v-if="!userId" class="mb-4" @click="handleLogin()" color="primary">Login</v-btn>
        <v-btn v-if="userId && !applied" class="mb-4" @click="handleApply()" color="primary">Apply</v-btn>
        <v-btn v-if="userId && applied && !accepted" class="mb-4" disabled><icon name="hourglass-o" class="mr-2"></icon>Already applied</v-btn>
        <v-btn v-if="userId && applied && accepted" class="mb-4 green white--text"><icon name="check" class="mr-2"></icon>Accepted</v-btn>
        <v-card>
            <v-card-title>
                <h3 class="headline mb-2">{{ donator.Name }}</h3>
            </v-card-title>
            <v-card-title class="pt-0">
                <v-flex xs2>
                    <v-avatar
                            class="grey lighten-4"
                    >
                        <img src="https://lorempixel.com/180/180/cats/" alt="avatar">
                    </v-avatar>
                </v-flex>
                <v-flex xs9 offset-xs1 v-html="randomText.text_out">
                </v-flex>
            </v-card-title>
            <v-card-text primary-title class="pt-0">
                <h3>{{ donator.Address }}</h3>
                <a :href="'http://' + donator.Website" target="_blank">{{ donator.Website }}</a>
            </v-card-text>
        </v-card>
    </div>
</template>
<script>
    import { getChallenge, getDonator, getRandomText, setApplication } from '../../api/api';

    export default {
		name: 'ChallengeDetail',
        data() {
    		return {
    		    idChallenge: this.$route.params.id,
                challenge: { Applications: [] },
                donator: {},
                randomText: '',
                hasApplied: false,
                userId: window.localStorage.getItem('userId')
            }
        },
        mounted() {
            getChallenge(this.idChallenge).then(response => {
                this.challenge = response.data;
            }).then(() => {
                getDonator(this.challenge.DonatorID).then(response => {
                    this.donator = response.data;
                });
            });

            getRandomText('10-20').then(response => {
                this.randomText = response.data;
            });
        },
        methods: {
            handleSignUp() {
                window.localStorage.setItem('idChallenge', this.idChallenge);
                this.$router.push({
                    path: '/refugee/register'
                });
            },

            handleLogin() {
                window.localStorage.setItem('idChallenge', this.idChallenge);
                this.$router.push({
                    path: '/refugee/login'
                });
            },

            handleApply() {
                setApplication(this.challenge.ID, this.userId).then(
                	response => {
                	    this.hasApplied = true;
                    }
                );
            }
        },
        computed: {
            applied() {
            	if (this.hasApplied) {
            		return true;
                }

                return this.challenge.Applications.filter((application) => {
                    return application.TechfugeeID === parseInt(this.userId);
                }).length > 0;
            },

            accepted() {
                return this.challenge.Applications.filter((application) => {
                    return (application.TechfugeeID === parseInt(this.userId)) && application.Accepted === true;
                }).length > 0;
            }
        }
    }
</script>