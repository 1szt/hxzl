<template>
    <div class="page">
        <div class="lang-bar">
            <a v-for="lang in languages" :key="lang.code" :href="lang.url" class="mx-2 p-1 label label-secondary">
                {{ lang.name }}
            </a>
        </div>

        <div class="hero">
            <div class="hero-body">
                <h1 class="text-center">
                    <a class="text-primary" href="https://github.com/bs-community/blessing-skin-server">
                        mcup-server
                    </a>
                </h1>
                <div class="divider"></div>
                <div class="content-center-left">
                    <h3>欢迎</h3>
                    <div class="divider"></div>
                    <p>欢迎使用 mcup-server v0.0.1！</p>
                </div>
                <p class="step">
                    <a href="https://skin.1szt.com/setup/database" class="btn btn-primary">
                        下一步
                    </a>
                </p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
const languages = [
    { code: 'zh_CN', name: '中文 (简体)', url: 'https://skin.1szt.com/setup?lang=zh_CN' },
    { code: 'zh_TW', name: '中文 (正體)', url: 'https://skin.1szt.com/setup?lang=zh_TW' },
    { code: 'en', name: 'English', url: 'https://skin.1szt.com/setup?lang=en' },
    { code: 'es_ES', name: 'Español', url: 'https://skin.1szt.com/setup?lang=es_ES' },
    { code: 'ru_RU', name: 'Русский язык', url: 'https://skin.1szt.com/setup?lang=ru_RU' },
]
</script>

<style scoped>
.page {
    background: #f1f1fc;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
}

.lang-bar {
    text-align: right;
    margin: 0.4rem;
    padding: 0.4rem;
}

.mx-2 {
    margin-left: 0.4rem !important;
    margin-right: 0.4rem !important;
}

.p-1 {
    padding: 0.2rem !important;
}

.label {
    background: #eef0f3;
    border-radius: 0.1rem;
    color: #455060;
    display: inline-block;
    line-height: 1.25;
    padding: 0.1rem 0.2rem;
}

.label.label-secondary {
    background: #f1f1fc;
    color: #5755d9;
}

.hero {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 2rem;
}

.hero-body {
    max-width: 480px;
    margin: 0 auto;
    padding: 2rem 0;
}

.text-center {
    text-align: center;
}

.content-center-left {
    text-align: center;
    padding-left: 1rem;
}

.divider {
    border-top: 0.05rem solid #f1f3f5;
    display: block;
    height: 0.05rem;
    margin: 0.4rem 0;
    position: relative;
}

.text-primary {
    color: #5755d9 !important;
}

a.text-primary:hover {
    color: #4240d4;
}

a.text-primary:visited {
    color: #6c6ade;
}

.step {
    padding: 0.2rem;
    text-align: center;
    margin-top: 1rem;
}

.btn {
    -webkit-appearance: none;
    appearance: none;
    background: #fff;
    border: 0.05rem solid #5755d9;
    border-radius: 0.1rem;
    color: #5755d9;
    cursor: pointer;
    display: inline-block;
    font-size: 0.8rem;
    height: 1.8rem;
    line-height: 1.2rem;
    outline: 0;
    padding: 0.25rem 0.4rem;
    text-align: center;
    text-decoration: none;
    transition: background 0.2s, border 0.2s, box-shadow 0.2s, color 0.2s;
    user-select: none;
    vertical-align: middle;
    white-space: nowrap;
}

.btn:focus {
    box-shadow: 0 0 0 0.1rem rgba(87, 85, 217, 0.2);
}

.btn:focus,
.btn:hover {
    background: #f1f1fc;
    border-color: #4b48d6;
    text-decoration: none;
}

.btn.btn-primary {
    background: #5755d9;
    border-color: #4b48d6;
    color: #fff;
}

.btn.btn-primary:focus,
.btn.btn-primary:hover {
    background: #4240d4;
    border-color: #3634d2;
    color: #fff;
}

.btn.btn-primary:active {
    background: #3a38d2;
    border-color: #302ecd;
    color: #fff;
}
</style>
