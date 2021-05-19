import { getLanguge, setLanguge } from '@/utils/auth'
import { addRoutes } from '@/router'

const state = {
  appname: "Lego",
  currlanguage: getLanguge() ? getLanguge() : 'zh',
  routers: null,
}

const mutations = {
  Set_currlanguage: (state, lang) => {
    state.currlanguage = lang
    setLanguge(lang)
  },
  Set_routers: (state, routers) => {
    state.routers = routers
  },
}

const actions = {
  setcurrlanguage ({ commit }, lang) {
    commit('Set_currlanguage', lang)
  },
  setrouters ({ commit }, rule) {
    const routers = addRoutes(rule)
    commit('Set_routers', routers)
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}