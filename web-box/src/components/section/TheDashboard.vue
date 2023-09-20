<script setup lang="ts">
import { useUserStore } from '@/stores/userStore';
import { storeToRefs } from 'pinia'
import { ref, inject, computed, onMounted } from "vue";
import { fetchTransactions } from '@/service/appServices';
import echoAudio from '@/assets/echo.mp3';
import BalanceLoader from '@/components/skeleton/BalanceLoader.vue'
import TransactionLoader from '../skeleton/TransactionLoader.vue';

const userStore = useUserStore();
const { signOut } = useUserStore();
const { userInfo } = storeToRefs(userStore)
interface ITransactions {
    channel: string
    created_dt: string
    amount: string
    currency: string
}
const transactionsArray = ref<Array<ITransactions>>([]);
const balance = ref("");
const account = ref("");
const refresh = ref(false)
const firstTimeLoad = ref(true)
const audio: HTMLAudioElement = new Audio(echoAudio);
const speakMessage = ref('')
const speakLang = ref('')
let synth = window.speechSynthesis;
let greetingSpeech = new window.SpeechSynthesisUtterance()


const newAmount = ref("");


const pusher: any = inject('pusher');
var channel = pusher.subscribe(import.meta.env.VITE_PUSHER_APP_CHANNEL);

channel.bind(userInfo.value.email, function (data: any) {
    console.log(userInfo.value.email)
    console.log(data)
    const parseInfo: any = JSON.parse(data.meta)
    newAmount.value=data.message
    fetchBalanceAndTransactions(true);
    if (parseInfo.sound == false)
        return
    if (parseInfo.lang == "en") {
        speakStart("You received " + data.message + " XRP on Web Box", "en-UK");
    } else if (parseInfo.lang == "sp") {
        speakStart("Recibiste " + data.message + " XRP en Web Box", "es-ES");
    } else if (parseInfo.lang == "fr") {
        speakStart("Vous avez reçu " + data.message + " XRP sur  Web Box", "fr-FR");
    } else if (parseInfo.lang == "hi") {
        speakStart("Web Box per, " + data.message + " XRP praapth hue", "hi-IN");
    } else {
        speakStart("You received " + data.message + " XRP on Web Box", "en-UK");
    }
});


const fetchBalanceAndTransactions = async (calledFromNotification: boolean) => {
    if (calledFromNotification) {
        refresh.value = true;
    } else {
        refresh.value = false;
    }
    try {
        const response = await fetchTransactions(userInfo.value.email);
        console.log(response);
        if (response.status == 200 && response.data.data) {
            transactionsArray.value = response.data.data as Array<ITransactions>;
            console.log("transactions");
        }
        firstTimeLoad.value=false;

    } catch (error) {
        console.log(error);
        firstTimeLoad.value=false;
    }
    refresh.value = false;
}

const configDateFormat: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'short',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: true,
    weekday: 'short',
};

const formatTransactions = computed(() => {
    return transactionsArray.value.map((val) => {
        var d = new Date(val.created_dt)
        return {
            channel: val.channel, amount: val.amount, created_dt: d.toLocaleString('en-US', configDateFormat), currency:val.currency
        };
    });
});

audio.addEventListener('ended', () => {
    speak(speakMessage.value, speakLang.value);
});

const speakStart = (message: string, lang: string) => {
    audio.play();
    speakMessage.value = message;
    speakLang.value = lang;
}

const playSoungTest = () => {
    speakStart("Welcome to Web Box!", "en-US");
}

const speak = (text: string, lang: String) => {
    const voice = synth.getVoices().filter(x => x.lang == lang);
    greetingSpeech.voice = voice[0]
    greetingSpeech.text = text
    greetingSpeech.pitch = 1
    greetingSpeech.rate = 1
    synth.speak(greetingSpeech)
    setTimeout(function(){
        newAmount.value=""
    },10000)
}

onMounted(() => {
    fetchBalanceAndTransactions(false);
})
</script>
<template>
    <div v-if="refresh" class="indeterminate-progress-bar">
        <div class="indeterminate-progress-bar__progress bg-yellow-300 dark:text-gray-600 "></div>
    </div>
    <!-- <div class="w-full mt-4 xl:grid-cols-2 2xl:grid-cols-3">
        <div v-if="account != ''"
            class="items-center justify-between p-4 bg-white border border-gray-200 rounded-lg shadow-sm sm:flex dark:border-gray-700 sm:p-6 dark:bg-gray-800">
            <div class="w-full">
                <div>
                    <h3 class="text-base font-normal text-gray-500 dark:text-gray-400">Your account balance</h3>
                    <span v-if="account != ''"
                        class="text-2xl font-bold leading-none text-gray-900 sm:text-3xl dark:text-white">XRP {{ balance
                        }}</span>
                    <p class="flex items-center text-base font-normal text-gray-500 dark:text-gray-400">
                        <span class="flex items-center mr-1.5 text-sm text-green-500 dark:text-green-400">
                            {{ account }}
                        </span>
                    </p>
                </div>
            </div>
        </div>
        <BalanceLoader v-if="account == ''" />
    </div> -->
    <div class="flex items-center p-4 mb-4 text-sm text-green-800 border border-green-300 rounded-lg bg-green-50 dark:bg-gray-800 dark:text-green-400 dark:border-green-800"
        role="alert">
        <svg class="flex-shrink-0 inline w-4 h-4 mr-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
            fill="currentColor" viewBox="0 0 20 20">
            <path
                d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z" />
        </svg>
        <span class="sr-only">Info</span>
        <div>
            You need to manually test sound first to voice notify, on every page reload. Tested on Chrome, Safari and Edge.
        </div>
    </div>

    <div class="flex justify-between">
        <button @click="playSoungTest"
            class="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-pink-500 to-orange-400 group-hover:from-pink-500 group-hover:to-orange-400 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-pink-200 dark:focus:ring-pink-800">
            <span
                class="relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
                Sound / Speaker Test
            </span>
        </button>
        <!--  <button type="button" @click="fetchBalanceAndTransactions(true)"
            class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">Reload
            Page</button>
 -->
        <button @click="signOut"
            class="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-blue-200 via-blue-300 to-blue-200 group-hover:from-red-200 group-hover:via-red-300 group-hover:to-yellow-200 dark:text-white dark:hover:text-gray-900 focus:ring-4 focus:outline-none focus:ring-red-100 dark:focus:ring-red-400">
            <span class="relative px-5 py-2.5 transition-all ease-in duration-75 rounded-md group-hover:bg-opacity-0">
                Logout
            </span>
        </button>

    </div>
    <!-- component -->
    <div class="bg-gray-100" v-if="newAmount!=''">
        <div class="bg-gray-100 p-6  md:mx-auto">
            <svg viewBox="0 0 24 24" class="text-green-600 w-16 h-16 mx-auto my-6">
                <path fill="currentColor"
                    d="M12,0A12,12,0,1,0,24,12,12.014,12.014,0,0,0,12,0Zm6.927,8.2-6.845,9.289a1.011,1.011,0,0,1-1.43.188L5.764,13.769a1,1,0,1,1,1.25-1.562l4.076,3.261,6.227-8.451A1,1,0,1,1,18.927,8.2Z">
                </path>
            </svg>
            <div class="text-center">
                <h3 class="md:text-xl text-base text-gray-900 text-center">Payment Received on Ripple!</h3>
                <p class="text-gray-900 my-2 md:text-2xl font-semibold ">Amount: XRP {{newAmount}}</p>
            </div>
        </div>
    </div>
    <TransactionLoader v-if="firstTimeLoad" />
    <div class="min-h-70 h-70 overflow-y-scroll pr-4" id="about" role="tabpanel" aria-labelledby="about-tab"
    v-if="!firstTimeLoad">
        <ul role="list" class="divide-y divide-gray-200 dark:divide-gray-700" v-auto-animate>
            <li class="py-3 sm:py-4" v-for="(txn, index) in formatTransactions" :key="txn.created_dt">
                <div class="flex items-center space-x-4">
                    <div
                        class="relative inline-flex items-center justify-center w-10 h-10 overflow-hidden bg-gray-100 rounded-full dark:bg-gray-600">
                        <span class="font-medium text-gray-600 dark:text-gray-300">{{ index + 1 }}</span>
                    </div>
                    <div class="flex-1 min-w-0">
                        <p class="font-medium text-gray-900 truncate dark:text-white">
                            {{ txn.created_dt }}
                        </p>
                        <p class="text-sm text-gray-500 truncate dark:text-gray-400">
                            {{ txn.channel }}
                        </p>
                    </div>
                    <div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">
                        {{ txn.currency }} {{ txn.amount }}
                    </div>
                </div>
            </li>
        </ul>
    </div>
</template>
<style scoped>
.indeterminate-progress-bar {
    border-radius: 9999px;
    height: 0.5rem;
    position: relative;
    overflow: hidden;
}

.indeterminate-progress-bar__progress {
    border-radius: 9999px;
    position: absolute;
    bottom: 0;
    top: 0;
    width: 70%;
    animation-duration: 2s;
    animation-iteration-count: infinite;
    animation-name: indeterminate-progress-bar;
}

@keyframes indeterminate-progress-bar {
    from {
        left: -50%;
    }

    to {
        left: 100%;
    }
}
</style>
    