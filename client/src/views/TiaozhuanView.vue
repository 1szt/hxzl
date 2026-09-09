<template>
    <div>
        <h1>这是 / 根页面</h1>
    </div>

    <div class="menu">
        <router-link to="/">标签跳转</router-link>
        <router-link to="/" target="_blank">在新标签页中打开</router-link>


        <button @click="goToHome">编程式跳转</button>

        <button @click="goToBack">返回</button>\


        <button @click="goToRepeat">替换 不可返回</button>


    </div>
</template>
<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const goToHome = () => {
    router.push('/')
}

const goToBack = () => {
    router.go(-1)
}

const goToRepeat = () => {
    router.replace('/')
}

</script>