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
              <v-col md="12">
                <v-text-field
                  v-model="urlObject.original_url"
                  required
                  filled
                  color="blue-grey lighten-2"
                  label="Original URL"
                ></v-text-field>
              </v-col>
              <v-col md="3">
                <v-menu
                  ref="datePicker"
                  v-model="datePicker"
                  :close-on-content-click="false"
                  transition="scale-transition"
                  offset-y
                  max-width="290px"
                  min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="dateFormatted"
                      label="Expire Date"
                      persistent-hint
                      prepend-icon="mdi-calendar"
                      v-bind="attrs"
                      @blur="date = parseDate(dateFormatted)"
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="date"
                    no-title
                    @input="datePicker = false"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col md="9">
                <v-text-field
                  filled
                  color="blue-grey lighten-2"
                  label="Custome URL"
                  v-model="urlObject.short_url"
                  required
                ></v-text-field>
              </v-col>
              <v-col md="3">
                <v-btn
                  color="cyan"
                  class="white--text"
                  style="height: 56px"
                  @click="ShortenURL"
                >
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
              v-if="(loggedIn == false)"
              filled
              type="password"
              color="blue-grey lighten-2"
              label="Password"
              v-model="user.password"
              :rules="passRules"
              required
            ></v-text-field>
            <v-row>
              <v-col md="3" v-if="loggedIn == false">
                <v-btn
                  @click="RegisterUser"
                  color="success"
                  class="white--text"
                  style="height: 50px"
                >
                  Register
                </v-btn>
              </v-col>
              <v-col v-if="loggedIn == false">
                <v-btn
                  @click="LoginUser"
                  color="indigo lighten-1"
                  class="white--text"
                  style="height: 50px"
                >
                  Login
                </v-btn>
              </v-col>
              <v-col v-if="loggedIn == true">
                <v-btn
                  @click="LogoutUser"
                  color="error"
                  class="white--text"
                  style="height: 50px"
                >
                  Log out
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
    shortened: "",
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
    urlObject: {
      original_url: "",
      short_url: "",
      expire_at: null,
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
    picker: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
      .toISOString()
      .substr(0, 10),
    date: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
      .toISOString()
      .substr(0, 10),
    dateFormatted: null,
    datePicker: false,
    loggedIn: false,
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
          this.$cookies.set("email", response.data.data.email, end);
          this.snackbar = true;
          this.snackbarColor = "success";
          this.snackbarText = "User registered successfully";
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
          this.loggedIn = true;
        })
        .catch((error) => {
          console.log(error);
          this.snackbar = true;
          this.snackbarColor = "error";
          this.snackbarText = error;
          return;
        });
    },
    async LogoutUser() {
      this.$cookies.remove("email");
      this.$cookies.remove("token");
      this.loggedIn = false;
      this.user = null;
    },
    async ShortenURL() {
      if (this.urlObject.original_url == "") {
        this.snackbar = true;
        this.snackbarColor = "error";
        this.snackbarText = "Please fill the form";
        return;
      }

      this.urlObject.short_url =
        "http://localhost:8080/" + this.urlObject.short_url;
      await this.axios
        .post(this.$api + "/url", this.urlObject)
        .then((response) => {
          this.shortened = response.data.data.short_url;

          this.urlObject = null;
        })
        .catch((error) => {
          console.log(error);
          this.snackbar = true;
          this.snackbarColor = "error";
          this.snackbarText = error;
          return;
        });
    },
    async GetUser(savedEmail, savedToken) {
      let config = {
        headers: {
          Authorization: "Bearer " + savedToken,
        },
      };
      this.user.email = savedEmail;
      console.log("user: ", this.user);
      await this.axios
        .get(this.$api + "/user/" + this.user.email, this.user, config)
        .then((response) => {
          this.user = response.data.data;
          this.loggedIn = true;
        })
        .catch((error) => {
          console.log(error);
          this.snackbar = true;
          this.snackbarColor = "error";
          this.snackbarText = error;
          return;
        });
    },
    formatDate(date) {
      if (!date) return null;

      const [year, month, day] = date.split("-");
      return `${month}/${day}/${year}`;
    },
    parseDate(date) {
      if (!date) return null;

      const [month, day, year] = date.split("/");
      return `${year}-${month.padStart(2, "0")}-${day.padStart(2, "0")}`;
    },
  },
  mounted() {
    var savedEmail = this.$cookies.get("email");
    var savedToken = this.$cookies.get("token");
    if (savedEmail != null && savedToken != null) {
      this.GetUser(savedEmail, savedToken);
    } else {
      this.loggedIn = false
      console.log(this.loggedIn)
    }
  },
  watch: {
    date(val) {
      if (val != null) {
        this.dateFormatted = this.formatDate(this.date);
        var hms = "01:12:33";
        var target = new Date(Date.parse(val + " " + hms));
        this.urlObject.expire_at = target;
      }
    },
    snackbar(val) {
      if (val == true) {
        val = false;
      }
    },
  },
};
</script>

<style scoped>
</style>
