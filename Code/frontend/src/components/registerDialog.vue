<template>
  <el-dialog
    append-to-body
    title="新用户注册"
    :visible.sync="regVisible"
    width="40%"
    :before-close="handleClose"
    :close-on-click-modal="false"
  >
    <el-form
      ref="ruleForm"
      :model="form"
      label-position="left"
      label-width="80px"
    >
      <el-form-item label="头像">
        <el-row>
          <!--读取数据库存储的头像 路径写在src里-->
          <el-upload
            class="avatar-uploader"
            action="#"
            :auto-upload="false"
            :show-file-list="false"
            :file-list="form.header"
            :on-change="imgHandleChange"
            :before-upload="beforeAvatarUpload"
          >
            <img
              v-if="imageUrl"
              :src="require('../assets/' + this.imageUrl)"
              class="avatar"
            />
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            <!--action写上传的url路径-->
          </el-upload>
        </el-row>
      </el-form-item>
      <el-form-item label="昵称">
        <el-input
          placeholder="请输入昵称"
          v-model="form.name"
          clearable
          maxlength="10"
          show-word-limit
        >
        </el-input>
      </el-form-item>
      <el-form-item label="性别">
        <el-radio-group v-model="form.sex">
          <el-radio label="男"></el-radio>
          <el-radio label="女"></el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="手机号">
        <el-col>
          <el-input
            v-model="form.phone"
            placeholder="请输入手机号"
            clearable
          ></el-input
        ></el-col>
      </el-form-item>

      <el-form-item label="生日">
        <el-col :span="5">
          <el-date-picker
            v-model="form.birth"
            type="date"
            placeholder="选择日期"
            value-format="yyyy-MM-dd"
            :picker-options="pickerOptions"
          >
          </el-date-picker>
        </el-col>
      </el-form-item>
      <el-form-item label="简介">
        <el-input
          placeholder="请输入个人简介"
          v-model="form.sign"
          clearable
          maxlength="20"
          show-word-limit
        >
        </el-input>
      </el-form-item>
      <el-form-item label="密码" prop="pass">
        <el-input
          type="password"
          v-model="form.pass"
          autocomplete="off"
        ></el-input>
      </el-form-item>
      <!-- <el-form-item label="确认密码" 
            prop="checkPass">
            <el-input type="password" v-model="form.checkPass" autocomplete="off"></el-input>
        </el-form-item> -->
      <el-form-item>
        <!-- <el-button @click="regVisible.set()">取 消</el-button> -->
        <el-button type="primary" @click="submitReg()">注 册</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script>
import axios from "axios";
export default {
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        if (this.ruleForm.checkPass !== "") {
          this.$refs.ruleForm.validateField("checkPass");
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      form: {
        name: "",
        sex: "",
        birth: "",
        phone: "",
        pass: "",
        // checkPass: '',
        id: "",
        sign: "",
        header: [],
      },
      rules: {
        pass: [{ validator: validatePass, trigger: "blur" }],
        checkPass: [{ validator: validatePass2, trigger: "blur" }],
      },
      imageUrl: "momo.jpg",
    };
  },
  name: "registerDialog",
  props: {
    Visible: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    regVisible: {
      get() {
        return this.Visible;
      },
      set(val) {
        this.$emit("updateregVisible", val);
      },
    },
  },
  methods: {
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },

    imgHandleChange(file, fileList) {
      if (!/\.(jpg|JPG|png|PNG)$/.test(file.name)) {
        this.$message({
          type: "warning",
          message: "只支持格式为jpg/png的文件！",
        });
        fileList.pop();
        return false;
      }
      this.form.header = fileList;
    },

    //   头像上传
    handleAvatarSuccess(res, file) {
      this.imageUrl = file.name;
      this.form.header = file.raw;
      console.log(file.raw);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      return isJPG && isLt2M;
    },
    submitReg() {
      var FormData = require("form-data");
      var data = new FormData();
      data.append("birthday", this.form.birth);
      data.append("gender", this.form.sex);
      data.append("introduction", this.form.sign);
      data.append("password", this.form.pass);
      data.append("phoneNumber", this.form.phone);
      this.form.header.forEach((val, index) => {
        data.append("file", val.raw);
      });
      // data.append('file', this.form.header[0].raw);
      data.append("userName", this.form.name);
      axios.post("http://localhost:8080/register", data).then((response) => {
        console.log(response);
      });
    },
  },
};
</script>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
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
