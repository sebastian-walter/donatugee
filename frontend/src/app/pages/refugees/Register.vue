<template>
    <div>
        <v-flex xs12>
            <h1>
                Sign up
            </h1>
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
                >
                    Sign up
                </v-btn>
            </v-form>
        </v-flex>
    </div>
</template>
<script>
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
				]
			}
		},
        methods: {
			signUp() {

            }
        }
	};
</script>
<style lang="scss" type="text/scss">
    .sign-up-button {
        width: 100%;
    }
</style>