<template>
  <div class="home">
    <p class="text">Welcome to <span class="glitch" data-text="Zeronet">Zeronet</span></p>
    <Board />
  </div>
</template>

<script>
// @ is an alias to /src
import '@/assets/_home.scss';
import Board from '@/components/Board.vue';

export default {
  name: 'Home',
  components: {
    Board
  }
};
</script>
