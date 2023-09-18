<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";

const username = ref('')
const password = ref('')
const errorMessage = ref('')


const router = useRouter();

const signIn = async () => {
    console.log("clicked")
    window.location="https://pdn-sgyfszamhvxwhp5fskkyabgyl3bipbwq.login.aws.us.pangea.cloud/authorize?state=xxxxxxxxxxxxx"
}

const updateFormValue = (updateType, value) => {

    console.log("value", value)

    if (updateType == 'username') {
        username.value = value;
    } else if (updateType == 'password') {
        password.value = value;
    }
    errorMessage.value = '';
}
</script>

<template>
    <h2 className='text-2xl font-semibold mb-2 text-center'>Merchant Admin Login</h2>
          <BaseButtons class="w-full">
            <BaseButton  class="w-full" type="button" color="custom" label="Login Securely" @click="signIn"/>
          </BaseButtons>

    <p class="py-1 mt-2"> <span class="font-semibold">Username:</span> demo</p>
    <p class="py-1 mt-2"> <span class="font-semibold">Password:</span> demo@1234</p>
</template>