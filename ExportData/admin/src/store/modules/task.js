import { get, add } from '@/api/task'
import { getToken, setToken, removeToken, getTokenExpire, setTokenExpire, removeTokenExpire } from '@/utils/auth'
import { resetRouter } from '@/router'

const state = {
  token: getToken(),
  name: '',
  avatar: '',
  tokenExpire: getTokenExpire()
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_TOKENEXPIRE: (state, token) => {
    state.tokenExpire = token
  }
}

const actions = {
  get({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      get({ username: username.trim(), password: password }).then(response => {
        const data = response
        commit('SET_TOKEN', data.token)
        commit('SET_TOKENEXPIRE', data.expire)
        setToken(data.token)
        setTokenExpire(data.expire)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  add({ commit }, taskInfo) {
    const { username, password, confirmPassword } = userInfo
    return new Promise((resolve, reject) => {
      add({ username: username.trim(), password: password, confirmPassword: confirmPassword }).then(response => {
        const data = response
        commit('SET_TOKEN', data.token)
        commit('SET_TOKENEXPIRE', data.expire)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

