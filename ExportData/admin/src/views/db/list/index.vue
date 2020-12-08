<template>
  <div class="app-container">

    <el-form ref="form" :model="form">
      <el-form-item label="Task Code:">
        <el-select v-model="form.region" placeholder="please select your taskCode">
          <el-option label="Zone one" value="shanghai"/>
          <el-option label="Zone two" value="beijing"/>
        </el-select>
      </el-form-item>
      <el-form-item label="Handle Time:">
        <el-col :span="5">
          <el-date-picker v-model="form.date1" type="date" placeholder="start time" style="width: 100%;"/>
        </el-col>
        <el-col :span="0.01">
          &nbsp;--&nbsp;
        </el-col>
        <el-col :span="5">
          <el-date-picker v-model="form.date1" type="date" placeholder="end time" style="width: 100%;"/>
        </el-col>
      </el-form-item>
    </el-form>

    <el-table
      :data="tableData"
      stripe
      style="width: 100%">
      <el-table-column
        prop="date"
        label="Task Code"
        width="180">
      </el-table-column>
      <el-table-column
        prop="name"
        label="Task Information"
        width="180">
      </el-table-column>
      <el-table-column
        prop="address"
        label="Create Time">
      </el-table-column>
    </el-table>

  </div>
</template>

<script>
// import { getList } from '@/api/list'

export default {
  data() {
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      },
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1518 弄'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1517 弄'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1519 弄'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        address: '上海市普陀区金沙江路 1516 弄'
      }]
    }
  }

  // filters: {
  //   statusFilter(Status) {
  //     const statusMap = {
  //       published: 'success',
  //       draft: 'gray',
  //       deleted: 'danger'
  //     }
  //     return statusMap[Status]
  //   }
  // },
  // data() {
  //   return {
  //     list: null,
  //     listLoading: true
  //   }
  // },
  // created() {
  //   this.fetchData()
  // },
  // methods: {
  //   fetchData() {
  //     this.listLoading = true
  //     getList().then(response => {
  //       this.list = response.data
  //       this.listLoading = false
  //     })
  //   }
  // }
}
</script>
