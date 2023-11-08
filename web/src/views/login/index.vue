<script lang="ts" setup>
import { Lock, User } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'

/** 登录表单元素的引用 */
const loginFormRef = ref<FormInstance | null>(null)

/** 登录按钮 Loading */
const loading = ref(false)

/** 登录表单数据 */
const loginFormData: any = reactive({
  username: 'admin',
  password: '123456',
})

/** 登录表单校验规则 */
const loginFormRules: FormRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 16, message: '长度在 6 到 16 个字符', trigger: 'blur' },
  ],
}

/** 登录逻辑 */
const handleLogin = () => {
  loginFormRef.value?.validate((valid: boolean, fields) => {
    if (valid) {
      loading.value = true
      console.error('登录成功', fields)
      loading.value = false
    } else {
      console.error('表单校验不通过', fields)
    }
  })
}
</script>

<template>
  <div class="flex h-screen w-full items-center justify-center">
    <div class="w-120 rounded-4 overflow-hidden bg-white shadow-xl">
      <div class="flex h-20 items-center justify-center font-mono text-4xl font-bold">Carot Admin</div>
      <div class="px-10 pb-10 pt-4">
        <el-form ref="loginFormRef" :model="loginFormData" :rules="loginFormRules" @keyup.enter="handleLogin">
          <el-form-item prop="username">
            <el-input
              v-model.trim="loginFormData.username"
              placeholder="用户名"
              type="text"
              tabindex="1"
              :prefix-icon="User"
              size="large"
            />
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model.trim="loginFormData.password"
              placeholder="密码"
              type="password"
              tabindex="2"
              :prefix-icon="Lock"
              size="large"
              show-password
            />
          </el-form-item>
          <el-button class="w-full" :loading="loading" type="primary" size="large" @click.prevent="handleLogin">
            登 录
          </el-button>
        </el-form>
      </div>
    </div>
  </div>
</template>
