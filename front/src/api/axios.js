import axios from 'axios';
import { BASE_URL } from '../constants';
// import { saveSession, clearSession, getLocalAccessToken } from '../utils';
//
// let requestsToRefresh = [];
//
// let isRefreshRequesting = false;

let axiosInstance = axios.create({
  baseURL: BASE_URL + '/api/v1',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
});

// axiosInstance.interceptors.response.use(null, err => {
//   const { response, config } = err;
//   const currentToken = getLocalAccessToken();
//
//   if (response.status === 401) {
//     if (!currentToken) {
//       return Promise.reject(err);
//     }
//
//     if (!isRefreshRequesting) {
//       isRefreshRequesting = true;
//
//       axios
//         .post('/refresh')
//         .then(({ data }) => {
//           saveSession(data);
//           requestsToRefresh.forEach(cb => cb(data.content.token));
//         })
//         .catch(() => {
//           clearSession();
//           requestsToRefresh.forEach(cb => cb(null));
//         })
//         .finally(() => {
//           requestsToRefresh = [];
//           isRefreshRequesting = false;
//         });
//     }
//
//     return new Promise((resolve, reject) => {
//       requestsToRefresh.push(token => {
//         if (token) {
//           config.headers.Authorization = 'Bearer ' + token;
//           resolve(axios(config));
//         }
//
//         reject(Promise.reject(err));
//       });
//     });
//   }
//
//   return Promise.reject(err);
// });

export default axiosInstance;
