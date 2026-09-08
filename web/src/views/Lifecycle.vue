<template>
    <div id="lifecycle">
        <p>{{ num }}</p>
        <button @click="num++">增加</button>
        <button @click="num--">减少</button>
    </div>

</template>

<script setup lang="ts">
import { onBeforeUpdate, ref } from 'vue'

const num = ref(0)

console.log('组件创建')

onBeforeUpdate(() => {
    console.log('组件更新前')
})

</script>



<style scoped>

</style>
