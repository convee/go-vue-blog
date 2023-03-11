<template>
  <div>
    <n-page-header subtitle="" @back="handleBack">
      <n-grid :cols="5">
        <n-gi>
          <n-statistic label="文章" v-model:value="stats.articleCount" />
        </n-gi>
        <n-gi>
          <n-statistic label="分类" v-model:value="stats.categoryCount" />
        </n-gi>
        <n-gi>
          <n-statistic label="标签" v-model:value="stats.tagCount" />
        </n-gi>
        <n-gi>
          <n-statistic label="页面"  v-model:value="stats.pageCount" />
        </n-gi>
      </n-grid>
      <template #title>
        <a
            href="https://convee.cn/"
            style="text-decoration: none; color: inherit"
        >Go vue blog</a>
      </template>
    </n-page-header>
  </div>
</template>

<script setup>

import {AdminStore} from '../../stores/AdminStore'
import {inject, onMounted, reactive, ref} from 'vue'
import {useRoute, useRouter} from 'vue-router'

const router = useRouter()
const route = useRoute()

const message = inject("message")
const dialog = inject("dialog")
const axios = inject("axios")

const adminStore = AdminStore()
const stats = reactive({
  articleCount:0,
  tagCount:0,
  categoryCount:0,
  pageCount:0
})


onMounted(() => {
  loadStats()
})

const loadStats = async () => {
  let res = await axios.get("/backend/article/stat")
  console.log(res)
  stats.articleCount = res.data.data.articleCount
  stats.tagCount = res.data.data.tagCount
  stats.categoryCount = res.data.data.categoryCount
  stats.pageCount = res.data.data.pageCount
}

</script>

<style lang="scss" scoped>
</style>