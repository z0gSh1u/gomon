import { ref, shallowRef } from 'vue'
import _axios, { AxiosResponse, AxiosRequestConfig } from 'axios'

export interface UseAxiosOptions<T> {
  immediate?: boolean
  onSuccess?: (data: T) => void
  onError?: (error: unknown) => void
}

/**
 * Hook for Axios where <T> is type of data in response.
 */
export function useAxios<T>(
  url: string,
  config?: AxiosRequestConfig,
  options?: UseAxiosOptions<T>
) {
  const axios = _axios
  const response = shallowRef<AxiosResponse<T>>()
  const data = ref<T>()
  const error = ref<unknown>()
  const isFinished = ref(false)

  const run = () =>
    axios(url, config)
      .then((r) => {
        response.value = r
        const _data = r.data // hold
        data.value = _data
        options?.onSuccess?.(_data)
      })
      .catch((e: unknown) => {
        error.value = e
        options?.onError?.(e)
      })
      .finally(() => {
        isFinished.value = true
      })

  return {
    run,
    response,
    data,
    error,
    isFinished,
  }
}
