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
                        label="Name"
                        v-model="name"
                        :rules="nameRules"
                        required
                ></v-text-field>
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
                       @click="signUp"
                >
                    Sign up
                </v-btn>
            </v-form>
        </v-flex>
    </div>
</template>
<script>
    import { createProfile } from '../../api/api';

	export default {
		name: 'Register',
		data () {
			return {
				valid: false,
				name: '',
				nameRules: [
					(v) => !!v || 'Name is required',
				],
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
			signUp() {
                createProfile({
                    name: this.name,
                    email: this.email,
                    password: this.password
                }).then(response => {
                    	if (response.status !== 200) {
                    		this.errorMessage = 'User cannot be created';
                        }

                        window.localStorage.setItem('userId', response.data.ID);
						window.localStorage.setItem('email', response.data.Email);
						window.localStorage.setItem('name', response.data.Name);
						window.localStorage.setItem('skills', response.data.Skills);
						window.localStorage.setItem('wrongAnswers', 0);

                        return this.$router.push({
                            path: '/interests/add/' + response.data.ID,
                        });
                    });
            }
        }
	};
</script>

<style lang="scss" type="text/scss">
    .sign-up-button {
        width: 100%;
    }
</style>