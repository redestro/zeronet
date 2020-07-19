/* eslint-disable */
<template>
  <div>
    <div class="gameStatus" :class="gameStatusColor">{{ gameStatusMessage }}</div>
    <table class="grid">
      <tr>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="0"></cell>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="1"></cell>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="2"></cell>
      </tr>
      <tr>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="3"></cell>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="4"></cell>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="5"></cell>
      </tr>
      <tr>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="6"></cell>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="7"></cell>
        <cell :activePlayer="this.activePlayer" @strike="updateBoard" name="8"></cell>
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
      console.log(this.activePlayer)
      this.gameStatusMessage = `${this.activePlayer}'s turn`;
    }
  },
  methods: {
    endSession() {
      console.log("hi")
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
      api.playGameAIvsHuman(currentMove, this.$store.getters.token).then(
        (event => {
          console.log(event)
          this.updateBoard(event.move)
          console.log(this.cells)
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