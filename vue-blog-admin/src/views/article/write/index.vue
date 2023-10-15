<script setup lang="ts">
import api from './api'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
const [router, route] = [useRouter(), useRoute()]
const categoryOptions = ref([]) // 分类选项
const tagOptions = ref([]) // 分类选项
interface ArticleForm {
  id: any
  title: string
  content: string
  categoryId: any
  tagIds: any
}
const articleForm = ref<ArticleForm>({
  id: null,
  title: '',
  content: '',
  categoryId: null,
  tagIds: null,
})

const btnLoading = ref<boolean>(false)

onMounted(async () => {
  api.getCategoryOption().then((resp: any) => {
    console.log(resp)
    if (resp.code == 0) {
      categoryOptions.value = resp.data.map((item: any) => {
        return {
          label: item.name,
          value: item.id,
        };
      });
    } else {
      categoryOptions.value = [];
    }
  });
  api.getTagOption().then((resp: any) => {
    console.log(resp)
    if (resp.code == 0) {
      tagOptions.value = resp.data.map((item: any) => {
        return {
          label: item.name,
          value: item.id,
        };
      });
    } else {
      tagOptions.value = [];
    }
  });
  await getArticleInfo()

})
async function handleSaveArticle() {
  const { id, title, content, categoryId, tagIds } = articleForm.value
  if (!title) {
    window.$message?.warning('请输入文章标题')
    return
  }
  if (!content) {
    window.$message?.warning('请输入文章内容')
    return
  }
  if (!categoryId) {
    window.$message?.warning('请选择文章分类')
    return
  }
  if (tagIds.length == 0) {
    window.$message?.warning('请选择文章标签')
    return
  }
  btnLoading.value = true
  try {
    if (id) {
      const res: any = await api.updateArticle({ id, title, content, categoryId, tagIds })
      console.log(res)
    } else {
      const res: any = await api.addArticle({ title, content, categoryId, tagIds })
      console.log(res)
    }

    window.$notification?.success({ title: '保存成功！', duration: 2500 })
    router.push('/article/list')
  } catch (error) {
    console.error(error)
  }
  btnLoading.value = false
}

async function getArticleInfo() {
  const id = route.params.id

  if (!id) {
    articleForm.value = { id: null, title: '', content: '', categoryId: null, tagIds: null }
    return
  }

  window.$loadingBar?.start()

  try {
    const res = await api.getArticleById(id)
    console.log(res)
    articleForm.value = res.data

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
  <CommonPage :show-header="false" show-footer title="写文章">
    <n-form ref="modalFormRef" label-placement="left" label-align="left" :label-width="80" :model="articleForm">
      <n-input v-model:value="articleForm.id" hidden />
      <div class="mb-10 pl-20 pr-20 items-center bg-white" style="width: 400px">
        <n-form-item label="文章标题" path="title" :rule="{
          required: true,
          message: '请输入文章标题',
          trigger: ['change', 'blur'],
        }">
          <n-input v-model:value="articleForm.title" placeholder="请输入文章标题" />
        </n-form-item>

      </div>
      <div class="mb-10 pl-20 pr-20  items-center bg-white" style="width: 300px">
        <n-form-item label="文章分类" path="categoryId" :rule="{
          required: true,
          message: '请输入文章分类',
          //trigger: ['input', 'blur'],
        }">
          <n-select v-model:value="articleForm.categoryId" clearable filterable placeholder="文章分类"
            :options="categoryOptions" />
        </n-form-item>
      </div>
      <div class="mb-10 pl-20 pr-20  items-center bg-white" style="width: 300px">
        <n-form-item label="文章标签" path="tagIds" :rule="{
          required: true,
          message: '请输入文章标签',
          //trigger: ['input', 'blur'],
        }">
          <n-select v-model:value="articleForm.tagIds" filterable clearable :options="tagOptions" multiple
            placeholder="标签名称" />
        </n-form-item>
      </div>

      <div class="mb-10 pl-20 pr-20 flex items-center bg-white">
        <MdEditor v-model="articleForm.content" style="height: calc(100vh - 305px)" />
      </div>

      <div class="mb-10 pl-20 pr-20 flex items-center bg-white">
        <n-button type="primary" style="width: 80px" :loading="btnLoading" @click="handleSaveArticle">
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
