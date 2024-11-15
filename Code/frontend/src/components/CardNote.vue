<template>
  <div>
    <el-card :body-style="{ padding: '0px', height: '100%' }">
      <div>
        <img :src="qiuqiuurl" class="image" />
      </div>
      <div class="div">
        <span>{{ card.title }}</span>
      </div>
      <div>
        <el-row type="flex" align="middle" class="bottom clearfix">
          <el-avatar
            :size="40"
            :src="require('../assets/' + card.sendurl)"
          ></el-avatar>
          <span class="senduser">{{ card.sendname }}</span>
          <el-button
            :icon="card.islike ? 'el-icon-solidheart' : 'el-icon-heart'"
            type="text"
            size="mini"
            style="width: 4rem"
            @click="changelike"
          ></el-button>
          <span class="likespan">{{ card.like }}</span>
        </el-row>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "cardNote",
  data() {
    return {
      card: {
        noteID: this.attrs.noteId,
        url: this.attrs.cover,
        title: this.attrs.title,
        sendurl: this.attrs.portrait, //笔记发送者的头像
        sendname: this.attrs.creatorName, //笔记发送者的用户名
        like: this.attrs.likedNum, //点赞数
        islike: false,
      },
      require: this.flag,
      qiuqiuurl: "",
    };
  },
  methods: {
    changelike() {
      this.card.islike = !this.card.islike;
      if (this.card.islike) this.card.like += 1;
      else this.card.like -= 1;
    },

    settitle(value) {
      this.card.title = value;
    },

    setfile(value) {
      if (this.require == true) {
        this.qiuqiuurl = value;
      } else {
        this.qiuqiuurl = require("../assets/" + this.card.url);
      }
    },
  },
  props: {
    attrs: {
      type: Object,
      default() {
        return {};
      },
    },
    flag: {
      type: Boolean,
      default() {
        return false;
      },
    },
  },
};
</script>
<style>
.div {
  height: 100px;
  line-height: 170px;
}
.el-icon-solidheart {
  background: url("../assets/solidheart.jpg") center no-repeat;
  height: 3rem;
  width: 3rem;
}
.el-icon-solidheart:before {
  content: "切换";
  font-size: 0;
  visibility: hidden;
}
.el-icon-heart {
  background: url("../assets/heart.jpg") center no-repeat;
  width: 3rem;
  height: 3rem;
}
.el-icon-heart:before {
  content: "切换";
  font-size: 0;
  visibility: hidden;
}
.bottom {
  margin-top: 13px;
  line-height: 12px;
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
  clear: both;
}
</style>
