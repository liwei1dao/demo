<template>
  <v-chart :option="option" />
</template>

<script>
export default {
  name: "MapEcharts",
  data () {
    return {
      option: {
        title: {
          text: "Map Echarts Demo",
          left: "center"
        },
        geo: {
          map: 'china'
        }
      }
    }
  }
}
</script>

<style>
</style>