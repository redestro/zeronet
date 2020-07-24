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
import { BoardMixins } from "./mixins/BoardMixins";
import Cell from "./Cell.vue";
import "@/assets/_grid.scss";

export default {
  name: "Board",
  components: { Cell },
  mixins: [BoardMixins]
};
</script>
