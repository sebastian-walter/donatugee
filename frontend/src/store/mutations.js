export default {
	COMPANY_CREATED(state, payload) {
		state.donator = payload;
	},
	DONATOR_RETRIEVED(state, payload) {
		state.donator = payload;
	},
	REFUGEE_RETRIEVED(state, payload) {
		state.refugee = payload;
	},
	DONATOR_CHALLENGES_RETRIEVED(state, payload) {
		state.companyChallenges = payload;
	}
};
