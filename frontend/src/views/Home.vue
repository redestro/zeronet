<template>
  <div class="home">
    <p class="text">Welcome to <span class="glitch" data-text="Zeronet">Zeronet</span></p>
    <OptionsScreen v-show="!$store.getters.token" />
    <Board v-show="$store.getters.token" />
  </div>
</template>

<script>
// @ is an alias to /src
import '@/assets/_home.scss';
import Board from '@/components/Board.vue';
import OptionsScreen from '@/components/OptionsScreen.vue';

export default {
  name: 'Home',
  components: {
    Board,
    OptionsScreen
  }
};
</script>
