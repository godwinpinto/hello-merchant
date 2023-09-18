<script setup lang="ts">
import { useUserStore } from '@/stores/userStore';
import { onBeforeMount } from "vue";

const userStore = useUserStore();
const { signInWithGoogle, asyncSetUser } = useUserStore();

onBeforeMount(() => {
    asyncSetUser();
});


const signIn = async () => {
    console.log("clicked")
    window.location.assign("https://pdn-sgyfszamhvxwhp5fskkyabgyl3bipbwq.login.aws.us.pangea.cloud/authorize?state=xxxxxxxxxxxxx&redirect_uri="+encodeURIComponent(import.meta.env.VITE_REDIRECT_URI||''))
}

</script>
<template>
    <h2 class="text-2xl font-bold text-gray-900 dark:text-white flex justify-center items-center">
        Web Box - Login
    </h2>
    <div class="text-sm font-medium text-gray-500 dark:text-gray-400 text-center flex justify-center items-center">
        Use your Hello Merchant login to receive sound alerts for payments
    </div>
    <div class="flex justify-center items-center">
        <button type="button" @click="signIn"
            class="text-white bg-blue-500 hover:bg-blue-300 border border-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-600 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:focus:ring-blue-600 dark:bg-blue-800 dark:border-blue-700 dark:text-white dark:hover:bg-blue-700 mr-2 mb-2">
            Click here to Login
        </button>
    </div>
</template>