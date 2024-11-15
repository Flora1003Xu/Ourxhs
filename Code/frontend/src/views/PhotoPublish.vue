<template>
  <div class="photoPublish">
    <el-container>
      <el-header>
        <el-col :span="1" :offset="0">
          <div>
            <img
              src="../assets/xhs_logo.png"
              width="120"
              height="50"
              alt="小红书"
            />
          </div>
        </el-col>
      </el-header>
      <el-container>
        <el-aside width="300px">
          <!-- 预览 -->
          <cardNote :attrs="image" :flag="true" ref="children"></cardNote>
        </el-aside>

        <el-main>
          <el-row style="margin-top: -20px">
            <el-col :span="3"
              ><div style="font-size: 25px; white-space: nowrap">
                发布图文<i class="el-icon-picture-outline"></i></div
            ></el-col> </el-row
          ><br />

          <el-row>
            <el-col :span="3"
              ><div style="font-size: 20px; transform: scale(1, 0.975)">
                上传图片
              </div></el-col
            > </el-row
          ><br />

          <el-row>
            <el-form ref="form" :model="form" label-width="10px">
              <el-form-item style="display: flex">
                <el-upload
                  class="avatar-uploader"
                  style="float: left; text-align: left"
                  action="#"
                  multiple
                  :auto-upload="false"
                  :on-change="imgHandleChange"
                  :on-remove="handleRemove"
                  :file-list="questionForm.files"
                  drag
                  list-type="picture-card"
                  :on-exceed="handleExceed"
                  :limit="6"
                >
                  <i class="el-icon-plus avatar-uploader-icon"></i>
                  <div slot="tip" class="el-upload__tip">
                    支持批量上传，上传格式为jpg/png文件，只能上传不超过6张图片
                  </div>
                </el-upload>

                <el-dialog :visible.sync="dialogVisible">
                  <img width="100%" :src="dialogImageUrl" alt="" />
                </el-dialog>
              </el-form-item>
            </el-form>
          </el-row>

          <el-row
            ><el-col
              ><div style="width: 500px">
                <el-input
                  type="text"
                  placeholder="请输入标题"
                  v-model="title"
                  maxlength="20"
                  show-word-limit
                  @blur="topicFinished()"
                >
                </el-input></div
            ></el-col> </el-row
          ><br />

          <TextArea
            @change="handleChange"
            @aitelocationList="handleAitelocationList"
            @aiteList="handleaiteList"
            @aiteidList="handleaiteidList"
            ref="textchild"
          /><br /><br /><br /><br />

          <el-row>
            <el-col :span="3"
              ><div style="font-size: 20px; transform: scale(1, 0.975)">
                发布设置
              </div></el-col
            > </el-row
          ><br />

          <el-form
            ref="form"
            :model="form"
            label-width="85px"
            label-position="left"
          >
            <el-form-item label="自定义地点" prop="location">
              <el-input
                v-model="form.location"
                placeholder="请输入自定义地点"
                maxlength="20"
                show-word-limit
              ></el-input>
            </el-form-item>
            <el-form-item label="所属分栏" prop="'classi'">
              <el-checkbox-group v-model="form.classi">
                <el-checkbox label="美食" name="type"></el-checkbox>
                <el-checkbox label="彩妆" name="type"></el-checkbox>
                <el-checkbox label="穿搭" name="type"></el-checkbox>
                <el-checkbox label="影视" name="type"></el-checkbox>
                <el-checkbox label="职场" name="type"></el-checkbox>
                <el-checkbox label="情感" name="type"></el-checkbox>
                <el-checkbox label="家居" name="type"></el-checkbox>
                <el-checkbox label="游戏" name="type"></el-checkbox>
                <el-checkbox label="旅游" name="type"></el-checkbox>
                <el-checkbox label="健身" name="type"></el-checkbox>
                <el-checkbox label="其他" name="type"></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-form>

          <el-row style="text-align: center">
            <el-button round size="medium" @click="submit('form')"
              >发布</el-button
            > </el-row
          ><br />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
// import axios from 'axios';
import TextArea from "@/components/TextArea.vue";
import cardNote from "../components/CardNote.vue";
import GlobalToken from "../components/GlobalToken.vue";
export default {
  name: "PhotoPublish",
  data() {
    return {
      rules: {
        location: [
          { required: true, message: "请设置自定义地点", trigger: "blur" },
        ],
        classi: [{ required: true, message: "请选择分栏", trigger: "change" }],
      },
      dialogImageUrl: "",
      tempUrl: "",
      dialogVisible: false,
      AtInfos: [],
      NULL: "",
      // 之后传进来了东西再改,头像和用户名和用户id要从前面传进来
      image: {},
      authorid: 1,
      authorpicture: "1.jpg",
      authorname: "yzt",
      like: 0,
      i: 0,
      title: "", //标题！！！！！！
      textarea: "",
      aite: {
        aiteList: [],
        aitelocationList: [],
        aiteidList: [],
      },
      form: {
        location: "",
        classi: [],
      },
      time: "",

      questionForm: {
        files: [],
        creatorId: 1,
        createtime: "",
        title: "",
        location: "",
        tags: [],
        body: "",
      },
    };
  },
  components: {
    TextArea,
    cardNote,
  },
  created: function () {
    this.authorid = this.$route.query.index;
    this.authorname = this.$route.query.name;
    this.authorpicture = this.$route.query.head;
    this.image = {
      noteID: 1,
      cover: require("../assets/1.jpg"),
      title: "此为默认标题",
      portrait: this.authorpicture,
      creatorName: this.authorname,
      likedNum: 0,
    };
  },
  mounted: function () {
    this.$refs.children.setfile(require("../assets/1.jpg"));
    this.$refs.textchild.setiddd(this.authorid);
  },
  methods: {
    topicFinished() {
      //title
      this.image.title = this.title;
      this.$refs.children.settitle(this.title);
    },
    handleChange(value) {
      // 获取子组件传递过来的 value 值
      this.textarea = value;
    },
    handleAitelocationList(value) {
      this.aite.aitelocationList = value;
    },
    handleaiteList(value) {
      this.aite.aiteList = value;
    },
    handleaiteidList(value) {
      this.aite.aiteidList = value;
    },

    submit(formName1) {
      if (this.questionForm.files.length == 0) {
        this.$message.error("请上传笔记图片！");
        return false;
      }
      if (this.title.length == 0) {
        this.$message.error("请输入笔记标题！");
        return false;
      }
      if (this.textarea.length == 0) {
        this.$message.error("请输入笔记内容！");
        return false;
      }
      this.$refs[formName1].validate((valid) => {
        if (valid) {
          //时间
          var date = new Date(); //date:string类型
          var year = date.getFullYear();
          var month =
            date.getMonth() + 1 < 10
              ? "0" + (date.getMonth() + 1)
              : date.getMonth() + 1;
          var day = date.getDate() < 10 ? "0" + date.getDate() : date.getDate();
          var hours =
            date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
          var minutes =
            date.getMinutes() < 10
              ? "0" + date.getMinutes()
              : date.getMinutes();
          var seconds =
            date.getSeconds() < 10
              ? "0" + date.getSeconds()
              : date.getSeconds();
          this.time =
            year +
            "-" +
            month +
            "-" +
            day +
            " " +
            hours +
            ":" +
            minutes +
            ":" +
            seconds;
          this.questionForm.createtime = this.time;
          this.questionForm.title = this.title;
          this.questionForm.location = this.form.location;
          this.questionForm.tags = this.form.classi;
          this.questionForm.body = this.textarea;
          let form = new FormData();
          for (let j = 0; j < this.aite.aiteList.length; j++) {
            let unit = {};
            unit.atName = this.aite.aiteList[j];
            form.append("atName", this.aite.aiteList[j]);
            form.append("atLocation", this.aite.aitelocationList[j]);
            form.append("atUserId", this.aite.aiteidList[j]);
            unit.atLocation = this.aite.aitelocationList[j];
            this.AtInfos.push(unit);
          } //this.AtInfos是结构体数组
          this.aite.aiteidList = [];
          this.aite.aiteList = [];
          this.aite.aitelocationList = [];

          form.append("createtime", this.time);
          form.append("updatetime", this.time);
          this.time = "";
          form.append("title", this.title);
          this.title = "";
          form.append("location", this.form.location);
          this.form.location = [];

          if (this.form.classi.includes("美食")) form.append("tags", "美食");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("彩妆")) form.append("tags", "彩妆");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("穿搭")) form.append("tags", "穿搭");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("影视")) form.append("tags", "影视");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("职场")) form.append("tags", "职场");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("情感")) form.append("tags", "情感");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("家居")) form.append("tags", "家居");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("游戏")) form.append("tags", "游戏");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("旅游")) form.append("tags", "旅游");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("健身")) form.append("tags", "健身");
          else form.append("tags", this.NULL);
          if (this.form.classi.includes("其他")) form.append("tags", "其他");
          else form.append("tags", this.NULL);
          this.form.classi = [];

          form.append("body", this.textarea);
          this.textarea = "";

          this.questionForm.files.forEach((val, index) => {
            console.log(index);
            form.append("files", val.raw);
          });
          this.questionForm.files = [];

          this.$axios({
            headers: {
              "Content-Type":
                "multipart/form-data;boundary = " + new Date().getTime(),
              Authorization: GlobalToken.Token,
            },
            method: "post",
            url: "http://localhost:8081/" + this.authorid + "/publish",
            data: form,
          }).then((successResponse) => {
            if (successResponse.data.code === 200) {
              this.$message({
                type: "info",
                message: "保存成功",
              });
              this.$refs.textchild.delete_value();
              this.$refs.children.setfile(require("../assets/1.jpg"));
              this.$refs.textchild.setiddd(this.authorid);
              this.$refs.children.settitle("此为默认标题");
              this.$router.push({
                name: "home",
                params: { refresh: true },
              });
            } else {
              console.log("错误");
            }
          });
        } else {
          return false;
        }
      });
    },

    imgHandleChange(file, fileList) {
      this.i++;
      if (!/\.(jpg|JPG|png|PNG)$/.test(file.name)) {
        this.$message({
          type: "warning",
          message: "只支持格式为jpg/png的文件！",
        });
        fileList.pop();
        return false;
      }
      this.tempUrl = URL.createObjectURL(file.raw);
      this.questionForm.files = fileList; //每传一张图片触发一次，filelist是已上传的所有图片
      if (this.i == 1) {
        this.$refs.children.setfile(fileList[0].url);
      }
    },

    // 移除图片
    handleRemove(file, fileList) {
      this.questionForm.files = fileList;
      this.i--;
    },
    handleExceed(files, fileList) {
      this.$message.warning(
        `当前限制选择 6 个文件，本次选择了 ${files.length} 个文件，共选择了 ${
          files.length + fileList.length
        } 个文件`
      );
    },
  },
};
</script>

<style scoped>
.left {
  text-align: left;
}

.el-aside {
  color: #000000;
  text-align: center;
  /* 高度 */
  line-height: 500px;
}
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  float: left;
  position: absolute;
  left: 10px;
  top: 20px;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #000000;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 240px;
  height: 178px;
  line-height: 150px;
  text-align: left;
}
</style>
