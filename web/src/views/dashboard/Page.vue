<template>
  <n-tabs v-model:value="tabValue" justify-content="start" type="line">
    <n-tab-pane name="list" tab="页面列表">
      <div v-for="(page, index) in pageList" style="margin-bottom:15px">
        <n-card :title="page.title">
          {{ page.content }}

          <template #footer>
            <n-space align="center">
              <div>发布时间：{{ page.createdAt }}</div>
              <n-button @click="toUpdate(page)">修改</n-button>
              <n-button @click="toDelete(page)">删除</n-button>
            </n-space>
          </template>
        </n-card>
      </div>

      <n-space>
        <n-pagination @update:page="loadPages" v-model:page="pageInfo.page" :page-count="pageInfo.totalPage"/>
      </n-space>

    </n-tab-pane>
    <n-tab-pane name="add" tab="添加页面">

      <n-form>
        <n-form-item label="标题">
          <n-input v-model:value="addPage.title" placeholder="请输入标题"/>
        </n-form-item>
        <n-form-item label="标识">
          <n-input v-model:value="addPage.ident" placeholder="请输入标识"/>
        </n-form-item>
        <n-form-item label="内容">
          <rich-text-editor v-model="addPage.content"></rich-text-editor>
        </n-form-item>
        <n-form-item label="">
          <n-button @click="add">提交</n-button>
        </n-form-item>
      </n-form>

    </n-tab-pane>
    <n-tab-pane name="update" tab="修改">
      <n-form>
        <n-form-item label="标题">
          <n-input v-model:value="updatePage.title" placeholder="请输入标题"/>
        </n-form-item>
        <n-form-item label="标识">
          <n-input v-model:value="addPage.ident" placeholder="请输入标识"/>
        </n-form-item>
        <n-form-item label="内容">
          <rich-text-editor v-model="updatePage.content"></rich-text-editor>
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

//页面添加数据
const addPage = reactive({
  ident: "",
  title: "",
  content: "",
})

//页面修改数据
const updatePage = reactive({
  id: 0,
  ident: "",
  title: "",
  content: "",
})

//分类选项
const categoriesOptions = ref([])
const pageList = ref([])
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
  loadPages()
})

//读取博客列表
const loadPages = async (page = 0) => {
  if (page !== 0) {
    pageInfo.page = page;
  }
  let res = await axios.get(`/backend/page/list?page=${pageInfo.page}&pageSize=${pageInfo.pageSize}`)
  console.log(res)
  let temp_rows = res.data.data.data;
  for (let row of temp_rows) {
    row.content += "..."
    let d = new Date(row.createdAt)
    row.createdAt = `${d.getFullYear()}年${d.getMonth() + 1}月${d.getDate()}日`
  }
  pageList.value = temp_rows;
  pageInfo.total = res.data.data.total;
  pageInfo.totalPage = res.data.data.totalPage
  console.log(res)
}

const add = async () => {
  let res = await axios.post("/backend/page/add", addPage)
  if (res.data.code === 0) {
    message.info(res.data.msg)
  } else {
    message.error(res.data.msg)
  }
}

const toUpdate = async (page) => {
  tabValue.value = "update"
  let res = await axios.get("/backend/page/detail?id=" + page.id)
  updatePage.id = page.id
  updatePage.title = res.data.data.title
  updatePage.content = res.data.data.content
  updatePage.ident = res.data.data.ident
}

const update = async () => {
  let res = await axios.post("/backend/page/update", updatePage)
  if (res.data.code === 0) {
    message.info(res.data.msg)
    await loadPages()
    tabValue.value = "list"
  } else {
    message.error(res.data.msg)
  }
}

const toDelete = async (page) => {
  let res = await axios.post("/backend/page/delete", {id: page.id})
  if (res.data.code === 0) {
    message.info(res.data.msg)
    await loadPages()
  } else {
    message.error(res.data.msg)
  }
}

</script>

<style lang="scss" scoped>
</style>