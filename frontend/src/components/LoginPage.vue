<script setup lang="ts">
import { ref } from "vue";

const emit = defineEmits(['successful_login'])

// type Password_Res = {
//     status: string
// }

const username = ref<string>("");
const password = ref<string>("");
const login_button = async () => {
  console.log(username.value);
  console.log(password.value);
  const res = await fetch("http://localhost:3000/api/password", {
    method: "POST",
    body: JSON.stringify({
    user: username.value,
    password: password.value
  })})
  const data = await res.json();
  console.log(data)
  if( data.status === "true" ){
    emit('successful_login');
  }
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
