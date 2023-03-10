<template>
    <div>
        <button @click="testClick()">clickMe</button>
    </div>
</template>

<script setup>
import { useRouter, useRoute } from 'vue-router';
const router = useRouter()
const route = useRoute()//获取上层路由传递的参数
function testClick() {
    router.push({
        path: '/login',
        query: {
            name: "张三",
            id: 1
        }
    })
}
</script>

<style scoped></style>