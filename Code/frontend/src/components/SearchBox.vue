<template>
  <div>
    <el-input
      v-model="search"
      @focus="focus"
      @blur="blur"
      @keyup.enter.native="searchHandler"
      placeholder="搜索感兴趣的内容"
    >
      <el-button
        slot="append"
        icon="el-icon-search"
        id="search"
        @click="searchHandler"
      ></el-button>
    </el-input>
    <!---设置z-index优先级,防止被走马灯效果遮挡-->
    <el-card
      @mouseenter="enterSearchBoxHanlder"
      v-if="isSearch"
      class="box-card"
      id="search-box"
      style="position: relative; z-index: 15"
    >
      <dl v-if="isHistorySearch">
        <dt class="search-title" v-show="history">历史搜索</dt>
        <dt class="remove-history" v-show="history" @click="removeAllHistory">
          <i class="el-icon-delete"></i>清空历史记录
        </dt>
        <el-tag
          v-for="search in historySearchList"
          :key="search.id"
          closable
          :type="search.type"
          @close="closeHandler(search)"
          style="margin-right: 5px; margin-bottom: 5px"
          >{{ search.name }}</el-tag
        >
        <dt class="search-title">热门搜索</dt>
        <dd v-for="search in hotSearchList" :key="search.id">{{ search }}</dd>
      </dl>
      <dl v-if="isSearchList">
        <dd v-for="search in searchList" :key="search.id">{{ search }}</dd>
      </dl>
    </el-card>
  </div>
</template>

<script>
import axios from "axios";
import RandomUtil from "@/utils/randomUtil";
import Store from "@/utils/store";
export default {
  data() {
    return {
      search: "", //当前输入框的值，this.search
      isFocus: false, //是否聚焦
      hotSearchList: ["暂无热门搜索"], //热门搜索数据
      historySearchList: [], //历史搜索数据
      searchList: ["暂无数据"], //搜索返回数据,
      history: false,
      types: ["", "success", "info", "warning", "danger"], //搜索历史tag式样
    };
  },
  methods: {
    fetchImages() {
      //向后端发送 GET 请求
      axios
        .get("/explore", { params: { keyword: this.search } })
        .then((res) => {
          console.log(res);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    focus() {
      //鼠标移到这聚焦时显示
      this.isFocus = true;
      this.historySearchList =
        Store.loadHistory() == null ? [] : Store.loadHistory();
      this.history = this.historySearchList.length == 0 ? false : true;
    },
    blur() {
      var self = this;
      this.searchBoxTimeout = setTimeout(function () {
        self.isFocus = false;
      }, 300);
    },
    enterSearchBoxHanlder() {
      clearTimeout(this.searchBoxTimeout);
    },
    searchHandler() {
      //点击搜索，故向前端请求
      //随机生成搜索历史tag式样
      let n = RandomUtil.getRandomNumber(0, 5);
      let exist =
        this.historySearchList.filter((value) => {
          return value.name == this.search;
        }).length == 0
          ? false
          : true;
      if (!exist) {
        this.historySearchList.push({ name: this.search, type: this.types[n] });
        Store.saveHistory(this.historySearchList);
      }
      this.history = this.historySearchList.length == 0 ? false : true;
      //点击了搜索就向后端请求
      axios
        .get("http://localhost:8081/search/" + this.search)
        .then((res) => {
          // 将搜索结果通过事件传递给父组件
          this.$emit("handleSearch", res.data.data.notes);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    closeHandler(search) {
      this.historySearchList.splice(this.historySearchList.indexOf(search), 1);
      Store.saveHistory(this.historySearchList);
      clearTimeout(this.searchBoxTimeout);
      if (this.historySearchList.length == 0) {
        this.history = false;
      }
    },
    removeAllHistory() {
      Store.removeAllHistory();
    },
  },
  computed: {
    isHistorySearch() {
      return this.isFocus && !this.search;
    },
    isSearchList() {
      return this.isFocus && this.search;
    },
    isSearch() {
      return this.isFocus;
    },
  },
};
</script>

<style>
.threeline {
  width: 210px;
  height: 90px;
  margin-top: -250px;
  padding-top: -50px;
  margin-left: -1050px;
}
.left {
  margin-left: 200px;
}
.center {
  margin-top: 0px;
  padding-top: 0px;
  margin-left: 300px;
}
#search {
  background-color: #f53b3b00;
  border-radius: 0%;
}
.search-title {
  color: #bdbaba;
  font-size: 15px;
  margin-bottom: 5px;
}
.remove-history {
  color: #bdbaba;
  font-size: 15px;
  float: right;
  margin-top: -22px;
}
#search-box {
  width: 555px;
  height: 300px;
  margin-top: 0px;
  padding-bottom: 20px;
}
</style>
