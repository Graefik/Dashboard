<script setup lang="ts">
import type { ApexOptions } from "apexcharts";
import { ref } from "vue";
import Wrapper from "@/shared/components/Wrapper.vue";

const colors = [
  "#00c2b2",
  "#00e0d1",
  "#66d9ef",
  "#a6e22e",
  "#f92672",
  "#fd971f",
  "#ae81ff",
  "#e6db74",
];

interface SeriesData {
  name?: string;
  data: number[] | { x: string; y: number }[];
}

const baseChartConfig: Partial<ApexOptions> = {
  chart: {
    toolbar: {
      show: false,
    },
    background: "transparent",
    fontFamily: "Inter, sans-serif",
  },
  colors: colors,
  dataLabels: {
    enabled: false,
  },
  legend: {
    show: false,
  },
  tooltip: {
    theme: "dark",
    style: {
      fontSize: "1.2rem",
      fontFamily: "Inter, sans-serif",
    },
    cssClass: "dashboard-tooltip",
  },
  grid: {
    borderColor: "rgba(255, 255, 255, 0.1)",
    strokeDashArray: 3,
  },
  xaxis: {
    labels: {
      style: {
        colors: "#ffffff",
        fontSize: "1.2rem",
        fontFamily: "Inter, sans-serif",
      },
    },
    axisBorder: {
      color: "rgba(255, 255, 255, 0.1)",
    },
    axisTicks: {
      color: "rgba(255, 255, 255, 0.1)",
    },
  },
  yaxis: {
    labels: {
      style: {
        colors: "#ffffff",
        fontSize: "1.2rem",
        fontFamily: "Inter, sans-serif",
      },
    },
  },
};

const barSeries = ref<SeriesData[]>([
  {
    name: "Sales",
    data: [21, 22, 10, 28, 16, 21, 13, 30],
  },
]);

const barChartOptions = ref<ApexOptions>({
  ...baseChartConfig,
  chart: {
    ...baseChartConfig.chart,
    type: "bar",
  },
  plotOptions: {
    bar: {
      columnWidth: "45%",
      distributed: true,
    },
  },
  xaxis: {
    ...baseChartConfig.xaxis,
    categories: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug"],
  },
});

const lineSeries = ref<SeriesData[]>([
  {
    name: "Revenue",
    data: [31, 40, 28, 51, 42, 109, 100, 120],
  },
  {
    name: "Profit",
    data: [11, 32, 45, 32, 34, 52, 41, 80],
  },
]);

const lineChartOptions = ref<ApexOptions>({
  ...baseChartConfig,
  chart: {
    ...baseChartConfig.chart,
    type: "line",
  },
  stroke: {
    curve: "smooth",
    width: 2,
  },
  xaxis: {
    ...baseChartConfig.xaxis,
    categories: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug"],
  },
  legend: {
    show: true,
    position: "top",
    labels: {
      colors: "#ffffff",
    },
  },
});

const areaSeries = ref<SeriesData[]>([
  {
    name: "Users",
    data: [13, 23, 20, 8, 13, 27, 33, 42],
  },
]);

const areaChartOptions = ref<ApexOptions>({
  ...baseChartConfig,
  chart: {
    ...baseChartConfig.chart,
    type: "area",
  },
  fill: {
    type: "gradient",
    gradient: {
      shadeIntensity: 1,
      opacityFrom: 0.7,
      opacityTo: 0.3,
    },
  },
  xaxis: {
    ...baseChartConfig.xaxis,
    categories: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug"],
  },
});

const donutSeries = ref<number[]>([44, 55, 13, 43, 22]);

const donutChartOptions = ref<ApexOptions>({
  ...baseChartConfig,
  chart: {
    ...baseChartConfig.chart,
    type: "donut",
  },
  labels: ["Team A", "Team B", "Team C", "Team D", "Team E"],
  legend: {
    show: true,
    position: "bottom",
    labels: {
      colors: "#ffffff",
    },
  },
});
</script>

<template>
  <Wrapper class="dashboard">
    <div class="dashboard__header">
      <div class="dashboard__header-content">
        <div class="dashboard__title-section">
          <h1 class="dashboard__title">Dashboard Principal</h1>
          <p class="dashboard__subtitle">
            Votre dashboard principal pour un aperçu rapide de vos données et
            services en cours.
          </p>
        </div>
      </div>
    </div>

    <div class="dashboard__content">
      <div class="dashboard__grid">
        <div class="card">
          <div class="card__header">
            <div class="card__title">
              <h3 class="card__title-text">Ventes Mensuelles</h3>
            </div>
            <div class="card__actions"></div>
          </div>
          <div class="card__content">
            <apexchart
              width="100%"
              height="280"
              type="bar"
              :options="barChartOptions"
              :series="barSeries"
            ></apexchart>
          </div>
        </div>

        <div class="card">
          <div class="card__header">
            <div class="card__title">
              <h3 class="card__title-text">Revenus & Profits</h3>
            </div>
          </div>
          <div class="card__content">
            <apexchart
              width="100%"
              height="280"
              type="line"
              :options="lineChartOptions"
              :series="lineSeries"
            ></apexchart>
          </div>
        </div>

        <div class="card">
          <div class="card__header">
            <div class="card__title">
              <h3 class="card__title-text">Croissance Utilisateurs</h3>
            </div>
          </div>
          <div class="card__content">
            <apexchart
              width="100%"
              height="280"
              type="area"
              :options="areaChartOptions"
              :series="areaSeries"
            ></apexchart>
          </div>
        </div>

        <div class="card">
          <div class="card__header">
            <div class="card__title">
              <h3 class="card__title-text">Répartition Équipes</h3>
            </div>
          </div>
          <div class="card__content">
            <apexchart
              width="100%"
              height="280"
              type="donut"
              :options="donutChartOptions"
              :series="donutSeries"
            ></apexchart>
          </div>
        </div>
      </div>
    </div>
  </Wrapper>
</template>

<style scoped lang="scss">
.dashboard {
  &__header {
    margin-bottom: 4rem;
    background-color: $primary-fg;
    border-radius: 3rem;
    padding: 3rem;
    box-shadow: 0 2rem 4rem rgba(0, 0, 0, 0.3);
    border: 1px solid rgba(255, 255, 255, 0.1);

    &-content {
      display: flex;
      justify-content: space-between;
      align-items: center;
      flex-wrap: wrap;
      gap: 2rem;

      @media (max-width: 768px) {
        flex-direction: column;
        text-align: center;
      }
    }
  }

  &__title {
    font-size: 3rem;
    font-weight: 700;
    margin: 0 0 1rem 0;
    color: $white;
    background: linear-gradient(135deg, $white 0%, $cta 100%);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;

    @media (max-width: 768px) {
      font-size: 2.5rem;
    }
  }

  &__subtitle {
    color: rgba($white, 0.7);
    margin: 0;
    font-size: 1.4rem;
    font-weight: 400;
  }

  &__content {
    margin-top: 2rem;
  }

  &__grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(45rem, 1fr));
    gap: 2.5rem;

    @media (max-width: 768px) {
      grid-template-columns: 1fr;
      gap: 2rem;
    }
  }
}

.card {
  background-color: $primary-fg;
  border-radius: 2rem;
  overflow: hidden;
  box-shadow: 0 1rem 3rem rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &:hover {
    transform: translateY(-0.5rem);
    box-shadow: 0 2rem 4rem rgba(0, 0, 0, 0.3);
    border-color: rgba($cta, 0.5);
  }

  &__header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 2rem 2.5rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    background: linear-gradient(135deg, rgba($cta, 0.05) 0%, transparent 100%);
  }

  &__title {
    display: flex;
    align-items: center;
    gap: 1.5rem;
  }

  &__title-text {
    margin: 0;
    font-size: 1.6rem;
    color: $white;
    font-weight: 600;
  }

  &__content {
    padding: 2rem;
    background-color: $primary-fg;
  }
}
</style>
