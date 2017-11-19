<template>
    <div>
        <v-flex xs12 sm6 offset-sm3>
            <v-card>
                <v-card-title>
                    <h1>
                        Create your challenge
                    </h1>
                </v-card-title>
                <v-card-text>
                    <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
                        {{ this.errorMessage }}
                    </v-alert>
                    <v-form v-model="valid">
                        <v-text-field
                                label="Title of your challenge"
                                v-model="name"
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
                                required
                        ></v-checkbox>
                        <v-text-field
                                v-if="hardwareProvided"
                                name="laptopType"
                                label="What kind of hardware do you provide?"
                                v-model="laptopType"
                                required
                        ></v-text-field>
                        <v-select
                                :items="options"
                                v-model="duration"
                                label="How long should the challenge last?"
                                single-line
                                bottom
                        ></v-select>
                        <v-text-field
                                name="couponValue"
                                label="What is the amount of the Udemy coupon"
                                v-model="amount"
                                type="number"
                                required
                        ></v-text-field>
                    </v-form>
                       <v-divider>

                       </v-divider>
                        <v-btn class="full-width"
                               small
                               color="primary"
                               dark
                               large
                               @click="create"
                        >
                            Create
                        </v-btn>
                </v-card-text>
            </v-card>
        </v-flex>
    </div>
</template>
<script>
	import {createCompanyProfile} from '../../api/api';
	import {mapActions, mapState} from 'vuex';

	export default {
		name: 'CreateChallenge',
		data() {
			return {
				valid: false,
				name: '',
				hardwareProvided: false,
				laptopType: '',
                duration: '',
                amount: '',
                description: '',
				errorMessage: '',
                options: [
                    {
                    	text: '1 month',
                        value: '1 month',
                    },
					{
						text: '2 months',
						value: '2 months',
					},
					{
						text: '3 months',
						value: '3 months',
					},
					{
						text: '4 months',
						value: '4 months',
					},
                ]
			};
		},
		methods: {
			...mapActions([
				'createChallenge',
			]),
			create() {
				this.createChallenge({
                    idDonator: window.localStorage.getItem('companyId'),
					name: this.name,
                    description: this.description,
					hardwareProvided: this.hardwareProvided,
					laptopType: this.laptopType,
					duration: this.duration,
                    amount: this.amount
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
