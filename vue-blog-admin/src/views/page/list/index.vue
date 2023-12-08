<script setup lang="ts">
import { NButton } from 'naive-ui'
import { CrudModal, CrudTable, QueryBarItem, useCRUD } from '@zclzone/crud'
import api from './api'
import { formatDateTime, renderIcon } from '@/utils'
const router = useRouter()
const $table = ref<any>(null)
/** QueryBar筛选参数（可选） */
const queryItems = ref<any>({})
/** 补充参数（可选） */
const extraParams = ref<any>({})

const {
  modalVisible,
  modalAction,
  modalTitle,
  modalLoading,
  handleDelete,
  // handleEdit,
  // handleView,
  handleSave,
  modalForm,
  modalFormRef,
} = useCRUD({
  name: '页面',
  initForm: { author: 'admin' },
  doDelete: api.deletePage,
  refresh: () => $table.value?.handleSearch(),
})

const columns: any = [
  // { type: 'selection', fixed: 'left' },
  // {
  //   title: '发布',
  //   key: 'isPublish',
  //   width: 60,
  //   align: 'center',
  //   fixed: 'left',
  //   render(row: any) {
  //     return h(NSwitch, {
  //       size: 'small',
  //       rubberBand: false,
  //       value: row.isPublish,
  //       loading: !!row.publishing,
  //       onUpdateValue: () => handlePublish(row),
  //     })
  //   },
  // },
  { title: '标题', key: 'title', width: 150, ellipsis: { tooltip: true } },
  { title: '标识', key: 'ident', width: 50, ellipsis: { tooltip: true } },
  {
    title: '创建时间',
    key: 'createDate',
    width: 150,
    render(row: any) {
      return h('span', formatDateTime(row.createDate))
    },
  },
  {
    title: '更新时间',
    key: 'updateDate',
    width: 150,
    render(row: any) {
      return h('span', formatDateTime(row.updateDate))
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 140,
    align: 'center',
    fixed: 'right',
    hideInExcel: true,
    render(row: any) {
      return [
       
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: 'margin-left: 15px;',
            onClick: () => router.replace(`/page/write/${row.id}`), // 携带参数前往 写页面 页面
          },
          { default: () => '编辑', icon: renderIcon('material-symbols:edit-outline', { size: 14 }) },
        ),

        h(
          NButton,
          {
            size: 'small',
            type: 'error',
            style: 'margin-left: 15px;',
            onClick: () => handleDelete(row.id),
          },
          { default: () => '删除', icon: renderIcon('material-symbols:delete-outline', { size: 14 }) },
        ),
      ]
    },
  },
]

// 选中事件
function onChecked(rowKeys: string[]) {
  if (rowKeys.length)
    window.$message?.info(`选中${rowKeys.join(' ')}`)
}

// 发布
// function handlePublish(row: any) {
//   if (isNullOrUndef(row.id))
//     return

//   row.publishing = true
//   setTimeout(() => {
//     row.isPublish = !row.isPublish
//     row.publishing = false
//     window.$message?.success(row.isPublish ? '已发布' : '已取消发布')
//   }, 1000)
// }

onMounted(() => {
  $table.value?.handleSearch()
})
</script>

<template>
  <CommonPage show-footer title="页面">
    <template #action>
      <div>
        <NButton type="primary" secondary @click="$table?.handleExport()">
          <TheIcon icon="mdi:download" :size="18" class="mr-5" /> 导出
        </NButton>
      </div>
    </template>

    <CrudTable ref="$table" v-model:query-items="queryItems" :extra-params="extraParams" :scroll-x="1200"
      :columns="columns" :get-data="api.getPages" @on-checked="onChecked">
      <template #queryBar>
        <QueryBarItem label="标题" :label-width="50">
          <n-input v-model:value="queryItems.title" type="text" placeholder="请输入标题" />
        </QueryBarItem>
      </template>
    </CrudTable>
    <!-- 新增/编辑/查看 -->
    <CrudModal v-model:visible="modalVisible" :title="modalTitle" :loading="modalLoading"
      :show-footer="modalAction !== 'view'" @on-save="handleSave">
      <n-form ref="modalFormRef" label-placement="left" label-align="left" :label-width="80" :model="modalForm"
        :disabled="modalAction === 'view'">
        <n-form-item label="页面标题" path="title" :rule="{
          required: true,
          message: '请输入页面标题',
          trigger: ['input', 'blur'],
        }">
          <n-input v-model:value="modalForm.title" placeholder="请输入页面标题" />
        </n-form-item>
        <n-form-item label="页面内容" path="content" :rule="{
          required: true,
          message: '请输入页面内容',
          trigger: ['input', 'blur'],
        }">
          <n-input v-model:value="modalForm.content" placeholder="请输入页面内容" type="textarea" :autosize="{
            minRows: 3,
            maxRows: 5,
          }" />
        </n-form-item>
      </n-form>
    </CrudModal>
  </CommonPage>
</template>
