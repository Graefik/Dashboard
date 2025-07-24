import "normalize.css";
import { createApp } from "vue";
import VueApexCharts from "vue3-apexcharts";
import App from "./App.vue";
import router from "./router";
import "./scss/layout.scss";

const app = createApp(App);
app.use(router);
app.component("apexchart", VueApexCharts);
app.mount("#app");
