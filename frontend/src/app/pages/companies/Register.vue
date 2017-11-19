<template>
    <div>
        <v-flex xs12>
            <h1>
                Sign up
            </h1>
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
                <v-btn class="sign-up-button"
                       small
                       color="primary"
                       dark
                       large
                       @click="signUp"
                >
                    Sign up
                </v-btn>
            </v-form>
        </v-flex>
    </div>
</template>
<script>
    import { createCompanyProfile } from '../../api/api';
    import { mapActions, mapState } from 'vuex';

	export default {
		name: 'Register',
		data () {
			return {
				valid: false,
				companyName: '',
				nameRules: [
					(v) => !!v || 'Name is required',
				],
				email: '',
				emailRules: [
					(v) => !!v || 'E-mail is required',
					(v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
				],
				errorMessage: '',
                company: '',
                address: '',
                websiter: '',
			}
		},
		methods: {
            ...mapActions([
            	'createCompany',
            ]),
            signUp() {
				this.createCompany({
					name: this.name,
					email: this.email,
					password: this.password,
                    website: this.website,
                    address: this.website,
				}).then(response => {
					if (response.status !== 200) {
						this.errorMessage = 'User cannot be created';
					}

					window.localStorage.setItem('companyId', response.data.ID);

					return this.$router.push({
						path: '/interests/add/' + response.data.ID,
					});
				});
            }
		}
	}
</script>
