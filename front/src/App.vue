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
  <router-view />
    <!-- <v-main>
      <router-view />
    </v-main> -->
  </v-app>
</template>
