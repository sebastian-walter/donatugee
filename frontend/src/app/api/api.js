import axios from 'axios';

export const HTTP = axios.create({
  baseURL: `https://www.donatugee.de/api/v1/`
});

export const getChallenges = () => {
  return HTTP.get('challenges').then(response => {
    return response;
  }).catch(e => {
    return e;
  })
};