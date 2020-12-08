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
      <el-form-item label="">
        <el-button type="primary" icon="el-icon-search" plain>search</el-button>
      </el-form-item>
    </el-form>

    <el-table
      :data="tableData"
      style="width: 100%">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="商品名称">
              <span>{{ props.row.name }}</span>
            </el-form-item>
            <el-form-item label="所属店铺">
              <span>{{ props.row.shop }}</span>
            </el-form-item>
            <el-form-item label="商品 ID">
              <span>{{ props.row.id }}</span>
            </el-form-item>
            <el-form-item label="店铺 ID">
              <span>{{ props.row.shopId }}</span>
            </el-form-item>
            <el-form-item label="商品分类">
              <span>{{ props.row.category }}</span>
            </el-form-item>
            <el-form-item label="店铺地址">
              <span>{{ props.row.address }}</span>
            </el-form-item>
            <el-form-item label="商品描述">
              <span>{{ props.row.desc }}</span>
            </el-form-item>
          </el-form>

          <el-steps :active="1" finish-status="success" process-status="error" align-center class="step">
            <el-step title="步骤1" description="这是一段很长很长很长的描述性文字"></el-step>
            <el-step title="步骤2" description="这是一段很长很长很长的描述性文字"></el-step>
            <el-step title="步骤3" description="这是一段很长很长很长的描述性文字"></el-step>
            <el-step title="步骤4" description="这是一段很长很长很长的描述性文字"></el-step>
          </el-steps>

        </template>
      </el-table-column>
      <el-table-column
        label="Process No"
        prop="id">
      </el-table-column>
      <el-table-column
        label="Process description"
        prop="name">
      </el-table-column>
      <el-table-column
        label="Process type"
        prop="desc">
      </el-table-column>
      <el-table-column
        label="Process state"
        prop="desc">
      </el-table-column>
      <el-table-column
        label="Process Operate"
        prop="operate">
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
          id: '12987122',
          name: '好滋好味鸡蛋仔',
          category: '江浙小吃、小吃零食',
          desc: '荷兰优质淡奶，奶香浓而不腻',
          address: '上海市普陀区真北路',
          shop: '王小虎夫妻店',
          shopId: '10333'
        }, {
          id: '12987123',
          name: '好滋好味鸡蛋仔',
          category: '江浙小吃、小吃零食',
          desc: '荷兰优质淡奶，奶香浓而不腻',
          address: '上海市普陀区真北路',
          shop: '王小虎夫妻店',
          shopId: '10333'
        }, {
          id: '12987125',
          name: '好滋好味鸡蛋仔',
          category: '江浙小吃、小吃零食',
          desc: '荷兰优质淡奶，奶香浓而不腻',
          address: '上海市普陀区真北路',
          shop: '王小虎夫妻店',
          shopId: '10333'
        }, {
          id: '12987126',
          name: '好滋好味鸡蛋仔',
          category: '江浙小吃、小吃零食',
          desc: '荷兰优质淡奶，奶香浓而不腻',
          address: '上海市普陀区真北路',
          shop: '王小虎夫妻店',
          shopId: '10333'
        }]
      }
    }

  }
</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
  .step{
    margin-top: 20px;
  }
</style>
