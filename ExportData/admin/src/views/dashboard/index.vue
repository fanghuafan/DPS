<template>
  <div class="dashboard-container">
    <div class="dashboard-text">Data File Square</div>

    <template>
      <el-carousel :interval="4000" type="card" height="200px">
        <el-carousel-item v-for="item in bannerUris" :key="item">
          <el-image :src="item"></el-image>
          <h3 class="medium">EveryBody should have a dream!</h3>
        </el-carousel-item>
      </el-carousel>
    </template>

    <hr class="hr-class"/>

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

    <hr class="hr-class"/>

    <el-table
      :data="tableData"
      stripe
      style="width: 100%">
      <el-table-column
        prop="taskCode"
        label="task code"
        width="180">
      </el-table-column>
      <el-table-column
        prop="taskName"
        label="task name"
        width="100">
      </el-table-column>
      <el-table-column
        prop="taskInfomation"
        label="task infomation"
        width="180">
      </el-table-column>
      <el-table-column
        prop="fileName"
        label="file name">
      </el-table-column>
      <el-table-column
        prop="handleTime"
        label="handle time">
      </el-table-column>
      <el-table-column
        prop="operator"
        label="operator"
        width="100">
        <template slot-scope="scope">
        <el-tag :type="scope.row.Status | statusFilter">下载</el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
  import {mapGetters} from 'vuex'

  export default {
    data() {
      return {
        bannerUris: [
          'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
          'https://fuss10.elemecdn.com/1/34/19aa98b1fcb2781c4fba33d850549jpeg.jpeg',
          'https://fuss10.elemecdn.com/0/6f/e35ff375812e6b0020b6b4e8f9583jpeg.jpeg',
          'https://fuss10.elemecdn.com/9/bb/e27858e973f5d7d3904835f46abbdjpeg.jpeg',
          'https://fuss10.elemecdn.com/d/e6/c4d93a3805b3ce3f323f7974e6f78jpeg.jpeg',
          'https://fuss10.elemecdn.com/3/28/bbf893f792f03a54408b3b7a7ebf0jpeg.jpeg',
          'https://fuss10.elemecdn.com/2/11/6535bcfb26e4c79b48ddde44f4b6fjpeg.jpeg'
        ],
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
          taskCode: '2016-05-02',
          handleTime: '2016-05-02',
          operator: '下载',
          taskName: "SEM组合词",
          taskInfomation: '上海市普陀区金沙江路 1518 弄',
          fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
        }, {
          taskCode: '2016-05-04',
          handleTime: '2016-05-04',
          operator: '下载',
          taskName: "SEM组合词",
          taskInfomation: '上海市普陀区金沙江路 1517 弄',
          fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
        }, {
          taskCode: '2016-05-01',
          handleTime: '2016-05-01',
          operator: '下载',
          taskName: "SEM组合词",
          taskInfomation: '上海市普陀区金沙江路 1519 弄',
          fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
        }, {
          taskCode: '2016-05-03',
          handleTime: '2016-05-03',
          operator: '下载',
          taskName: "SEM组合词",
          taskInfomation: '上海市普陀区金沙江路 1516 弄',
          fileName: "export-2020-12-01(上海市普陀区金沙江路).xls"
        }]
      }
    },
    name: 'Dashboard',
    // computed: {
    //   ...mapGetters([
    //     'name'
    //   ])
    // }
  }
</script>

<style lang="scss" scoped>
.dashboard {

  &-container {
    margin: 30px;
  }

  &-text {
    font-size: 25px;
    line-height: 46px;
  }
}

  .hr-class {
    background-color: rgba(255, 255, 255, 0);
    border: 1px dashed #e2e2e2;
  }

.el-carousel__item h3 {
  color: #ffffff;
  font-size: 17px;
  opacity: 0;
  line-height: 200px;
  margin: 0;
  text-align: center;
}


</style>
