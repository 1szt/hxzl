<template>
  <div class="shop">
    <h1>我的vue奶茶店</h1>
    <p>剩余奶茶杯数: {{ remainingMilkTeaCups }}</p>
    <button @click="goToShop">奶茶-1</button>
  </div>
</template>


<style scoped>
h1 {
  color: #42b983;
}

.shop {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>


<script setup lang="ts">
import { ref } from 'vue';

const remainingMilkTeaCups = ref(10);
function goToShop() {
  if (remainingMilkTeaCups.value > 0) {
    remainingMilkTeaCups.value--;
  } else {
    alert('奶茶已售罄！');
  }
}
</script>