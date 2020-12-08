<template>
  <div class="app-container">
    <el-table
      :data="tableData"
      stripe
      style="width: 100%">
      <el-table-column
        prop="userName"
        label="User Name"
        width="250">
      </el-table-column>
      <el-table-column
        prop="userRole"
        label="Auth Role"
        width="250">
      </el-table-column>
      <el-table-column
        prop="createTime"
        label="create time"
        width="180">
      </el-table-column>
      <el-table-column
        prop="handleTime"
        label="handle time">
      </el-table-column>
      <el-table-column
        prop="auth"
        label="auth"
        width="100">
        <template slot-scope="scope">
          <el-tag :type="scope.row.Status | statusFilter">hit atuh</el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData: [{
        createTime: '2016-05-02',
        handleTime: '2016-05-02',
        operator: '下载',
        userName: "SEM组合词",
        userRole: '上海市普陀区金沙江路 1518 弄',
        fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
      }, {
        createTime: '2016-05-04',
        handleTime: '2016-05-04',
        operator: '下载',
        userName: "SEM组合词",
        userRole: '上海市普陀区金沙江路 1517 弄',
        fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
      }, {
        createTime: '2016-05-01',
        handleTime: '2016-05-01',
        operator: '下载',
        userName: "SEM组合词",
        userRole: '上海市普陀区金沙江路 1519 弄',
        fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
      }, {
        createTime: '2016-05-03',
        handleTime: '2016-05-03',
        operator: '下载',
        userName: "SEM组合词",
        userRole: '上海市普陀区金沙江路 1516 弄',
        fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
      }]
    }
  },
  methods: {
    onSubmit() {
      this.$message('submit!')
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

