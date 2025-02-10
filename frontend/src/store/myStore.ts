import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const myTokenStore = defineStore('myToken', () => {
    const token = ref('')
    // const removeToken = computed(() => token.value = '')
    function addToken(data: string) {
        console.log('save token')
        token.value = data
    }
    function removeToken() {
        token.value = ''
    }

    return { token, removeToken, addToken }
}, {
    persist: {
        key: 'my-custom-key',
        storage: localStorage,
    }
})
