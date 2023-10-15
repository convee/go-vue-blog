<template>
    <div class="nav">

        <n-space>
            <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions" />
            <n-input-group>
                <n-input :style="{ width: '70%' }" v-model:value="keyword" placeholder="请输入关键字" />
                <n-button type="primary" ghost @click="loadArticles(0)">搜索</n-button>
            </n-input-group>
        </n-space>
    </div>
</template>

<script setup>
import { computed, inject, onMounted, reactive, ref, defineComponent, h } from 'vue';
import { useRoute, useRouter, RouterLink } from 'vue-router';
import { NIcon } from "naive-ui";
import {
    PersonOutline as PersonIcon,
    HomeOutline as HomeIcon,
    MenuOutline as MenuIcon,
    PricetagsOutline as PricetagsIcon,
} from "@vicons/ionicons5";

const activeKey = ref(null)
function renderIcon(icon) {
    return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions = [
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    path: "/"
                }
            },
            { default: () => "首页" }
        ),
        key: "home",
        icon: renderIcon(HomeIcon)
    },
    {
        label: "分类",
        key: "category",
        icon: renderIcon(MenuIcon),
        children: [
            {
                label: "分类1",
                key: "category1",
            },
            {
                label: "分类2",
                key: "category2",
            }
        ]
    },
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    path: "/tags"
                }
            },
            { default: () => "标签" }
        ),
        key: "tags",
        icon: renderIcon(PricetagsIcon)
    }
    ,
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    path: "/page/about"
                }
            },
            { default: () => "关于我" }
        ),
        key: "about",
        icon: renderIcon(PersonIcon)
    }
];


const axios = inject("axios")
const router = useRouter()

// 选中的分类
const selectedCategory = ref(0)
// 分类选项
const categoryOptions = ref([])
// 搜索关键词
const keyword = ref("")

const categoryName = computed(() => {
    //获取选中的分类
    let selectedOption = categoryOptions.value.find((option) => {
        return option.value === selectedCategory.value
    })
    //返回分类的名称
    return selectedOption ? selectedOption.label : ""
})

onMounted(() => {
    loadCategories();
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

// 文章列表
const articleListInfo = ref([])
// 查询和分页数据
const pageInfo = reactive({
    page: 1,
    pageSize: 10,
    totalPage: 0,
    total: 0,
    categoryId: 0,
})
/**
 * 选中分类
 */
const searchByCategory = (categoryId) => {
    pageInfo.categoryId = categoryId;
    loadArticles()
}


/**
 * 获取文章列表
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

</script>

<style lang="scss" scoped>
.nav {
    display: flex;
    padding-top: 16px;
}
</style>