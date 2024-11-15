<template>
  <div id="textarea">
    <el-input
      v-model="textarea"
      class="chatText"
      resize="none"
      type="textarea"
      placeholder="请输入内容"
      id="texttexttext"
      maxlength="1000"
      show-word-limit
      rows="5"
      @keydown.native="listen($event)"
      @blur="handleInput()"
    >
    </el-input>

    <div class="chatIcon">
      <!-- el-popover是弹出框 -->
      <el-popover
        placement="top-start"
        width="400"
        trigger="click"
        class="emoBox"
      >
        <div class="emotionList">
          <a
            href="javascript:void(0);"
            @click="getEmo(index)"
            v-for="(item, index) in faceList"
            :key="index"
            class="emotionItem"
            >{{ item }}</a
          >
        </div>
        <el-button class="emotionSelect" slot="reference"
          >(◍•ᴗ•◍)表情</el-button
        >
      </el-popover>
    </div>

    <div class="others">
      <el-popover
        placement="top-start"
        width="200"
        trigger="click"
        class="othersBox"
      >
        <div class="othersList">
          <div
            href="javascript:void(0);"
            @click="getothers(index)"
            v-for="(item, index) in others"
            :key="index"
            class="othersone"
          >
            <el-row type="flex" align="middle">
              <el-avatar
                :size="40"
                :src="require('../assets/' + item.portrait)"
              ></el-avatar
              >&nbsp;&nbsp;{{ item.userName }}
            </el-row>
          </div>
        </div>
        <el-button class="emotionSelect" slot="reference">@用户</el-button>
      </el-popover>
    </div>
  </div>
</template>

<script>
import GlobalToken from "@/components/GlobalToken.vue";
const appData = require(`@/assets/emojis.json`);
export default {
  name: "TextArea",

  data() {
    return {
      userid: 1, //要从父组件传过来
      faceList: [], //所有表情
      facecodeList: [],
      textarea: "",
      others: [
        {
          userName: "Linyx",
          portrait: "1.jpg",
        },
        {
          userName: "yzt",
          portrait: "3.jpg",
        },
        {
          userName: "xy",
          portrait: "2.jpg",
        },
      ],
      aiteList: [], //已选的@的人
      aiteidList: [], //已选的@的人的ID
      aiteheaderList: [], //已选的@的人的头像
      aitelocationList: [],
    };
  },

  mounted() {
    for (let i in appData) {
      this.faceList.push(appData[i].char);
    }
  },

  methods: {
    setiddd(value) {
      this.userid = value;
      //请求数据
      this.$axios({
        method: "get",
        url: "http://localhost:8081/" + this.userid + "/follow",
        headers: {
          "Content-Type": false,
          Authorization: GlobalToken.Token,
        },
      }).then((response) => {
        this.others = response.data.data;
      });
    },

    delete_value() {
      this.textarea = "";
    },

    // 向父组件传值
    handleInput() {
      this.$emit("change", this.textarea);
      this.$emit("aitelocationList", this.aitelocationList);
      this.$emit("aiteList", this.aiteList);
      this.$emit("aiteidList", this.aiteidList);
    },

    getEmo(index) {
      //按下表情中的某一个才执行
      const textArea = document.getElementById("texttexttext");
      function changeSelectedText(obj, str) {
        if (window.getSelection) {
          // 非IE浏览器
          textArea.focus();
          textArea.setRangeText(str);
          // 在未选中文本的情况下，重新设置光标位置
          textArea.selectionStart += str.length;
          textArea.focus();
        } else if (document.selection) {
          // OE浏览器
          obj.focus();
          var sel = document.selection.createRange();
          sel.text = str;
        }
      }
      changeSelectedText(textArea, this.faceList[index]);

      this.textarea = textArea.value; // 要同步data中的数据
      return;
    },

    getothers(index) {
      const textArea = document.getElementById("texttexttext");
      function changeSelectedText(obj, str) {
        if (window.getSelection) {
          // 非IE浏览器
          textArea.focus();
          textArea.setRangeText(" @" + str + " ");
          // 在未选中文本的情况下，重新设置光标位置
          textArea.selectionStart += str.length + 3;
          textArea.focus();
        } else if (document.selection) {
          // IE浏览器
          obj.focus();
          var sel = document.selection.createRange();
          sel.text = str;
        }
      }
      this.aiteList.push(this.others[index].userName);
      this.aitelocationList.push(this.textarea.length);
      this.aiteidList.push(this.others[index].followAct);
      changeSelectedText(textArea, this.others[index].userName);
      this.textarea = textArea.value; // 要同步data中的数据
      return;
    },

    listen(event) {
      if (event.which === 8) {
        //删除时，@整体删除
        var textArea = document.getElementById("texttexttext");
        var msgLen = textArea.selectionStart;
        var lastword = this.textarea.substring(msgLen);
        for (let j in this.aiteList) {
          var loc1 = msgLen - this.aiteList[j].length - 2;
          var ailoc = 0;
          var loc2 = loc1;
          if (this.textarea[loc1] == "@") {
            loc1++;
            while (
              this.textarea[loc1] == this.aiteList[j][ailoc] &&
              ailoc < this.aiteList[j].length
            ) {
              loc1++;
              ailoc++;
            }
            if (ailoc == this.aiteList[j].length) {
              this.textarea =
                this.textarea.substring(0, loc2 - 1) + lastword + " ";
            }
          }
        }
      }
    },
  },
};
</script>

<style lang="scss">
/* el-popover是和app同级的，所以scoped的局部属性设置了无效 */
/* 需要设置全局style */
.el-popover {
  height: 200px;
  width: 400px;
  overflow: scroll;
  overflow-x: auto;
}
</style>

<style scoped>
.chatIcon {
  padding: 0 10px;
  font-size: 25px;
}
.emotionList {
  display: flex;
  flex-wrap: wrap;
  padding: 5px;
}

.topicList {
  display: block;
  flex-wrap: wrap;
}

.emotionSelect {
  float: left;
}
.emotionItem {
  width: 10%;
  font-size: 20px;
}

.othersone {
  margin-bottom: 5px;
  border-bottom: 1px solid #d3d3d3;
}

/*包含以下四种的链接*/
.emotionItem {
  text-decoration: none;
}
/*正常的未被访问过的链接*/
.emotionItem:link {
  text-decoration: none;
}
/*已经访问过的链接*/
.emotionItem:visited {
  text-decoration: none;
}
/*鼠标划过(停留)的链接*/
.emotionItem:hover {
  text-decoration: none;
}
/* 正在点击的链接*/
.emotionItem:active {
  text-decoration: none;
}
</style>
