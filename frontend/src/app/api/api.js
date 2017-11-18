import axios from 'axios';

export const HTTP = axios.create({
  baseURL: `http://donatugee.de/`
});

export const getChallenges = () => {
  return HTTP.get('/challenges').then(response => {
    return response;
  }).catch(e => {
    return e;
  })
};