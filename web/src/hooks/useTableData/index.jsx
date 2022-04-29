import { ref } from 'vue';

export const useTableData = (getApi, search) => {
  const loading = ref(false);

  return { loading, getApi, search };
};
