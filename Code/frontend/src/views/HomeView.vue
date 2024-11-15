<template>
  <div id="home">
    <!--放小红书logo-->
    <el-container>
      <el-aside style="width:180px;">
        <!-- 设置当前菜单状态跟路由保持一致 -->
        <el-menu :default-active="$route.path" :text-clor="color">
          <el-menu-item :class="index == number ? 'btn1' : 'btn'" @click="tab(index)" v-for="(item, index) in dataList"
            :key="index" style="text-align:center">{{ item.option }}
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header>

          <div>
            <!--elementui有24分栏，span等于几就占据了几分栏，offset为偏移量，el-col里面得有div分隔才生效-->
            <el-row>

              <el-col :span="2" :offset="0" class="left">
                <div><img src="./1.png" width="120" height="40" alt="小红书" /></div>
              </el-col>

              <!--设置搜索栏,组件插入，减少本页代码量-->
              <el-col :span="7" :offset="2" class="center">
                <div>
                  <SearchBox v-on:handleSearch="handleSearch"></SearchBox>
                </div>
              </el-col>

              <el-col :span="3" :offset="1" class="left">
                <el-button round icon="el-icon-s-home" size="medium" @click="finded" :plain="find" type="danger"> 发现
                </el-button>
              </el-col>

              <el-col :span="3" :offset="0" class="left">
                <el-button round icon="el-icon-circle-plus-outline" size="medium" @click="post" :plain="posted"
                  type="danger"> 发布 </el-button>
              </el-col>

              <el-col :span="3" :offset="0" class="left">
                <el-button round icon="el-icon-tableware" size="medium" @click="mine" :plain="my" type="danger" refs="my">
                  我的 </el-button>
              </el-col>

              <!-- 弹窗 -->
              <el-col :span="3" :offset="0" class="left">
                <!-- 点击button进入外层 -->
                <div id="log">
                  <el-button size="medium" @click="loginVisible = true" :plain="my" type="danger"> 登录 </el-button>
                </div>
                <!-- 这里是登录后换成头像 -->
                <el-avatar v-if="head_visible" :src="require('../assets/' + portrait)" :size="50"></el-avatar>
                <loginDialog :Visible="loginVisible" @change="handleChange($event)" @updateVisible="updateVisible" />
              </el-col>

            </el-row>
          </div>
        </el-header>

        <el-main class="tab">
          <!-- 11个div可以进行切换，number=0是默认第一个div -->

          <div v-show="number == 0">

            <!-- 下面的模块是走马灯 -->
            <div v-if="value_show">
              <el-carousel :interval="4000" type="card" height="300px" :autoplay="true">
                <el-carousel-item v-for="item in imgList" :key="item.id">
                  <div class="pic_item">
                    <div @click="zoumadeng(item.id)">
                      <img class="image" :src="require('../assets/' + item.cover)" />
                    </div>
                    <!-- 显示走马灯详情 -->
                    <DetailsCom :dailogshow="DetaildialogVisible" ref="children" />
                    <h3>{{ item.title }}</h3>
                  </div>
                </el-carousel-item>
              </el-carousel>
            </div>

            <!-- 下面的模块是图片三列排布之推荐页-->

            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[0]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image" ></noteCard>
                  <!-- <img :src="image.url" class="image"> -->
                </div>
              </div>
            </div>

          </div>

          <!-- 我的关注 -->
          <div v-show="number == 1">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[1]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>

          <!-- 美食 -->
          <div v-show="number == 2">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[2]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>

          <div v-show="number == 3">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[3]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>
          <div v-show="number == 4">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[4]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>
          <div v-show="number == 5">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[5]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>
          <div v-show="number == 6">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[6]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>
          <div v-show="number == 7">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[7]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>
          <div v-show="number == 8">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[8]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>
          <div v-show="number == 9">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[9]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>
          <div v-show="number == 10">
            <div style="display: flex" class="average_container">
              <div v-for="(column, index) in columns[10]" :key="index">
                <div v-for="image in column" :key="image.noteId">
                  <noteCard :attrs="image"></noteCard>
                </div>
              </div>
            </div>
          </div>

        </el-main>

      </el-container>
    </el-container>
  </div>
</template>

<script>
import loginDialog from "../components/loginDialog.vue"
import axios from "axios";
import SearchBox from '@/components/SearchBox.vue'
import noteCard from '@/components/NoteCard.vue'
import DetailsCom from '@/components/DetailsCom.vue'
import GlobalToken from '@/components/GlobalToken.vue'
export default {
  name: 'HomeView',
  components: {
    SearchBox,
    noteCard,
    loginDialog,
    DetailsCom,
  },
  data() {
    return {
      number: 0,//设置菜单的默认选中值
      find: false,
      posted: true,
      my: true,
      image: {},
      noteVisible: false,
      loginVisible: false,
      head_visible: false,
      DetaildialogVisible: false,
      //登录后获取id和用户名和头像
      userId: GlobalToken.userid,
      username: GlobalToken.username,
      portrait: GlobalToken.portrait,

      index_now: 0,
      value_show: true,
      dataList: [
        { option: '推荐' },
        { option: '我的关注' },
        { option: '美食' },
        { option: '彩妆' },
        { option: '穿搭' },
        { option: '影视' },
        { option: '情感' },
        { option: '职场' },
        { option: '家居' },
        { option: '游戏' },
        { option: '旅行' },
      ],
      //走马灯的
      list: [],
      imgList: [],
      screenWidth: 0,
      color: "black",
      //图片和列,note和columns的个数相对应
      note: [[], [], [], []],
      columns: [[[], [], []], [[], [], []], [[], [], []], [[], [], []]],
    };

  },

  methods: {
    zoumadeng(index) {
      this.DetaildialogVisible = true
      this.$refs.children[0].setvisvile(this.DetaildialogVisible, this.imgList[index].noteId, this.imgList[index].creatorId, this.imgList[index].creatorName, this.imgList[index].portrait);//改成setvalue
    },
    updateVisible() {
      this.loginVisible = !this.loginVisible
    },
    // 登录后更新个人信息
    async handleChange(newuserId) {
      GlobalToken.userid = newuserId;
      this.userId=GlobalToken.userid
      axios({
        method: 'get',
        url: 'http://localhost:8081/' + this.userId+'/' + this.userId +'/PersonalView',
        headers: {
          'Content-Type': false,
          'Authorization': GlobalToken.Token,
        },
      }).then(response => {
        GlobalToken.portrait= response.data.data.userInfo.portrait;
        GlobalToken.username = response.data.data.userInfo.userName;
        this.loginVisible = false;//
        var el = document.getElementById('log');
        el.remove();
      this.portrait=GlobalToken.portrait
      this.head_visible = true;
      })
    },
    //用于防止图片叠加
    delBlock(index) {
      this.columns.splice(index, 1);
    },
    addBlock(index) {
      this.columns.splice(index, 1, [[], [], []]);
    },
    //下面是走马灯图片自适应的第一个函数，通过图片宽度计算高度
    setSize: function () {
      // 通过浏览器宽度(图片宽度)计算高度
      this.bannerHeight = 400 / 1920 * this.screenWidth;
    },
    //走马灯点击跳转的事件
    //tab点击到类别就刷新
    tab(index) {
      this.index_now = index;
      this.delBlock(index);
      this.addBlock(index);
      this.number = index;
      if (index == 0) {
        this.value_show = true;
        //this.fetchImages();
        this.splitImagesIntoColumns(0);
      }
      if (index == 1) {
        //我的关注的请求接口
        axios({
          method: 'get',
          url: 'http://localhost:8081/' + this.userId + '/follow/notes',
          headers: {
            'Content-Type': false,
            'Authorization': GlobalToken.Token,
          },
        }).then(res => {
          this.note[1] = res.data.data.notes;
          this.splitImagesIntoColumns(1)
        })
          .catch(error => {
            console.log(error)
          })
      }

      if (index == 2) {
        this.classify('美食');
      }
      if (index == 3) {
        this.classify('彩妆');
      }
      if (index == 4) {
        this.classify('穿搭');
      }
      if (index == 5) {
        this.classify('影视');
      }
      if (index == 6) {
        this.classify('情感');
      }
      if (index == 7) {
        this.classify('职场');
      }
      if (index == 8) {
        this.classify('家居');
      }
      if (index == 9) {
        this.classify('游戏');
      }
      if (index == 10) {
        this.classify('旅行');
      }
    },

    // 即可实现点击按钮的跳转，对应路由里面的path
    finded() {
      this.find = false//背景为透明色
      this.posted = true
      this.my = true
    },
    post() {
      if (GlobalToken.Token != '') {
        this.find = true//背景为透明色
        this.posted = false
        this.my = true
        //this.$router.push('/post');//可实现跳转页面
        this.$router.push({ path: '/post?index=' + GlobalToken.userid+ '&name=' + GlobalToken.username + '&head=' + GlobalToken.portrait });
      }
      else{
        this.$message.error("请先登录！")
      }
    },
    // 跳转到个人主页
    mine() {
      if (GlobalToken.Token != '') {
      this.find = true
      this.posted = true
      this.my = false
      //this.$router.push('/person');
      this.$router.push({ path: '/person?index=' + GlobalToken.userid });}
      else{
        this.$message.error("请先登录！")
      }
    },

    //分类的请求
    classify(cla) {
      axios.get('http://localhost:8081/search/' + cla)
        .then(res => {
          if (cla == '美食') {
            this.note[2] = res.data.data.notes;
            this.splitImagesIntoColumns(2)
          }
          if (cla == '彩妆') {
            this.note[3] = res.data.data.notes;
            this.splitImagesIntoColumns(3)
          }
          if (cla == '穿搭') {
            this.note[4] = res.data.data.notes;
            this.splitImagesIntoColumns(4)
          }
          if (cla == '影视') {
            this.note[5] = res.data.data.notes;
            this.splitImagesIntoColumns(5)
          }
          if (cla == '情感') {
            this.note[6] = res.data.data.notes;
            this.splitImagesIntoColumns(6)
          }
          if (cla == '职场') {
            this.note[7] = res.data.data.notes;
            this.splitImagesIntoColumns(7)
          }
          if (cla == '家居') {
            this.note[8] = res.data.data.notes;
            this.splitImagesIntoColumns(8)
          }
          if (cla == '游戏') {
            this.note[9] = res.data.data.notes;
            this.splitImagesIntoColumns(9)
          }
          if (cla == '旅行') {
            this.note[10] = res.data.data.notes;
            this.splitImagesIntoColumns(10)
          }
        })
        .catch(error => {
          console.log(error)
        })
    },

    //images先定义为向本地请求的照片的名称，即阿拉伯数字1-10
    fetchImages() {
      //向后端发送 GET 请求
      axios.get('http://localhost:8081/explore')
        .then(res => {
          this.note[0] = res.data.data.notes;
          this.splitImagesIntoColumns(0)
        })
        .catch(error => {
          console.log(error)
        })
    },

    //note到时候变成从请求端去获取
    splitImagesIntoColumns(index) {
      let columnCount = 3;
      let notesPerColumn = Math.ceil(this.note[index].length / columnCount);
      let pre = notesPerColumn
      for (let i = 0; i < columnCount; i++) {
        if (i == 2) notesPerColumn = this.note[index].length - notesPerColumn * 2;
        for (let j = 0; j < notesPerColumn; j++) {
          let inde = i * notesPerColumn + j;
          if (i == 2) inde = i * pre + j;
          if (inde < this.note[index].length) {
            // 叫成cover才能只传文件名
            const image = this.note[index][inde];
            this.columns[index][i].push(image);
          }
        }
      }
    },
    //搜索框更新页面
    handleSearch(result) {
      // 接收子组件传递过来的搜索结果，刷新推荐页面的部分
      if (this.index_now == 0) { this.value_show = false }
      this.delBlock(this.index_now);
      this.addBlock(this.index_now);
      this.note[this.index_now] = [];
      this.note[this.index_now] = result;
      this.splitImagesIntoColumns(this.index_now);
    }
  },
  watch: {
    $route() {
      if (this.$route.params.refresh==true) { 
        this.delBlock(0);
        this.addBlock(0);
        this.fetchImages();
        // this.splitImagesIntoColumns(0);
      }
    }
  },
  //这个是走马灯的第二个图片自适应的函数
  mounted() {
    // 首次加载时,需要调用一次
    this.screenWidth = window.innerWidth;
    this.setSize();
    // 窗口大小发生改变时,调用一次
    window.onresize = () => {
      this.screenWidth = window.innerWidth;
      this.setSize();

    }

  },

  //函数里面的会先执行！在加载页面的时候，
  created: function () {
    // 一加载页面就请求一下走马灯
    //走马灯请求接口
    axios.get("http://localhost:8081/explore/tops")
      .then(res => {
        this.list = res.data.data;
        for (let i = 0; i < 4; i++) {
          const yuansu = this.list[i];
          this.$set(yuansu, 'id', i);
          this.imgList.push(yuansu);
        }

      })
      .catch(error => {
        console.log(error)
      })
    //首次加载时去请求所有的笔记
    this.fetchImages();
    this.splitImagesIntoColumns(0);
  }
}
</script>

<style scoped>
.el-container {
  height: 100%;
}

.el-header {
  position: absolute;
  line-height: 50px;
  top: 0px;
  left: 0px;
  right: 0px;
  background-color: #3881ce01;
  margin-top: 0px;
}

.el-footer {
  background-color: #bdbaba0e;
}

.el-aside {
  position: absolute;
  width: 200px;
  top: 50px;
  /* 距离上面50像素 */
  left: 0px;
  bottom: 0px;
  overflow-y: auto;
  /* 当内容过多时y轴出现滚动条 */
  background-color: #def3f927;

}

.el-main {
  position: absolute;
  top: 50px;
  left: 200px;
  bottom: 0px;
  right: 0px;
  /* 距离右边0像素 */
  padding: 10px;
  overflow-y: auto;
  /* 当内容过多时y轴出现滚动条 */
  /* background-color: red; */
}


/* 走马灯的style设置 */
.el-carousel__item h3 {
  color: #fef6f6f5;
  font-size: 15px;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #489ff7;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

img {
  /*设置图片宽度和浏览器宽度一致*/
  width: 100%;
  height: inherit;
}

.pic_item {
  position: relative;
  height: 100%;
}

.pic_item:hover {
  cursor: pointer;
}

.pic_item img {
  width: 100%;
  height: 100%;
}

.pic_item h3 {
  position: absolute;
  left: 0.5rem;
  bottom: -5rem;
}

/* 实现布局三列平分 */
.average_container>div {
  flex: 1;
}

.average_container div {
  border: solid 0px rgba(237, 226, 226, 0.174);
  margin-right: calc(20px / 3);
  box-sizing: border-box;
  margin-top: 10px;
  margin-bottom: 10px;
}</style>

<style>
/*选中时*/
.btn1 {
  background-color: #fff !important;
  color: #ec2929 !important;
  /* 设置选中时的字体颜色和背景颜色 */
}
</style>
