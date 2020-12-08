<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Task Name">
        <el-col :span="7">
          <el-input placeholder="Task Name" v-model="form.name" />
        </el-col>
      </el-form-item>
      <el-form-item label="Auth User">
        <el-col :span="7">
          <el-input placeholder="Auth User" v-model="form.user" />
        </el-col>
      </el-form-item>
      <el-form-item label="Description">
        <el-col :span="7">
          <el-input placeholder="Task Description" v-model="form.desc" type="textarea" :rows="7"/>
        </el-col>
      </el-form-item>
      <el-form-item label="Time Regex">
        <el-col :span="7">
          <el-input placeholder="Time Regex Expression" v-model="form.regex" />
        </el-col>
      </el-form-item>
      <el-form-item label="Task Timeout">
        <el-col :span="7">
          <el-input placeholder="Task Timeout" v-model="form.timeout" />
        </el-col>
      </el-form-item>
      <el-form-item label="State">
        <el-switch v-model="form.state" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button @click="onCancel">Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { get, add } from '@/api/task'
export default {

  data() {
    return {
      filterText: '',
      form: {
        name: 'task01',
        user: '1',
        desc: 'Testor test task',
        regex: '*/12/*/89',
        timeout: '100',
        state: '1',
      },
      defaultProps: {
        children: 'children',
        label: 'label'
      }
    }
  },
  watch: {
    filterText(val) {
      this.$refs.tree2.filter(val)
    }
  },

  methods: {
    onSubmit() {
      if(this.form.desc.length == 0
        || this.form.name.length == 0
        || this.form.regex.length == 0
        || this.form.user.length == 0
        || this.form.timeout.length == 0) {
        this.$message({
          message: 'Please task info!',
          type: 'warning'
        });
        return
      }

      // action add request
      this.form.state = this.form.state ? 1 : 0;
      // var resp = add(this.form)
      // console.log(resp)

      this.$store.dispatch('task/add', this.form).then(() => {
        this.$router.push({ path: this.redirect || '/task/taskList' })
        this.loading = false
      }).catch(() => {
        this.loading = false
      })

    },

    onCancel() {

      this.form = {
        name: '',
        user: '',
        desc: '',
        regex: '',
        timeout: '',
        state: '',
      }

      this.$message({
        message: 'Cancel info is ok!',
        type: 'success'
      })

    },

    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    }
  }
}
</script>

