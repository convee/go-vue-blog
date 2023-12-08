<script setup lang="ts">
import { NButton } from 'naive-ui'
import { CrudModal, CrudTable, QueryBarItem, useCRUD } from '@zclzone/crud'
import api from './api'
import { formatDateTime, renderIcon } from '@/utils'

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
  handleAdd,
  handleDelete,
  handleEdit,
  handleView,
  handleSave,
  modalForm,
  modalFormRef,
} = useCRUD({
  name: '分类',
  doCreate: api.addCategory,
  doDelete: api.deleteCategory,
  doUpdate: api.updateCategory,
  refresh: () => $table.value?.handleSearch(),
})

const columns: any = [
  { type: 'selection', fixed: 'left' },
  { title: 'ID', key: 'id', width: 80, ellipsis: { tooltip: true } },
  { title: '名称', key: 'name', width: 150, ellipsis: { tooltip: true } },
  {
    title: '创建时间',
    key: 'createdAt',
    width: 150,
    render(row: any) {
      return h('span', formatDateTime(row.createdAt))
    },
  },
  {
    title: '最后更新时间',
    key: 'updatedAt',
    width: 150,
    render(row: any) {
      return h('span', formatDateTime(row.updatedAt))
    },
  },
  {
    title: '操作',
    key: 'actions',
    width: 240,
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
            secondary: true,
            onClick: () => handleView(row),
          },
          { default: () => '查看', icon: renderIcon('majesticons:eye-line', { size: 14 }) },
        ),
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            style: 'margin-left: 15px;',
            onClick: () => handleEdit(row),
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

onMounted(() => {
  $table.value?.handleSearch()
})
</script>

<template>
  <CommonPage show-footer title="分类">
    <template #action>
      <div>
        <NButton type="primary" class="ml-16" @click="handleAdd">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" /> 新建分类
        </NButton>
      </div>
    </template>

    <CrudTable ref="$table" v-model:query-items="queryItems" :extra-params="extraParams" :scroll-x="1200"
      :columns="columns" :get-data="api.getCategory" @on-checked="onChecked">
      <template #queryBar>
        <QueryBarItem label="名称" :label-width="50">
          <n-input v-model:value="queryItems.name" type="text" placeholder="请输入名称"/>
        </QueryBarItem>
      </template>
    </CrudTable>
    <!-- 新增/编辑/查看 -->
    <CrudModal v-model:visible="modalVisible" :title="modalTitle" :loading="modalLoading"
      :show-footer="modalAction !== 'view'" @on-save="handleSave">
      <n-form ref="modalFormRef" label-placement="left" label-align="left" :label-width="80" :model="modalForm"
        :disabled="modalAction === 'view'">
        <n-form-item label="分类名称" path="name" :rule="{
          required: true,
          message: '请输入分类名称',
          trigger: ['input', 'blur'],
        }">
          <n-input v-model:value="modalForm.name" placeholder="请输入分类名称" />
        </n-form-item>
      </n-form>
    </CrudModal>
  </CommonPage>
</template>
