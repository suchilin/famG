export const fetchingCategory = (state, loading) => {
  state.loading = loading;
};

export const onGetCategoriesSuccess = (state, categories) => {
  state.loading = false;
  state.data = categories;
  state.error = null;
};

export const onGetCategorySuccess = (state, category) => {
  state.loading = false;
  state.data.push(category);
  state.error = null;
};

export const onGetCategoryFailure = (state, error) => {
  state.error = error;
};
