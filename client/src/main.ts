import "normalize.css";
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "./scss/layout.scss";

createApp(App).use(router).mount("#app");
