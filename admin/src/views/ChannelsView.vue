<script setup>
import { reactive, ref, onMounted } from "vue";
import { useRouter } from 'vue-router';
import { mdiBallotOutline, mdiAccount, mdiMail, mdiGithub } from "@mdi/js";
import SectionMain from "@/components/SectionMain.vue";
import CardBox from "@/components/CardBox.vue";
import FormCheckRadioGroup from "@/components/FormCheckRadioGroup.vue";
import FormFilePicker from "@/components/FormFilePicker.vue";
import FormField from "@/components/FormField.vue";
import FormControl from "@/components/FormControl.vue";
import BaseDivider from "@/components/BaseDivider.vue";
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import SectionTitle from "@/components/SectionTitle.vue";
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import SectionTitleLineWithButton from "@/components/SectionTitleLineWithButton.vue";
import NotificationBarInCard from "@/components/NotificationBarInCard.vue";
import axios from "axios";
const selectNetworkOptions = [
  { id: 1, label: "TestNet" }
];
const router = useRouter();


const form = reactive({
  accountNumber: "",
  reEnterAccountNumber: "",
  network: selectNetworkOptions[0],
});
const form2 = reactive({
  squareupURL: ""
});

const customElementsForm = reactive({
  checkbox: ["lorem"],
  radio: "one",
  switch: ["one"],
  file: null,
});


async function fetchXrpl() {
  try {
    const response = await axios.get(import.meta.env.VITE_SERVER_URL + "/config?type=XRPL");
    console.log("success", response)
    if (response.status == 200 && response.data.ac_no) {
      form.accountNumber=response.data.ac_no
      form.reEnterAccountNumber=response.data.ac_no
    }else if(response.data.status==500){
      router.push("/")
    }
  } catch (error) {
    console.error('Error fetching data:', error);
    return false
  }
}

/* async function fetchSquareUp() {
  try {
    const response = await axios.get(import.meta.env.VITE_SERVER_URL + "/config?type=SQUAREUP");
    console.log("success", response)
    if (response.status == 200 && response.data.ac_no) {
      form2.squareupURL="https://hello-squareup.coauth.dev/"+response.data.ac_no
    }else if(response.data.status==500){
      router.push("/")
    }
  } catch (error) {
    console.error('Error fetching data:', error);
    return false
  }
}  
 */
const submitXrpl = async() => {
  console.log("submit called")

  if(form.accountNumber!=form.reEnterAccountNumber){
    alert("account numbers dont match");
    return;
  }

  try {
    const response = await axios.get(import.meta.env.VITE_SERVER_URL + "/update_config?type=XRPL&ac_no="+form.accountNumber);
    console.log("success", response)
    if (response.status == 200 && response.data.status) {
      alert("Account number updated")
    }else if(response.data.status==500){
      router.push("/")
    }
  } catch (error) {
    console.error('Error fetching data:', error);
    return false
  }
  //
};
const submitSquare = () => {
  console.log("submit square called")
  //
};


onMounted(() => {
//  fetchSquareUp();
  fetchXrpl();
});

</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <SectionTitleLineWithButton title="Channels Configuration" main>
        <BaseButton :icon="mdiReload" color="whiteDark" />

      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submitXrpl">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="Ripple Ledger(XRPL)">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>

        <FormField label="Enter Your Ripple Account Number" help="XRPL account number starts with 'r' (mainNet coming soon)">
          <FormControl v-model="form.accountNumber" :icon="mdiMail" placeholder="Enter your ripple account number"
            type="password" />
          <FormControl v-model="form.reEnterAccountNumber" type="email" :icon="mdiMail"
            placeholder="Re-enter your ripple account number" />
          <FormControl v-model="form.network" :options="selectNetworkOptions" />
        </FormField>

        <template #footer>
          <BaseButtons>
            <BaseButton type="submit" color="custom" label="Submit" @click="submitXrpl"/>
          </BaseButtons>
        </template>
      </CardBox>
      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
    <!--   <CardBox form @submit.prevent="submitSquare">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="Square Up">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>


        <FormField label="Webhook URL" help="Copy this Webhook URL to your Squareup merchant dashboard">
          <FormControl v-model="form2.squareupURL" type="text" placeholder="Webhook URL" :disabled="true"/>
        </FormField>
        <template #footer>
          <BaseButtons>
            <BaseButton type="button" color="custom" label="Generate / Re-generate" />
          </BaseButtons>
        </template>
      </CardBox> -->
      <CardBox form @submit.prevent="submit">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="Square Up">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>
        Coming soon

      </CardBox>


      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="Paytm">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>
        Coming soon

      </CardBox>
      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="PayPal">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>
        Coming soon

      </CardBox>
      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="Stripe">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>
        Coming soon

      </CardBox>


      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="PhonePe">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>
        Coming soon

      </CardBox>

      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="Google Pay">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>
        Coming soon

      </CardBox>
      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
        <SectionTitleLineWithButton :icon="mdiChartPie" title="Coinbase">
          <BaseButton :icon="mdiReload" color="whiteDark" />
        </SectionTitleLineWithButton>
        Coming soon

      </CardBox>

    </SectionMain>


  </LayoutAuthenticated>
</template>
