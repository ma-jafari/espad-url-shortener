import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import axios from "axios";
import VueAxios from "vue-axios";
import VueCookies from "vue-cookies";

Vue.use(VueCookies);
Vue.use(VueAxios, axios);

Vue.config.productionTip = false;
Vue.prototype.$api = "http://localhost:7070";

new Vue({
  vuetify,
  render: (h) => h(App),
}).$mount("#app");
