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
                height="40"
                alt="小红书"
              />
            </div>
          </el-col>
        </el-header>
        <el-main>
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
          <!-- <el-row>
                          <el-col :span="3"><el-button type="primary"  size="small" v-if="isfollowshow" round @click="followButtonclick">{{followButtonText}}</el-button></el-col>
                        </el-row> -->
          <el-row :gutter="20">
            <el-col :span="8" :offset="10"
              ><div>
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
                </el-table></div
            ></el-col>
          </el-row>

          <el-row type="flex" justify="center" align="middle">
            <el-tabs v-model="activeName" @tab-click="handleClick" stretch>
              <el-tab-pane label="笔记" name="first"> </el-tab-pane>
              <el-tab-pane label="收藏" name="second"> </el-tab-pane>
              <el-tab-pane label="点赞" name="third"> </el-tab-pane>
            </el-tabs>
          </el-row>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import noteCard from "../components/NoteCard.vue";
import GlobalToken from "@/components/GlobalToken.vue";
export default {
  // 组件引用
  components: {
    noteCard,
  },
  data() {
    return {
      followButtonText: "",
      tableData: [
        {
          guanzhu: "0",
          fans: "0",
          liked: "0",
        },
      ],
      activeName: "first",
      avatarurl: "1.jpg",
      Visible: false,
      userId: 1,
      form: {
        name: "",
        id: 1,
        sign: "",
      },
      isfollow: false,
      watchId: 1, //要传的
    };
  },
  created() {
    this.userId = this.$route.query.other_userid; //加加！！！！
    this.watchId = this.$route.query.other_watchid; //加加！！！
    this.$axios({
      method: "get",
      url:
        "http://localhost:8081/" +
        this.userId +
        "/" +
        this.watchId +
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
      this.form.name = response.data.data.userInfo.userName;
      this.form.sign = response.data.data.userInfo.introduction;
      this.isfollow = response.data.isFollowed;
      this.isfollowshow();
      this.form.id = this.userId;
    });
  },

  methods: {
    isfollowshow() {
      if (this.isfollow) {
        this.followButtonText = "已关注";
      } else {
        this.followButtonText = "关注";
      }
      return this.isfollow;
    },

    followButtonclick() {
      if (this.isfollow) {
        this.followButtonText = "关注";
        this.$axios
          .delete(
            "http://localhost:8081/" + this.watchId + "/PersonalView/follow",
            {
              data: { followID: this.userId },
            }
          )
          .then((response) => {
            this.token = response.data.token;
          });
        this.isfollow = !this.isfollow;
      } else {
        this.followButtonText = "已关注";
        this.$axios
          .post(
            "http://localhost:8081/" + this.watchId + "/PersonalView/follow",
            {
              followID: this.userId,
            }
          )
          .then((response) => {
            this.token = response.data.token;
          });
        this.isfollow = !this.isfollow;
      }
    },

    handleClick(tab, event) {
      console.log(tab, event);
    },
    handleClose(done) {
      this.$confirm("确认关闭？")
        .then((_) => {
          done();
        })
        .catch((_) => {});
    },
    changelike() {
      this.card.islike = !this.card.islike;
      if (this.card.islike) this.card.like += 1;
      else this.card.like -= 1;
    },
  },
};
</script>

<style scoped>
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
::v-deep .el-tabs__nav-scroll {
  width: 50%;
  margin: 0 auto;
}
</style>
