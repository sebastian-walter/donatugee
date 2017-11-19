import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

import * as actions from './actions';
import mutations from './mutations';

const state = {
	company: null,
	companyChallenges: [],
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
