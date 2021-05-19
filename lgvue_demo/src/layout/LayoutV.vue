<template>
  <v-app>
    <v-app-bar app
               :color="theme.primary"
               dark
               hide-on-scroll
               prominent
               flat
               scroll-target="#scrolling-techniques-4">
      <v-app-bar-nav-icon>
        <v-btn icon
               @click.stop="drawer = !drawer">
          <v-icon>mdi-menu</v-icon>
        </v-btn>
      </v-app-bar-nav-icon>
      <v-toolbar-title>{{info.title}}</v-toolbar-title>
      <v-spacer></v-spacer>
      <select-languge></select-languge>
    </v-app-bar>
    <v-navigation-drawer v-model="drawer"
                         absolute
                         temporary>
      <navigation :routers="routers"></navigation>
    </v-navigation-drawer>
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import Navigation from "@/components/navigation/"
import SelectLanguge from "@/components/selectlanguge/"
import { mapGetters } from 'vuex'
export default {
  name: "LayoutV",
  components: {
    Navigation,
    SelectLanguge,
  },
  computed: {
    ...mapGetters([
      'routers',
    ]),
  },
  props: {
    theme: {
      type: Object,
      default: function () {
        return {
          primary: '#1976D2',
          secondary: '#424242',
          accent: '#82B1FF',
          error: '#FF5252',
          info: '#2196F3',
          success: '#4CAF50',
          warning: '#FFC107',
        }
      }
    },
    info: {
      type: Object,
      default: function () {
        return {
          title: 'LgVue',
        }
      }
    }
  },
  data: () => {
    return {
      drawer: false,
    }
  },
}
</script>

<style>
</style>