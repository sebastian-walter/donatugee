<template>
    <div>
        <v-flex xs12>
            <h1>
                Log In
            </h1>
            <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
                {{ this.errorMessage }}
            </v-alert>
            <v-form v-model="valid">
                <v-text-field
                        label="E-mail"
                        v-model="email"
                        :rules="emailRules"
                        required
                ></v-text-field>
                <v-text-field
                        label="Password"
                        v-model="password"
                        :rules="passwordRules"
                        type="password"
                        required
                ></v-text-field>
                <v-btn class="sign-up-button"
                       small
                       color="primary"
                       dark
                       large
                       @click="login"
                >
                    Log In
                </v-btn>
            </v-form>
        </v-flex>
    </div>
</template>
<script>
	import { loginDonator } from '../../api/api';

	export default {
		name: 'Login',
		data () {
			return {
				valid: false,
				email: '',
				emailRules: [
					(v) => !!v || 'E-mail is required',
					(v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
				],
				errorMessage: '',
				password: '',
				passwordRules: [
					(v) => !!v || 'Password is required',
				],
			}
		},
		computed: {
			disabled() {
				return !this.valid;
			}
		},
		methods: {
			login() {
				loginDonator({
					email: this.email,
				}).then(response => {
					if (response.status !== 200) {
						this.errorMessage = 'There is no such user';
						return
					}

					window.localStorage.setItem('companyId', response.data.ID);
					return this.$router.push({
						path: '/company/your-challenges',
					});
				})
			}
		}
	};
</script>

<style lang="scss" type="text/scss">
    .sign-up-button {
        width: 100%;
    }
</style>
