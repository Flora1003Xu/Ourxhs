<!-- 登陆注册 -->
<template>
  <el-dialog title="登录" width="40%" :before-close="handleClose" :visible.sync="isVisible" :close-on-click-modal="false"
    @submit.prevent="submitlogin">

    <el-form ref="form" label-position="left" label-width="80px">
      <el-form-item label="手机号">
        <el-col>
          <el-input v-model="phone" placeholder="请输入手机号" clearable></el-input></el-col>
      </el-form-item>
      <el-form-item label="密码">
        <el-input type="password" v-model="password" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="$emit('input', false), submitLogin()">登 录</el-button>
        <el-button @click="regVisible = true,clearForm()">新用户注册</el-button>

        <!-- 注册dialog -->
        <el-dialog append-to-body title="新用户注册" :visible.sync="regVisible" width="40%" 
        :limit="1"
        :before-close="handleClose"
          :close-on-click-modal="false">
          <el-form :rules="rules" ref="form" :model="form" label-position="left" label-width="80px">
            <el-form-item label="头像" prop="header">
              <el-row>
                <el-upload class="avatar-uploader" action="#" :auto-upload="false" :show-file-list="false"
                  :file-list="form.header" :on-change="imgHandleChange">
                  <img v-if="imageUrl" :src="this.imageUrl" class="avatar">
                  <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                </el-upload>
              </el-row>
            </el-form-item>
            <el-form-item label="昵称" prop="name">
              <el-input placeholder="请输入昵称" v-model="form.name" clearable maxlength="10" show-word-limit>
              </el-input>
            </el-form-item>
            <el-form-item label="性别" prop="sex">
              <el-radio-group v-model="form.sex">
                <el-radio label="男"></el-radio>
                <el-radio label="女"></el-radio>
              </el-radio-group>
            </el-form-item>

            <el-form-item label="手机号" prop="phone">
              <el-col>
                <el-input v-model="form.phone" placeholder="请输入用于登录的手机号" clearable></el-input></el-col>
            </el-form-item>

            <el-form-item label="生日" prop="birth">
              <el-col :span="5">
                <el-date-picker v-model="form.birth" type="date" placeholder="选择日期" value-format="yyyy-MM-dd"
                  :picker-options="pickerOptions">
                </el-date-picker>
              </el-col>
            </el-form-item>
            <el-form-item label="简介" prop="sign">
              <el-input placeholder="请输入个人简介" v-model="form.sign" clearable maxlength="20" show-word-limit>
              </el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pass">
              <el-input type="password" v-model="form.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitReg('form')">注 册</el-button>
            </el-form-item>
          </el-form>
        </el-dialog>
      </el-form-item>
    </el-form>
    <div v-if="error">{{ error }}</div>
  </el-dialog>
</template>

<script>
import axios from 'axios'
import GlobalToken from '@/components/GlobalToken.vue'
export default {
  data() {
    var phoneReg = /^[1][3,4,5,7,8][0-9]{9}$/
    var validatePhone = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('号码不能为空!!'))
      }
      setTimeout(() => {
        if (!phoneReg.test(value)) {
          callback(new Error('格式有误'))
        } else {
          callback()
        }
      }, 100)
    };
    return {
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now();
        },
      },
      phone: '',
      password: '',
      error: '',
      regVisible: false,
      form: {
        name: '',
        sex: '',
        birth: '',
        phone: '',
        pass: '',
        id: '',
        sign: '',
        header: [],
      },
      rules: {
        pass: [
          { required:true,message:"请设置密码", trigger: 'blur' }
        ],
        name:[
          {required:true,message:"请设置用户名", trigger: 'blur'}
        ],
        birth:[
          {required:true,message:"请选择生日日期", trigger: 'change'}
        ],
        sex:[
          {required:true,message:"请选择性别", trigger: 'change'}
        ],
        sign:[
          {required:true,message:"请设置个人简介", trigger: 'blur'}
        ],
        phone:[
          {required:true,message:"请输入正确的手机号码", validator: validatePhone,trigger: 'blur'}
        ],
        header:[
          {required:true,message:"请选择头像"}
        ]
      },
      imageUrl: '',
      token: '',
      err: '',
      userId:1,
    }
  },
  name: 'loginDialog',
  props: ["Visible"],
  computed: {
    isVisible: {
      get() {
        return this.Visible
      },
      set(val) {
        this.$emit('updateVisible', val)
      }
    }
  },
  methods: {
    // 注册表单每次点进清空
    clearForm(){
      this.$refs.form.resetFields();
      console.log(this.form.header)
      if(this.form.header.length>0)
      {
        this.form.header.pop();
      }
      this.imageUrl='';  
    },
    // 提交登录
    submitLogin() {
      try {
        axios.post('http://localhost:8081/login', {
          phoneNumber: this.phone,
          password: this.password
        }).then(response => {
          if(response.data.code==200)
            this.$message.success('登陆成功！')
          this.token = response.data.data;
          console.log('1231',this.token)
          GlobalToken.Token = this.token;
          if(this.token=='')
            this.$message.error('登陆失败！')
          console.log('asdad',GlobalToken.Token)
          this.userId = response.data.uid;
          var id = this.userId;
          this.$emit('change',id)
          console.log(this.userId)
          this.phone=''
          this.password=''
          this.imageUrl='';
         })
      } catch (error) {
        this.error = error.response.data.message
      }
      console.log(this.token);
      
    },
    handleClose(done) {
      
      this.$confirm('确认关闭？')
        .then(_ => {
          done();
        })
        .catch(_ => { });
        
    },
    imgHandleChange(file, fileList) {
      if (!/\.(jpg|JPG|png|PNG)$/.test(file.name)) {
        this.$message({
          type: 'warning',
          message: '只支持格式为jpg/png的文件!'
        })
        fileList.pop()
        return false
      }
      this.form.header = fileList
      // 改的
      this.imageUrl=URL.createObjectURL(file.raw)
      this.$refs.form.clearValidate('header')
    },
    submitReg(formName) {   // 同步注册
      this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
            var FormData = require('form-data');
      var data = new FormData();
      this.form.header.forEach(
        (val, index) => {
          data.append("file", val.raw);
        }
      );
      data.append('birthday', this.form.birth);
      data.append('gender', this.form.sex);
      data.append('introduction', this.form.sign);
      data.append('password', this.form.pass);
      data.append('phoneNumber', this.form.phone);    
      data.append('userName', this.form.name);
      try {
        axios({
              method:'post',
              url: 'http://localhost:8081/register',
              headers: {'Content-Type': false},
              data
              }).then(
          response => {
            console.log(response)
          this.token = response.data.token;
          GlobalToken.Token = this.token;
          console.log('asdad',GlobalToken.Token)
        })
      } catch (error) {
        this.error = error.response.data.message
      }
      this.$refs.form.resetFields();
      this.regVisible = false;
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      
    }
  }
}
</script>
<style>
.likespan {
  float: right;
}

.senduser {
  font-size: 13px;
  color: #999;
}

.el-card {
  height: 80%;
  min-height: 300px;
}

.el-icon-heart {
  background: url('../assets/heart.jpg') center no-repeat;
  font-size: 3rem;
  width: 80%
}

.el-icon-solidheart {
  background: url('../assets/solidheart.jpg') center no-repeat;
  font-size: 3rem;
  width: 80%
}

.el-icon-heart:before {
  content: "切换";
  font-size: 3rem;
  visibility: hidden;
}

.el-icon-solidheart:before {
  content: "切换";
  font-size: 3rem;
  visibility: hidden;
}

.bottom {
  margin-top: 10px;
  line-height: 25px;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both
}

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 120px;
  height: 120px;
  line-height: 120px;
  text-align: center;
}

.avatar {
  width: 120px;
  height: 120px;
  display: block;
}
</style>