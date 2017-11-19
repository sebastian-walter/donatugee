<template>
    <v-list-tile v-if="techfugee !== null" avatar>
        <v-list-tile-avatar>
            <img :src="'http://lorempixel.com/' + 80 + Math.round((20* Math.random())) + '/80/people/'">
        </v-list-tile-avatar>
        <v-list-tile-content @click="goToRefugeeProfile">
            <v-list-tile-title>
                {{ techfugee.Name }}
            </v-list-tile-title>
            <v-list-tile-sub-title>
                {{ techfugee.Introduction.substring(0, 30) + '...' }}
            </v-list-tile-sub-title>
        </v-list-tile-content>
        <v-list-tile-action>
            <v-btn icon ripple @click="acceptApplication">
                <v-icon v-if="!accepted && !application.Accepted"
                        color="grey lighten-1">check</v-icon>
                <v-icon v-else color="success">check_circle</v-icon>
            </v-btn>
        </v-list-tile-action>
    </v-list-tile>
</template>

<script>
    import { getRefugee } from '../api/api';
    import { mapActions } from 'vuex';

    export default {
    	name: 'Applicant',
        props: {
    		application: {
    			type: Object,
            }
        },
        data() {
    		return {
    			techfugee: null,
                accepted: false,
            }
        },
        methods: {
            ...mapActions([
            	'acceptApplicant',
            ]),
			goToRefugeeProfile() {
			    this.$router.push({
                    path: '/refugee/profile/' + this.techfugee.ID,
                })
            },
			acceptApplication() {
			    this.acceptApplicant({ id: this.application.ID }).then(response => {
			    	if (response.status !== 200) {
			    		this.errorMessage = 'OOooops, sorry!';
						return;
                    }

                    this.accepted = true;
                })
            }
        },
        mounted() {
				getRefugee({id: this.application.TechfugeeID}).then(response => {
					this.techfugee = response.data;
				});
        }
    }
</script>
