import { login, logout } from '../../../api/auth';
import router from '../../../router';
import { getLocalAccessToken, log } from '../../../utils';
export const doLogin = async ({ commit }, credentials) => {
  commit('authFetching', true);
  try {
    const { username, password } = credentials;
    await login(username, password);
    commit('onAuthSuccess');
    commit('authFetching', false);
    router.push('/home');
  } catch (error) {
    commit('onAuthFailure', error);
    commit('authFetching', false);
  }
};

export const fetchLocalAccessToken = ({ commit }) => {
  const token = getLocalAccessToken();
  if (token) {
    commit('onLocalLogin');
  }
};

export const doLogout = async ({ commit }) => {
  try{
    await logout()
    commit('onLogout');
    router.push('/login');
  }catch(err){
    log('ERROR ON LOGOUT', err)
  }
}
