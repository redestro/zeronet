/* eslint-disable */
<template>
  <td class="cell" @click="strike">{{ mark }}</td>
</template>

<script>
import "@/assets/_grid.scss";

export default {
  props: { activePlayer: String, name: String, mark: String },
  data() {
    return {
      frozen: false
    };
  },
  methods: {
    strike() {
      if (!this.frozen && !this.$store.getters.freeze) {
        this.mark = this.activePlayer;
        this.$emit('strike', this.name);
      }
    }
  },
  created() {
    this.$on("clearCell", () => {
      this.mark = "";
      this.frozen = false;
    });
  }
};
</script>
