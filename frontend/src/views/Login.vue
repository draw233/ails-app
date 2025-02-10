<script setup lang="ts">
import { ElCheckbox, ElForm, ElFormItem, ElInput, ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
import { Login } from '../../wailsjs/go/main/App'
import router from '../router'
import { myTokenStore } from '../store/myStore'
import { SUCESS_CODE } from '../util'
const errorMsg = ref('')
const loading = ref(false)

const loginForm = reactive({
    username: 'admin',
    password: '123456',
    remember: false
})

const rules = reactive({
    username: [
        { required: true, message: 'Please input username', trigger: 'blur' }
    ],
    password: [
        { required: true, message: 'Please input password', trigger: 'blur' },
        { min: 6, message: 'Password must be at least 6 characters', trigger: 'blur' }
    ]
})


const userStore = myTokenStore()
const formRef = ref()
const handleLogin = async () => {
    try {
        await formRef.value.validate()
        loading.value = true

        // 调用Go后端登录接口
        const response = await Login({
            username: loginForm.username,
            password: loginForm.password
        })

        if (response.code === SUCESS_CODE) {
            userStore.addToken("your_token_here")

            if (loginForm.remember) {
                localStorage.setItem('rememberedUser', loginForm.username)
            }
            router.push({ name: 'Home' });
        }
    } catch (error) {
        errorMsg.value = '连接服务失败'
        ElMessage.error('Invalid credentials')
    } finally {
        loading.value = false
    }
}

</script>

<template>
    <div class="login-container">
        <el-card class="login-card">
            <template #header>
                <div class="login-header">
                    <h2>Welcome to My App</h2>
                    <el-icon :size="30" color="#409EFF">
                        <UserFilled />
                    </el-icon>
                </div>
            </template>

            <el-form :model="loginForm" :rules="rules" ref="formRef" @submit.prevent="handleLogin">
                <el-form-item prop="username">
                    <el-input v-model="loginForm.username" placeholder="Username" prefix-icon="User" />
                </el-form-item>

                <el-form-item prop="password">
                    <el-input v-model="loginForm.password" type="password" placeholder="Password" prefix-icon="Lock"
                        show-password />
                </el-form-item>

                <el-form-item>
                    <el-checkbox v-model="loginForm.remember">Remember me</el-checkbox>
                </el-form-item>

                <el-button type="primary" native-type="submit" :loading="loading" class="login-btn">
                    登录
                </el-button>
            </el-form>
        </el-card>
    </div>
</template>

<style scoped>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
    width: 400px;
    border-radius: 10px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.login-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
}

.login-btn {
    width: 100%;
    margin-top: 10px;
}
</style>