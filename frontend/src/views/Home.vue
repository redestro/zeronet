<template>
  <div class="home">
    <img src="@/assets/left.png" class="left" />
    <p class="text">Welcome to <span class="glitch" data-text="Zeronet">Zeronet</span></p>
    <OptionsScreen v-show="!$store.getters.token" />
    <Board v-show="$store.getters.token" />
    <img src="@/assets/right.png" class="right" />
  </div>
</template>

<script>
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
