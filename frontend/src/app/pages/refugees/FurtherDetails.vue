<template>
    <div>
        <v-flex xs12>
            <h1>
                Tell us something more about you
            </h1>
            <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
                {{ this.errorMessage }}
            </v-alert>
            <v-form>
                <v-text-field
                        label="In which city are you at the moment?"
                        v-model="city"
                        required
                ></v-text-field>
                <v-text-field
                        name="description"
                        label="Tell us a bit about yourself"
                        v-model="description"
                        multi-line
                ></v-text-field>
                <v-btn class="sign-up-button"
                       small
                       color="primary"
                       dark
                       large
                       @click="save"
                >
                    Save and continue
                </v-btn>
            </v-form>
        </v-flex>
    </div>
</template>
<script>
    import { saveFurtherDetails } from '../../api/api';

    export default {
    	name: 'FurtherDetails',
        data() {
    		return {
    			city: '',
                description: '',
                errorMessage: '',
            }
        },
        methods: {
    		save() {
    			saveFurtherDetails({
                    id: window.localStorage.getItem('userId'),
                    city: this.city,
                    introduction: this.description
    			}).then(response => {
    				if (response.status === 200) {
    					debugger;
    					let idChallenge = window.localStorage.getItem('idChallenge');
    					if (idChallenge !== null) {
							this.$router.push({
								path: '/challenge/' + idChallenge,
							});
                            return;
                        }
						this.$router.push({
							path: 'challenges',
						});
    					return;
                    }
                    this.errorMessage = 'Something went wrong, please try again.';
                })
            }
        }
    }
</script>