/* eslint-disable */
<template>
  <div>
    <div class="gameStatus" :class="gameStatusColor">{{ gameStatusMessage }}</div>
    <table class="grid">
      <tr>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[0]" @strike="updateBoard" name="0"></cell>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[1]" @strike="updateBoard" name="1"></cell>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[2]" @strike="updateBoard" name="2"></cell>
      </tr>
      <tr>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[3]" @strike="updateBoard" name="3"></cell>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[4]" @strike="updateBoard" name="4"></cell>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[5]" @strike="updateBoard" name="5"></cell>
      </tr>
      <tr>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[6]" @strike="updateBoard" name="6"></cell>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[7]" @strike="updateBoard" name="7"></cell>
        <cell :activePlayer="this.activePlayer" :mark="this.cells[8]" @strike="updateBoard" name="8"></cell>
      </tr>
    </table>
    <button v-on:click="endSession">End Game Session</button>
  </div>
</template>

<script>
import api from '@/api/playGameAPI';
import endSessionApi from '@/api/endSessionAPI';
import Cell from './Cell.vue';
import '@/assets/_grid.scss';

export default {
  name: 'Board',
  components: { Cell },
  data() {
    return {
      activePlayer: 'X',
      token: '',
      gameStatus: 'turn',
      gameStatusMessage: "O's turn",
      moves: 0,
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
      },
      winConditions: [
        [0, 1, 2],
        [3, 4, 5],
        [6, 7, 8],
        [0, 3, 6],
        [1, 4, 7],
        [2, 5, 8],
        [0, 4, 8],
        [2, 4, 6]
      ]
    };
  },
  computed: {
    nonActivePlayer() {
      if (this.activePlayer === 'O') {
        return 'X';
      }
      return 'O';
    }
  },
  watch: {
    gameStatus() {
      if (this.gameStatus === 'win') {
        this.gameStatusColor = 'statusWin';
        return;
      }
      if (this.gameStatus === 'draw') {
        this.gameStatusColor = 'statusDraw';
        this.gameStatusMessage = 'Draw !';

        return;
      }
      this.gameStatusMessage = `${this.activePlayer}'s turn`;
    }
  },
  methods: {
    endSession() {
      this.$store.commit("updateToken", "");
      endSessionApi.endSession();
    },
    updateBoard(cellNumber) {
      this.cells[cellNumber] = this.activePlayer;
      this.moves++;
      this.gameStatus = this.changeGameStatus();
      this.changePlayer();
      this.getMove(cellNumber)
    },
    getMove(currentMove) {
      const payload = {
        player: '1',
        move: currentMove,
        symbol: 'O'
      }
      api.playGameAIvsHuman(payload, this.$store.getters.token).then(
        (event => {
          this.cells[event.data.move] = this.activePlayer;
          this.moves++;
          this.gameStatus = this.changeGameStatus();
          this.changePlayer();
        }).bind(this)
      )
    },
    changePlayer() {
      this.activePlayer = this.nonActivePlayer;
    },
    checkForWin() {
      for (let i = 0; i < this.winConditions.length; i++) {
        const wc = this.winConditions[i];
        const cells = this.cells;
        if (this.areEqual(cells[wc[0]], cells[wc[1]], cells[wc[2]])) {
          return true;
        }
      }
      return false;
    },
    gameIsWon() {
      this.$emit('win', this.activePlayer);
      this.gameStatusMessage = `${this.activePlayer} Wins !`;
      this.$emit('freeze');
      return 'win';
    },
    changeGameStatus() {
      if (this.checkForWin()) {
        return this.gameIsWon();
      } else if (this.moves === 9) {
        return 'draw';
      }
      return 'turn';
    },
    areEqual() {
      var len = arguments.length;
      for (var i = 1; i < len; i++) {
        if (arguments[i] === '' || arguments[i] !== arguments[i - 1]) {
          return false;
        }
      }
      return true;
    },
  },
  created() {
    Event.$on('gridReset', () => {
      Object.assign(this.$data, this.$options.data());
    });
  }
};
</script>
