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
                  v-model="tempShortURL"
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
            <v-divider v-if="loggedIn"></v-divider>
            <h4 class="mt-5" v-if="loggedIn">Prevoius Links:</h4>
            <v-list three-line v-if="loggedIn">
              <template v-for="item in histoyLinks">
                <v-list-item :key="item.title">
                  <v-list-item-avatar>
                    <v-chip
                      class="ma-2"
                      color="red"
                      text-color="white"
                      v-if="item.is_expired"
                    >
                    </v-chip>
                    <v-chip
                      class="ma-2"
                      color="success"
                      text-color="white"
                      v-else
                    >
                    </v-chip>
                  </v-list-item-avatar>

                  <v-list-item-content>
                    <v-list-item-title
                      v-html="item.original_url"
                    ></v-list-item-title>
                    <v-list-item-subtitle
                      v-html="'http://localhost:8080/' + item.short_url"
                    ></v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
              </template>
            </v-list>
            <span
              v-if="!loggedIn"
              style="padding: 10px; border-radius: 5px; color: white"
              class="success"
              >To see links history, log in</span
            >
          </v-card>
          <v-card flat v-else>
            <v-text-field
              :disabled="loggedIn"
              filled
              color="blue-grey lighten-2"
              label="Name"
              v-model="user.name"
              :rules="nameRules"
              required
            ></v-text-field>
            <v-text-field
              filled
              :disabled="loggedIn"
              color="blue-grey lighten-2"
              label="E-Mail"
              v-model="user.email"
              :rules="emailRules"
              required
            ></v-text-field>
            <v-text-field
              v-if="loggedIn == false"
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
      v-model="snackbar"
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
    tempShortURL: "",
    shortened: "",
    previousLink: null,
    histoyLinks: [
      {
        original_url: "",
        shortened_url: "",
        is_expired: false,
      },
    ],
    user: {
      name: "",
      email: "",
      password: "",
    },
    userID: "",
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
      try {
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
            this.loggedIn = true;
            this.userID = response.data.data.id;
          });
      } catch (error) {
        this.snackbar = true;
        this.snackbarColor = "error";
        this.snackbarText = error;
        return;
      }
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
          this.$cookies.set("email", response.data.data.email, end);
          this.snackbar = true;
          this.snackbarColor = "success";
          this.snackbarText = "User logged in successfully";
          this.loggedIn = true;
          this.userID = response.data.data.id;
        })
        .catch((error) => {
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
      this.user.name = "";
      this.user.email = "";
      this.user.password = "";
      this.userID = "";
    },
    async ShortenURL() {
      if (this.urlObject.original_url == "") {
        this.snackbar = true;
        this.snackbarColor = "error";
        this.snackbarText = "Please fill the form";
        return;
      }

      this.urlObject.user_id = this.userID;
      this.urlObject.short_url = this.tempShortURL;
      await this.axios
        .post(this.$api + "/url", this.urlObject)
        .then((response) => {
          this.shortened =
            "http://localhost:8080/" + response.data.data.short_url;

          // this.urlObject = null;
        })
        .catch((error) => {
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
      await this.axios
        .get(this.$api + "/user/" + this.user.email, this.user, config)
        .then((response) => {
          this.user = response.data.data;
          this.loggedIn = true;
          this.userID = response.data.data.id;
        })
        .catch((error) => {
          this.snackbar = true;
          this.snackbarColor = "error";
          this.snackbarText = error;
          return;
        });
    },
    async GetUserHistory() {
      await this.axios
        .get(this.$api + "/url/history/" + this.userID)
        .then((response) => {
          this.histoyLinks = response.data.data.url_history;
        })
        .catch((error) => {
          this.snackbar = true;
          this.snackbarColor = "error";
          this.snackbarText = error;
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
  async mounted() {
    var savedEmail = this.$cookies.get("email");
    var savedToken = this.$cookies.get("token");
    if (savedEmail != null && savedToken != null) {
      await this.GetUser(savedEmail, savedToken);
      await this.GetUserHistory();
    } else {
      this.loggedIn = false;
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
    async tempShortURL(val) {
      // let object = {
      //   short_url: val,
      // };
      // let request = JSON.stringify(object);
      await this.axios.get(this.$api + "/url/check/" + val).catch((error) => {
        this.snackbar = true;
        this.snackbarColor = "error";
        this.snackbarText = error;
      });
    },
  },
};
</script>

<style scoped>
</style>
