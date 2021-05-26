<template>
  <v-container class="d-flex justify-center mt-6 fill-height ">
    <v-card height="100%"
            width="100%"
            flat>
      <v-card-text>
        <v-row justify="center">
          <v-col cols="8">
            <v-card flat>
              <v-card flat>
                <v-card-title class="text-h4">{{article.Title}}</v-card-title>
                <v-card-text class="d-flex justify-start">
                  <v-avatar color="primary text-h4 text-uppercase font-weight-black white--text"
                            size="56">{{shouname}}</v-avatar>
                  <v-card class="ml-2"
                          flat>
                    <v-card flat
                            height="36"
                            class="d-flex justify-start">
                      <p class="mt-2 text-subtitle-2">{{article.AuthorName}}</p>
                      <v-btn icon
                             color="pink">
                        <v-icon>mdi-flower-poppy</v-icon>
                      </v-btn>
                      <v-btn class="ml-1 mt-1"
                             small
                             rounded
                             outlined>关注</v-btn>
                    </v-card>
                    <v-card class="d-flex justify-start"
                            flat>
                      <v-icon small
                              color="red">
                        mdi-diamond-stone
                      </v-icon>
                      <v-card class="ml-2 text-caption text--disabled transparent"
                              flat>{{article.CreationTime | formatDate}}</v-card>

                      <v-card class="ml-2 text-caption text--secondary transparent"
                              flat>字数 {{article.Content.length}}</v-card>

                      <v-card class="ml-2 text-caption text--secondary transparent"
                              flat>阅读数 0</v-card>
                    </v-card>
                  </v-card>
                </v-card-text>
              </v-card>
              <v-card class="mt-2"
                      flat>
                <mavon-editor class="v-note-wrapper"
                              :value="article.Content"
                              :editable="false"
                              :subfield="false"
                              defaultOpen="preview"
                              :toolbarsFlag="false"
                              :boxShadow="false"
                              :scrollStyle="false">

                </mavon-editor>
              </v-card>
            </v-card>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card class="v-btn--example transparent d-flex flex-column"
              flat>
        <v-btn fab
               outlined>
          <v-icon>mdi-thumb-up</v-icon>
        </v-btn>
        <p class="pt-3 text-center">{{greatnum}}赞</p>
      </v-card>

    </v-card>

  </v-container>

</template>

<script>
import { getarticle } from '@/api/article.js'
import moment from 'moment'
export default {
  components: {
  },
  data: () => {
    return {
      fab: false,
      article: {
        Id: 0,
        Title: "",
        Content: "",
        ShortContent: "",
        Images: [],
        CreationTime: 0,
        AuthorName: "",
      }
    }
  },
  filters: {
    formatDate (time) {
      var date = new Date(time * 1000);
      return moment(date).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  computed: {
    greatnum: function () {
      if (this.article.GreatNum) {
        return this.article.GreatNum
      } else {
        return 0
      }
    },
    shouname: function () {
      return this.article.AuthorName.substring(0, 1)
    }
  },
  created () {
    getarticle({ ArticleId: Number(this.$route.params.id) }).then(response => {
      this.article = response.data
    })
  },
}
</script>

<style>
.v-note-wrapper {
  z-index: 1 !important;
}
.v-btn--example {
  position: fixed;
  top: 30%;
  left: 20%;
  z-index: 100;
}
</style>