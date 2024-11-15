<template>
  <div class="details">

    <!-- 弹窗 -->
    <el-dialog class="abow_dialog" :visible.sync="dialogVisible" style="border-radius: 100px;" width="80%"
      :append-to-body="true" @close="dialogclose()">


      <el-container>

        <!-- 侧边放笔记图片 -->
        <el-aside width="50%">
          <!-- width="50%" style="display: block;  position: absolute;  left: 0;  top: 0;  bottom: 0;" -->
          <div class="photoshow">
            <!-- 走马灯 -->
            <el-carousel trigger="click" height="500px">
              <el-carousel-item style="display: flex;justify-content: center;align-items: center;" v-for="item in imgList"
                :key="item">
                <!-- <img :src="require('../assets/'+item)" id = "toEnlargeImg" class="image" > -->
                <el-image :src="item" :preview-src-list="imgList"></el-image>
              </el-carousel-item>
            </el-carousel>
          </div>
        </el-aside>

        <el-container>

          <!-- 头像，笔记者名字，我关注TA了吗 -->
          <el-header style=" position: relative; width: 100%; height: 60px;">
            <el-row type="flex" align="middle">
              <el-col :span="3"><el-avatar :size="40" :src="require('../assets/' + headimg)"></el-avatar></el-col>
              <el-col :span="12">
                <div style="font-size: 18px">{{ authorname }}</div>
              </el-col>
              <div v-if="not_me">
                <el-col :span="3"><el-button v-if="isfollowshow" round @click="followButtonclick">{{ followButtonText
                }}</el-button></el-col>
              </div>
            </el-row>
          </el-header>


          <el-main style="position: absolute;  left: 50%;  right: 0;  top: 35px;  bottom: 70px;  overflow-y: scroll;">
            <!-- style="position: absolute;  left: 2000px;  right: 0;  top: 60px;  bottom: 0;  overflow-y: scroll;" -->
            <!-- 标题 -->
            <el-row>
              <el-col :span="24">
                <h2>{{ title }}</h2>
              </el-col>
            </el-row>
            <!-- 正文 -->
            <el-row>
              <p style="font-size: 18px">{{ contentList[0] }}</p>
              <div v-for="(item, index) in aiteList" :key="index" style="display:flex;flex-wrap: wrap;font-size: 18px">
                <el-link style="font-size: 18px" :underline="false" @click="To_personalView(index)">@{{ item }}</el-link>
                <p>{{ contentList[index + 1] }}</p>
              </div>
            </el-row>
            <!-- 发布时间 -->
            <el-row>
              <div class="timeclass">{{ createtime }}&nbsp;&nbsp;{{ location }}</div>
            </el-row>
            <!-- 分割线 -->
            <el-divider></el-divider>

            <!-- 评论 -->
            <!-- 评论数 -->
            <el-row>
              <div class="timeclass">共{{ (comments) ? comments.length : 0 }}条评论</div>
            </el-row><br>
            <!-- 一条条评论 -->
            <div v-for="(item, index) in comments" :key="index">
              <el-container>
                <el-aside width="40px">
                  <div class="block"><el-avatar :size="40" :src="require('../assets/' + item.portrait)"></el-avatar></div>
                </el-aside>
                <el-container>
                  <el-header>
                    <el-row type="flex" align="middle">
                      <h4>{{ item.commentatorName }}</h4>
                    </el-row>
                  </el-header>
                  <el-main style="padding:0px;margin-left:18px">
                    <span>{{ item.content }}</span>

                    <!-- <p style="font-size: 18px">{{commentList[index][0]}}</p>
                <div  v-for="(item, littleindex) in comments[index].atName" :key="littleindex" style = "display:flex;flex-wrap: wrap;font-size: 18px">
                  <el-link style="font-size: 18px" :underline="false">@{{ item }}</el-link>
                  <p>{{ commentList[index][littleindex+1] }}</p>
                </div> -->
                    <!-- <div v-if="is_new_value_function(index)">
                <p style="font-size: 8px">{{pinglun.neirong}}</p>
                <div v-if="comments[index].atName">
                <div  v-for="(item, littleindex) in comments[index].atName" :key="littleindex" style = "display:flex;flex-wrap: wrap;font-size: 18px">
                  <el-link style="font-size: 18px" :underline="false">@{{ item }}</el-link>
                  <p>{{get_next_comment(index,littleindex)}}</p>
                </div>
                </div>
                </div> -->

                    <el-row type="flex" align="middle">
                      <el-col :span="24">
                        <span style="font-size:10px">{{ item.commentTime }}</span>
                      </el-col>
                      <!-- <el-col :span="2">
                      <el-button :icon="item.cislike ? 'el-icon-solidheart' : 'el-icon-heart'"  type = "text" size="mini" @click="commentchangelike(item)" ></el-button>
                    </el-col>
                    <el-col :span="2">
                      <span class="likespan">{{item.clikenum}}</span>
                    </el-col> -->

                    </el-row>

                    <el-divider></el-divider>
                  </el-main>
                </el-container>
              </el-container>
            </div>

          </el-main>

          <el-footer style="position: absolute;	bottom: 50px;width:50%; background-color:white ;">
            <el-row type="flex" align="middle">
              <el-col :span="1">
                <el-button :icon="islike ? 'el-icon-solidheart' : 'el-icon-heart'" type="text" size="mini"
                  @click="changelike"></el-button>
              </el-col>
              <el-col :span="2" :offset="0">
                <span class="likespan">{{ like }}</span>
              </el-col>
              <el-col :span="1" :offset="1">
                <el-button :icon="iscollect ? 'el-icon-star-on' : 'el-icon-star-off'" type="text" size="mini"
                  style="font-size:22px" @click="changecollect"></el-button>
              </el-col>
              <el-col :span="2" :offset="0">
                <span class="likespan">{{ collect }}</span>
              </el-col>
              <el-col :span="2" :offset="1">
                <el-button type="text" style="font-size:22px" icon="el-icon-chat-dot-square"
                  @click="commentclick"></el-button>
              </el-col>
              <span class="likespan">{{ (this.comments) ? this.comments.length : 0 }}</span>
            </el-row>

            <el-row type="flex" align="middle">
              <el-col :span="16">
                <el-input v-model="mycomment" class="commentclass" resize="none" placeholder="✎说点什么" id='commentid'
                  maxlength="30" show-word-limit rows="1" ref="commentinput" @keydown.native="listen($event)">
                </el-input>
              </el-col>
              <el-col :span="2">
                <div class="chatIcon">
                  <el-popover placement="top-start" width="400" trigger="click" class="emoBox">
                    <div class="emotionList">
                      <a href="javascript:void(0);" @click="getEmo(index)" v-for="(item, index) in faceList" :key="index"
                        class="emotionItem">{{ item }}</a>
                    </div>
                    <el-button style="padding-left:10px;font-size: 25px;" slot="reference" type="text"
                      class="emotionSelect">☺</el-button>
                  </el-popover>
                </div>
              </el-col>
              <el-col :span="2">
                <div class="others">
                  <el-popover placement="top-start" width="400" trigger="click" class="othersBox">
                    <div class="othersList">
                      <div href="javascript:void(0);" @click="getothers(index)" v-for="(item, index) in others"
                        :key="index" class="othersone">
                        <el-row type="flex" align="middle">
                          <el-avatar :size="40" :src="require('../assets/' + item.portrait)"></el-avatar>&nbsp;&nbsp;{{
                            item.userName }}
                        </el-row>
                      </div>
                    </div>
                    <el-button style="padding-left:5px;font-size: 18px;" slot="reference" type="text">@</el-button>
                  </el-popover>
                </div>
              </el-col>
              <el-col :span="4">
                <el-button style="display:block;margin:0 auto ;align:middle" round @click="commentsend">发送</el-button>
              </el-col>
            </el-row>

          </el-footer>

        </el-container>

      </el-container>
    </el-dialog>




  </div>
</template>

<script>
import GlobalToken from '@/components/GlobalToken.vue'
const appData = require(`@/assets/emojis.json`);
export default {
  name: 'DetailsCom',
  data() {
    return {
      dialogVisible: false,
      not_me: true,
      followButtonText: "",
      contentText: "",
      fit: 'contain',//头像
      newCommentNum: 0,
      contentList: [],
      commentList: [],
      faceList: [],

      others: [
        {
          userName: "Linyx",
          portrait: '1.jpg',
        },
        {
          userName: "yzt",
          portrait: '2.jpg',
        },
        {
          userName: "xy",
          portrait: '3.jpg',
        }
      ],

      //需要从别的页面传过来
      noteID: 127,
      mycomment: "",//这个不用传
      commentaiteList: [],//这个不用传
      commentaitelocationList: [],//这个不用传
      myid: GlobalToken.userid,//看评论的用户ID
      myname: GlobalToken.username,
      myportrait: GlobalToken.portrait,
      athorid: 1,
      authorname: "Linyx",
      headimg: "4.jpg",


      imgList: [require('../assets/4.jpg'), 'https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg', 'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg', 'https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg',
      ],
      isfollow: true,
      title: "",
      content: "",
      aiteList: [],//已选的@的人
      aitelocationList: [],
      aiteidList:[],
      createtime: "",
      location: "",
      like: 15,//点赞数
      islike: false,
      collect: 10,
      iscollect: false,
      commentnum: 21,

      comments: [
        {
          commentId: 2,
          commentatorName: "",
          portrait: "2.jpg",
          content: "",
          commentTime: "",
          atName: [],
          atLocation: [],
        },
      ]
    }
  },



  mounted() {
    for (let i in appData) {
      this.faceList.push(appData[i].char);
    }
  },

  methods: {

    To_personalView(index)
      {
        var personalid=this.aiteidList[index]
        //跳到个人主页，传this.myid过去
        if(this.myid==personalid)
        {
          this.dialogVisible=false
          this.$router.push({path: '/person?detailId='+this.myid});
        }
        //跳到otherPersonal,传personalid,和this.myid
        else
        {
          this.dialogVisible=false
          this.$router.push({path: '/otherPerson?other_watchid='+this.myid+'&other_userid='+personalid});
        }
      },
    listen(event) {
      if (event.which === 8) { //删除时，@整体删除
        var textArea = document.getElementById("commentid");
        var msgLen = textArea.selectionStart;
        var lastword = this.mycomment.substring(msgLen);
        if (!this.commentaiteList) return
        for (let j in this.commentaiteList) {
          var loc1 = msgLen - this.commentaiteList[j].length - 2;
          var ailoc = 0;
          var loc2 = loc1;
          if (this.mycomment[loc1] == "@") {
            loc1++;
            while (this.mycomment[loc1] == this.commentaiteList[j][ailoc] && ailoc < this.commentaiteList[j].length) {
              loc1++;
              ailoc++;
            }
            if (ailoc == this.commentaiteList[j].length) {
              this.mycomment = this.mycomment.substring(0, loc2 - 1) + lastword + " ";
            }
          }
        }
      }
    },

    setvisvile(value, noteid,  biji_id, biji_name, biji_head) {
      this.noteID = noteid,
        this.myid = GlobalToken.userid,
        this.myname = GlobalToken.username,
        this.myportrait = GlobalToken.portrait,
        this.athorid = biji_id,
        this.authorname = biji_name
        this.headimg = biji_head
        this.load_data(value)
    },

    getEmo(index) {//按下表情中的某一个才执行
      const commentArea = document.getElementById("commentid");
      function changeSelectedText(obj, str) {
        if (window.getSelection) {
          // 非IE浏览器
          commentArea.focus();
          commentArea.setRangeText(str);
          // 在未选中文本的情况下，重新设置光标位置
          commentArea.selectionStart += str.length;
          commentArea.focus()
        } else if (document.selection) {
          // OE浏览器
          obj.focus();
          var sel = document.selection.createRange();
          sel.text = str;
        }
      }
      changeSelectedText(commentArea, this.faceList[index]);

      this.mycomment = commentArea.value;// 要同步data中的数据
      return;
    },

    getothers(index) {
      const commentArea = document.getElementById("commentid");
      function changeSelectedText(obj, str) {
        if (window.getSelection) {
          // 非IE浏览器
          commentArea.focus();
          commentArea.setRangeText(" @" + str + " ");
          // 在未选中文本的情况下，重新设置光标位置
          commentArea.selectionStart += (str.length + 3);
          commentArea.focus()
        } else if (document.selection) {
          // IE浏览器
          obj.focus();
          var sel = document.selection.createRange();
          sel.text = str;
        }
      }
      this.commentaiteList.push(this.others[index].userName);
      this.commentaitelocationList.push(this.mycomment.length);
      changeSelectedText(commentArea, this.others[index].userName);
      this.mycomment = commentArea.value;// 要同步data中的数据
      return;
    },

    async load_data(value) {
      this.comments = []
      this.imgList = []
      await this.$axios.get('http://localhost:8081/' + this.myid + '/explore/' + this.noteID)
        .then(response => {
          console.log("response", response.data.data)
          this.authorid = response.data.data.NoteInfo.creatorId;
          this.title = response.data.data.NoteInfo.title;
          this.content = response.data.data.NoteInfo.body;
          this.createtime = response.data.data.NoteInfo.createtime;
          this.location = response.data.data.NoteInfo.location;
          var plist = response.data.data.PicsOfNote;
          this.aiteList = response.data.data.NoteInfo.atName
          this.aitelocationList = response.data.data.NoteInfo.atLocation
          this.aiteidList=response.data.data.NoteInfo.atUserId
          this.like = response.data.data.NoteInfo.likedNum
          this.islike = response.data.data.NoteInfo.isLiked
          this.collect = response.data.data.NoteInfo.collectNum
          this.iscollect = response.data.data.NoteInfo.isCollected
          this.isfollow = response.data.data.NoteInfo.isFollowed

          if (this.authorid == this.myid) {
            this.not_me = false
          }
          else {
            this.not_me = true
          }

          //设置按钮内容
          if (this.isfollow) {
            this.followButtonText = "已关注";
          }
          else {
            this.followButtonText = "关注"
          }

          this.contentList = []//要清空，不然会读到初始值
          //设置contentText//分段
          if (this.aitelocationList) {
            var num = this.aitelocationList.length;
            var loc1 = -1;
            for (let i = 0; i < num; i++) {
              var loc2 = this.aitelocationList[i];
              this.contentList.push(this.content.substring(loc1 + 1, loc2))
              loc1 = loc2 + this.aiteList[i].length + 1;
            }
            this.contentList.push(this.content.substring(loc1 + 1))
          }
          else {
            this.contentList.push(this.content)
          }

          for (let i = 0; i < plist.length; i++) {
            this.imgList.push(require('../assets/' + plist[i]))

          }
        })

      this.$axios.get('http://localhost:8081/comment/' + this.noteID)
        .then(response => {
          console.log(response)
          this.comments = response.data.data;
          this.is_new_value = true;
          if (this.comments) this.is_new_value_add = this.comments.length;

        })

      if (GlobalToken.Token != '') {
        this.$axios({
          method: 'get',
          url: 'http://localhost:8081/' + this.myid + '/follow',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
        }).then(response => {
          console.log(response)
          this.others = response.data.data;
        })
      }
      else {
        this.$message.error("请先登录！")
        this.others = []
      }

      this.dialogVisible=value

    },

    dialogclose() {
      this.dialogVisible = false

    },

    async getCommentdata() {
      await this.$axios.get('http://localhost:8081/comment/' + this.noteID)
        .then(response => {
          console.log(response)
          this.comments = response.data.data;
          this.is_new_value = true
        })
    },

    async commentsend() {
      if (GlobalToken.Token == '')//没权限
      {
        this.$message.error("请先登录！")
        return
      }
      var date = new Date();
      var year = date.getFullYear();
      var month = date.getMonth() + 1 < 10 ? "0" + (date.getMonth() + 1) : date.getMonth() + 1;
      var day = date.getDate() < 10 ? "0" + date.getDate() : date.getDate();
      var hours = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
      var minutes = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
      var seconds = date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
//忘加token!!!!!!
      let form =new FormData()
      form.append('commentatorId',this.myid)
      form.append('commentatorName',this.myname)
      form.append('content',this.mycomment)
      form.append('commentTime',year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds)
      form.append('portrait',this.myportrait)
      form.append('atName',this.commentaiteList)
      form.append('atLocation',this.commentaitelocationList)
    this.$axios({
         method: 'post',
        url: 'http://localhost:8081/comment/' + this.noteID,
        headers: {
          'Content-Type': false,
          'Authorization': GlobalToken.Token,
        },
        data:form
      }).
      then(response => {
        this.token = response.data.token;
        this.is_new_value = false;
        this.mycomment = "";
        this.getCommentdata()
      })
    },

    commentclick() {
      if (GlobalToken.Token == '')//没权限
      {
        this.$message.error("请先登录！")
        return
      }
      this.$refs.commentinput.focus();
    },

    changelike() {
      if (GlobalToken.Token == '')//没权限
      {
        this.$message.error("请先登录！")
        return
      }
      var date = new Date();
      var year = date.getFullYear();
      var month = date.getMonth() + 1 < 10 ? "0" + (date.getMonth() + 1) : date.getMonth() + 1;
      var day = date.getDate() < 10 ? "0" + date.getDate() : date.getDate();
      var hours = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
      var minutes = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
      var seconds = date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
      let form = new FormData()
      this.islike = !this.islike;
      if (this.islike) {
        this.like += 1;
        form.append('userAct',this.myid )
        form.append('likeTime', year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds)
        this.$axios({
          method: 'post',
          url: 'http://localhost:8081/explore/' + this.noteID + '/like',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
          data:form

        }).then(response => {
          this.token = response.data.token;
          console.log(response)
        })
      }
      else {
        this.like -= 1;
        form.append('userAct',this.myid )
        this.$axios({
          method: 'delete',
          url: 'http://localhost:8081/explore/' + this.noteID + '/like',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
          data: form
        }).then(response => {
          this.token = response.data.token;
          console.log(response)
        })
      }

    },

    commentchangelike(item) {
      item.cislike = !item.cislike;
      if (item.cislike)
        item.clikenum += 1;
      else
        item.clikenum -= 1;
    },
    changecollect() {
      if (GlobalToken.Token == '')//没权限
      {
        this.$message.error("请先登录！")
        return
      }
      let form=new FormData()
      this.iscollect = !this.iscollect;
      if (this.iscollect) {
        this.collect += 1;
        form.append('userAct',this.myid)
        this.$axios({
          method: 'post',
          url: 'http://localhost:8081/explore/' + this.noteID + '/collect',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
          data:form,
        }).then(response => {
          this.token = response.data.token;
          console.log(response)
        })
      }
      else {
        this.collect -= 1;
        form.append('userAct',this.myid)
        this.$axios({
          method: 'delete',
          url: 'http://localhost:8081/explore/' + this.noteID + '/collect',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
          data: form,
        }).then(response => {
          this.token = response.data.token;
          console.log(response)
        })
      }

    },

    isfollowshow() {
      if (this.isfollow) {
        this.followButtonText = "已关注"
      }
      else {
        this.followButtonText = "关注"
      }
      return this.isfollow;
    },
    followButtonclick() {
      if (GlobalToken.Token == '')//没权限
      {
        this.$message.error("请先登录！")
        return
      }
      
      let form = new FormData();
      form.append('followID',this.authorid)

      if (this.isfollow) {
        this.followButtonText = "关注"
        this.$axios({
          method: 'delete',
          url: 'http://localhost:8081/' + this.myid + '/PersonalView/follow',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
          data:form
        }).then(response => {
          this.token = response.data.token;
          console.log(response)
        })
        this.isfollow = !this.isfollow
      }
      else {
this.followButtonText = "已关注"
this.$axios({
          method: 'post',
          url: 'http://localhost:8081/' + this.myid + '/PersonalView/follow',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
          data:form
        }).then(response => {
          this.token = response.data.token;
          console.log(response)
        })
this.isfollow = !this.isfollow}
    },

  }
}
</script>
<style>
.el-icon-solidheart {
  background: url('../assets/solidheart.jpg') center no-repeat;
  width: 2rem;
  height: 1.5rem;
}

.el-icon-solidheart:before {
  content: '切换';
  font-size: 0;
  visibility: hidden;
}

.el-icon-heart {
  background: url('../assets/heart.jpg') center no-repeat;
  width: 2rem;
  height: 1.5rem;
}

.el-icon-heart:before {
  content: '切换';
  font-size: 0;
  visibility: hidden;
}

.abow_dialog {
  display: flex;
  justify-content: center;
  align-items: Center;
  overflow: hidden;

  .el-dialog {
    margin: 0 auto !important;
    height: 550px;
    overflow: hidden;

    .el-dialog__body {
      position: absolute;
      left: 0;
      top: 20px;
      bottom: 0;
      right: 0;
      padding: 0;
      z-index: 1;
      overflow: hidden;
      overflow-y: auto;
    }
  }
}
</style>
<style scoped>
.el-row {
  display: flex;
  flex-wrap: wrap;
}

.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

.photoshow {
  width: 100%;
  height: 100%;
}

/* img{
width: 100%;
height: 100%;
object-fit: contain;
} */
commentclass {
  background-color: #bfabab
}

.el-input__inner {
  background-color: #bfabab
}

.user_name {
  letter-spacing: 1px;
  margin-top: 0px;
  margin-bottom: 1px;
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

.emotionSelect {
  float: left;
}

.emotionItem {
  width: 10%;
  font-size: 20px;
}


.othersone {
  margin-bottom: 5px;
  border-bottom: 1px solid #D3D3D3;
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