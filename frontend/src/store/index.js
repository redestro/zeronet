import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    token: '',
    gameMode: '',
  },
  mutations: {
    updateToken(state, token) {
      state.token = token;
    },
    updateGameMode(state, gameMode) {
      state.gameMode = gameMode;
    },
  },
  getters: {
    token: state => state.token,
    gameMode: state => state.gameMode,
  },
  modules: {
  },
});
