<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="User Name">
        <el-select v-model="form.region" placeholder="please select your zone">
          <el-option label="MySQL" value="MySQL" />
          <el-option label="MogoDB" value="MogoDB" />
        </el-select>
      </el-form-item>
      <el-form-item label="User Role">
        <el-select v-model="form.region" placeholder="please select your zone">
          <el-option label="MySQL" value="MySQL" />
          <el-option label="MogoDB" value="MogoDB" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
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
      }
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

