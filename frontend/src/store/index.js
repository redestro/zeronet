import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    freeze: false,
    token: '',
    gameMode: '',
    cells: Array(9).fill('')
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
    },
    startSession(state, payload) {
        state.token = payload.token;
        state.gameMode = payload.gameMode;
    },
    freezeSession(state) {
      state.freeze = true;
    },
    endSession(state) {
      state.freeze = false;
      state.token = '';
      state.gameMode = '';
      state.cells.fill('');
    }
  },
  getters: {
    token: state => state.token,
    gameMode: state => state.gameMode,
    cells: state => state.cells,
    freeze: state => state.freeze
  },
  modules: {
  }
});
