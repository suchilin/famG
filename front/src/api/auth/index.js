import axios from '../axios';
import { log, saveSession } from '../../utils';

export const login = async (username, password) => {
  try {
    const { data } = await axios.post('/auth/login', {
      username,
      password
    });
    saveSession(data);
  } catch (err) {
    log('ERROR ON LOGIN', err.response);
    throw new Error(err.response.data.message);
  }
};

export const logout = async ()=>{
  return axios.get('/auth/logout')
}

export const isAuthenticated = async ()=>{
  return axios.get('/auth/isauth')
}
