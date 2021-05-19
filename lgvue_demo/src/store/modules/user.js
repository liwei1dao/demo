import { loginbycaptcha, loginbypassword, getuserinfo, loginout } from '@/api/user'
import { paramsign } from '@/utils/md5'
import { setToken, getToken, removeToken } from '@/utils/auth'
import { resetRouter } from '@/router'

const state = {
  token: getToken(),
  userinfo: null
}

const mutations = {
  Set_info: (state, userinfo) => {
    if (userinfo != null) {
      state.userinfo = userinfo
      state.token = userinfo.Token
      setToken(userinfo.Token)
    } else {
      state.userinfo = null
      state.token = null
      removeToken()
    }
  },
}

const actions = {
  loginbycaptcha ({ commit }, form) {
    console.log("loginbycaptcha:" + form)
    return new Promise((resolve, reject) => {
      loginbycaptcha(paramsign(form)).then(response => {
        const { data } = response
        commit("Set_info", data)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },
  loginbypassword ({ commit }, form) {
    console.log("loginbypassword req:" + form)
    return new Promise((resolve, reject) => {
      loginbypassword(paramsign(form)).then(response => {
        const { data } = response
        commit("Set_info", data)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  getuserinfo ({ commit }) {
    return new Promise((resolve, reject) => {
      console.log("loginbytoken:")
      getuserinfo().then(response => {
        const { data } = response
        commit("Set_info", data)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  loginout ({ commit }) {
    return new Promise((resolve, reject) => {
      loginout().then(() => {
        commit('Set_info', null)
        resetRouter()
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  resetToken ({ commit }) {
    return new Promise(resolve => {
      commit('Set_info', null)
      resetRouter()
      resolve()
    })
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}