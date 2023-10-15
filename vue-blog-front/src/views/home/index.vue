<template>
  <div class="container">
    <AppHeader />

    <n-divider />

    <n-list hoverable clickable>
      <n-list-item v-for="(article, index) in articleListInfo">
        <n-thing :title="article.title" @click="toDetail(article)" content-style="margin-top: 10px;">
          <template #description>
            <n-space size="small" style="margin-top: 4px">
              <n-tag :bordered="false" type="info" size="small">
                暑夜
              </n-tag>
              <n-tag :bordered="false" type="info" size="small">
                晚春
              </n-tag>
            </n-space>
          </template>
          奋勇呀然后休息呀<br>
          完成你伟大的人生
        </n-thing>
      </n-list-item>
    </n-list>
    <n-divider />
    <n-pagination @update:page="loadArticles" v-model:page="pageInfo.page" :page-count="pageInfo.totalPage" />
    <n-divider />
    <AppFooter />
  </div>
</template>

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
const articleListInfo = ref([])

// 查询和分页数据
const pageInfo = reactive({
  page: 1,
  pageSize: 10,
  totalPage: 0,
  total: 0,
  categoryId: 0,
})

onMounted(() => {

  loadArticles()
})

/**
 * 获取文章列表
 */
const loadArticles = async (page = 0) => {
  if (page !== 0) {
    pageInfo.page = page;
  }
  let res = await axios.get(`/api/article/list?keyword=${pageInfo.keyword}&page=${pageInfo.page}&pageSize=${pageInfo.pageSize}&categoryId=${pageInfo.categoryId}`)
  console.log(res)
  let temp_rows = res.data.data.data;
  // 处理获取的文章列表数据
  for (let row of temp_rows) {
    row.content += "..."
    // 把时间戳转换为年月日
    let d = new Date(row.createdAt)
    row.create_time = `${d.getFullYear()}年${d.getMonth() + 1}月${d.getDate()}日`
  }
  articleListInfo.value = temp_rows;
  pageInfo.total = res.data.data.total;
  //计算分页大小
  pageInfo.totalPage = res.data.data.totalPage
  console.log(res)
}

//页面跳转
const toDetail = (article) => {
  router.push({ path: "/detail/" + article.id })
}

</script>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-direction: column;
  width: 900px;
  height: 1000px;
  margin: 0 auto
}
</style>