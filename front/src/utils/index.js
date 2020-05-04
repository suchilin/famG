export const log = (message, data = null) => {
  if (process.env.NODE_ENV == 'development') {
    console.log(message, data);
  }
};
export const saveSession = (/* { */
  // access_token,
  // refresh_token,
  // expires_at,
  // token_type
/* } */) => {
  // localStorage.setItem('access_token', access_token);
  // localStorage.setItem('refresh_token', refresh_token);
  // localStorage.setItem('expires_at', expires_at);
  // localStorage.setItem('token_type', token_type);
};

export const clearSession = () => {
  // localStorage.removeItem('access_token');
  // localStorage.removeItem('refresh_token');
  // localStorage.removeItem('expires_at');
  // localStorage.removeItem('token_type');
};

export const getLocalAccessToken = () => localStorage.getItem('access_token');

export const getLocalRefreshToken = () => localStorage.getItem('refresh_token');
