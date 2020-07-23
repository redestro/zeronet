import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    token: '',
    gameMode: '',
    cells: {
      0: '',
      1: '',
      2: '',
      3: '',
      4: '',
      5: '',
      6: '',
      7: '',
      8: ''
    }
  },
  mutations: {
    updateToken(state, token) {
      state.token = token;
    },
    updateGameMode(state, gameMode) {
      state.gameMode = gameMode;
    },
    updateCells(state, payload) {
      state.cells[payload.cellNumbers] = payload.activePlayer;
    }
  },
  getters: {
    token: state => state.token,
    gameMode: state => state.gameMode,
    cells: state => state.cells
  },
  modules: {
  }
});
