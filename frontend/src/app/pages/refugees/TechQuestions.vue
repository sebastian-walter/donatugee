<template>
    <div>
        <v-card>
            <v-card-title primary-title>
                <div>
                    <h3 class="headline mb-0">{{ techQuestions.question }}</h3>
                </div>
            </v-card-title>
            <ul class="tech-questions">
                <template v-for="(answer, index) in techQuestions.answers">
                    <li :class="['item', getClass(index)]"
                        @click="select(index)"
                    >
                        {{ answer }}
                    </li>
                </template>
            </ul>
            <v-divider></v-divider>
            <v-card-actions>
                <v-btn flat
                       color="primary"
                       @click="nextQuestion"
                >
                    Next question
                </v-btn>
            </v-card-actions>
        </v-card>
    </div>
</template>

<script>
	import {questions} from '../../../library/questions';

	export default {
		name: 'TechQuestions',
		data() {
			return {
				selectedAnswer: null,
			};
		},
		computed: {
			techQuestions() {
				return questions[ this.step - 1 ];
			},
			step() {
				return this.$route.params.step;
			},
		},
		methods: {
			select(index) {
				this.selectedAnswer = index;
			},
            getClass(index) {
				if (this.selectedAnswer === index) {
					return 'active';
                }
            },
			nextQuestion() {
				this.evaluate();
			},
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
        }
    }
</style>