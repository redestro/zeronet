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
import api from "@/api/startSessionAPI";

export default {
  methods: {
    startSession(type) {
      api.startNewSession(type).then(
        ((token) => {
          this.$store.commit("startSession", { token, gameMode: type });
          Event.$emit("aiMove", { type, token });
        }).bind(this)
      );
    },
  },
};
</script>
