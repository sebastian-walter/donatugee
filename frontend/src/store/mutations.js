export default {
	COMPANY_CREATED(state, payload) {
		state.donator = payload;
		state.isLoggedIn = true;
		state.isRefugee = false;
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
		state.isLoggedIn = true;
		state.refugee = payload;
		state.isRefugee = true;
	},
	REFUGEE_LOGGED_OUT(state, payload) {
		state.isLoggedIn = false;
		state.isRefugee = false;
	},
	COMPANY_LOGGED_OUT(state, payload) {
		state.isLoggedIn = false;
	},
	COMPANY_LOGGED_IN(state, payload) {
		state.isLoggedIn =  true;
		state.isRefugee = false;
	}
};
