<template>
    <button class="cute-btn" :class="[`cute-btn--${type}`]" @click="$emit('click')">
        <span class="cute-btn__text">{{ text }}</span>
    </button>
</template>

<script setup lang="ts">
defineProps<{
    text?: string
    type?: 'primary' | 'success' | 'warning' | 'danger'
}>()

defineEmits<{
    click: []
}>()
</script>

<style scoped>
.cute-btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    padding: 0.5rem 1.25rem;
    border: none;
    border-radius: 999px;
    font-size: 0.9rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.25s cubic-bezier(0.34, 1.56, 0.64, 1);
    position: relative;
    overflow: hidden;
}

.cute-btn:active {
    transform: scale(0.92);
}

.cute-btn::after {
    content: '';
    position: absolute;
    inset: 0;
    border-radius: inherit;
    opacity: 0;
    transition: opacity 0.3s;
    background: linear-gradient(135deg, rgba(255, 255, 255, 0.3), transparent);
}

.cute-btn:hover::after {
    opacity: 1;
}

/* primary */
.cute-btn--primary {
    background: linear-gradient(135deg, #a78bfa, #7c3aed);
    color: #fff;
    box-shadow: 0 4px 14px rgba(124, 58, 237, 0.4);
}

.cute-btn--primary:hover {
    box-shadow: 0 6px 20px rgba(124, 58, 237, 0.6);
    transform: translateY(-2px);
}

/* success */
.cute-btn--success {
    background: linear-gradient(135deg, #6ee7b7, #10b981);
    color: #fff;
    box-shadow: 0 4px 14px rgba(16, 185, 129, 0.4);
}

.cute-btn--success:hover {
    box-shadow: 0 6px 20px rgba(16, 185, 129, 0.6);
    transform: translateY(-2px);
}

/* warning */
.cute-btn--warning {
    background: linear-gradient(135deg, #fcd34d, #f59e0b);
    color: #fff;
    box-shadow: 0 4px 14px rgba(245, 158, 11, 0.4);
}

.cute-btn--warning:hover {
    box-shadow: 0 6px 20px rgba(245, 158, 11, 0.6);
    transform: translateY(-2px);
}

/* danger */
.cute-btn--danger {
    background: linear-gradient(135deg, #fca5a5, #ef4444);
    color: #fff;
    box-shadow: 0 4px 14px rgba(239, 68, 68, 0.4);
}

.cute-btn--danger:hover {
    box-shadow: 0 6px 20px rgba(239, 68, 68, 0.6);
    transform: translateY(-2px);
}
</style>
