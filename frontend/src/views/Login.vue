<script setup lang="ts">
import { ref } from 'vue'
import { Login } from '../../wailsjs/go/main/App'
import router from '../router'
import { Sleep, SUCESS_CODE } from '../util'
import { myTokenStore } from '../store/myStore'
const username = ref('admin')
const password = ref('123456')
const errorMsg = ref('')
const isLoading = ref(false)

const store = myTokenStore()
const handleLogin = async () => {
    if (!username.value || !password.value) {
        errorMsg.value = '请输入用户名和密码'
        return
    }
    isLoading.value = true
    // 模拟网路操作
    await Sleep(1000)
    try {
        const response = await Login({
            username: username.value,
            password: password.value
        })
        if (response.code === SUCESS_CODE) {
            // 触发路由切换
            // localStorage.setItem('token', 'your_token_here'); // 存储 token
            store.addToken("your_token_here")
            router.push({ name: 'Home' }); // 登录成功后跳转到主页
        } else {
            errorMsg.value = response.message
        }
    } catch (err) {
        errorMsg.value = '连接服务失败'
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
    <div class="login-box">
        <h2>登录</h2>
        <input v-model="username" placeholder="用户名">
        <input v-model="password" type="password" placeholder="密码">
        <button @click="handleLogin" :disabled="isLoading">
            {{ isLoading ? '登录中...' : '登录' }}
        </button>
        <p v-if="errorMsg" class="error">{{ errorMsg }}</p>
    </div>
</template>
