<script setup lang="ts">
import { ref } from "vue";
import BudgetPage from "./components/BudgetPage.vue";
import LoginPage from "./components/LoginPage.vue";

const has_logged_in = ref<boolean>(false);
const toggle_has_logged_in = () => {
  has_logged_in.value = !has_logged_in.value;
};
</script>

<template>
  <h1>Budget App</h1>
  <LoginPage v-if="!has_logged_in" @successful_login="toggle_has_logged_in" />
  <BudgetPage v-else />
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
