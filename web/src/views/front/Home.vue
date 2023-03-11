<template>
  <div class="container">
    <div class="nav">
      <div @click="homePage">首页</div>
      <div>
        <n-popselect @update:value="searchByCategory" v-model:value="selectedCategory" :options="categoryOptions"
                     trigger="click">
          <div>分类<span>{{ categoryName }}</span></div>
        </n-popselect>
      </div>
      <div @click="tags">标签</div>
      <div @click="about">关于</div>


    </div>
    <n-divider/>

    <div>
      <n-space class="search">
        <n-input v-model:value="pageInfo.keyword" :style="{ width: '300px'}" placeholder="请输入关键字"/>
        <n-button type="primary" ghost @click="loadArticles(0)"> 搜索</n-button>
      </n-space>
    </div>
    <div v-for="(article, index) in articleListInfo" style="margin-bottom:15px;cursor: pointer;">
      <n-card :title="article.title" @click="toDetail(article)">
        {{ article.content }}

        <template #footer>
          <n-space align="center">
            <div>发布时间：{{ article.create_time }}</div>
          </n-space>
        </template>
      </n-card>
    </div>

    <n-pagination @update:page="loadArticles" v-model:page="pageInfo.page" :page-count="pageInfo.totalPage"/>

    <n-divider/>
    <div class="footer">
      <div>Copyright © 2023.convee All rights reserved</div>
    </div>
  </div>
</template>

<script setup>
import {computed, inject, onMounted, reactive, ref} from 'vue'
import {useRoute, useRouter} from 'vue-router'

// 路由
const router = useRouter()
const route = useRoute()

const message = inject("message")
const dialog = inject("dialog")
const axios = inject("axios")

// 选中的分类
const selectedCategory = ref(0)
// 分类选项
const categoryOptions = ref([])
// 文章列表
const articleListInfo = ref([])

// 查询和分页数据
const pageInfo = reactive({
  page: 1,
  pageSize: 10,
  totalPage: 0,
  total: 0,
  keyword: "",
  categoryId: 0,
})

onMounted(() => {
  loadCategories();
  loadArticles()
})

/**
 * 获取博客列表
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

const categoryName = computed(() => {
  //获取选中的分类
  let selectedOption = categoryOptions.value.find((option) => {
    return option.value === selectedCategory.value
  })
  //返回分类的名称
  return selectedOption ? selectedOption.label : ""
})

/**
 * 获取分类列表
 */
const loadCategories = async () => {
  let res = await axios.get("/api/category/list")
  console.log(res)
  categoryOptions.value = res.data.data.map((item) => {
    return {
      label: item.name,
      value: item.id
    }
  })
  console.log(categoryOptions.value)
}

/**
 * 选中分类
 */
const searchByCategory = (categoryId) => {
  pageInfo.categoryId = categoryId;
  loadArticles()
}

//页面跳转
const toDetail = (article) => {
  router.push({path: "/detail", query: {id: article.id}})
}

const homePage = () => {
  router.push("/")
}

const tags = () => {
  router.push("/tags")
}

const about = () => {
  router.push("/about")
}

</script>

<style lang="scss" scoped>

.search {
  margin-bottom: 15px;
}

.container {
  width: 1200px;
  margin: 0 auto;
}

.nav {
  display: flex;
  font-size: 20px;
  padding-top: 20px;
  color: #64676a;

  div {
    cursor: pointer;
    margin-right: 15px;

    &:hover {
      color: #f60;
    }

    span {
      font-size: 12px;
    }
  }
}

.footer {
  text-align: center;
  line-height: 25px;
  color: #64676a;
}
</style>