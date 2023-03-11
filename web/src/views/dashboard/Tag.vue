<template>
  <div>
    <n-table :bordered="false" :single-line="false">
      <thead>
      <tr>
        <th>编号</th>
        <th>名称</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(tag, index) in tagList">
        <td>{{ tag.id }}</td>
        <td>{{ tag.name }}</td>
      </tr>

      </tbody>
    </n-table>

    <n-modal v-model:show="showAddModel" preset="dialog" title="Dialog">
      <template #header>
        <div>添加分类</div>
      </template>
      <div>
        <n-input v-model:value="addCategory.name" type="text" placeholder="请输入名称"/>
      </div>
      <template #action>
        <div>
          <n-button @click="add">提交</n-button>
        </div>
      </template>
    </n-modal>

    <n-modal v-model:show="showUpdateModel" preset="dialog" title="Dialog">
      <template #header>
        <div>修改分类</div>
      </template>
      <div>
        <n-input v-model:value="updateCategory.name" type="text" placeholder="请输入名称"/>
      </div>
      <template #action>
        <div>
          <n-button @click="update">提交</n-button>
        </div>
      </template>
    </n-modal>


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

const showAddModel = ref(false)
const showUpdateModel = ref(false)

const tagList = ref([])
const addCategory = reactive({
  name: ""
})

const updateCategory = reactive({
  id: 0,
  name: ""
})

onMounted(() => {
  loadTags()
})

const loadTags = async () => {
  let res = await axios.get("/backend/tag/list?page=1&pageSize=1000")
  tagList.value = res.data.data.data
}

</script>

<style lang="scss" scoped>
</style>