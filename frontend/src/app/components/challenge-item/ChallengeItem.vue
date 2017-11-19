<template>
    <div class="mb-4 challenge-item">
        <v-flex xs12>
            <v-card :to="{ path: 'challenge/' + challenge.ID }">
                <v-card-title>
                    <div class="card-title">
                        <h3 class="headline mb-2">{{ challenge.Name }}</h3>
                        <p>{{ challenge.Description }}</p>
                        <v-chip color="orange" text-color="white">
                            <v-icon left>laptop</v-icon>
                            {{ challenge.LaptopType }}
                        </v-chip>
                        <v-chip color="indigo" text-color="white">
                            <v-icon left>code</v-icon>
                            {{ challenge.Amount }} € Udemy Coupon
                        </v-chip>
                    </div>
                </v-card-title>
                <v-divider></v-divider>
                <v-card-title primary-title class="pt-3">
                    <icon v-if="!accepted && applied" class="applied-icon" name="hourglass-o"></icon>
                    <icon v-if="accepted && applied" class="check-icon green--text" name="check"></icon>
                    <v-flex xs2>
                        <v-avatar
                                class="grey lighten-4"
                        >
                            <img :src="'https://lorempixel.com/' + (180 + challenge.ID) + '/' + (180 + challenge.ID) + '/people/'" alt="avatar">
                        </v-avatar>
                    </v-flex>
                    <v-flex xs9 offset-xs1>
                        <h3 class="mb-0">{{ donator.Name }}</h3>
                        <div>{{ donator.Address }}</div>
                    </v-flex>
                </v-card-title>
            </v-card>
        </v-flex>
    </div>
</template>

<script>
    import { getDonator } from '../../api/api';


    export default {
        name: 'ChallengeItem',
        props: {
            challenge: { type: Object, required: true },
        },
        data() {
            return {
                donator: { type: Object, required: true },
                userId: window.localStorage.getItem('userId')
            }
        },
        mounted() {
            getDonator(this.challenge.DonatorID).then(response => {
                this.donator = response.data;
            });
        },
        computed: {
            applied() {
            	if (this.$route.path === '/your-challenges') {
            		return true;
                }
                if (this.challenge.Applications === null) {
            		return false;
                }
                return (this.challenge.Applications.filter((application) => {
                    return application.TechfugeeID === parseInt(this.userId);
                }).length > 0);
            },

            accepted() {
				if (this.challenge.Applications === null) {
					return false;
				}
                return (this.challenge.Applications.filter((application) => {
                    return (application.TechfugeeID === parseInt(this.userId)) && application.Accepted === true;
                }).length > 0);
            },
        }
    }
</script>

<style scoped lang="scss" text="text/scss">
    @import '../../../assets/_variables.scss';

    .card__title {
        position: relative;
    }

    .applied-icon, .check-icon {
        position: absolute;
        top: 50%;
        right: 16px;

        transform: translate(0, -50%);
    }
</style>