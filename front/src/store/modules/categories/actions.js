import { log } from '../../../utils';

export const addCategory = ({ commit }, category) => {
  try {
    commit('onGetCategorySuccess', category);
  } catch (err) {
    log('ERROR ON ADD CATEGORY', err);
  }
};
