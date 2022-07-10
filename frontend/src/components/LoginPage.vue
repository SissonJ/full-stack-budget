<script setup lang="ts">
import { ref } from "vue";

const emit = defineEmits(['successful_login'])

const username = ref<string>("");
const password = ref<string>("");
const login_button = () => {
  console.log(username.value);
  console.log(password.value);
  emit('successful_login');
};
</script>
<template>
  <label>Username: </label>
  <input
    v-model="username"
    type="text"
    id="username"
    name="username"
  /><br /><br />
  <label>Password: </label>
  <input
    v-model="password"
    type="text"
    id="password"
    name="password"
  /><br /><br />
  <button @click="login_button">Log in</button>
</template>
<style></style>
