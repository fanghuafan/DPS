<template>
  <div>
    <div class="addProcess">
      <el-button @click="dialog = true; mountedCodeEdit()" type="primary" icon="el-icon-plus" plain>Add Process</el-button>
      <el-select v-model="value" placeholder="Please Select Hnadler or Expoter">
        <el-option-group
          v-for="group in processType"
          :key="group.label"
          :label="group.label">
          <el-option
            v-for="item in group.options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-option-group>
      </el-select>
    </div>

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

          <!--<el-steps :active="1" finish-status="success" process-status="error" align-center class="step">-->
            <!--<el-step title="步骤1" description="这是一段很长很长很长的描述性文字"></el-step>-->
            <!--<el-step title="步骤2" description="这是一段很长很长很长的描述性文字"></el-step>-->
            <!--<el-step title="步骤3" description="这是一段很长很长很长的描述性文字"></el-step>-->
            <!--<el-step title="步骤4" description="这是一段很长很长很长的描述性文字"></el-step>-->
          <!--</el-steps>-->

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

        <template slot-scope="scope">
          <el-button type="primary" size="mini" icon="el-icon-edit" plain></el-button>

          <el-button type="danger" size="mini" icon="el-icon-delete" plain></el-button>
        </template>

      </el-table-column>
    </el-table>

    <el-divider></el-divider>

    <div class="block">
      <el-timeline>
        <el-timeline-item
          v-for="(activity, index) in activities"
          :key="index"
          :icon="activity.icon"
          :type="activity.type"
          :color="activity.color"
          :size="activity.size"
          :timestamp="activity.timestamp">
          {{activity.content}}
        </el-timeline-item>
      </el-timeline>
    </div>

    <el-drawer
      size="60%"
      title="Add Process Info"
      :before-close="handleClose"
      :visible.sync="dialog"
      direction="rtl"
      :destroy-on-close="true"
      custom-class="demo-drawer"
      ref="drawer">
      <div class="demo-drawer__content" >
        <el-form :model="form" >
          <el-form-item label="Process Name" :label-width="formLabelWidth">
            <el-input placeholder="Please enter name by process" v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Process Desc" :label-width="formLabelWidth">
            <el-input
              :rows="2" type="textarea" placeholder="Please enter description by process" ></el-input>
          </el-form-item>
          <el-form-item label="Process Type" :label-width="formLabelWidth">
            <el-select placeholder="Select Process Type" v-model="processType">
              <el-option label="区域一" value="shanghai"></el-option>
              <el-option label="区域二" value="beijing"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Process DB" :label-width="formLabelWidth">
            <el-select placeholder="Select DB Type" v-model="processType">
              <el-option label="MySql" value="mysql"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="Process State" :label-width="formLabelWidth">
            <el-switch>
            </el-switch>
          </el-form-item>
          <el-form-item label="Process Script" :label-width="formLabelWidth">

            <template>
              <div class="in-coder-panel">
                <textarea class="code-edit" ref="codeEdit"></textarea>
                <el-select size="mini" class="code-mode-select" v-model="mode"
                           @change="changeMode">
                  <el-option v-for="mode in modes"
                             :key="mode.value" :label="mode.label" :value="mode.value">
                  </el-option>
                </el-select>
              </div>
            </template>

          </el-form-item>

          <el-form-item style="display: block;" label="Process Callback" :label-width="formLabelWidth">
            <el-input placeholder="Please enter callback url by process" v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>

        </el-form>
        <div class="demo-drawer__footer">
          <el-button class="l-btn" @click="cancelForm">Cancel</el-button>
          <el-button class="r-btn" type="primary" @click="$refs.drawer.closeDrawer()" :loading="loading">{{ loading ? 'Commit...' : 'Confirm' }}</el-button>
        </div>
      </div>
    </el-drawer>

  </div>
</template>

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
  .addProcess{
    margin: 20px;
  }
  .demo-drawer__footer {
    display: flex;
    margin: 0px;
    justify-content: center;
  }
  .demo-drawer__content {
    display: flex;
    flex-direction: column;
    height: 100%;
    margin: 0 20px 0 20px;
  }
  .l-btn{
    width: 150px;
  }
  .r-btn{
    margin-left: 100px;
    width: 150px;
  }
  :focus {
    outline: 0;
  }

  .in-coder-panel {
    flex-grow:1;
    display:flex;
    position: relative;
  }

  .CodeMirror {
    flex-grow: 1;
    z-index: 1;
    overflow-y: auto !important;
  }

  .CodeMirror-code {
    line-height: 19px;
  }

  .code-mode-select {
    position: absolute;
    z-index: 2;
    right: 10px;
    top: 0px;
    max-width: 77px;

  }
  .code-edit {

  }

</style>

<script>
  // 引入全局实例
  import _CodeMirror from 'codemirror'

  // 核心样式
  import 'codemirror/lib/codemirror.css'
  // 引入主题后还需要在 options 中指定主题才会生效
  import 'codemirror/theme/idea.css'

  // 需要引入具体的语法高亮库才会有对应的语法高亮效果
  // codemirror 官方其实支持通过 /addon/mode/loadmode.js 和 /mode/meta.js 来实现动态加载对应语法高亮库
  // 但 vue 貌似没有无法在实例初始化后再动态加载对应 JS ，所以此处才把对应的 JS 提前引入
  import 'codemirror/mode/javascript/javascript.js'
  import 'codemirror/mode/css/css.js'
  import 'codemirror/mode/xml/xml.js'
  import 'codemirror/mode/clike/clike.js'
  import 'codemirror/mode/markdown/markdown.js'
  import 'codemirror/mode/python/python.js'
  import 'codemirror/mode/r/r.js'
  import 'codemirror/mode/shell/shell.js'
  import 'codemirror/mode/sql/sql.js'
  import 'codemirror/mode/swift/swift.js'
  import 'codemirror/mode/vue/vue.js'

  // 尝试获取全局实例
  const CodeMirror = window.CodeMirror || _CodeMirror

  export default {
    name: 'in-coder',
    props: {
      // 外部传入的内容，用于实现双向绑定
      value: String,
      // 外部传入的语法类型
      language: {
        type: String,
        default: null
      }
    },
    data() {
      return {
        activities: [{
          content: '支持使用图标',
          timestamp: '2018-04-12 20:46',
          size: 'large',
          type: 'primary',
          icon: 'el-icon-more'
        }, {
          content: '支持自定义颜色',
          timestamp: '2018-04-03 20:46',
          color: '#0bbd87'
        }, {
          content: '支持自定义尺寸',
          timestamp: '2018-04-03 20:46',
          size: 'large'
        }, {
          content: '默认样式的节点',
          timestamp: '2018-04-03 20:46'
        }],
        isFirst: true,
        // 内部真实的内容
        code: '',
        // 默认的语法类型
        mode: 'sql',
        // 编辑器实例
        coder: null,
        // 默认配置
        options: {
          // 缩进格式
          tabSize: 2,
          // 主题，对应主题库 JS 需要提前引入
          theme: 'idea',
          mode: 'text/x-sql',
          // 显示行号
          lineNumbers: true,
          line: true
        },
        // 支持切换的语法高亮类型，对应 JS 已经提前引入
        // 使用的是 MIME-TYPE ，不过作为前缀的 text/ 在后面指定时写死了
        modes: [{
          value: 'x-sql',
          label: 'sql'
        },{
          value: 'x-python',
          label: 'python'
        }],
        processType:[
          {
            label: "Handler",
            options: [{
              label: "SQL Handler (Script)",
              value: "1"
            },{
              label: "Remote Sys Handler (Link Callback)",
              value: "2"
            },{
              label: "Pipeline Handler (Support Python3)",
              value: "3"
            }]
          },
          {
            label: "Exporter",
            options: [{
              label: "SQL Exporter (For Excel)",
              value: "1"
            }]
          }

        ],
        drawerProp:{
          direction: 'ltr',
          drawer: false,
          withHeader: false,
          appendToBody: true,
          model: false
        },
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
        loading: false,
        dialog: false,
        formLabelWidth: '130px',
        timer: null,
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
        }],
        // options: [{
        //   value: '选项1',
        //   label: '黄金糕'
        // }, {
        //   value: '选项2',
        //   label: '双皮奶'
        // }, {
        //   value: '选项3',
        //   label: '蚵仔煎'
        // }, {
        //   value: '选项4',
        //   label: '龙须面'
        // }, {
        //   value: '选项5',
        //   label: '北京烤鸭'
        // }],
        // value: ''
      }
    },
    methods: {
      mountedCodeEdit () {
        if(!this.isFirst) {
          return
        }
        this.isFirst = false
        // 初始化
        this._initialize()
      },
      handleClose(done) {
        if (this.loading) {
          return;
        }
        this.$confirm('确定要提交表单吗？')
          .then(_ => {
            this.loading = true;
            this.timer = setTimeout(() => {
              done();
              // 动画关闭需要一定的时间
              setTimeout(() => {
                this.loading = false;
              }, 400);
            }, 2000);
          })
          .catch(_ => {});
      },
      cancelForm() {
        this.loading = false;
        this.dialog = false;
        clearTimeout(this.timer);
      },

      // 初始化
      _initialize() {
        
        this.$nextTick(() => {
          // 初始化编辑器实例，传入需要被实例化的文本域对象和默认配置
          this.coder = CodeMirror.fromTextArea(this.$refs.codeEdit, this.options)
          // 编辑器赋值
          this.coder.setValue(this.value || this.code)
          this.coder.setSize('auto','300px')
          // 支持双向绑定
          this.coder.on('change', (coder) => {
            this.code = coder.getValue()

            if (this.$emit) {
              this.$emit('input', this.code)
            }
          })

          // 尝试从父容器获取语法类型
          if (this.language) {
            // 获取具体的语法类型对象
            let modeObj = this._getLanguage(this.language)

            // 判断父容器传入的语法是否被支持
            if (modeObj) {
              this.mode = modeObj.label
            }
          }

        })

      },
      // 获取当前语法类型
      _getLanguage(language) {
        // 在支持的语法类型列表中寻找传入的语法类型
        return this.modes.find((mode) => {
          // 所有的值都忽略大小写，方便比较
          let currentLanguage = language.toLowerCase()
          let currentLabel = mode.label.toLowerCase()
          let currentValue = mode.value.toLowerCase()

          // 由于真实值可能不规范，例如 java 的真实值是 x-java ，所以讲 value 和 label 同时和传入语法进行比较
          return currentLabel === currentLanguage || currentValue === currentLanguage
        })
      },
      // 更改模式
      changeMode(val) {
        // 修改编辑器的语法配置
        this.coder.setOption('mode', `text/${val}`)

        // 获取修改后的语法
        let label = this._getLanguage(val).label.toLowerCase()

        // 允许父容器通过以下函数监听当前的语法值
        this.$emit('language-change', label)
      }
    },
    mounted () {
      // 初始化
      this._initialize()
    },

  }
</script>
