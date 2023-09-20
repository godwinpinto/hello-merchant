<script setup>
import { ref, onBeforeMount } from 'vue';
import { useRouter } from 'vue-router';
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import axios from 'axios';

const errorMessage = ref('')


const router = useRouter();

const signIn = async () => {
    console.log("clicked")
    window.location = import.meta.env.VITE_PANGEA_AUTHN_URL
}

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop()?.split(';').shift();
}


async function asyncSetUser() {
    try {

        let emailCookieValue = getCookie('admin_email');

        if (emailCookieValue) {
            console.log("Logged In")
            try {
                const response = await axios.post(import.meta.env.VITE_SERVER_URL + "/validate");
                if (response.data.status == 200) {
                    console.log("success",response)
                    router.push("/dashboard")
                } else {
                    errorMessage.value=response.data.message;
                    console.log("failure")
                }
            } catch (error) {
                console.error('Error fetching data:', error);
                return false
            }
        } else {
            console.log("Not Logged In")
        }
    } catch (error) {
        console.log("Something went wrong in authentication", error)
    }
}

onBeforeMount(() => {
    asyncSetUser();
});

</script>

<template>
    <h2 className='text-2xl font-semibold mb-2 text-center'>Merchant Admin Login</h2>
    <BaseButtons class="w-full">
        <BaseButton class="w-full" type="button" color="custom" label="Login Securely" @click="signIn" />
    </BaseButtons>
    <p class="py-1 mt-2 color text-red-800"> {{ errorMessage }}</p>

    <p class="py-1 mt-2"> <span class="font-semibold">Username:</span> hello.merchant@coauth.dev</p>
    <p class="py-1 mt-2"> <span class="font-semibold">Password:</span> Demo@1234</p>
</template>