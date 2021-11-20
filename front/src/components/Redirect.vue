<template>
  <div class="hello">123</div>
</template>

<script>
export default {
  name: "Home",
  data: () => ({}),
  async mounted() {
    let path = window.location.pathname;
    if (path == "" || path == "/") {
      window.location.replace(this.$panel);
    } else {
    window.location.replace(this.$api + path);

    }
  },
};
</script>

<style scoped>
</style>
