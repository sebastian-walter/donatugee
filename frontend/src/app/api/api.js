import axios from 'axios';

export const HTTPS = axios.create({
  baseURL: `https://www.donatugee.de/api/v1/`,
});

export const createProfile = ({ name, email }) => {
	return HTTPS.get('insert-techfugee', {
		params: {
			name,
			email,
			skills: {},
		}
	}).then(response => response).catch(error => error.response);
};

export const addSkills = ({ skills, id }) => {
	return HTTPS.get('add-skills', {
		params: {
			skills: JSON.stringify(skills),
			id,
		}
	}).then(response => response).catch(error => error.response);
};

export const getChallenges = () => {
  return HTTP.get('challenges').then(response => {
    return response;
  }).catch(e => {
    return e;
  })
};
