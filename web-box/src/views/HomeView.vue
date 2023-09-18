<script setup lang="ts">
import TheLogin from '../components/section/TheLogin.vue'
import TheDashboard from '../components/section/TheDashboard.vue'
import { ref, watch } from 'vue';
import { useUserStore } from '@/stores/userStore';
import { storeToRefs } from 'pinia'

const userStore = useUserStore();
const { stepIndicator } = storeToRefs(userStore)
const stepIndicatorLocal = ref(0);

watch(stepIndicator, (newVal, oldVal) => {
  stepIndicatorLocal.value = 1;
});
</script>
<template>
  <div class="flex flex-col items-center justify-center px-6 pt-8 mx-auto min-h-screen pt:mt-0 dark:bg-gray-900 bg-gray-300  pb-10">
    <a class="flex items-center justify-center mb-8 text-2xl font-semibold lg:mb-10 dark:text-white">
      <img src="@/assets/logo.png" class="mr-4 h-11" alt="FlowBite Logo">
    </a>
    <div class="w-full max-w-xl p-6 space-y-8 sm:p-8 bg-white rounded-lg shadow dark:bg-gray-800">
      <TheLogin v-if="stepIndicator == 0" />
      <TheDashboard v-if="stepIndicator == 4" />
    </div>
  </div>
</template>
