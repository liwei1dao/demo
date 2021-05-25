<template>
  <v-container class="d-flex justify-center mt-6">

  </v-container>
</template>

<script>
export default {
  components: {
    VueMarkdown
  },
  data () {
    return {
      //保存解析后的md文件对象
      msg: 'ad',
      //为了让组件强制渲染
      key: 0

    }
  },
}
</script>

<style>
</style>