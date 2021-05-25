<template>
  <v-container class="d-flex justify-center mt-6">
    <v-row justify="center">
      <v-col cols="6">
        <article-item v-for="item in articles"
                      :key="item.Id"
                      :article="item"></article-item>
      </v-col>
      <v-col cols="2">
        <v-btn class="mb-4 text-h6 float-left cyan cyan--text text--darken-4"
               height="58px"
               block>

          日更挑战
          <v-icon class="ml-4"
                  color="teal darken-4"
                  large
                  right>
            mdi-view-day
          </v-icon>
        </v-btn>
        <v-btn class="mb-4 text-h6 float-left orange lighten-2  orange--text text--darken-4"
               height="58px"
               block>

          简书会员
          <v-icon class="ml-4"
                  color="orange accent-4"
                  large
                  right>
            mdi-account-supervisor
          </v-icon>
        </v-btn>
        <v-btn class="mb-4 text-h6 float-left orange lighten-4 orange--text text--accent-3"
               height="58px"
               block>

          优选连载
          <v-icon class="ml-4"
                  color="orange accent-2"
                  large
                  right>
            mdi-clipboard-text-multiple
          </v-icon>
        </v-btn>
        <v-btn class="mb-4 text-h6 float-left teal lighten-4  teal--text text--darken-3"
               height="58px"
               block>

          简书版权
          <v-icon class="ml-4"
                  color="teal darken-4"
                  large
                  right>
            mdi-domain
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { getlastarticles } from '@/api/article.js'
import ArticleItem from "./ArticleItem.vue"
import { mapGetters } from 'vuex'
export default {
  name: "Index",
  components: {
    ArticleItem
  },
  computed: {
    ...mapGetters([
      'userinfo',
    ]),
  },
  data: () => {
    return {
      articles: [],
    }
  },
  created () {
    getlastarticles(null).then(response => {
      this.articles = response.data
    })
  },
}
</script>

<style>
</style>