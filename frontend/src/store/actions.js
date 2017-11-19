import {
	createCompanyProfile,
	getDonator,
	getRefugee,
	retrieveChallengesForDonator,
	createChallengeForDonator,
	doAcceptApplicant, createProfile,
} from '../app/api/api';

export const createCompany = ({state, commit}, data) => {
	return createCompanyProfile(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('COMPANY_CREATED', response.data);
		return response;
	});
};

export const getDonatorData = ({state, commit}, data) => {
	return getDonator(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('DONATOR_RETRIEVED', response.data);
		return response;
	});
};

export const getRefugeeData = ({state, commit}, data) => {
	return getRefugee(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('REFUGEE_RETRIEVED', response.data);
		return response;
	});
};

export const getChallengesForDonator = ({state, commit}, data) => {
	return retrieveChallengesForDonator(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('DONATOR_CHALLENGES_RETRIEVED', response.data);
		return response;
	});
};

export const createChallenge = ({state, commit}, data) => {
	return createChallengeForDonator(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('DONATOR_CHALLENGES_RETRIEVED', response.data);
		return response;
	});
};

export const acceptApplicant = ({state, commit}, data) => {
	return doAcceptApplicant(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		return response;
	});
};

export const doCreateProfile = ({state, commit}, data) => {
	return createProfile(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}

		commit('REFUGEE_PROFILE_CREATED', response.data);
		return response;
	});
};

export const logoutRefugee = ({state, commit}, data) => {
	commit('REFUGEE_LOGGED_OUT');
};

export const logoutCompany = ({state, commit}, data) => {
	commit('COMPANY_LOGGED_OUT');
};
