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
    log('ERROR ON LOGIN', err);
    throw new Error(err.message);
  }
};

export const logout = async ()=>{
  return axios.get('/auth/logout')
}
