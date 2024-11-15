<template>
  <div>
    <el-card :body-style="{ padding: '0px', height: '100%' }">
      <div @click="changeDetaildialog">
        <img :src="require('../assets/' + this.card.url)" class="image" />
      </div>
      <div style="padding: 14px">
        <span>{{ card.title }}</span>
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
    <DetailsCom :dailogshow="DetaildialogVisible" ref="children" />
  </div>
</template>

<script>
import DetailsCom from "@/components/DetailsCom.vue";
import GlobalToken from "@/components/GlobalToken.vue";
export default {
  name: "noteCard",
  components: {
    DetailsCom,
  },
  data() {
    return {
      DetaildialogVisible: false,
      card: {
        noteID: this.attrs.noteId,
        url: this.attrs.cover,
        title: this.attrs.title,
        sendurl: this.attrs.portrait, //笔记发送者的头像
        sendname: this.attrs.creatorName, //笔记发送者的用户名
        sendid: this.attrs.creatorId,
        like: this.attrs.likedNum, //点赞数
        islike: false,
      },
      //下面这三个从外面传进来
      userid: GlobalToken.userid,
      username: GlobalToken.username,
      user_head: GlobalToken.portrait,
      require: this.flag,
      //qiuqiuurl:require('../assets/'+this.attrs.cover),
    };
  },
  methods: {
    changeDetaildialog() {
      this.DetaildialogVisible = true;
      this.$refs.children.setvisvile(
        this.DetaildialogVisible,
        this.card.noteID,
        this.card.sendid,
        this.card.sendname,
        this.card.sendurl
      ); //改成setvalue
    },

    changelike() {
      console.log(this.card.title);
      this.card.islike = !this.card.islike;
      if (this.card.islike) this.card.like += 1;
      else this.card.like -= 1;
    },
    settitle(value) {
      this.card.title = value;
    },

    // setfile(value)
    // {
    //   if(this.require==true)
    //   {
    //     this.qiuqiuurl=value
    //     console.log("三点点滴艾迪阿娘")
    //   }
    //   else{
    //     this.uurl='../assets/'+this.card.url,
    //     this.qiuqiuurl=require(this.uurl);
    //   }
    // },
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
