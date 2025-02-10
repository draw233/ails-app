import { createRouter, createWebHashHistory } from 'vue-router'
import Index from '../App.vue'

const routes = [
    {
        path: '/',
        name: 'Index',
        component: Index
        // component: () => import('../App.vue')
    },
    {
        path: '/hello',
        name: 'Hello',
        component: () => import("../components/HelloWorld.vue")
        // component: () => import('../App.vue')
    },
    {
        path: '/home',
        name: 'Home',
        component: () => import("../views/Home.vue")
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import("../views/Login.vue")
    }, {
        path: '/page1',
        name: 'Page1',
        component: () => import("../pages/Page1.vue")
    }, {
        path: '/page2',
        name: 'Page2',
        component: () => import("../pages/Page2.vue")
    }
]

// Wails 需使用 createWebHistory('/') 而非默认路径
const router = createRouter({
    // linkActiveClass: "active",
    history: createWebHashHistory('/'),
    routes
})

import { myTokenStore } from '../store/myStore'

// 路由守卫
router.beforeEach((to, from, next) => {
    // const token = localStorage.getItem('token'); // 假设 token 存储在 localStorage
    const store = myTokenStore()
    const token = store.token
    if (to.name !== 'Login' && !token) {
        next({ name: 'Login' }); // 如果没有 token，重定向到登录页面
    } else if (to.name === 'Login' && token) {
        next({ name: 'Home' }); // 如果已经登录，重定向到主页
    } else {
        next(); // 继续导航
    }
});

export default router