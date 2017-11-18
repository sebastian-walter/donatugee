<template>
    <div>
        <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
            {{ this.errorMessage }}
        </v-alert>
        <v-card>
            <v-card-title primary-title>
                <div>
                    <h3 class="headline mb-0">Tell us what skills you have already</h3>
                </div>
            </v-card-title>
            <v-list>
                <v-list-tile v-for="skill in skills" :key="skill.name">
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
                <v-btn flat color="orange" @click="saveSkills">Save</v-btn>
            </v-card-actions>
        </v-card>
    </div>
</template>

<script>
	import {addSkills} from '../../api/api';

	export default {
		name: 'Interests',
		data() {
			return {
				errorMessage: '',
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
						name: 'Go',
						value: false,
					},
					{
						name: 'Ruby',
						value: false,
					},
					{
						name: 'C',
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
			saveSkills() {
				const skills = this.skills.slice().filter(skill => skill.value).map(skill => skill.name);
				addSkills({
					skills: skills,
					id: this.$route.params.idUser,
				})
					.then(response => {
						if (response.status !== 200) {
							this.errorMessage = 'Skills could not be added';
						}

						window.localStorage.setItem('skills', response.data.Skills);
						const authenticated = JSON.parse(window.localStorage.getItem('authenticated'));

						if (!authenticated) {
							return this.$router.push({
								path: '/tech-questions/1',
							});
						}

						if (window.localStorage.getItem('introduction') === '' ||
                            window.localStorage.getItem('city') === '') {
							return this.$router.push({
								path: 'refugee/further-actions',
							});
						}

						const idChallenge = window.localStorage.getItem('idChallenge');
						if (idChallenge === null) {
							this.$router.push({
								path: '/challenges',
							});
						}

						return this.$router.push({
							path: '/challenge/' + idChallenge,
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