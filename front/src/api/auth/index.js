import axios from '../axios';

export const login = async (username, password) => {
  return axios.post('/auth/login', {
    username,
    password,
  });
};

export const logout = async () => {
  return axios.get('/auth/logout');
};

export const isAuthenticated = async () => {
  return axios.get('/auth/isauth');
};
