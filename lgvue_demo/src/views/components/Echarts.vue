<template>
  <v-container class="pa-0 d-flex justify-center fill-height"
               fluid>
    <v-card class="ma-3"
            height="450px"
            width="450px">
      <bar-echarts></bar-echarts>
    </v-card>
    <v-card class="ma-3"
            height="450px"
            width="450px">
      <line-echarts></line-echarts>
    </v-card>
    <v-card class="ma-3"
            height="450px"
            width="450px">
      <line-echarts :smooth="true"></line-echarts>
    </v-card>
    <v-card class="ma-3"
            height="450px"
            width="450px">
      <pie-echarts></pie-echarts>
    </v-card>
    <v-card class="ma-3"
            height="450px"
            width="450px">
      <gauge-echarts></gauge-echarts>
    </v-card>
    <v-card class="ma-3"
            height="450px"
            width="650px">
      <map-echarts></map-echarts>
    </v-card>
  </v-container>
</template>

<script>
import { BarEcharts, LineEcharts, PieEcharts, GaugeEcharts, MapEcharts } from "@/components/echarts"
export default {
  components: {
    BarEcharts,
    LineEcharts,
    PieEcharts,
    GaugeEcharts,
    MapEcharts,
  },

};
</script>

<style>
</style>