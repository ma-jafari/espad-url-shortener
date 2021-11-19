<template>
  <v-app
    class="app-container"
    v-bind:style="{
      'background-image': 'url(' + require('./assets/background.png') + ')',
      'background-position': 'center',
      'background-repeat': 'no-repeat',
      'background-size': 'cover',
    }"
  >
    <v-main>
      <Home />
    </v-main>
  </v-app>
</template>

<script>
import Home from "./components/Home";

export default {
  name: "App",

  components: {
    Home,
  },
};
</script>