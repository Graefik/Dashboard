<script setup lang="ts">
import type { ApexOptions } from "apexcharts";
import { ref } from "vue";

const colors = [
  "#008FFB",
  "#00E396",
  "#FEB019",
  "#FF4560",
  "#775DD0",
  "#546E7A",
  "#26a69a",
  "#D10CE8",
];

interface SeriesData {
  name?: string;
  data: number[];
}

const series = ref<SeriesData[]>([
  {
    name: "Score",
    data: [21, 22, 10, 28, 16, 21, 13, 30],
  },
]);

const chartOptions = ref<ApexOptions>({
  chart: {
    height: 350,
    type: "bar",
    toolbar: {
      show: false,
    },
    events: {
      click: function (chart, w, e) {
        console.log(chart, w, e);
      },
    },
  },
  colors: colors,
  plotOptions: {
    bar: {
      columnWidth: "45%",
      distributed: true,
    },
  },
  dataLabels: {
    enabled: false,
  },
  legend: {
    show: false,
  },
  tooltip: {
    theme: "dark",
    style: {
      fontSize: "12px",
    },
  },
  xaxis: {
    categories: [
      ["John", "Doe"],
      ["Joe", "Smith"],
      ["Jake", "Williams"],
      "Amber",
      ["Peter", "Brown"],
      ["Mary", "Evans"],
      ["David", "Wilson"],
      ["Lily", "Roberts"],
    ],
    labels: {
      style: {
        colors: colors,
        fontSize: "12px",
      },
    },
  },
  yaxis: {
    labels: {
      style: {
        colors: "#ffffff",
        fontSize: "12px",
      },
    },
  },
});
</script>

<template>
  <div>
    <h1>Dashboard</h1>
    <p>Welcome to the dashboard!</p>
    <div>
      <apexchart
        width="500"
        type="bar"
        :options="chartOptions"
        :series="series"
      ></apexchart>
    </div>
    <router-link to="/login">Go to Login</router-link>
  </div>
</template>
