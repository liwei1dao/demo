<template>
  <v-app>
    <v-app-bar app
               :color="theme.primary"
               dark
               hide-on-scroll
               prominent
               scroll-target="#scrolling-techniques-4">
      <v-toolbar-title>{{info.title}}</v-toolbar-title>
      <v-spacer></v-spacer>
      <slot name="appbarbutts"></slot>
    </v-app-bar>

    <v-main>
      <v-container class="d-flex justify-center mb-6 fill-height"
                   fluid>
        <slot></slot>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import SelectLanguge from "@/components/selectlanguge/"
export default {
  name: "LayoutS",
  components: {
    SelectLanguge,
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