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
                <v-btn class="sign-up-button"
                       small
                       color="primary"
                       dark
                       large
                       @click="signUp"
                       :disabled="disabled"
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
					(v) => v.length <= 10 || 'Name must be less than 10 characters'
				],
				email: '',
				emailRules: [
					(v) => !!v || 'E-mail is required',
					(v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
				],
                errorMessage: '',
			}
		},
        methods: {
			signUp() {
                createProfile({ name: this.name, email: this.email })
                    .then(response => {
                    	if (response.status !== 200) {
                    		this.errorMessage = 'User cannot be created';
                        }

                        return this.$router.push({
                            path: '/authentication',
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