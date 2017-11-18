<template>
    <v-card>
        <v-card-title primary-title>
            <div>
                <h3 class="headline mb-0">Tell us what skills you have already</h3>
            </div>
        </v-card-title>
        <v-list>
            <v-list-tile v-for="skill in skills">
                <v-list-tile-action>
                    <v-checkbox
                            v-model="skill.value"
                    ></v-checkbox>
                </v-list-tile-action>
                <v-list-tile-content>
                    <v-list-tile-title>{{ skill.name }}</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
        </v-list>
        <v-card-actions>
            <v-btn flat color="orange">Save</v-btn>
        </v-card-actions>
    </v-card>
</template>
<script>
	import {createProfile} from '../../api/api';

	export default {
		name: 'Interests',
		data() {
			return {
				skills: [
					{
						name: 'Javascript',
						value: false,
					},
					{
						name: 'PHP',
						value: false,
					},
					{
						name: 'UX/Design',
						value: false,
					},
				],
				video: '',
			};
		},
		computed: {
			disabled() {
				return !this.valid;
			},
		},
		methods: {
			signUp() {
				createProfile({name: this.name, email: this.email})
					.then(response => {
						if (response.status !== 200) {
							this.errorMessage = 'User cannot be created';
						}

						return this.$router.push({
							path: '/authentication',
						});
					});
			},
		},
	};
</script>

<style scoped lang="scss" type="text/scss">
    .interests {
        ul {
            li {
                padding: 0;
            }
        }
    }
</style>