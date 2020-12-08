<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="DB Name">
        <el-col :span="7">
          <el-input placeholder="DB Name" v-model="form.name" />
        </el-col>
      </el-form-item>
      <el-form-item label="DB Type">
        <el-select v-model="form.region" placeholder="please select your zone">
          <el-option label="MySQL" value="MySQL" />
          <el-option label="MongoDB" value="MongoDB" />
        </el-select>
      </el-form-item>
      <el-form-item label="DB Host">
        <el-col :span="7">
          <el-input placeholder="ip host" v-model="form.name" />
        </el-col>
        <el-col :span="0.01">
          &nbsp;-&nbsp;
        </el-col>
        <el-col :span="2">
          <el-input v-model="form.name" placeholder="port"/>
        </el-col>
      </el-form-item>
      <el-form-item label="DB User">
        <el-col :span="7">
          <el-input placeholder="DB User" v-model="form.name" />
        </el-col>
      </el-form-item>
      <el-form-item label="DB Password">
        <el-col :span="7">
          <el-input placeholder="DB Password" v-model="form.name" />
        </el-col>
      </el-form-item>
      <el-form-item label="DB State">
        <el-switch v-model="form.delivery" />
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

