<template>
    <div>
        <v-layout mb-4>
            <v-flex xs4>
                <v-avatar size="90px" class="grey lighten-4">
                    <img src="https://lorempixel.com/180/180/cats/" alt="avatar">
                </v-avatar>
            </v-flex>

            <v-flex xs8>
                <h1 class="mb-0">{{ donator.Name }}</h1>
                <div>{{ donator.City }}</div>
            </v-flex>
        </v-layout>

        <v-layout>
            <v-flex xs12>
                <p>{{ donator.Introduction }}</p>
                <h2>{{ donator.Email }}</h2>
            </v-flex>
        </v-layout>

        <v-layout>
            <v-flex xs12>
                <v-btn color="primary">Edit profile</v-btn>
                <v-btn color="error">Log out</v-btn>
            </v-flex>
        </v-layout>
    </div>
</template>
<script>
	import { getDonator } from '../../api/api';

	export default {
		name: 'Profile',
        data() {
			return {
				id: this.$route.params.id,
                donator: {},
            }
        },
		mounted() {
            getDonator(this.id).then(response => {
            	this.donator = response.data;
            })
		},
	};
</script>
