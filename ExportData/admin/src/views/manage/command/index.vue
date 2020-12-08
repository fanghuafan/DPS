<template>
  <div>
    <div class="in-coder-panel-first">
      <el-form>
        <el-form-item label="">
          <el-button type="primary" plain>Run Script...</el-button>
          <el-button type="danger" plain>Stop Script...</el-button>
          <el-tag size="" class="danger-tips" type="danger">>>> Don't use system root!!!</el-tag>
        </el-form-item>
      </el-form>
    </div>

    <div class="in-coder-panel">
      <textarea class="code-edit" ref="codeEdit">
      </textarea>

      <el-select size="mini" class="code-mode-select" v-model="mode"
                 @change="changeMode">
        <el-option v-for="mode in modes"
                   :key="mode.value" :label="mode.label" :value="mode.value">
        </el-option>
      </el-select>
    </div>

    <div class="in-coder-panel">
      <el-input class="run-logs" disabled="true"
        :rows="30" type="textarea" placeholder="Command run logs..." ></el-input>
    </div>

  </div>
</template>

<style>
  .danger-tips {
    margin-left: 3px;
    vertical-align: bottom;
  }

  .in-coder-panel-first {
    flex-grow:1;
    display:flex;
    position: relative;
    margin-top: 20px;
    margin-left: 20px;
    margin-right: 20px;
  }

  .in-coder-panel {
    flex-grow:1;
    display:flex;
    position: relative;
    margin: 20px;
  }

  .CodeMirror {
    flex-grow: 1;
    z-index: 1;
    overflow-y: auto !important;
    height: 577px !important;
  }

  .CodeMirror-code {
    line-height: 19px;
  }

  .code-mode-select {
    position: absolute;
    z-index: 2;
    right: 10px;
    top: 10px;
    max-width: 177px;

  }
  .code-edit {
    height: 100%;
    font-size : 17px !important;
  }

  .run-logs {
    padding-bottom: 30px;
  }

  .el-textarea.is-disabled .el-textarea__inner {
    background-color: beige !important;
  }

</style>

<script>
  // 引入全局实例
  import _CodeMirror from 'codemirror'

  // 核心样式
  import 'codemirror/lib/codemirror.css'
  // 引入主题后还需要在 options 中指定主题才会生效
  import 'codemirror/theme/idea.css'
  import 'codemirror/theme/shadowfox.css'
  import 'codemirror/theme/erlang-dark.css'

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
        isFirst: true,
        // 内部真实的内容
        code: '',
        // 默认的语法类型
        mode: 'shell',
        // 编辑器实例
        coder: null,
        // 默认配置
        options: {
          value: "# please enter you code \n",
          // 缩进格式
          tabSize: 2,
          // 主题，对应主题库 JS 需要提前引入
          theme: 'shadowfox',
          mode: 'text/x-sh',
          // 显示行号
          lineNumbers: true,
          line: true,
          styleActiveLine: true,
        },
        // 支持切换的语法高亮类型，对应 JS 已经提前引入
        // 使用的是 MIME-TYPE ，不过作为前缀的 text/ 在后面指定时写死了
        modes: [{
          value: 'x-sh',
          label: 'shell'
        },{
          value: 'x-python',
          label: 'python3'
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
