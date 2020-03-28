<template>
  <div class="app-container">
    <dynamic-form
      v-model="form"
      ref="dataForm" :model="form"
      lang="en_US"
      :descriptors="descriptors"
      :show-outer-error="false">
      <template slot="operations">
        <el-button type="primary" @click="onSubmit" plain>Update</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </template>
    </dynamic-form>
  </div>
</template>

<script>
import { updateDomain } from '@/api/table'
import { Button } from 'element-ui'
import DynamicForm from '../packages/dynamic-form/index'
import descriptors from './descriptor'
import { getToken } from '@/utils/auth'

const state = {
  token: getToken(),
  name: '',
  avatar: ''
}


export default {
  components: {
    DynamicForm,
    ElButton: Button
  },
  data() {
    return {
      form: {
        token: state.token
      },
      descriptors,
    }
  },
  methods: {
    onSubmit() {
      this.$refs.dataForm.validate((valid) => {
        if (valid) {
          updateDomain(this.form).then(() => {
            this.dialogFormVisible = false
            this.$message({
              message: 'Update Success',
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

