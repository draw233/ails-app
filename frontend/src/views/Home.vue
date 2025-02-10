<template>
    <main>
        <div class="home">
            <h1>欢迎回来, {{ username }}!</h1>
            <button @click="logout">退出登录</button>
        </div>
        <nav>
            <RouterLink to="page1">page1</RouterLink>
            <RouterLink to="page2">page2</RouterLink>
        </nav>
    </main>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import router from '../router';
import { myTokenStore } from '../store/myStore'
const store = myTokenStore()
const username = ref('admin') // 实际应从状态管理获取
const logout = () => {
    // localStorage.removeItem('token'); // 清除 token
    store.removeToken()
    router.push('/login')
}
</script>