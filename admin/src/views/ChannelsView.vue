<script setup>
import { reactive, ref } from "vue";
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

const selectNetworkOptions = [
  { id: 1, label: "TestNet" }
];

const form = reactive({
  accountNumber: "rDKH6NniQpqoAJNBh4bTf7y9rXigwkyZHa",
  reEnterAccountNumber: "rDKH6NniQpqoAJNBh4bTf7y9rXigwkyZHa",
  network: selectNetworkOptions[0],
});
const form2 = reactive({
  squareupURL: "https://hello-squareup.coauth.dev/6bca6b57-b18b-412c-8bb9-11cfb75c5932"
});

const customElementsForm = reactive({
  checkbox: ["lorem"],
  radio: "one",
  switch: ["one"],
  file: null,
});

const submit = () => {
  //
};

const formStatusWithHeader = ref(true);

const formStatusCurrent = ref(0);

const formStatusOptions = ["info", "success", "danger", "warning"];

const formStatusSubmit = () => {
  formStatusCurrent.value = formStatusOptions[formStatusCurrent.value + 1]
    ? formStatusCurrent.value + 1
    : 0;
};
</script>

<template>
  <LayoutAuthenticated>
    <SectionMain>
      <SectionTitleLineWithButton title="Channels Configuration" main>
        <BaseButton :icon="mdiReload" color="whiteDark" />

      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
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
            <BaseButton type="submit" color="custom" label="Submit" />
          </BaseButtons>
        </template>
      </CardBox>
      <SectionTitleLineWithButton :icon="mdiChartPie" title="" :class="pt - 0">
        <BaseButton :icon="mdiReload" color="whiteDark" />
      </SectionTitleLineWithButton>
      <CardBox form @submit.prevent="submit">
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
