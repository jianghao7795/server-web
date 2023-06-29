import { isReactive, isRef, reactive } from 'vue';
import type { Ref } from 'vue';

function setLoading(loading?: Ref | { loading?: boolean }, val?: boolean) {
  if (loading != undefined && isRef(loading)) {
    loading.value = val;
  } else if (loading != undefined && isReactive(loading)) {
    loading.loading = val;
  }
}

export const useAsync = async (func: Promise<any>, loading: any): Promise<any> => {
  setLoading(loading, true);

  return await func.finally(() => setLoading(loading, false));
};
