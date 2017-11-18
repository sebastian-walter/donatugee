<template>
    <div>
        <v-alert v-if="errorMessage !== ''" color="error" icon="warning" value="true">
            {{ this.errorMessage }}
        </v-alert>
        <v-alert v-if="successMessage !== ''" color="success" icon="check" value="true">
            {{ this.successMessage }}
        </v-alert>
        <v-card>
            <v-card-title primary-title>
                <div>
                    <h3 class="headline mb-0">{{ techQuestions.question }}</h3>
                </div>
            </v-card-title>
            <ul class="tech-questions">
                <li v-for="(answer, index) in techQuestions.answers" :key="index"
                    :class="['item', getClass(index)]"
                    @click="select(index)"
                >
                    {{ answer }}
                </li>
            </ul>
            <v-divider></v-divider>
            <v-card-actions>
                <template v-if="noMoreQuestions && !isAnswered">
                    <v-btn flat
                           color="primary"
                           @click="evaluate"
                           v-if="!isAnswered"
                    >
                        Submit Answer
                    </v-btn>
                    <v-btn flat
                           color="primary"
                           @click="nextQuestion"
                           v-if="isAnswered"
                    >
                        Next question
                    </v-btn>

                </template>
                <v-btn flat
                       color="primary"
                       @click="nextSection"
                       v-else
                >
                    Continue
                </v-btn>
            </v-card-actions>
        </v-card>
    </div>
</template>

<script>
	import {questions, maxNumberOfIncorrectAnswers} from '../../../library/questions';

	export default {
		name: 'TechQuestions',
		data() {
			return {
				selectedAnswer: null,
				successMessage: '',
				errorMessage: '',
                isAnswered: false,
			};
		},
		computed: {
			techQuestions() {
				return questions[ this.step - 1 ];
			},
			step() {
				return this.$route.params.step;
			},
            noMoreQuestions() {
				return this.steps === this.techQuestions.length
            },
		},
		methods: {
			select(index) {
				this.selectedAnswer = index;
			},
			getClass(index) {
				let classes = [];
				if (this.isAnswered && this.selectedAnswer === index) {
					classes.push('disabled');
					return classes;
                }
				if (this.selectedAnswer === index) {
					classes.push('active');
				}
				return classes;
			},
			evaluate() {
				this.isAnswered = true;
				debugger;

				if (this.selectedAnswer !== this.techQuestions.correctAnswerIndex) {
					let wrongAnswers = parseInt(window.localStorage.getItem('wrongAnswers'));
					window.localStorage.setItem('wrongAnswers', wrongAnswers + 1);
					this.errorMessage = 'Ooooops, that answer was not correct.';
					if (wrongAnswers > maxNumberOfIncorrectAnswers) {
						this.$router.push({
							path: '/rejected',
						});
					}
					return;
				}

				this.successMessage = 'Congrats, that was correct.';
			},
            nextQuestion() {
				let nextStep = parseInt(this.step) + 1;

				this.successMessage = '';
				this.errorMessage = '';
				this.selectedAnswer = null;
				this.isAnswered = false;

				this.$router.push({
					params: {
						step: nextStep,
					},
				});
            },
            nextSection() {
				this.$router.push({
                    path: '/further-details',
                })
            }
		},
	};
</script>

<style lang="scss" type="text/scss">
    @import '../../../assets/variables';

    .tech-questions {
        list-style: none;
        li {
            padding: 1rem 1.5rem;

            &.active {
                color: white;
                background-color: $primary-color;
            }
            &.disabled {
                color: white;
                background-color: $grey-light;
            }
        }
    }
</style>