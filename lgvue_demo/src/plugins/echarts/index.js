import Vue from 'vue'
import ECharts from 'vue-echarts'
import { use, registerMap } from 'echarts/core'
// 手动引入 ECharts 各模块来减小打包体积
import {
  CanvasRenderer
} from 'echarts/renderers'

import {
  LineChart,
  BarChart,
  PieChart,
  GaugeChart,
  MapChart,
  ScatterChart,
  CustomChart,
  EffectScatterChart,
} from 'echarts/charts'

import {
  GridComponent,
  TooltipComponent,
  ToolboxComponent,
  TitleComponent,
  VisualMapComponent,
  TimelineComponent,
  CalendarComponent,
  LegendComponent,
  GeoComponent,
} from 'echarts/components'

use([
  CanvasRenderer,
  LineChart,
  BarChart,
  PieChart,
  GaugeChart,
  MapChart,
  ScatterChart,
  CustomChart,
  EffectScatterChart,
  GridComponent,
  TooltipComponent,
  TitleComponent,
  ToolboxComponent,
  VisualMapComponent,
  TimelineComponent,
  CalendarComponent,
  LegendComponent,
  GeoComponent,
]);

// Map of China
import chinaMap from "./china.json";
// registering map data
registerMap("china", chinaMap);

// 全局注册组件（也可以使用局部注册）
Vue.component('v-chart', ECharts)