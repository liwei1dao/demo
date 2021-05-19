<template>
  <v-container class="d-flex flex-column align-center ma-6"
               fluid>

    <v-card class="col-8 text-center text--secondary pt-12 pb-12"
            flat>
      <p class="text-h6"> {{$t('HostInfo')}}</p>
    </v-card>

    <v-card v-if="hostinfo"
            class="col-8 mb-6">
      <v-toolbar class="elevation-0 primary"
                 dark>
        {{$t('system')}}
      </v-toolbar>
      <v-card-text>
        <v-simple-table>
          <template v-slot:default>
            <tbody>
              <tr>
                <td>{{$t('hostid')}}</td>
                <td>{{hostinfo.HostID}}</td>
              </tr>
              <tr>
                <td>{{$t('hostname')}}</td>
                <td>{{hostinfo.Hostname}}</td>
              </tr>
              <tr>
                <td>{{$t('hostuptime')}}</td>
                <td>{{hostinfo.Uptime | secondToDate}}</td>
              </tr>
              <tr>
                <td>{{$t('hostboottime')}}</td>
                <td>{{hostinfo.BootTime | formatDate}}</td>
              </tr>
              <tr>
                <td>{{$t('hostprocsnum')}}</td>
                <td>{{hostinfo.Procs}}</td>
              </tr>
              <tr>
                <td>{{$t('hostos')}}</td>
                <td>{{hostinfo.OS}}</td>
              </tr>
              <tr>
                <td>{{$t('hostplatform')}}</td>
                <td>{{hostinfo.Platform}}</td>
              </tr>
              <tr>
                <td>{{$t('hostplatformfamily')}}</td>
                <td>{{hostinfo.PlatformFamily}}</td>
              </tr>
              <tr>
                <td>{{$t('hostplatformversion')}}</td>
                <td>{{hostinfo.PlatformVersion}}</td>
              </tr>
              <tr>
                <td>{{$t('hostplatformkernelarch')}}</td>
                <td>{{hostinfo.KernelArch}}</td>
              </tr>
              <tr>
                <td>{{$t('hostVirtualizationsystem')}}</td>
                <td>{{hostinfo.VirtualizationSystem}}</td>
              </tr>
              <tr>
                <td>{{$t('hostVirtualizationrole')}}</td>
                <td>{{hostinfo.VirtualizationRole}}</td>
              </tr>
            </tbody>
          </template>
        </v-simple-table>
      </v-card-text>
    </v-card>

    <v-card v-if="cpuinfo"
            class="col-8 mb-6">
      <v-toolbar class="elevation-0 orange"
                 dark>
        <p>{{$t('cpu')}}</p>
        <v-spacer></v-spacer>
        <v-tabs class="pl-6"
                v-model="cputab"
                center-active
                dark>
          <v-tabs-slider color="yellow"></v-tabs-slider>
          <v-tab v-for="(cpu,index) in cpuinfo"
                 :key="index"
                 :href="'#cputab-' + index"
                 class="text-capitalize">
            {{$t('cpu')+":"+index}}
          </v-tab>
        </v-tabs>
      </v-toolbar>
      <v-card-text class="pa-0 ma-0">
        <v-tabs-items v-model="cputab">
          <v-tab-item v-for="(cpu,index) in cpuinfo"
                      :key="index"
                      :value="'cputab-' + index">
            <v-card class="pa-0 ma-0">
              <v-carousel v-model="cpumodule"
                          height="675px"
                          light>
                <v-carousel-item>
                  <v-card>
                    <v-simple-table>
                      <template v-slot:default>
                        <tbody>
                          <tr>
                            <td>{{$t('cpuindex')}}</td>
                            <td>{{index}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpuvendorid')}}</td>
                            <td>{{cpu.VendorID}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpufamily')}}</td>
                            <td>{{cpu.Family}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpumodel')}}</td>
                            <td>{{cpu.Model}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpustepping')}}</td>
                            <td>{{cpu.Stepping}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpuphysicalid')}}</td>
                            <td>{{cpu.PhysicalID}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpucoreid')}}</td>
                            <td>{{cpu.CoreID}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpucores')}}</td>
                            <td>{{cpu.Cores}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpumodelname')}}</td>
                            <td>{{cpu.ModelName}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpumhz')}}</td>
                            <td>{{cpu.Mhz}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpucachesize')}}</td>
                            <td>{{cpu.CacheSize}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpuflags')}}</td>
                            <td>{{cpu.Flags}}</td>
                          </tr>
                          <tr>
                            <td>{{$t('cpumicrocode')}}</td>
                            <td>{{cpu.Microcode}}</td>
                          </tr>
                        </tbody>
                      </template>
                    </v-simple-table>
                  </v-card>
                </v-carousel-item>
                <v-carousel-item v-if="hostmonitordata">
                  <line-echarts :keys="hostmonitordata.Keys"
                                :values="hostmonitordata.Cpu">
                  </line-echarts>
                </v-carousel-item>
              </v-carousel>
            </v-card>
          </v-tab-item>
        </v-tabs-items>
      </v-card-text>
    </v-card>

    <v-card v-if="memoryinfo"
            class="col-8 mb-6">
      <v-toolbar class="elevation-0 brown"
                 dark>
        <p>{{$t('memory')}}</p>
      </v-toolbar>
      <v-card-text class="pa-0 ma-0">
        <v-card class="pa-0 ma-0">
          <v-carousel v-model="memorymodule"
                      height="480px"
                      light>
            <v-carousel-item>
              <v-card>
                <v-simple-table>
                  <template v-slot:default>
                    <tbody>
                      <tr>
                        <td>{{$t('memorytotal')}}</td>
                        <td>{{memoryinfo.Total | memoryconver}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memoryavailable')}}</td>
                        <td>{{memoryinfo.Available | memoryconver}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memoryused')}}</td>
                        <td>{{memoryinfo.Used | memoryconver}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memoryusedpercent')}}</td>
                        <td>{{memoryinfo.UsedPercent+"%"}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memoryfree')}}</td>
                        <td>{{memoryinfo.Free | memoryconver}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memoryactive')}}</td>
                        <td>{{memoryinfo.Active}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memoryinactive')}}</td>
                        <td>{{memoryinfo.Inactive}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memorywired')}}</td>
                        <td>{{memoryinfo.Wired}}</td>
                      </tr>
                      <tr>
                        <td>{{$t('memorylaundry')}}</td>
                        <td>{{memoryinfo.Laundry}}</td>
                      </tr>
                    </tbody>
                  </template>
                </v-simple-table>
              </v-card>
            </v-carousel-item>
            <v-carousel-item v-if="hostmonitordata">
              <line-echarts :keys="hostmonitordata.Keys"
                            :values="hostmonitordata.Memory">
              </line-echarts>
            </v-carousel-item>
          </v-carousel>
        </v-card>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import { LineEcharts, GaugeEcharts, } from "@/components/echarts"
import moment from 'moment'
import { mapGetters } from 'vuex'
export default {
  components: {
    LineEcharts,
    GaugeEcharts
  },
  computed: {
    ...mapGetters([
      'hostinfo',
      'cpuinfo',
      'memoryinfo',
      'hostmonitordata',
    ]),
  },
  data: function () {
    return {
      cputab: null,
      cputabindex: null,
      memoryindex: null,
      cpumodule: 0,
      memorymodule: 0,
      colors: [
        'primary',
        'secondary',
        'yellow darken-2',
        'red',
        'orange',
      ],
    }
  },
  filters: {
    formatDate (time) {
      var date = new Date(time * 1000);
      return moment(date).format("YYYY-MM-DD HH:mm:ss");
    },
    //秒转化成 时分秒
    secondToDate (result) {
      var h = Math.floor(result / 3600);
      var m = Math.floor((result / 60 % 60));
      var s = Math.floor((result % 60));
      return result = h + ":" + m + ":" + s;
    },
    memoryconver (limit) {
      var size = "";
      if (limit < 0.1 * 1024) { //如果小于0.1KB转化成B
        size = limit.toFixed(2) + "B";
      } else if (limit < 0.1 * 1024 * 1024) {//如果小于0.1MB转化成KB
        size = (limit / 1024).toFixed(2) + "KB";
      } else if (limit < 0.1 * 1024 * 1024 * 1024) { //如果小于0.1GB转化成MB
        size = (limit / (1024 * 1024)).toFixed(2) + "MB";
      } else { //其他转化成GB
        size = (limit / (1024 * 1024 * 1024)).toFixed(2) + "GB";
      }

      var sizestr = size + "";
      var len = sizestr.indexOf(".");
      var dec = sizestr.substr(len + 1, 2);
      if (dec == "00") {//当小数点后为00时 去掉小数部分
        return sizestr.substring(0, len) + sizestr.substr(len + 3, 2);
      }
      return sizestr;
    }
  },
  created () {
    // this.starttimer()
    this.$store.dispatch('console/getprojectinfo')
    this.$store.dispatch('console/gethostinfo')
    this.$store.dispatch('console/getcpuinfo')
    this.$store.dispatch('console/getmemoryinfo')
    this.$store.dispatch('console/gethostmonitordata')
  },
}
</script>

<style>
</style>