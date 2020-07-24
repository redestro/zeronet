/* eslint-disable */
<template>
  <div>
    <div class="gameStatus" :class="gameStatusColor">{{ gameStatusMessage }}</div>
    <table class="grid">
      <tr>
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[0]"
          @strike="updateBoard"
          name="0"
        />
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[1]"
          @strike="updateBoard"
          name="1"
        />
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[2]"
          @strike="updateBoard"
          name="2"
        />
      </tr>
      <tr>
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[3]"
          @strike="updateBoard"
          name="3"
        />
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[4]"
          @strike="updateBoard"
          name="4"
        />
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[5]"
          @strike="updateBoard"
          name="5"
        />
      </tr>
      <tr>
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[6]"
          @strike="updateBoard"
          name="6"
        />
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[7]"
          @strike="updateBoard"
          name="7"
        />
        <cell
          :activePlayer="this.activePlayer"
          :mark="this.$store.getters.cells[8]"
          @strike="updateBoard"
          name="8"
        />
      </tr>
    </table>
    <div class="recommendations">Recommendation for current player</div>
    <div class="recommendations">Row: {{ this.row }} Col: {{ this.col }}</div>
    <button v-on:click="endSession">End Game Session</button>
  </div>
</template>

<script>
import api from "@/api/playGameAPI";
import endSessionApi from "@/api/endSessionAPI";
import { MOVES_ARR } from "../utils/utils";
import Cell from "./Cell.vue";
import "@/assets/_grid.scss";

export default {
  name: "Board",
  components: { Cell },
  data() {
    return {
      activePlayer: "O",
      gameStatus: "turn",
      gameStatusColor: "",
      gameStatusMessage: "O's turn",
      moves: 0,
      winConditions: [
        [0, 1, 2],
        [3, 4, 5],
        [6, 7, 8],
        [0, 3, 6],
        [1, 4, 7],
        [2, 5, 8],
        [0, 4, 8],
        [2, 4, 6],
      ],
      recommendations: "",
      row: "1",
      col: "1",
    };
  },
  watch: {
    gameStatus(newValue) {
      if (newValue === "win") {
        this.gameStatusColor = "statusWin";
        return;
      } else if (newValue === "draw") {
        this.gameStatusColor = "statusDraw";
        this.gameStatusMessage = "Draw !";
        return;
      }
      this.gameStatusMessage = `${this.activePlayer}'s turn`;
    },
  },
  methods: {
    endSession() {
      this.$store.commit("endSession");
      this.updateBoard();
      this.activePlayer = "O";
      this.moves = 0;
      this.gameStatusColor = "";
      this.row = 1;
      this.col = 1;
      this.gameStatusMessage = "O's turn";
    },
    updateBoard(cellNumber, gameMode) {
      this.$store.commit("updateCells", {
        cellNumbers: cellNumber,
        activePlayer: this.activePlayer,
      });
      this.moves++;
      this.gameStatus = this.changeGameStatus();
      this.changePlayer();
      if (gameMode !== "NA") {
        this.getMove(cellNumber, this.$store.getters.gameMode || gameMode);
      }
    },
    changeGameStatus() {
      if (this.checkForWin()) {
        return this.gameIsWon();
      } else if (this.moves === 9) {
        return "draw";
      }
      return "turn";
    },
    checkForWin() {
      for (let i = 0; i < this.winConditions.length; i++) {
        const wc = this.winConditions[i];
        const cells = this.$store.getters.cells;
        if (this.areEqual(cells[wc[0]], cells[wc[1]], cells[wc[2]])) {
          return true;
        }
      }
      return false;
    },
    gameIsWon() {
      this.gameStatusMessage = `${this.activePlayer} Wins !`;
      this.$store.commit("freezeSession");
      return "win";
    },
    getMove(currentMove, gameMode) {
      const payload = {
        player: "1",
        move: currentMove,
        symbol: "O",
      };
      const revPayload = {
        player: "2",
        move: currentMove,
        symbol: "X",
      };
      switch (gameMode) {
        case "human":
          api.playGameAIvsHuman(payload, this.$store.getters.token).then(
            ((event) => {
              this.$store.commit("updateCells", {
                cellNumbers: event.data.move,
                activePlayer: this.activePlayer,
              });
              this.moves++;
              this.recommendations = event.data.recoMove + 1;
              this.row = Math.floor(
                Number.isInteger(this.recommendations / 3)
                  ? this.recommendations / 3
                  : this.recommendations / 3 + 1
              );
              this.col =
                this.recommendations % 3 === 0 ? 3 : this.recommendations % 3;
              this.gameStatus = this.changeGameStatus();
              this.changePlayer();
            }).bind(this)
          );
          break;
        case "revhuman":
          api.playGameAIvsHuman(revPayload, this.$store.getters.token).then(
            ((event) => {
              this.$store.commit("updateCells", {
                cellNumbers: event.data.move,
                activePlayer: this.activePlayer,
              });
              this.moves++;
              this.recommendations = event.data.recoMove + 1;
              this.row = Math.floor(
                Number.isInteger(this.recommendations / 3)
                  ? this.recommendations / 3
                  : this.recommendations / 3 + 1
              );
              this.col =
                this.recommendations % 3 === 0 ? 3 : this.recommendations % 3;
              this.gameStatus = this.changeGameStatus();
              this.changePlayer();
            }).bind(this)
          );
          break;
        case "human2":
          api
            .playGameAIvsHuman(payload, this.$store.getters.token)
            .then((event) => {
              this.recommendations = event.data.recoMove + 1;
              this.row = Math.floor(
                Number.isInteger(this.recommendations / 3)
                  ? this.recommendations / 3
                  : this.recommendations / 3 + 1
              );
              this.col =
                this.recommendations % 3 === 0 ? 3 : this.recommendations % 3;
            });
          break;
        default:
          break;
      }
    },
    changePlayer() {
      if (this.activePlayer === "O") {
        this.activePlayer = "X";
      } else {
        this.activePlayer = "O";
      }
    },
    areEqual(...args) {
      const len = args.length;
      for (let i = 1; i < len; i++) {
        if (args[i] === "" || args[i] !== args[i - 1]) {
          return false;
        }
      }
      return true;
    },
  },
  created() {
    Event.$on("aiMove", ({ type, token }) => {
      if (type === "revhuman") {
        const payload = {
          player: 0,
          move: 0,
          symbol: "O",
        };
        api.firstAiMove(payload, token).then((event) => {
          this.recommendations = event.data.recoMove + 1;
          this.row = Math.floor(
            Number.isInteger(this.recommendations / 3)
              ? this.recommendations / 3
              : this.recommendations / 3 + 1
          );
          this.col =
            this.recommendations % 3 === 0 ? 3 : this.recommendations % 3;
        });
        this.updateBoard(0, "NA");
        this.activePlayer = "X";
      }
      if (type === "ai") {
        for (let i = 0; i < MOVES_ARR.length; i++) {
          setTimeout(() => {
            this.updateBoard(MOVES_ARR[i], "NA");
          }, i * 1000);
        }
      }
    });
  },
};
</script>
