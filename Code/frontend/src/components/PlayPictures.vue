<template>
  <div>
    <div>
      <div style="display: flex" class="average_container">
        <div v-for="(column, index) in columns" :key="index">
          <div v-for="image in column" :key="image.id">
            <img :src="image.url" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      //图片和列
      images: [], //封面
      header: [], //头像
      headline: [], //标题
      name: [], //用户名
      nums: [], //点赞量
      columns: [[], [], []],
    };
  },
  methods: {
    splitImagesIntoColumns() {
      this.images = ["beauty.jpg", "2.jpg", "3.jpg", "4.jpg", "beauty.jpg"];
      let columnCount = 3;
      let imagesPerColumn = Math.ceil(this.images.length / columnCount);
      for (let i = 0; i < columnCount; i++) {
        for (let j = 0; j < imagesPerColumn; j++) {
          const image = {
            id: j,
            url: require("../views/" + this.images[i * imagesPerColumn + j]),
            // header_url: require('../views/'+this.header[i * imagesPerColumn + j]),
            // headline: this.headline[i * imagesPerColumn + j],
            // name:this.name[i * imagesPerColumn + j],
            // num: this.num[i * imagesPerColumn + j]
          };
          if (image) {
            this.columns[i].push(image);
          }
        }
      }
    },
    mounted() {
      this.splitImagesIntoColumns();
    },
    created: function () {
      this.splitImagesIntoColumns();
    },
  },
};
</script>

<style>
/* 实现布局三列平分 */
.average_container > div {
  flex: 1;
}
.average_container div {
  border: solid 1px rgba(237, 226, 226, 0.174);
  margin-right: calc(80px / 3);
  box-sizing: border-box;
  margin-top: 20px;
  margin-bottom: 20px;
}
</style>
