<template>
  <v-card width="650px">
    <v-toolbar dense>
      <v-toolbar-title>{{service.Info.ServiceId}}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon
             @click="ishow = !ishow">
        <v-icon>{{ishow?"mdi-chevron-up":"mdi-chevron-down"}} </v-icon>
      </v-btn>
    </v-toolbar>
    <v-divider></v-divider>
    <v-card-text class="pa-0">
      <v-card flat>
        <v-card-title class="d-flex justify-start">
          <v-card width="250"
                  flat>
            <p>{{$t('serviceid')}}:<strong>{{service.Info.ServiceId}}</strong></p>
          </v-card>
          <v-card width="250"
                  flat>
            <p>{{$t('servicetype')}}:<strong>{{service.Info.ServiceType}}</strong></p>
          </v-card>
          <v-card width="250"
                  flat>
            <p>{{$t('servicetag')}}:<strong>{{service.Info.ServiceTag}}</strong></p>
          </v-card>
          <v-card width="250"
                  flat>
            <p>{{$t('version')}}:<strong>{{service.Info.ServiceVersion}}</strong></p>
          </v-card>
          <v-card width="250"
                  flat>
            <p>{{$t('servicecategory')}}:<strong>{{service.Info.ServiceCategory}}</strong></p>
          </v-card>
          <v-card width="250"
                  flat>
            <p>{{$t('processid')}}:<strong>{{service.Info.Pid}}</strong></p>
          </v-card>
          <v-card width="250"
                  flat>
            <p>{{$t('processname')}}:<strong>{{service.Info.Pname}}</strong></p>
          </v-card>
          <template v-if="service.Info.Setting"
                    v-for="(k,v) in service.Info.Setting">
            <v-card width="250"
                    flat>
              <p>{{k}}:<strong>{{v}}</strong></p>
            </v-card>
          </template>
        </v-card-title>
      </v-card>
      <v-expand-transition>
        <v-card v-show="ishow"
                flat>
          <v-card-text>
            <v-card class="mt-2">
              <v-card-title>
                <v-card-title>{{$t('SysSetting')}}</v-card-title>
              </v-card-title>
              <v-divider></v-divider>
              <v-card-text>
                <v-tabs v-model="systab"
                        height="35px">
                  <v-tab class="text-capitalize"
                         v-for="(item,index) in service.Info.SysSetting"
                         :key="index"
                         :href="'#systab-' + index">{{index}}</v-tab>
                </v-tabs>
                <v-tabs-items class="pt-4"
                              v-model="systab">
                  <v-tab-item v-for="(item,index) in service.Info.SysSetting"
                              :key="index"
                              :value="'systab-' + index">
                    <template v-if="item.Setting && Object.keys(item.Setting).length > 0">
                      <v-row class="ma-2"
                             no-gutters>
                        <v-col cols="12"
                               sm="6"
                               v-for="(item,name) in item.Setting"
                               :key="name">
                          <v-text-field v-if="!item.IsWrite"
                                        class="pl-2 setting-centext_field text--disabled"
                                        :value="item.Data"
                                        :label="name"
                                        disabled
                                        dense
                                        outlined
                                        readonly></v-text-field>
                          <v-text-field v-else
                                        class="setting-centext_field"
                                        v-model="item.NewData"
                                        :value="item.Data"
                                        :label="name"
                                        dense
                                        outlined
                                        :append-icon="'mdi-content-save'"
                                        @click:append="modifyModuleSetting(index,name,item.Data,item.NewData)"></v-text-field>
                        </v-col>
                      </v-row>
                    </template>
                    <template v-else>
                      <div class="modules-centext-nodes">
                        {{$t('noconfiguration')}}
                      </div>
                    </template>
                  </v-tab-item>
                </v-tabs-items>
              </v-card-text>

            </v-card>
            <v-card class="mt-2">
              <v-card-title>{{$t('ModuleSetting')}}</v-card-title>
              <v-divider></v-divider>
              <v-card-text>
                <v-tabs v-model="tab"
                        height="35px">
                  <v-tab class="text-capitalize"
                         v-for="(item,index) in service.Info.ModuleMonitor"
                         :key="index"
                         :href="'#tab-' + index">{{index}}</v-tab>
                </v-tabs>
                <v-tabs-items class="pt-4"
                              v-model="tab">
                  <v-tab-item v-for="(item,index) in service.Info.ModuleMonitor"
                              :key="index"
                              :value="'tab-' + index">
                    <template v-if="item.Setting && Object.keys(item.Setting).length > 0">
                      <v-row class="ma-2"
                             no-gutters>
                        <v-col cols="12"
                               sm="6"
                               v-for="(item,name) in item.Setting"
                               :key="name">
                          <v-text-field v-if="!item.IsWrite"
                                        class="pl-2 setting-centext_field text--disabled"
                                        :value="item.Data"
                                        :label="name"
                                        disabled
                                        dense
                                        outlined
                                        readonly></v-text-field>
                          <v-text-field v-else
                                        class="setting-centext_field"
                                        v-model="item.NewData"
                                        :value="item.Data"
                                        :label="name"
                                        dense
                                        outlined
                                        :append-icon="'mdi-content-save'"
                                        @click:append="modifyModuleSetting(index,name,item.Data,item.NewData)"></v-text-field>
                        </v-col>
                      </v-row>
                    </template>
                    <template v-else>
                      <div class="modules-centext-nodes">
                        {{$t('noconfiguration')}}
                      </div>
                    </template>
                  </v-tab-item>
                </v-tabs-items>
              </v-card-text>

            </v-card>
            <v-card class="mt-2">
              <v-card-title>{{$t('Monitor')}}</v-card-title>
              <v-divider></v-divider>
              <v-card-text>
                <v-tabs v-model="tab1"
                        height="35px">
                  <v-tab href="#tab1-0">{{$t('servicepreweight')}}</v-tab>
                  <v-tab href="#tab1-1">{{$t('servicegoroutine')}}</v-tab>
                  <v-tab href="#tab1-2">{{$t('cpu')}}</v-tab>
                  <v-tab href="#tab1-3">{{$t('memory')}}</v-tab>
                </v-tabs>
                <v-tabs-items v-model="tab1">
                  <v-tab-item value="tab1-0">
                    <v-card height="300px">
                      <line-echarts :keys="service.Monitor.Keys"
                                    :values="service.Monitor.PreWeight">
                      </line-echarts>
                    </v-card>
                  </v-tab-item>
                  <v-tab-item value="tab1-1">
                    <v-card height="300px">
                      <line-echarts :keys="service.Monitor.Keys"
                                    :values="service.Monitor.Goroutine"></line-echarts>
                    </v-card>
                  </v-tab-item>
                  <v-tab-item value="tab1-2">
                    <v-card height="300px">
                      <line-echarts :keys="service.Monitor.Keys"
                                    :values="service.Monitor.Cpu"></line-echarts>
                    </v-card>
                  </v-tab-item>
                  <v-tab-item value="tab1-3">
                    <v-card height="300px">
                      <line-echarts :keys="service.Monitor.Keys"
                                    :values="service.Monitor.Memory"></line-echarts>
                    </v-card>
                  </v-tab-item>
                </v-tabs-items>
              </v-card-text>

            </v-card>
          </v-card-text>
        </v-card>
      </v-expand-transition>
    </v-card-text>
  </v-card>
</template>

<script>
import { LineEcharts, } from "@/components/echarts"
export default {
  name: "ServiceItem",
  props: {
    service: {
      type: Object,
      default: function () {
        return {
          Info: {
            serviceId: "demo_1",
            serviceType: "demo",
            serviceTag: "demo",
            serviceVersion: 1.0,
            serviceCategory: "demoservice",
            Pid: "1010",
            Pname: "demo",
            Setting: {},
            SysSetting: {},
            ModuleMonitor: {},
          },
          Monitor: {}
        }
      }
    }
  },
  components: {
    LineEcharts,
  },
  data: function () {
    return {
      ishow: false,
      tab: null,
      systab: null,
      tab1: null,
    }
  },
}
</script>

<style>
</style>