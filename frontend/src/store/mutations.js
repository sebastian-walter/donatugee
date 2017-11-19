export default {
	COMPANY_CREATED(state, payload) {
		state.donator = payload;
		state.isLoggedIn = true;
	},
	DONATOR_RETRIEVED(state, payload) {
		state.donator = payload;
	},
	REFUGEE_RETRIEVED(state, payload) {
		state.refugee = payload;
	},
	DONATOR_CHALLENGES_RETRIEVED(state, payload) {
		state.companyChallenges = payload;
	},
	REFUGEE_PROFILE_CREATED(state, payload) {
		state.refugee = payload;
	}
};
