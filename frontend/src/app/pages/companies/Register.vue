<template>
    <div>
        <v-flex xs12 sm6 offset-sm3>
            <v-card>
                <v-card-title>
                    <h1>
                        Sign up
                    </h1>
                </v-card-title>
                <v-card-text>
                    <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
                        {{ this.errorMessage }}
                    </v-alert>
                    <v-form v-model="valid">
                        <v-text-field
                                label="Company Name"
                                v-model="companyName"
                                :rules="nameRules"
                                required
                        ></v-text-field>
                        <v-text-field
                                label="Company website"
                                v-model="website"
                        ></v-text-field>
                        <v-text-field
                                label="E-mail"
                                v-model="email"
                                :rules="emailRules"
                                required
                        ></v-text-field>
                        <v-text-field
                                name="Address"
                                label="Company Address"
                                v-model="address"
                                required
                                multi-line
                        ></v-text-field>
                        <v-text-field
                                name="Password"
                                label="Password"
                                v-model="password"
                                type="password"
                                required
                        ></v-text-field>
                        <v-btn class="full-width"
                               small
                               color="primary"
                               dark
                               large
                               @click="signUp"
                        >
                            Sign up
                        </v-btn>
                        <v-btn class="full-width"
                               small
                               dark
                               large
                               @click="signUp"
                        >
                            Login
                        </v-btn>
                    </v-form>
                </v-card-text>
            </v-card>
        </v-flex>
    </div>
</template>
<script>
	import {createCompanyProfile} from '../../api/api';
	import {mapActions, mapState} from 'vuex';

	export default {
		name: 'Register',
		data() {
			return {
				valid: false,
				companyName: '',
				nameRules: [
					(v) => !!v || 'Name is required',
				],
				email: '',
				emailRules: [
					(v) => !!v || 'E-mail is required',
					(v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid',
				],
				errorMessage: '',
				company: '',
				address: '',
				website: '',
				password: '',
			};
		},
		methods: {
			...mapActions([
				'createCompany',
			]),
			signUp() {
				this.createCompany({
					name: this.companyName,
					email: this.email,
					password: this.password,
					website: this.website,
					address: this.address,
				}).then(response => {
					if (response.status !== 200) {
						this.errorMessage = 'There is already an account with this email address. Please try to login';
						return;
					}

					window.localStorage.setItem('companyId', response.data.ID);

					return this.$router.push({
						path: '/company/your-challenges',
					});
				});
			},
		},
	};
</script>

<style lang="scss" type="text/scss">
    .full-width {
        width: 100%;
        margin-bottom: 1rem;
    }
</style>
