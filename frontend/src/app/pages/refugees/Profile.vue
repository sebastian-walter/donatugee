<template>
    <div>
        <v-layout mb-4>
            <v-flex xs4>
                <v-avatar size="90px" class="grey lighten-4">
                    <img :src="'https://lorempixel.com/180/180/people/' + (techfugee.ID%10)" alt="avatar">
                </v-avatar>
            </v-flex>

            <v-flex xs8>
                <h1 class="mb-0">{{ techfugee.Name }}</h1>
                <div>{{ techfugee.City }}</div>
            </v-flex>
        </v-layout>

        <v-layout>
            <v-flex xs12>
                <p>{{ techfugee.Introduction }}</p>
                <v-chip v-for="skill in skills" key="skill" color="indigo" text-color="white">
                    <v-icon left>code</v-icon>
                    {{ skill }}
                </v-chip>
                <v-divider class="mt-2 mb-4"></v-divider>
            </v-flex>
        </v-layout>

        <v-layout v-if="isRefugee">
            <v-flex xs12>
                <v-btn color="primary">Edit profile</v-btn>
                <v-btn color="error" @click="logout">Log out</v-btn>
            </v-flex>
        </v-layout>
    </div>
</template>
<script>
	import {getTechfugee} from '../../api/api';
	import { mapActions } from 'vuex';

	export default {
		name: 'Profile',
		data() {
			return {
				id: this.$route.params.id,
				techfugee: {type: Object, required: true},
			};
		},
		computed: {
			skills() {
				if (this.techfugee.Skills === '' || typeof this.techfugee.Skills === 'undefined' || this.techfugee === null) {
					return [];
				}
				return JSON.parse(this.techfugee.Skills);
			},
			isRefugee() {
				return window.localStorage.getItem('userId') !== null;
            }
		},
		mounted() {
			getTechfugee(this.id).then(response => {
				this.techfugee = response.data;
			});
		},
		methods: {
            ...mapActions([
            	'logoutRefugee',
            ]),
			logout() {
				window.localStorage.removeItem('idChallenge');
				window.localStorage.removeItem('userId');
				window.localStorage.removeItem('wrongAnswers');
				window.localStorage.removeItem('authenticated');
				this.logoutRefugee();
				this.$router.push({
					path: '/',
				});

			},
		},
	};
</script>
