import {
	createCompanyProfile,
	getDonator,
	getRefugee,
	retrieveChallengesForDonator
} from '../app/api/api';

export const createCompany = ({state, commit}, data) => {
	return createCompanyProfile(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('COMPANY_CREATED', response.data);
		return response;
	})
};

export const getDonatorData = ({state, commit}, data) => {
	return getDonator(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('DONATOR_RETRIEVED', response.data);
		return response;
	})
};

export const getRefugeeData = ({state, commit}, data) => {
	return getRefugee(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('REFUGEE_RETRIEVED', response.data);
		return response;
	})
};

export const getChallengesForDonator = ({state, commit}, data) => {
	return retrieveChallengesForDonator(data).then((response) => {
		if (response.status !== 200) {
			return response;
		}
		commit('DONATOR_CHALLENGES_RETRIEVED', response.data);
		return response;
	})
};
