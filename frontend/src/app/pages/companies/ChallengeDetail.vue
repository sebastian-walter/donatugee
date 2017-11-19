<template>
    <div>
        <h1 class="mb-2">Challenge {{ challenge.Name }}</h1>
        <h3 class="mb-2">powered by you</h3>

        <p class="mb-4">{{ challenge.Description }}</p>

        <v-card class="mb-4">
            <v-card-title>
                <h3 class="title mb-2 display-1">What you provide</h3>
            </v-card-title>
            <v-divider></v-divider>
            <v-card-text>
                <div>
                    <v-chip v-if="challenge.HardwareProvided === 'true'" color="orange" text-color="white">
                        <v-icon left>laptop</v-icon>
                        {{ challenge.LaptopType }}
                    </v-chip>
                    <v-chip color="indigo" text-color="white">
                        <v-icon left>code</v-icon>
                        {{ challenge.Amount }} € Udemy Coupon
                    </v-chip>
                </div>
            </v-card-text>
            <v-list two-line>
                <v-subheader v-if="challenge.Applications !== null && challenge.Applications.length > 0">
                    Applicants
                </v-subheader>
                <template v-for="application in challenge.Applications">
                    <v-divider  v-bind:inset="true"></v-divider>
                    <applicant :application="application"></applicant>
                </template>
            </v-list>
        </v-card>
    </div>
</template>
<script>
	import Applicant from '../../components/Applicant.vue';
    import { getChallenge, getDonator, getRandomText, setApplication, getRefugee } from '../../api/api';
	import { mapState, mapActions } from 'vuex';

	export default {
		name: 'ChallengeDetail',
        components: {
			Applicant,
        },
		data() {
			return {
				idChallenge: this.$route.params.id,
				challenge: { Applications: [] },
				randomText: '',
				hasApplied: false,
				userId: window.localStorage.getItem('userId'),
                loading: true,
			}
		},
		mounted() {
			getChallenge(this.idChallenge).then(response => {
				this.challenge = response.data;
				this.loading = false;
			});

			getRandomText('10-30').then(response => {
				this.randomText = response.data;
			});
		},
		computed: {
            ...mapState({
                donator: state => state.donator,
            }),
			applied() {
				if (this.hasApplied) {
					return true;
				}

				return this.challenge.Applications.filter((application) => {
					return application.TechfugeeID === parseInt(this.userId);
				}).length > 0;
			}
		}
	}
</script>