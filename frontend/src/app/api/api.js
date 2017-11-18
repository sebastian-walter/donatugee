import axios from 'axios';

export const createProfile = ({ name, email }) => {
	return HTTP.post('/refugees/create/', {
		name,
		email,
	}).then(response => response).catch(error => error.response);
};