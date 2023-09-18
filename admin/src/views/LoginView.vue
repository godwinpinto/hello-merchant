<script setup>
import { reactive } from "vue";
import { useRouter } from "vue-router";
import { mdiAccount, mdiAsterisk } from "@mdi/js";
import SectionFullScreen from "@/components/SectionFullScreen.vue";
import CardBox from "@/components/CardBox.vue";
import FormCheckRadio from "@/components/FormCheckRadio.vue";
import FormField from "@/components/FormField.vue";
import FormControl from "@/components/FormControl.vue";
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import LayoutGuest from "@/layouts/LayoutGuest.vue";
import LandingIntro from '@/components/auth/LandingIntro.vue'
import TheAuth from '@/components/auth/TheAuth.vue'


const form = reactive({
  login: "john.doe",
  pass: "highly-secure-password-fYjUw-",
  remember: true,
});

const router = useRouter();

const submit = () => {
  router.push("/dashboard");
};

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

const myCookieValue = getCookie('token');
if(myCookieValue){
  router.push("/dashboard");
}

if (myCookieValue) {
    console.log('Value of myCookie:', myCookieValue);
} else {
    console.log('myCookie not found');
}
</script>

<template>
  <LayoutGuest>
    <SectionFullScreen v-slot="{ cardClass }" bg="purplePink">
      <CardBox :class="cardClass" is-form @submit.prevent="submit">


            <div class="grid sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-3 xl:grid-cols-3 bg-base-100 rounded-xl">
                <div class='lg:col-span-2 xl:col-span-2 md:col-span-2'>
                    <LandingIntro />
                </div>
                <div class='py-24 px-10'>
                    <TheAuth/>
                </div>
            </div>
      </CardBox>
    </SectionFullScreen>
  </LayoutGuest>
</template>
