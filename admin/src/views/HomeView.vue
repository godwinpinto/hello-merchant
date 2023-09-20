<script setup>
import { computed, ref, onMounted } from "vue";
import { useMainStore } from "@/stores/main";
import {
  mdiAccountMultiple,
  mdiCartOutline,
  mdiChartTimelineVariant,
  mdiMonitorCellphone,
  mdiReload,
  mdiGithub,
  mdiChartPie,
} from "@mdi/js";
import { useRouter } from 'vue-router';
import * as chartConfig from "@/components/Charts/chart.config.js";
import LineChart from "@/components/Charts/LineChart.vue";
import SectionMain from "@/components/SectionMain.vue";
import CardBoxWidget from "@/components/CardBoxWidget.vue";
import CardBox from "@/components/CardBox.vue";
import BaseButton from "@/components/BaseButton.vue";
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import SectionTitleLineWithButton from "@/components/SectionTitleLineWithButton.vue";
import axios from "axios";

const router = useRouter();

const chartData = ref(null);

const rippleCount = ref(0);
const squareUpCount = ref(0);

const fillChartData = (data) => {
  chartData.value = sampleChartData(data);
};

const chartColors = {
  default: {
    primary: "#00D1B2",
    info: "#209CEE",
    danger: "#FF3860",
  },
};


const datasetObject = (color, data, node) => {


  const filteredArray = data
    .filter(item => item.channel === node)
    .map(item => item.record_count);
  return {
    fill: false,
    borderColor: chartColors.default[color],
    borderWidth: 2,
    borderDash: [],
    borderDashOffset: 0.0,
    pointBackgroundColor: chartColors.default[color],
    pointBorderColor: "rgba(255,255,255,0)",
    pointHoverBackgroundColor: chartColors.default[color],
    pointBorderWidth: 20,
    pointHoverRadius: 4,
    pointHoverBorderWidth: 15,

    pointRadius: 4,
    data: filteredArray,
    tension: 0.5,
    cubicInterpolationMode: "default",
  };
};

function formatDate(date) {
  const day = date.getDate().toString().padStart(2, '0');
  const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
  const month = monthNames[date.getMonth()];
  return `${day}-${month}`;
}


const sampleChartData = (data) => {
  const labels = [];

  const points = 5


  for (let i = points - 1; i >= 0; i--) {
    const currentDate = new Date();
    currentDate.setDate(currentDate.getDate() - i);
    const formattedDate = formatDate(currentDate);
    //    console.log(formattedDate);
    labels.push(formattedDate);
  }


  return {
    labels,
    datasets: [
      datasetObject("info", data, "Ripple"),
      datasetObject("primary", data, "SquareUp"),
    ],
  };
};

async function fetchDashboard() {
  try {
    const response = await axios.post(import.meta.env.VITE_SERVER_URL + "/dashboard");
    console.log("success", response)
    if (response.status == 200) {
      rippleCount.value = response.data.ripple_count
      squareUpCount.value = response.data.squareup_count
      fillChartData(response.data.transactions)
    } else if (response.data && response.data.status == 500) {
      router.push("/")
    }
  } catch (error) {
    console.error('Error fetching data:', error);
    return false
  }

}


onMounted(() => {
  fetchDashboard();
  //  fillChartData();
});

const mainStore = useMainStore();

const clientBarItems = computed(() => mainStore.clients.slice(0, 4));

const transactionBarItems = computed(() => mainStore.history);
</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <SectionTitleLineWithButton :icon="mdiChartPie" title="Overview" main>
        <BaseButton :icon="mdiReload" color="whiteDark" />

      </SectionTitleLineWithButton>

      <div class="grid grid-cols-1 gap-6 lg:grid-cols-3 mb-6">
        <CardBoxWidget trend-type="up" color="text-emerald-500" :icon="mdiChartTimelineVariant" :number="rippleCount"
          label="Ripple Request" />
        <CardBoxWidget trend-type="down" color="text-blue-500" :icon="mdiChartTimelineVariant" :number="squareUpCount"
          label="SquareUp Request" />
        <CardBoxWidget trend-type="alert" color="text-red-500" :icon="mdiChartTimelineVariant"
          :number="rippleCount + squareUpCount" label="Total Request" />
      </div>



      <SectionTitleLineWithButton :icon="mdiChartPie" title="Trends overview">
        <BaseButton :icon="mdiReload" color="whiteDark" @click="fillChartData" />
      </SectionTitleLineWithButton>

      <CardBox class="mb-6">
        <div v-if="chartData">
          <line-chart :data="chartData" class="h-96" />
        </div>
      </CardBox>

    </SectionMain>
  </LayoutAuthenticated>
</template>
