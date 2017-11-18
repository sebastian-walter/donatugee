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
  return HTTPS.get('challenges').then(response => response).catch(error => error.response);
};

export const getChallenge = (id) => {
	return HTTPS.get('challenge', {
		params: {
			id
		}
	}).then(response => response).catch(error => error.response);
};

export const getTechfugee = (id) => {
	return HTTPS.get('techfugee', {
		params: {
			id
		}
	}).then(response => response).catch(error => error.response);
};

export const saveFurtherDetails = ({ id, city, introduction }) => {
	return HTTPS.get('update-techfugee', {
		params: {
			id,
			city,
			introduction
		}
	}).then(response => response).catch(error => error.response);
};


export const getDonator = (id) => {
	return HTTPS.get('donator', {
		params: {
			id,
		}
	}).then(response => response).catch(error => error.response);
};


export const getRandomText = (length) => {
	return axios.get('http://www.randomtext.me/api/gibberish/p-1/' + length).then(response => response).catch(error => error.response);
};

export const techfugeeAuthenticated = ({ id, passed }) => {
	return HTTPS.get('update-auth', {
		params: {
			id: id,
			passed: passed
		}
	}).then(response => response).catch(error => error.response)
};
