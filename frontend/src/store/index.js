import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

import * as actions from './actions';
import mutations from './mutations';

const isLoggedIn = () => {
	const refugeeId = window.localStorage.getItem('userId');
	const companyId = window.localStorage.getItem('companyId');
	if ((refugeeId !== null && typeof refugeeId !== 'undefined') ||
		(companyId !== null && companyId !== 'undefined')) {
		return true;
	}
	return false;
};

const isRefugee = () => {
	const refugeeId = window.localStorage.getItem('userId');
	if (refugeeId !== null && typeof refugeeId !== 'undefined') {
		return true;
	}
	return false;
};

const state = {
	company: null,
	companyChallenges: [],
	isLoggedIn: isLoggedIn(),
	isRefugee: isRefugee(),
};

const store = {
	state,
	actions,
	mutations,
};

export default new Vuex.Store({
	...store,
	strict: debug,
});
