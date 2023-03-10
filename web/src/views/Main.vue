<template>
    <div style="flex">
        <div @click="gotoPage('page1')">page1</div>
        <div @click="gotoPage('page2')">page2</div>
    </div>
    <div>
        <router-view></router-view>
    </div>
</template>

<script setup>
import { useRouter, useRoute } from 'vue-router';
const router = useRouter()
const route = useRoute()//获取上层路由传递的参数
function gotoPage(page) {
    router.push({
        path: "/main/" + page
    })
}
</script>

<style scoped></style>