<template>
    <div class="container containerRef">
        <AppHeader />
        <n-divider />
        <main flex-1>
            <n-grid x-gap="15" cols="12" max-w-1200 mt-380 mb-50 mx-auto px-5 class="card-fade-up" responsive="screen"
                item-responsive lg:mt-440>

                <n-gi span="12 m:9">

                    <!-- 标题 -->
                    <n-h1>{{ blogInfo.title }}</n-h1>
                    <!-- 文章内容 -->
                    <MdPreview :editorId="state.id" :modelValue="blogInfo.content" :theme="state.theme" />
                </n-gi>
                <n-gi span="0 m:3">
                    <n-scrollbar style="max-height: 800px" >
                        <MdCatalog :editorId="state.id" :scrollElement="scrollElement" :theme="state.theme" />
                    </n-scrollbar>
                </n-gi>
            </n-grid>

        </main>

        <n-back-top :right="100" />
        <n-divider />
        <AppFooter />
    </div>
</template>

<script setup>
import { ref, reactive, inject, onMounted, computed, defineComponent } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { MdPreview, MdCatalog } from 'md-editor-v3';
// preview.css相比style.css少了编辑器那部分样式
import 'md-editor-v3/lib/preview.css';

const id = 'preview-only';
const state = reactive({
    // theme: 'dark',
    text: 'heading',
    id: 'my-editor'
});
const scrollElement = document.documentElement;

const router = useRouter()
const route = useRoute()
const blogInfo = ref({})
const axios = inject("axios")
const containerRef = ref(void 0);

onMounted(() => {
    loadBlog()
})

/**
 * 读取文章详情
 */
const loadBlog = async () => {
    let res = await axios.get("/api/article/" + route.params.id)
    console.log(res)
    blogInfo.value = res.data.data;
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