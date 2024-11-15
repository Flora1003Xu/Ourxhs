<template>
  <div>
    <el-backtop></el-backtop>
    <div class="personal">
      <el-container>
        <el-header>
          <el-col :span="2" :offset="0" class="left">
            <div>
              <img
                src="../assets/xhs_logo.png"
                width="120"
                height="50"
                alt="小红书"
              />
            </div>
          </el-col>
          <el-badge :is-dot="isdot" style="float: right">
            <el-button
              type="danger"
              icon="el-icon-message"
              style="font-size: 23px; float: right"
              @click="(mesVisible = true), (isdot = false)"
            ></el-button
          ></el-badge>
          <!-- 消息列表 -->
          <mesDialog
            :Visible="mesVisible"
            :mesList="message"
            :likeList="likeList"
            :userId="userId"
            :userName="form.name"
            :portrait="imageUrl"
            @updateVisible="updatemesVisible"
          />
        </el-header>
        <el-main>
          <el-dialog
            title="资料修改"
            :visible.sync="dialogVisible"
            width="40%"
            :before-close="handleClose"
          >
            <el-form
              ref="ruleForm"
              :model="form"
              label-position="left"
              label-width="80px"
            >
              <el-form-item label="头像">
                <el-row>
                  <el-upload
                    class="avatar-uploader"
                    action="#"
                    ref="upload"
                    :limit="1"
                    :show-file-list="true"
                    :auto-upload="false"
                    :on-change="imgHandleChange"
                    :file-list="form.uploadUrl"
                  >
                    <img v-if="imageUrl" :src="this.imageUrl" class="avatar" />
                    <i v-else class="el-icon-plus avatar-uploader-icon"></i>
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
                <el-radio-group v-model="form.sex" :style="{ color: 'red' }">
                  <el-radio label="男"></el-radio>
                  <el-radio label="女"></el-radio>
                </el-radio-group>
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
              <el-form-item>
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button
                  type="danger"
                  @click="(dialogVisible = false), submitForm('ruleForm')"
                  >保 存</el-button
                >
              </el-form-item>
            </el-form>
          </el-dialog>

          <el-dialog
            append-to-body
            title="修改密码"
            :visible.sync="editpassVisible"
            width="40%"
            :limit="1"
            :before-close="handleClose"
            :close-on-click-modal="false"
          >
            <el-form
              ref="ruleForm"
              :model="form"
              label-position="left"
              label-width="80px"
            >
              <el-form-item label="原密码" prop="oldpass">
                <el-input
                  placeholder="请输入原密码"
                  v-model="form.oldpass"
                  autocomplete="off"
                  clearable
                >
                </el-input>
              </el-form-item>
              <el-form-item label="新密码" prop="newpass">
                <el-input
                  placeholder="请输入新密码"
                  v-model="form.newpass"
                  autocomplete="off"
                  clearable
                >
                </el-input>
              </el-form-item>
              <el-form-item label="确认密码" prop="checkpass">
                <el-input
                  placeholder="请重新输入新密码"
                  v-model="form.checkpass"
                  autocomplete="off"
                  clearable
                >
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button @click="editpassVisible = false">取 消</el-button>
                <el-button
                  type="primary"
                  @click="(editpassVisible = false), submitpass('ruleForm')"
                  >保 存</el-button
                >
              </el-form-item>
            </el-form>
          </el-dialog>

          <div class="block">
            <el-avatar
              :size="180"
              :src="require('../assets/' + this.avatarurl)"
            ></el-avatar>
          </div>
          <h3 class="user_name">{{ form.name }}</h3>
          <span class="id">小红书号:</span>
          <span class="id" src="form.id">{{ form.id }}</span>
          <el-row>
            <span class="sign">{{ form.sign }}</span>
          </el-row>
          <el-row>
            <el-button
              type="danger"
              icon="el-icon-edit"
              round
              class="edit"
              size="small"
              @click="(dialogVisible = true), getperson()"
              >资料修改</el-button
            >
          </el-row>
          <el-row :gutter="20">
            <el-col :span="8" :offset="10">
              <div>
                <el-table :data="tableData" style="width: 210px">
                  <el-table-column
                    prop="guanzhu"
                    label="关注"
                    width="60"
                    align="center"
                  >
                  </el-table-column>
                  <el-table-column
                    prop="fans"
                    label="粉丝"
                    width="60"
                    align="center"
                  >
                  </el-table-column>
                  <el-table-column
                    prop="liked"
                    label="获赞"
                    width="90"
                    align="center"
                  >
                  </el-table-column>
                </el-table>
              </div>
            </el-col>
          </el-row>

          <el-row type="flex" justify="center">
            <el-tabs v-model="activeName" @tab-click="handleClick">
              <!-- 笔记 -->
              <el-tab-pane label="笔记" name="first">
                <div style="display: flex" class="average_container">
                  <div v-for="(column, index) in columns[0]" :key="index">
                    <div v-for="image in column" :key="image.noteId">
                      <noteCard :attrs="image"></noteCard>
                    </div>
                  </div>
                </div>
              </el-tab-pane>

              <!-- 收藏 -->
              <el-tab-pane label="收藏" name="second">
                <div style="display: flex" class="average_container">
                  <div v-for="(column, index) in columns[1]" :key="index">
                    <div v-for="image in column" :key="image.noteId">
                      <noteCard :attrs="image"></noteCard>
                    </div>
                  </div>
                </div>
              </el-tab-pane>

              <!-- 点赞 -->
              <el-tab-pane label="点赞" name="third">
                <div style="display: flex" class="average_container">
                  <div v-for="(column, index) in columns[2]" :key="index">
                    <div v-for="image in column" :key="image.noteId">
                      <noteCard :attrs="image"></noteCard>
                    </div>
                  </div>
                </div>
              </el-tab-pane>
            </el-tabs>
          </el-row>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import loginDialog from "../components/loginDialog.vue";
import mesDialog from "../components/mesDialog.vue";
import noteCard from "@/components/NoteCard.vue";
import GlobalToken from "@/components/GlobalToken.vue";
export default {
  // 组件引用
  components: {
    loginDialog,
    mesDialog,
    noteCard,
  },
  data() {
    var phoneReg = /^[1][3,4,5,7,8][0-9]{9}$/;
    var validatePhone = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("号码不能为空!!"));
      }
      setTimeout(() => {
        if (!phoneReg.test(value)) {
          callback(new Error("格式有误"));
        } else {
          callback();
        }
      }, 100);
    };

    return {
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now();
        },
      },
      ruleForm: {
        phone: [
          {
            required: true, // required就是这个框必须填写
            validator: validatePhone, // 规则
            trigger: "blur", // blur失去焦点，事件何时触发
          },
        ],
        oldpass: [{ required: true, message: "请输入原密码", trigger: "blur" }],
        newpass: [{ required: true, message: "请输入新密码", trigger: "blur" }],
        checkpass: [{ required: true, message: "不能为空", trigger: "blur" }],
      },
      tableData: [
        {
          guanzhu: "0",
          fans: "0",
          liked: "0",
        },
      ],
      activeName: "first",
      dialogVisible: false,
      mesVisible: false,
      // 数据库里只存文件名，全部存到assets里
      imageUrl: "",
      editpassVisible: false,
      avatarurl: "",
      Visible: false,
      loginVisible: false,
      isdot: false,
      userId: 1,
      note: [[], [], []],
      columns: [
        [[], [], []],
        [[], [], []],
        [[], [], []],
      ],
      form: {
        name: "",
        sex: "",
        birth: "",
        phone: "",
        id: 1,
        sign: "",
        password: "",
        uploadUrl: [],
        uploadflag: false,
        oldpass: "",
        newpass: "",
        checkpass: "",
      },
      message: [
        {
          commentatorId: 1,
          commentatorName: "",
          commentId: 1,
          commentTime: "",
          content: "",
          portrait: "",
          state: 1,
          noteId: 1,
        },
      ],
      likeList: [
        {
          likeId: 1,
          userName: "",
          portrait: "",
          likeTime: "",
          noteId: 1, //
          state: 1,
        },
      ],
    };
  },
  created() {
    this.userId = GlobalToken.userid;
    axios({
      method: "get",
      url:
        "http://localhost:8081/" +
        this.userId +
        "/" +
        this.userId +
        "/PersonalView",
      headers: {
        "Content-Type": false,
        Authorization: GlobalToken.Token,
      },
    }).then((response) => {
      this.tableData[0].guanzhu = response.data.data.userInfo.followNum;
      this.tableData[0].fans = response.data.data.userInfo.fansNum;
      this.tableData[0].liked = response.data.data.userInfo.likedNum;
      this.avatarurl = response.data.data.userInfo.portrait;
      this.imageUrl = response.data.data.userInfo.portrait;
      this.imageUrl = require("../assets/" + this.imageUrl);
      this.form.uploadUrl[0] = response.data.data.userInfo.portrait;
      this.form.birth = response.data.data.userInfo.birthday;
      this.form.name = response.data.data.userInfo.userName;
      this.form.sex = response.data.data.userInfo.gender;
      this.form.phone = response.data.data.userInfo.phoneNumber;
      this.form.sign = response.data.data.userInfo.introduction;
      this.form.password = response.data.data.userInfo.password;
      this.form.id = this.userId;
      this.note[0] = response.data.data.notes; //笔记
      this.note[1] = response.data.data.collects; //收藏
      this.note[2] = response.data.data.likes; //点赞
      this.splitImagesIntoColumns(0);
      this.splitImagesIntoColumns(1);
      this.splitImagesIntoColumns(2);
    });
    axios({
      method: "get",
      url: "http://localhost:8081/messages/" + this.userId + "/comments",
      headers: {
        "Content-Type": false,
        Authorization: GlobalToken.Token,
      },
    }).then((response) => {
      this.message = response.data.data;
      this.isdot = !response.data.totalState;
    });
    axios({
      method: "get",
      url: "http://localhost:8081/messages/" + this.userId + "/likes",
      headers: {
        "Content-Type": false,
        Authorization: GlobalToken.Token,
      },
    }).then((response) => {
      this.likeList = response.data.data;
      this.isdot = !response.data.totalState;
    });
  },

  methods: {
    submitpass(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.form.oldpass != this.form.password)
            return this.$message.error("密码错误！");
          else {
            if (this.form.checkpass != this.form.newpass)
              return this.$message.error("密码不一致！");
            else {
              if (this.form.newpass == this.form.oldpass)
                return this.$message.error("新密码与旧密码相同！");
              else {
                var FormData = require("form-data");
                var data = new FormData();
                data.append("oldPassword", this.form.oldpass);
                data.append("newPassword", this.form.newpass);
                axios({
                  method: "put",
                  url:
                    "http://localhost:8081/" +
                    this.userId +
                    "/PersonalView/password",
                  headers: { "Content-Type": false },
                  data,
                }).then((res) => {
                  if (res.data.code == 200) {
                    this.$message.success("密码修改成功！");
                    // 跳转到登陆前的主页
                  }
                });
              }
            }
          }
        }
      });
    },
    //note到时候变成从请求端去获取
    splitImagesIntoColumns(index) {
      let columnCount = 3;
      let notesPerColumn = Math.ceil(this.note[index].length / columnCount);
      for (let i = 0; i < columnCount; i++) {
        for (let j = 0; j < notesPerColumn; j++) {
          const inde = i * notesPerColumn + j;
          if (inde < this.note[index].length) {
            // 叫成cover才能只传文件名
            const image = this.note[index][inde];
            this.columns[index][i].push(image);
          }
        }
      }
    },
    getMes() {
      axios({
        method: "get",
        url: "http://localhost:8081/messages/" + this.userId + "/comments",
        headers: {
          "Content-Type": false,
          Authorization: GlobalToken.Token,
        },
      }).then((response) => {
        this.message = response.data.data;
      });
    },
    // 点击修改个人资料时重新加载，避免上次没成功的残留
    getperson() {
      axios({
        method: "get",
        url:
          "http://localhost:8081/" +
          this.userId +
          "/" +
          this.userId +
          "/PersonalView",
        headers: {
          "Content-Type": false,
          Authorization: GlobalToken.Token,
        },
      }).then((response) => {
        this.imageUrl = response.data.data.userInfo.portrait;
        this.imageUrl = require("../assets/" + this.imageUrl);
        this.form.uploadUrl[0] = response.data.data.userInfo.portrait;
        this.form.birth = response.data.data.userInfo.birthday;
        this.form.name = response.data.data.userInfo.userName;
        this.form.sex = response.data.data.userInfo.gender;
        this.form.phone = response.data.data.userInfo.phoneNumber;
        this.form.sign = response.data.data.userInfo.introduction;
        this.form.password = response.data.data.userInfo.password;
      });
    },
    imgHandleChange(file, fileList) {
      if (!/\.(jpg|JPG|png|PNG)$/.test(file.name)) {
        this.$message({
          type: "warning",
          message: "只支持格式为jpg/png的文件!",
        });
        fileList.pop();
        return false;
      }
      this.form.uploadUrl = fileList;
      this.imageUrl = URL.createObjectURL(file.raw);
      this.form.uploadflag = true;
    },
    handleClick(tab, event) {},
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    submitForm(formName) {
      //修改资料
      this.$refs[formName].validate((valid) => {
        if (valid) {
          var FormData = require("form-data");
          var data = new FormData();
          var puturl = "";
          data.append("birthday", this.form.birth);
          data.append("gender", this.form.sex);
          data.append("introduction", this.form.sign);
          data.append("userName", this.form.name);
          data.append("isHost", "true");
          if (this.form.uploadflag == true) {
            this.form.uploadUrl.forEach((val, index) => {
              data.append("file", val.raw);
            });
            this.form.uploadflag == false;
            puturl = "http://localhost:8081/" + this.userId + "/PersonalView";
          } else {
            puturl = "http://localhost:8081/" + this.userId + "/PersonalView2";
          }

          axios({
            method: "put",
            url: puturl,

            headers: {
              "Content-Type": false,
              Authorization: GlobalToken.Token,
            },
            data,
          }).then((res) => {
            if (res.data.code != 200) {
              return this.$message.error("更改用户信息失败！");
            } else {
              this.$message.success("用户信息修改成功");
              this.$refs.upload.clearFiles();
              axios
                .get("http://localhost:8081/" + this.userId + "/PersonalView")
                .then((response) => {
                  this.tableData[0].guanzhu =
                    response.data.data.userInfo.followNum;
                  this.tableData[0].fans = response.data.data.userInfo.fansNum;
                  this.tableData[0].liked =
                    response.data.data.userInfo.likedNum;
                  this.avatarurl = response.data.data.userInfo.portrait;
                  this.imageUrl = response.data.data.userInfo.portrait;
                  this.imageUrl = require("../assets/" + this.imageUrl);
                  this.form.uploadUrl[0] = response.data.data.userInfo.portrait;
                  this.form.birth = response.data.data.userInfo.birthday;
                  this.form.name = response.data.data.userInfo.userName;
                  this.form.sex = response.data.data.userInfo.gender;
                  this.form.phone = response.data.data.userInfo.phoneNumber;
                  this.form.sign = response.data.data.userInfo.introduction;
                  this.form.password = response.data.data.userInfo.password;
                });
            }
          });
        } else {
          return false;
        }
      });
    },
    updateVisible() {
      this.loginVisible = !this.loginVisible;
    },
    changelike() {
      this.card.islike = !this.card.islike;
      if (this.card.islike) this.card.like += 1;
      else this.card.like -= 1;
    },
    updatemesVisible() {
      this.mesVisible = !this.mesVisible;
    },
  },
};
</script>

<style scoped>
.el-radio__input.is-checked .el-radio__inner {
  background: #28d4c1 !important;
  border-color: #28d4c1 !important;
}

.el-header,
.el-footer {
  background-color: #ffffff;
  color: #303133;
  text-align: center;
  line-height: 60px;
}

.el-main {
  background-color: #ffffff;
  color: #303133;
  text-align: center;
  line-height: 30px;
}

body > .el-container {
  margin-bottom: 40px;
}

.user_name {
  letter-spacing: 1px;
  margin-top: 0px;
  margin-bottom: 1px;
}

.id {
  color: #909399;
  letter-spacing: 1px;
  margin-top: 0px;
  margin-bottom: 1px;
  font-weight: 500;
  font-size: 10px;
}

.sign {
  letter-spacing: 1px;
  margin-top: 0px;
  margin-bottom: 1px;
  font-weight: 500;
  font-size: 13px;
}

.input {
  el-form: inline=true;
}

.el-row {
  margin-bottom: 20px;

  &:last-child {
    margin-bottom: 0;
  }
}

.el-col {
  border-radius: 0px;
}

.el-row {
  height: 100%;
}

.row-bg {
  padding: 10px 0;
}

.el-table {
  border: none;
  fit: "true";
  size: "mini";
  text-align: center;
  color: #303133;
  font-weight: 400;
}

.group > .el-table--enable-row-hover .el-table__body tr:hover > td {
  background-color: transparent !important;
}

.inline-block {
  display: inline-block;
  vertical-align: middle;
}
</style>

<style lang="scss" scoped>
//去掉el-tab-pane底部灰色线条
.el-tabs__nav-wrap::after {
  height: 0 !important;
}

.el-tabs /deep/.el-tabs__nav-wrap::after {
  background-color: #fff;
}

//修改样式
/deep/.el-tabs__item {
  padding: 0 20px 0 0;
}

/deep/.el-tabs__active-bar {
  position: absolute;
  bottom: 0;
  left: 0;
  height: 3px;
  width: 70px;
  background-color: rgb(232, 28, 41);
  z-index: 1;
  color: rgb(214, 33, 33);
}

/deep/.el-tabs__item.is-active {
  color: black;
  font-weight: bold;
}

.edit {
  color: white;
  // background-color:red;
}

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 5px;
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

/* 实现布局三列平分 */
.average_container > div {
  flex: 1;
}

.average_container div {
  border: solid 0px rgba(237, 226, 226, 0.174);
  margin-right: calc(20px / 3);
  box-sizing: border-box;
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>
<style>
.el-menu {
  background-color: #fbfafacc;
  color: #030303cc !important;
}
</style>
