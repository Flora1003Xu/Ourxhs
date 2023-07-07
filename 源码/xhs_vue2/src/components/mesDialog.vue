<template>
<el-dialog
          :visible.sync="mesIsVisible"
          width="35%"
          :before-close="handleClose"
          :close-on-click-modal="false"
          @open="getMes()"
          class="mesDialog">
          <el-tabs v-model="activeName" 
                  @tab-click="handleClick">
                    <el-tab-pane label="评论" name="first">
                      <div v-for="(mes,index) in mesList" :key="index">
            <el-container>
          <el-aside width="60px">
            <div class="block" style="margin-top: 10px;margin-left: 0px;padding-left: 0px;">
              <el-badge :is-dot="mes.state==0">
              <el-avatar :size="45" :src="require('../assets/'+mes.portrait)" style="padding-left: 0px;"></el-avatar></el-badge></div>
          </el-aside>
          <el-container style="height:100%">
            <el-header style="padding-top: 0px;">
              <div  @click="mes.state=1,openDetail(mes)">
              <el-row type="flex" align="middle">
                <h3 style="margin: 0px;">{{mes.commentatorName}}</h3>
              </el-row></div>
            </el-header>

            <el-main style="padding: 0px;margin-left: 18px;">
              <div  @click="mes.state=1,openDetail(mes)">
              <el-row type="flex" align="left">
                <span class="span" style="font-size: 14px;">{{ mes.content }}</span>
              </el-row> </div>                 
            </el-main>

            <el-footer style="padding-bottom: 0px;height: 100%;">
              <el-row type="flex" align="left">
                <span class="date" style="font-size: 12px;">{{ mes.commentTime }}</span>
              </el-row>
              
            </el-footer>
          </el-container>
        </el-container>
        <div class="bottomLine"></div>
        </div>
                    </el-tab-pane>
                    <el-tab-pane label="点赞" name="second">
                      <div v-for="(like,index) in likeList" :key="index">
            <el-container>
          <el-aside width="60px">
            
            <div class="block" style="margin-top: 10px;">
              <el-badge :is-dot="like.state==0">
              <el-avatar :size="45" :src="require('../assets/'+like.portrait)" style="padding-left: 0px;"></el-avatar></el-badge></div>
          </el-aside>
          <el-container style="height:100%">
            <el-header style="padding-top: 0px;">
              <div  @click="like.state=1,likeDetail(like)">
              <el-row type="flex" align="middle">
                <h3 style="margin: 0px;">{{like.userName}}</h3>
              </el-row></div>
            </el-header>

            <el-main style="padding: 0px;margin-left: 18px;">
              <div  @click="like.state=1,likeDetail(like)">
              <el-row type="flex" align="left">
                <span class="span" style="font-size: 14px;">点赞了你的笔记</span>
              </el-row></div> 
               <!-- 笔记详情 -->
            <DetailsCom :dailogshow="DetaildialogVisible" ref="children"/>                 
            </el-main>


           

            <el-footer style="padding-bottom: 0px;height: 100%;">
              <el-row type="flex" align="left">
                <span class="date" style="font-size: 12px;">{{ like.likeTime }}</span>
              </el-row>
              
            </el-footer>
          </el-container>
        </el-container>
        <div class="bottomLine"></div></div>
                    </el-tab-pane>
                  </el-tabs>
          
        </el-dialog>
</template>

<script>
import axios from 'axios'
import DetailsCom from '../components/DetailsCom.vue'
export default{
  components:{
    DetailsCom
  },
  data(){
    return{
      activeName:"first",
      DetaildialogVisible:false,

    }
  },
  name:'mesDialog',
  props:["mesList","Visible","likeList","userId","userName","portrait"],
  computed:{
    mesIsVisible:{
      get(){
        return this.Visible
      },
      set(val){
        this.$emit('updateVisible',val)
      }
    }
  },
  methods:{
    getMes(){
        axios.get('http://localhost:8080/messages/'+this.userId+'/comments')
        .then(response=>{
          this.message=response.data.data;
          console.log(response);
        })
      },
    handleClose(done) {
      console.log(this.mesList)
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {});
          
      },
      handleClick(tab, event) {
        console.log(tab, event);
      },
      openDetail(mes){
        console.log(mes)
        this.DetaildialogVisible=true;
        console.log(111,this.DetaildialogVisible)
        console.log(this.$refs.children)
        //this.$refs.children.setvisvile(this.DetaildialogVisible);

        //this.$nextTick(() => {
      this.$refs.children[0].setvisvile(this.DetaildialogVisible);
        axios({
              method:'put',
              url: 'http://localhost:8080/messages/'+this.userId+'/comments/'+mes.commentId,
              }).then(res => {})

      },
      likeDetail(like){
        this.DetaildialogVisible=true;
        this.$refs.children[0].setvisvile(this.DetaildialogVisible);
        console.log("传值有无人体？",this.like.noteId,this.userId,this.userName,this.portrait)
        this.$refs.children[0].setvalue(this.like.noteId,this.userId,this.userName,this.portrait)
       // noteID=like.noteId;
      //  userId
//userName
      //  portrait
        axios({
              method:'put',
              url: 'http://localhost:8080/messages/'+this.userId+'/likes/'+like.likeId,
              }).then(res => {})
      }
  }
}
</script>

<style scoped>
  .bottomLine {
    width: 400px;
    border: 0.5px solid rgba(230, 224, 224, 0.966);
    margin-top: 3px;
    margin-left:40px;
    margin-bottom: 0px;
  }
  .span {
    text-align: left;
            display: inline-block;
            word-break: break-all;
            white-space: small;
            line-height: 18px;
        }
  .date{
    text-align: left;
    margin-top: 2px;
    color: rgb(129, 138, 148);
    line-height: 20px;
  }

</style>
<style>
  .mesDialog.el-dialog__body{
    padding: 0px;
  }
</style>