export const authFetching = (state, loading) => {
  state.loading = loading;
};

export const onAuthSuccess = (state) => {
  state.loading = false;
  state.authenticated = true;
  state.error = null;
};

export const onAuthFailure = (state, error) => {
  state.error = error;
};

export const onLocalLogin = (state) => {
  state.authenticated = true;
  state.error = null;
};

export const onLogout = (state) => {
  state.authenticated = false;
};
