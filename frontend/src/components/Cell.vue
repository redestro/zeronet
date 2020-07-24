/* eslint-disable */
<template>
  <td class="cell" @click="strike">{{ mark }}</td>
</template>

<script>
import "@/assets/_grid.scss";

export default {
  props: { activePlayer: String, name: String, mark: String },
  methods: {
    strike() {
      if (!this.$store.getters.freeze && !this.$store.getters.cells[this.name]) {
        this.mark = this.activePlayer;
        this.$emit('strike', this.name);
      }
    }
  }
};
</script>
