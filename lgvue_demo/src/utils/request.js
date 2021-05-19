import axios from 'axios'
import store from '@/store'
import router from '@/router'
import Message from '@/components/message/'
import { paramsign } from '@/utils/md5'
import { getToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  request => {
    if (request.data != null) {
      console.log(request.data)
      request.data = paramsign(request.data)
    }
    if (store.getters.token != null) {
      request.headers['X-Token'] = store.getters.token
      console.log("X-Token" + request.headers['X-Token'])
    }
    return request
  },
  error => {
    console.log(error) //for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  response => {
    const res = response.data
    console.log('response:' + JSON.stringify(response.baseURL) + "|res:" + JSON.stringify(res))
    if (res.code !== 0) {
      if (res.code == 101) {
        async () => {
          await store.dispatch('user/resetToken')
          router.push({ path: '/' })
        }
      } else {
        Message.error(res.message || 'Error')
      }
      return Promise.reject(new Error(res.message || 'Error'))
    } else {
      return res
    }
  },
  error => {
    Message.error(error.message)
    return Promise.reject(error)
  }
)

export default service