<template>
  <v-chart :option="option" />
</template>

<script>
export default {
  name: "BarEcharts",
  props: {
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
          text: "Bar Echarts Demo",
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
          type: 'bar'
        }]
      }
    }
  }
}
</script>

<style>
</style>