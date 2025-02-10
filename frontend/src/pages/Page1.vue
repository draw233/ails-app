<script lang="ts" setup>
import router from '../router';
 const comeBackCompany = () => {
  router.back();
}
</script>

<template>
      <button  @click="comeBackCompany">回退</button>
    Page1
</template>

<style scoped>
</style>
