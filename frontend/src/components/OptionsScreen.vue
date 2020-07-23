/* eslint-disable */
<template>
  <div>
    <Button v-on:click="startSession('human2')">Human vs Human</Button>
    <Button v-on:click="startSession('human')">Human vs AI</Button>
    <Button v-on:click="startSession('revhuman')">AI vs Human</Button>
    <Button v-on:click="startSession('ai')">AI vs AI</Button>
  </div>
</template>

<script>
import "@/assets/_grid.scss";
import api from '@/api/startSessionAPI';

export default {
  props: {},
  data() {
    return {};
  },
  methods: {
    startSession(type) {
      api.startNewSession(type).then(
        (event => {
          this.$store.commit("updateToken", event);
        }).bind(this)
      );
      this.$store.commit("updateCells", { cellNumbers: 0, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 1, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 2, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 3, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 4, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 5, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 6, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 7, activePlayer: '' });
      this.$store.commit("updateCells", { cellNumbers: 8, activePlayer: '' });
      this.$store.commit('updateGameMode', type);
    }
  }
};
</script>
