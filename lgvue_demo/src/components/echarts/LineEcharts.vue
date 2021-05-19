<!-- 
    Echarts 线性图标
    Demo:https://echarts.apache.org/examples/zh/editor.html?c=gauge-grade
 -->
<template>
  <v-chart :option="option"
           :autoresize="true" />
</template>

<script>
export default {
  name: "LineEcharts",
  props: {
    smooth: {//false折线 true曲线
      type: Boolean,
      default: false,
    },
    keys: {
      type: Array,
      default: () => ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    values: {
      type: Array,
      default: () => [150, 230, 224, 218, 135, 147, 260]
    }
  },
  data () {
    return {
      option: {
        title: {
          text: "Line Echarts Demo",
          left: "center"
        },
        xAxis: {
          type: 'category',
          data: this.keys
        },
        yAxis: {
          type: 'value'
        },
        series: [{
          data: this.values,
          type: 'line',
          smooth: this.smooth               //false折线 true曲线
        }]
      }
    }
  }
}
</script>

<style>
</style>
