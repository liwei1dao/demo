<!-- 
    Echarts 评分仪表盘组件
    Demo:https://echarts.apache.org/examples/zh/editor.html?c=gauge-grade
 -->
<template>
  <v-chart :option="option" />
</template>

<script>
// import echarts from 'echarts'
// require('echarts/theme/macarons') // echarts theme
export default {
  name: "GaugeEcharts",
  data () {
    return {
      option: {
        title: {
          text: "Gauge Echarts Demo",
          left: "center"
        },
        tooltip: {
          formatter: '{a} <br/>{b} : {c}%'
        },
        series: [{
          name: 'Pressure',
          type: 'gauge',
          detail: {
            formatter: '{value}'
          },
          data: [{
            value: 50,
            name: 'SCORE'
          }]
        }]
      }
    }
  }
}
</script>

<style>
</style>