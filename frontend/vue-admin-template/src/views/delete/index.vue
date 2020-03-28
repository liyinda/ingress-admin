<template>
  <div class="app-container">
    <el-form ref="dataForm" :model="form" label-width="120px">
      <el-form-item label="Ingress Name">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Confirm Delete</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { delDomain } from '@/api/table'
import { getToken } from '@/utils/auth'

const state = {
  token: getToken(),
  name: '',
  avatar: ''
}

export default {
  data() {
    return {
      form: {
        name: '',
        token: state.token
      }
    }
  },
  methods: {
    onSubmit() {
      this.$refs.dataForm.validate((valid) => {
        if (valid) {
          delDomain(this.form).then(() => {
            this.dialogFormVisible = false
            this.$message({
              message: 'Delete Success',
              type: 'warning'
            })
          })
        }
      })
    },
    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

