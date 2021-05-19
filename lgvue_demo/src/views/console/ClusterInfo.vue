<template>
  <v-container class="ma-6"
               fluid>
    <v-card flat>
      <v-card-title class="d-flex justify-center mb-6">
        <v-card class="col-8 text-center text--secondary pt-12 pb-12"
                flat>
          <p class="text-h6"> {{$t('HostInfo')}}</p>
        </v-card>
      </v-card-title>
      <v-card-text class="d-flex justify-center ">
        <template v-if="clusterservice"
                  v-for="item in clusterservice">
          <service-item :service="item"></service-item>
        </template>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>

import ServiceItem from "./ServiceItem.vue"
import { mapGetters } from 'vuex'
export default {
  components: {
    ServiceItem
  },
  computed: {
    ...mapGetters([
      'clusterservice'
    ])
  },
  created () {
    // this.starttimer()
    this.$store.dispatch('console/getconsolecluster')
  },
}
</script>

<style>
</style>