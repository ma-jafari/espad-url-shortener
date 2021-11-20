import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import axios from "axios";
import VueAxios from "vue-axios";
import VueCookies from "vue-cookies";
import VueRouter from 'vue-router'
import router from './router'

Vue.use(VueRouter)
Vue.use(VueCookies);
Vue.use(VueAxios, axios);

Vue.config.productionTip = false;
Vue.prototype.$api = "http://localhost:6010";
Vue.prototype.$panel = "http://localhost:8080/home";

new Vue({
  vuetify,
  router,
  render: (h) => h(App)
}).$mount("#app");
