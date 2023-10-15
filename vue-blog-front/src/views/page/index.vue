<template>
    <div class="container containerRef">
        <AppHeader />
        <n-divider />
        <main flex-1>
            <MdPreview :editorId="data.id" :modelValue="data.content" :theme="state.theme" />
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

const state = reactive({
    // theme: 'dark',
    text: 'heading',
    id: 'my-editor'
});

const router = useRouter()
const route = useRoute()
const data = ref({})
const axios = inject("axios")

onMounted(() => {
    loadPage()
})

/**
 * 页面详情
 */
const loadPage = async () => {
    let res = await axios.get("/api/page/" + route.params.ident)
    console.log(res)
    data.value = res.data.data;
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