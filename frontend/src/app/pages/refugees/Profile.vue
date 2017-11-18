<template>
    <div>
        <v-layout mb-4>
            <v-flex xs4>
                <v-avatar size="90px" class="grey lighten-4">
                    <img src="https://lorempixel.com/180/180/cats/" alt="avatar">
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
                <v-chip v-for="skill in JSON.parse(techfugee.Skills)" key="skill" color="indigo" text-color="white">
                    <v-icon left>code</v-icon>
                    {{ skill }}
                </v-chip>
                <v-divider class="mt-2 mb-4"></v-divider>
            </v-flex>
        </v-layout>

        <v-layout>
            <v-flex xs12>
                <v-btn color="primary">Edit profile</v-btn>
                <v-btn color="error" @click="logout">Log out</v-btn>
            </v-flex>
        </v-layout>
    </div>
</template>
<script>
	import {getTechfugee} from '../../api/api';

	export default {
		name: 'Profile',
		data() {
			return {
				id: this.$route.params.id,
				techfugee: {type: Object, required: true},
			};
		},
		mounted() {
			getTechfugee(this.id).then(response => {
				this.techfugee = response.data;
			});
		},
		methods: {
			logout() {
				window.localStorage.removeItem('idChallenge');
				window.localStorage.removeItem('userId');
				window.localStorage.removeItem('email');
				window.localStorage.removeItem('name');
				window.localStorage.removeItem('skills');
				window.localStorage.removeItem('wrongAnswers');
				window.localStorage.removeItem('city');
				window.localStorage.removeItem('introduction');
				window.localStorage.removeItem('authenticated');
				this.$router.push({
					path: '/',
				});
			},
		},
	};
</script>
