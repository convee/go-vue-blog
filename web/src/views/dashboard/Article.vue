<template>
  <n-tabs v-model:value="tabValue" justify-content="start" type="line">
    <n-tab-pane name="list" tab="文章列表">
      <div v-for="(article, index) in articleList" style="margin-bottom:15px">
        <n-card :title="article.title">
          {{ article.content }}

          <template #footer>
            <n-space align="center">
              <div>发布时间：{{ article.createdAt }}</div>
              <n-button @click="toUpdate(article)">修改</n-button>
              <n-button @click="toDelete(article)">删除</n-button>
            </n-space>
          </template>
        </n-card>
      </div>

      <n-space>
        <n-pagination @update:page="loadArticles" v-model:page="pageInfo.page" :page-count="pageInfo.totalPage"/>
      </n-space>

    </n-tab-pane>
    <n-tab-pane name="add" tab="添加文章">

      <n-form>
        <n-form-item label="标题">
          <n-input v-model:value="addArticle.title" placeholder="请输入标题"/>
        </n-form-item>
        <n-form-item label="分类">
          <n-select v-model:value="value" :options="categoriesOptions" placeholder="请选择分类"/>
        </n-form-item>
        <n-form-item label="内容">
          <rich-text-editor v-model="addArticle.content"></rich-text-editor>
        </n-form-item>
        <n-form-item label="">
          <n-button @click="add">提交</n-button>
        </n-form-item>
      </n-form>

    </n-tab-pane>
    <n-tab-pane name="update" tab="修改">
      <n-form>
        <n-form-item label="标题">
          <n-input v-model:value="updateArticle.title" placeholder="请输入标题"/>
        </n-form-item>
        <n-form-item label="分类">
          <n-select v-model:value="updateArticle.categoryId" :options="categoriesOptions"/>
        </n-form-item>
        <n-form-item label="内容">
          <rich-text-editor v-model="updateArticle.content"></rich-text-editor>
        </n-form-item>
        <n-form-item label="">
          <n-button @click="update">提交</n-button>
        </n-form-item>
      </n-form>
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import {AdminStore} from '../../stores/AdminStore'
import {inject, onMounted, reactive, ref} from 'vue'
import {useRoute, useRouter} from 'vue-router'
import RichTextEditor from '../../components/RichTextEditor.vue'

const router = useRouter()
const route = useRoute()

const message = inject("message")
const dialog = inject("dialog")
const axios = inject("axios")

const adminStore = AdminStore()

//文章添加数据
const addArticle = reactive({
  categoryId: 0,
  title: "",
  content: "",
})

//文章修改数据
const updateArticle = reactive({
  id: 0,
  categoryId: 0,
  title: "",
  content: "",
})

//分类选项
const categoriesOptions = ref([])
const articleList = ref([])
//标签页
const tabValue = ref("list")

//分页数据
const pageInfo = reactive({
  page: 1,
  pageSize: 3,
  totalPage: 0,
  total: 0,
})

onMounted(() => {
  loadArticles()
  loadCategories()
})

//读取博客列表
const loadArticles = async (page = 0) => {
  if (page !== 0) {
    pageInfo.page = page;
  }
  let res = await axios.get(`/backend/article/list?page=${pageInfo.page}&pageSize=${pageInfo.pageSize}`)
  console.log(res)
  let temp_rows = res.data.data.data;
  for (let row of temp_rows) {
    row.content += "..."
    let d = new Date(row.createdAt)
    row.createdAt = `${d.getFullYear()}年${d.getMonth() + 1}月${d.getDate()}日`
  }
  articleList.value = temp_rows;
  pageInfo.total = res.data.data.total;
  pageInfo.totalPage = res.data.data.totalPage
  console.log(res)
}

//读取分类
const loadCategories = async () => {
  let res = await axios.get("/backend/category/list?page=1&pageSize=1000")
  console.log(res)
  categoriesOptions.value = res.data.data.data.map((item) => {
    return {
      label: item.name,
      value: item.id
    }
  })
}

const add = async () => {
  let res = await axios.post("/backend/article/add", addArticle)
  if (res.data.code === 0) {
    message.info(res.data.msg)
  } else {
    message.error(res.data.msg)
  }
}

const toUpdate = async (article) => {
  tabValue.value = "update"
  let res = await axios.get("/backend/article/detail?id=" + article.id)
  updateArticle.id = article.id
  updateArticle.title = res.data.data.title
  updateArticle.content = res.data.data.content
  updateArticle.categoryId = res.data.data.categoryId
}

const update = async () => {
  let res = await axios.post("/backend/article/update", updateArticle)
  if (res.data.code === 0) {
    message.info(res.data.msg)
    await loadArticles()
    tabValue.value = "list"
  } else {
    message.error(res.data.msg)
  }
}

const toDelete = async (article) => {
  let res = await axios.post("/backend/article/delete", {id: article.id})
  if (res.data.code === 0) {
    message.info(res.data.msg)
    await loadArticles()
  } else {
    message.error(res.data.msg)
  }
}

</script>

<style lang="scss" scoped>
</style>