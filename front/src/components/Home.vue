<template>
  <div class="hello">
    <v-card class="mx-auto my-12" max-width="600">
      <v-toolbar color="cyan" dark flat>
        <template v-slot:extension>
          <v-tabs
            v-model="tab"
            align-with-title
            style="display: flex; flex-direction: row; justify-content: center"
          >
            <v-tab v-for="item in tabItems" :key="item">
              <span v-if="item != 'shortener'">
                <v-icon> {{ item }} </v-icon>
              </span>
              <span v-else>
                {{ item }}
              </span>
            </v-tab>
          </v-tabs>
        </template>
      </v-toolbar>

      <v-tabs-items v-model="tab" style="padding: 50px">
        <v-tab-item v-for="item in tabItems" :key="item">
          <v-card flat v-if="item == 'shortener'">
            <v-row>
              <v-col md="3">
                <v-select
                  required
                  :items="typeItems"
                  label="Type"
                  outlined
                ></v-select>
              </v-col>
              <v-col md="9">
                <v-text-field
                  required
                  filled
                  color="blue-grey lighten-2"
                  label="Original URL"
                ></v-text-field>
              </v-col>

              <v-col md="3">
                <v-btn color="cyan" class="white--text" style="height: 56px">
                  <v-icon left> mdi-cloud-upload </v-icon>

                  Shorten
                </v-btn>
              </v-col>
              <v-col md="9">
                <v-text-field
                  v-model="shortened"
                  :disabled="true"
                  filled
                  color="blue-grey lighten-2"
                  label="Shortened URL"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-divider class="mb-8"></v-divider>
            <v-select
              :items="previousLinkItems"
              label="Previous links"
              outlined
              @change="PreviousLink"
            ></v-select>
          </v-card>
          <v-card flat v-else>
            <v-text-field
              filled
              color="blue-grey lighten-2"
              label="Name"
              v-model="user.name"
              :rules="nameRules"
              required
            ></v-text-field>
            <v-text-field
              filled
              color="blue-grey lighten-2"
              label="E-Mail"
              v-model="user.email"
              :rules="emailRules"
              required
            ></v-text-field>
            <v-text-field
              filled
              type="password"
              color="blue-grey lighten-2"
              label="Password"
              v-model="user.password"
              :rules="passRules"
              required
            ></v-text-field>
            <v-row>
              <v-col md="3">
                <v-btn
                  @click="RegisterUser"
                  color="success"
                  class="white--text"
                  style="height: 50px"
                >
                  Register
                </v-btn>
              </v-col>
              <v-col>
                <v-btn
                  @click="LoginUser"
                  color="indigo lighten-1"
                  class="white--text"
                  style="height: 50px"
                >
                  Login
                </v-btn>
              </v-col>
            </v-row>
          </v-card>
        </v-tab-item>
      </v-tabs-items>
    </v-card>
    <v-snackbar
      :color="snackbarColor"
      :value="snackbar"
      :timeout="3000"
      absolute
      left
      top
      shaped
    >
      {{ snackbarText }}
    </v-snackbar>
  </div>
</template>

<script>
export default {
  name: "Home",
  data: () => ({
    loading: false,
    selection: 1,
    tab: null,
    tabItems: ["shortener", "mdi-account"],
    typeItems: ["Private", "Public"],
    shortened: "test",
    previousLink: null,
    previousLinkItems: ["google.com", "yahoo.com"],
    histoyLinks: [
      {
        original_url: "google.com",
        shortened_url: "test.testcom",
      },
    ],
    user: {
      name: "",
      email: "",
      password: "",
    },
    emailRules: [
      (v) => !!v || "E-mail is required",
      (v) => /.+@.+/.test(v) || "E-mail must be valid",
    ],
    nameRules: [
      (v) => !!v || "Name is required",
      (v) => v.length >= 0 || "Name cannot be empty",
    ],
    passRules: [
      (v) => !!v || "pass is required",
      (v) => v.length >= 6 || "Password must be more than 6 characters",
    ],
    snackbar: false,
    snackbarColor: "",
    snackbarText: "",
    token: "",
  }),
  methods: {
    PreviousLink(value) {
      console.log(value);
    },
    async RegisterUser() {
      if (
        this.user.name == "" ||
        this.user.email == "" ||
        this.user.password == ""
      ) {
        this.snackbar = true;
        this.snackbarColor = "error";
        this.snackbarText = "Please fill the form";
        return;
      }

      await this.axios
        .post(this.$api + "/user/register", this.user)
        .then((response) => {
          var end = new Date();
          end.setUTCHours(23, 59, 59, 999);
          this.$cookies.set("token", response.data.data.token, end);
          this.snackbar = true;
          this.snackbarColor = "success";
          this.snackbarText = "User registered successfully";
          this.token = response.data.data.token;
        })
        .catch((error) => {
          console.log(error);
          this.snackbar = true;
          this.snackbarColor = "error";
          this.snackbarText = error;
          return;
        });
    },
    async LoginUser() {
      if (
        this.user.name == "" ||
        this.user.email == "" ||
        this.user.password == ""
      ) {
        this.snackbar = true;
        this.snackbarColor = "error";
        this.snackbarText = "Please fill the form";
        return;
      }

      await this.axios
        .post(this.$api + "/user/login", this.user)
        .then((response) => {
          var end = new Date();
          end.setUTCHours(23, 59, 59, 999);
          this.$cookies.set("token", response.data.data.token, end);
          this.snackbar = true;
          this.snackbarColor = "success";
          this.snackbarText = "User logged in successfully";
          this.token = response.data.data.token;
        })
        .catch((error) => {
          console.log(error);
          this.snackbar = true;
          this.snackbarColor = "error";
          this.snackbarText = error;
          return;
        });
    },
  },
  mounted() {
    if (this.token == "") {
      var savedToken = this.$cookies.get("token");
      if (savedToken != "") {
        this.token = savedToken;
      }
    }
  },
  watch: {
    token: function (val) {
      console.log("watch: ", val);
    },
  },
};
</script>

<style scoped>
</style>
