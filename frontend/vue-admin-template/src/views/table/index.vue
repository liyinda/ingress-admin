<template>
  <div class="app-container">
    <tree-table ref="recTree" :list.sync="list" @actionFunc="actionFunc" @deleteFunc="deleteFunc" @handlerExpand="handlerExpand" @orderByFunc="orderByFunc"></tree-table>
  </div>
</template>

<script>
import { getList } from '@/api/table'
import treeTable from '@/components/pc/tree-table.vue'
import { getToken } from '@/utils/auth'

const state = {
  token: getToken(),
  name: '',
  avatar: ''
}

export default {
  components: {
    treeTable
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    orderByFunc(val) {
      alert('排序')
      alert(val)
    },
    actionFunc(m) {
      alert('edit')
    },
    deleteFunc(m) {
      alert('Please check name carefully !')
      window.location.href="/#/delete/index"
    },
    handlerExpand(m) {
      console.log('展开/收缩')
      m.isExpand = !m.isExpand
    },
    fetchData() {
      this.listLoading = true
      getList(state.token).then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>

<style>
.app-container {
    height:800px;
    overflow: scroll;
}
</style>
