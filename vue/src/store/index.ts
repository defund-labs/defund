import { createStore } from 'vuex'

import init from './config'

const store = createStore({
  state: {
      stakePopup: false
  },
  mutations: {
    setStakePopup(state, value) {
      state.stakePopup = value
    }
  },
  getters: {
    getStakePopup: (state) => {
      return state.stakePopup
    }
  },
  actions: {},
})
init(store)
export default store
