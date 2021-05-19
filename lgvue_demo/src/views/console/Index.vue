<template>
  <v-container class="pa-0 d-flex justify-center"
               fluid>
    <v-card class="float-none"
            flat>
      <v-card-text v-if="projectinfo">
        <v-card class="text-center text--secondary pt-12 pb-12"
                flat>
          <p class="text-h6"> {{$t('ProjjectInfo')}}</p>
        </v-card>
        <v-card class="text-center"
                flat>
          <p class="text-h4"> {{projectinfo.ProjectName}}</p>
          <p class="text-h6"> {{projectinfo.ProjectDes}}</p>
        </v-card>
        <v-card class=" ma-6">
          <v-card-title>
            {{$t('ProjectMember')}}
          </v-card-title>
          <v-card-text>
            <v-simple-table>
              <template v-slot:default>
                <thead>
                  <tr>
                    <th class="text-left">
                      {{$t('Position')}}
                    </th>
                    <th class="text-left">
                      {{$t('Name')}}
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="(member,key) in projectinfo.ProjectMember"
                      :key="key">
                    <td>{{ member.ProjectPosition }}</td>
                    <td>{{ member.MemberName }}</td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
          </v-card-text>
        </v-card>
        <v-card class="pt-6"
                flat>
          <p class="text--disabled text-right">{{$t('version')+':'+projectinfo.ProjectVersion}}</p>
          <p class="text--disabled text-right">{{projectinfo.ProjectTime}}</p>
        </v-card>

      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  computed: {
    ...mapGetters([
      'projectinfo',
    ]),
  },
  filters: {
    formatDate (time) {
      var date = new Date(time * 1000);
      return moment(date).format("YYYY-MM-DD HH:mm:ss");
    }
  },
  created () {
    // this.starttimer()
    this.$store.dispatch('console/getprojectinfo')
    // this.$store.dispatch('console/gethostinfo')
    // this.$store.dispatch('console/getcpuinfo')
    // this.$store.dispatch('console/getmemoryinfo')
  },
}
</script>

<style>
</style>