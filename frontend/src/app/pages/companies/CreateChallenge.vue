<template>
    <div>
        <v-flex xs12>
            <h1>
                Create your challenge
            </h1>
            <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
                {{ this.errorMessage }}
            </v-alert>
            <v-form v-model="valid">
                <v-text-field
                        label="Title of your challenge"
                        v-model="title"
                        required
                ></v-text-field>
                <v-text-field
                        label="Describe your challenge"
                        v-model="description"
                        multi-line
                ></v-text-field>
                <v-checkbox
                        label="Do you provide any hardware?"
                        v-model="hardwareProvided"
                        :rules="emailRules"
                        required
                ></v-checkbox>
                <v-text-field
                        name="couponValue"
                        label="What is the amount of the Udemy coupon"
                        v-model="couponValue"
                        type="number"
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
        </v-flex>
    </div>
</template>
<script>
	import { createCompanyProfile } from '../../api/api';
	import { mapActions, mapState } from 'vuex';

	export default {
		components: {VCheckbox},
		name: 'CreateChallenge',
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
				website: '',
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
					debugger;
					if (response.status !== 200) {
						this.errorMessage = 'There is already an account with this email address. Please try to login';
						return;
					}

					window.localStorage.setItem('companyId', response.data.ID);

					return this.$router.push({
						path: '/company/your-challenges',
					});
				});
			}
		}
	}
</script>

<style lang="scss" type="text/scss">
    .full-width {
        width: 100%;
        margin-bottom: 1rem;
    }
</style>
