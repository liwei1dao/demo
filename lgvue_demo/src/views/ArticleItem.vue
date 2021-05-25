<template>
  <v-card class="mb-6 pl-4"
          width="100%"
          @click="gotoArticle"
          flat>
    <v-row>
      <v-col :cols="image?'9':'12'">
        <v-card class="text-h5 transparent"
                flat>{{article.Title}}</v-card>
        <v-card class="mt-1 text-justify text--disabled transparent"
                width="100%"
                min-height="28px"
                flat>{{article.ShortContent | ellipsis}}</v-card>
        <v-spacer></v-spacer>
        <v-card class="mt-2 d-flex justify-start transparent"
                flat>
          <v-icon small
                  color="red">
            mdi-diamond-stone
          </v-icon>
          <v-card class="ml-4 text-caption text--disabled transparent"
                  flat>{{article.AuthorName}}</v-card>
          <v-icon class="ml-2"
                  small>
            mdi-message
          </v-icon>
          <v-card class="ml-1 text-caption text--disabled transparent"
                  flat>{{messagenum}}</v-card>
          <v-icon class="ml-2"
                  small>
            mdi-cards-heart
          </v-icon>
          <v-card class="ml-1 text-caption text--disabled transparent"
                  flat>{{greatnum}}</v-card>
        </v-card>
      </v-col>
      <v-col :cols="image?'3':'0'">
        <v-img v-if="image"
               :src="image"
               height="140"
               contain></v-img>
      </v-col>
    </v-row>

    <v-divider class="mt-2"></v-divider>
  </v-card>
</template>

<script>
export default {
  name: "ArticleItem",
  props: {
    article: {
      type: Object,
      default () {
        return {
          Id: 0,
          Title: "",
          ShortContent: "",
          CreationTime: 0,
          AuthorName: "",
          Images: [],
          Diamond: 0,
          GreatNum: 0,
          Messagenum: 0,
        }
      }
    }
  },
  computed: {
    image: function () {
      if (this.article.Images && this.article.Images.length > 0) {
        return this.article.Images[0]
      } else {
        return null
      }
    },
    messagenum: function () {
      if (this.article.Messagenum) {
        return this.article.Messagenum
      } else {
        return 0
      }
    },
    greatnum: function () {
      if (this.article.GreatNum) {
        return this.article.GreatNum
      } else {
        return 0
      }
    },
  },
  filters: {
    ellipsis (value) {
      if (!value) return "";
      if (value.length > 160) {
        return value.slice(0, 160) + "...";
      }
      return value;
    }
  },
  methods: {
    gotoArticle () {
      this.$router.push({
        path: `/articleI`,
      })
    }
  }
}
</script>

<style>
</style>