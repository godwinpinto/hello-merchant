<script setup lang="ts">
import { verifyAccountNo } from '@/service/appServices';
import { useUserStore } from '@/stores/userStore';
import { storeToRefs } from 'pinia'
import { ref } from "vue";

const userStore = useUserStore();
const { signOut } = useUserStore();
const loading = ref(false);
const { stepIndicator, registrationInput } = storeToRefs(userStore)
const xrpl_ac_no_error = ref(false);
const xrpl_ac_no_error_exists = ref(false);

const goToNextPage = async () => {
    loading.value = true
    xrpl_ac_no_error_exists.value = false;
    try {
        if (registrationInput.value.account_no == "" || !registrationInput.value.account_no.startsWith("r") || registrationInput.value.account_no.length < 15) {
            xrpl_ac_no_error.value = true;
        } else {
            const response = await verifyAccountNo(registrationInput.value.account_no);
            const body = response.data.response;
            if (body.status == 200 && body.data.result && body.data.result == true) {
                xrpl_ac_no_error.value = false;
                stepIndicator.value++
            } else {
                xrpl_ac_no_error_exists.value = true;
                xrpl_ac_no_error.value = false;
            }
        }
    } catch (error) {
        console.log(error)
    }
    loading.value = false
}
</script>
<template>
    <div>
        <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your email</label>
        <input type="email" name="origin_id" id="origin_id" disabled :value="registrationInput.origin_id"
            class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
            required>
    </div>
    <div>
        <label for="account_no" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your XRPL account
            no</label>
        <input type="text" name="account_no" id="account_no" v-model="registrationInput.account_no"
            placeholder="Your XRPL account that you want notification for"
            class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
            required>
        <p class="mt-2 text-sm text-red-600 dark:text-red-500" v-if="xrpl_ac_no_error">Enter a valid XRPL account number</p>
        <p class="mt-2 text-sm text-red-600 dark:text-red-500" v-if="xrpl_ac_no_error_exists">Enter a valid XRPL account
            that exists on the ledger</p>
    </div>
    <div>
        <h3 class="mb-4 font-medium text-sm text-gray-900 dark:text-white">Notification configuration</h3>
        <ul
            class="items-center w-full text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-lg sm:flex dark:bg-gray-700 dark:border-gray-600 dark:text-white">
            <li class="w-full border-b border-gray-200 sm:border-b-0 sm:border-r dark:border-gray-600">
                <div class="flex items-center pl-3">
                    <input type="radio" value="true" name="sound" v-model="registrationInput.msg_meta_info.sound"
                        class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500">
                    <label for="horizontal-list-radio-license"
                        class="w-full py-3 ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Speak Aloud</label>
                </div>
            </li>
            <li class="w-full dark:border-gray-600">
                <div class="flex items-center pl-3">
                    <input type="radio" value="false" name="sound" v-model="registrationInput.msg_meta_info.sound"
                        class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500">
                    <label for="horizontal-list-radio-passport"
                        class="w-full py-3 ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Silent</label>
                </div>
            </li>
        </ul>
    </div>
    <div>
        <h3 class="mb-4 font-medium text-sm text-gray-900 dark:text-white">Language Preference</h3>
        <ul
            class="items-center w-full text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-lg sm:flex dark:bg-gray-700 dark:border-gray-600 dark:text-white">
            <li class="w-full border-b border-gray-200 sm:border-b-0 sm:border-r dark:border-gray-600">
                <div class="flex items-center pl-3">
                    <input id="horizontal-list-radio-license" type="radio" value="en" name="lang"
                        v-model="registrationInput.msg_meta_info.lang"
                        class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500">
                    <label for="horizontal-list-radio-license"
                        class="w-full py-3 ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">English</label>
                </div>
            </li>
            <li class="w-full border-b border-gray-200 sm:border-b-0 sm:border-r dark:border-gray-600">
                <div class="flex items-center pl-3">
                    <input id="horizontal-list-radio-license" type="radio" value="sp" name="lang"
                        v-model="registrationInput.msg_meta_info.lang"
                        class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500">
                    <label for="horizontal-list-radio-license"
                        class="w-full py-3 ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Spanish</label>
                </div>
            </li>
            <li class="w-full dark:border-gray-600">
                <div class="flex items-center pl-3">
                    <input id="horizontal-list-radio-passport" type="radio" value="fr" name="lang"
                        v-model="registrationInput.msg_meta_info.lang"
                        class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500">
                    <label for="horizontal-list-radio-passport"
                        class="w-full py-3 ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">French</label>
                </div>
            </li>
        </ul>
    </div>
    <div class="flex justify-between">
        <button @click="signOut"
            class="relative inline-flex items-center justify-center p-0.5 mb-2 mr-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-pink-500 to-orange-400 group-hover:from-pink-500 group-hover:to-orange-400 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-pink-200 dark:focus:ring-pink-800">
            <span
                class="relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
                Back
            </span>
        </button>
        <button type="button" @click="goToNextPage"
            class="text-white font-medium rounded-lg text-sm px-5 py-2.5 text-center mr-2 text-gray-900 rounded-lg group bg-gradient-to-br from-pink-500 to-orange-400 group-hover:from-pink-500 group-hover:to-orange-400 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-pink-200 dark:focus:ring-pink-800 inline-flex items-center">
            <svg v-if="loading" aria-hidden="true" role="status" class="inline w-4 h-4 mr-3 text-white animate-spin"
                viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path
                    d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                    fill="#E5E7EB" />
                <path
                    d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                    fill="currentColor" />
            </svg>
            Continue
        </button>
    </div>
</template>