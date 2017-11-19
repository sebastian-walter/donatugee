import { createCompanyProfile } from '../app/api/api';

export const createCompany = ({state, commit}, data) => {
	return createCompanyProfile(data).then((response) => {
		debugger;
		commit('COMPANY_CREATED', response.data);
	})
};
