<template>
    <div class="container">
        <AppHeader />
        <n-divider />
        <n-space>
            <n-tag :bordered="false" v-for="(tag, index) in tags" type="success">
                {{ tag.name }}
            </n-tag>
        </n-space>
        <n-divider />
        <AppFooter />
    </div>
</template>

<style lang="scss" scoped>
.container {
    display: flex;
    flex-direction: column;
    width: 900px;
    height: 1000px;
    margin: 0 auto
}
</style>

<script setup>
import { computed, inject, onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

// 路由
const router = useRouter()
const route = useRoute()

const message = inject("message")
const dialog = inject("dialog")
const axios = inject("axios")



// 文章列表
const tags = ref([])


onMounted(() => {

    loadTags()
})

const loadTags = async () => {

    let res = await axios.get(`/api/tag/list`)
    console.log(res)
    tags.value = res.data.data;
}
</script>