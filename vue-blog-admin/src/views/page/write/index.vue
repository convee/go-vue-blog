<script setup lang="ts">
import api from './api'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
const [router, route] = [useRouter(), useRoute()]
interface PageForm {
  id: any
  title: string
  ident: string
  content: string
}
const pageForm = ref<PageForm>({
  id: null,
  title: '',
  ident: '',
  content: '',
})

const btnLoading = ref<boolean>(false)

onMounted(async () => {
  await getPageInfo()
})
async function handleSavePage() {
  const { id, title, ident, content } = pageForm.value
  if (!title) {
    window.$message?.warning('请输入页面标题')
    return
  }
  if (!ident) {
    window.$message?.warning('请输入页面标识')
    return
  }
  if (!content) {
    window.$message?.warning('请输入页面内容')
    return
  }
  btnLoading.value = true
  try {
    if (id) {
      const res: any = await api.updatePage({ id, title, ident, content })
      console.log(res)
    } else {
      const res: any = await api.addPage({ title, ident, content })
      console.log(res)
    }
    window.$notification?.success({ title: '保存成功！', duration: 2500 })
    router.push('/page/list')
  } catch (error) {
    console.error(error)
  }
  btnLoading.value = false
}

async function getPageInfo() {
  const id = route.params.id

  if (!id) {
    pageForm.value = { id: null, title: '', ident: '', content: '' }
    return
  }

  window.$loadingBar?.start()

  try {
    const res = await api.getPageById(id)
    console.log(res)
    pageForm.value = res.data

    window.$loadingBar?.finish()
    window.$message?.success('加载成功')
  }
  catch (err) {
    window.$loadingBar?.error()
    window.$message?.error('加载失败')
  }
}
</script>

<template>
  <CommonPage :show-header="false" show-footer title="新增页面">
    <n-form ref="modalFormRef" label-placement="left" label-align="left" :label-width="80" :model="pageForm">
      <n-input v-model:value="pageForm.id" hidden />
      <div class="mb-10 pl-20 pr-20 items-center bg-white" style="width: 400px">
        <n-form-item label="页面标题" path="title" :rule="{
          required: true,
          message: '请输入页面标题',
          trigger: ['change', 'blur'],
        }">
          <n-input v-model:value="pageForm.title" placeholder="请输入页面标题" />
        </n-form-item>

      </div>
      <div class="mb-10 pl-20 pr-20 items-center bg-white" style="width: 400px">
        <n-form-item label="页面标识" path="ident" :rule="{
          required: true,
          message: '请输入页面标识',
          trigger: ['change', 'blur'],
        }">
          <n-input v-model:value="pageForm.ident" placeholder="请输入页面标识" />
        </n-form-item>

      </div>
      <div class="mb-10 pl-20 pr-20 flex items-center bg-white">
        <MdEditor v-model="pageForm.content" style="height: calc(100vh - 305px)" />
      </div>

      <div class="mb-10 pl-20 pr-20 flex items-center bg-white">
        <n-button type="primary" style="width: 80px" :loading="btnLoading" @click="handleSavePage">
          保存
        </n-button>
      </div>
    </n-form>

  </CommonPage>
</template>

<style lang="scss" scoped>
.md-preview {

  ul,
  ol {
    list-style: revert;
  }
}
</style>
