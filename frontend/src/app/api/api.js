import axios from 'axios';

export const HTTP = axios.create({
  baseURL: `https://www.donatugee.de/api/v1/`
});

export const createProfile = ({ name, email }) => {
	return HTTP.post('/refugees/create/', {
		name,
		email,
	}).then(response => response).catch(error => error.response);
};

export const getChallenges = () => {
  return HTTP.get('challenges').then(response => {
    return response;
  }).catch(e => {
    return e;
  })
};