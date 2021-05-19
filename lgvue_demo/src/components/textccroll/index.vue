<template>
  <v-card flat>
    <transition class=""
                name="slide"
                mode="out-in">
      <p :key="text.id">{{text.val}}</p>
    </transition>
  </v-card>
</template>

<script>

export default {
  name: 'TextScroll',
  props: {
    unitime: {
      type: Number,
      default: 5000,
    },
    dataList: {
      type: Array,
      default () {
        return [
          "生活是一面鏡子。你對它笑，它就對你笑；你對它哭，它也對你哭",
          "活著一天，就是有福氣，就該珍惜。當我哭泣我沒有鞋子穿的時候，我發現有人卻沒有腳"
        ];
      },
    },
  },
  data () {
    return {
      count: 0, // 当前索引
      intervalId: null, // 定时器ID
    };
  },
  computed: {
    text () {
      return {
        id: this.count,
        val: this.dataList[this.count],
      };
    },
  },
  created () {
    this.intervalId = setInterval(() => { // 定义定时器
      this.getText();
    }, this.unitime);
  },
  methods: {
    getText () {
      const len = this.dataList.length; // 获取数组长度
      if (len === 0) {
        return; // 当数组为空时，直接返回
      }
      if (this.count === len - 1) { // 当前为最后一个时
        this.count = 0; // 从第一个开始
      } else {
        this.count++; // 自增
      }
    },
  },
  destroyed () {
    clearInterval(this.intervalId); // 清除定时器
  },
};
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: all 1s;
  opacity: 1;
}

.slide-enter {
  transform: translateY(25px);
  opacity: 0;
}

.slide-leave-to {
  transform: translateY(-25px);
  opacity: 0;
}
</style>