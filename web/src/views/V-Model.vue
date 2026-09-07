<template>
    <div>
        input: <input v-model="inputValue" type="text" placeholder="请输入内容" />
        <p>输入的内容是: {{ inputValue }}</p>
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'

const inputValue = ref('')
</script>