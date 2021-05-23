<template>
  <v-card class="mb-6"
          width="100%"
          flat>
    <v-card class="text-h5"
            flat>{{title}}</v-card>
    <v-card class="mt-1 text-justify text--disabled "
            width="100%"
            flat>{{content | ellipsis}}</v-card>
    <v-card class="mt-2 d-flex justify-start"
            flat>
      <v-icon small
              color="red">
        mdi-diamond-stone
      </v-icon>
      <v-card class="ml-4 text-caption text--disabled"
              flat>{{authorname}}</v-card>
      <v-icon class="ml-2"
              small>
        mdi-message
      </v-icon>
      <v-card class="ml-1 text-caption text--disabled"
              flat>{{messagenum}}</v-card>
      <v-icon class="ml-2"
              small>
        mdi-cards-heart
      </v-icon>
      <v-card class="ml-1 text-caption text--disabled"
              flat>{{heartnum}}</v-card>
    </v-card>
    <v-divider class="mt-4"></v-divider>
  </v-card>
</template>

<script>
export default {
  name: "ArticleItem",
  props: {
    title: {
      type: String,
      default: "怀孕的感觉是什么"
    },
    content: {
      type: String,
      default: "一转眼认识小狐狸也有半年多了，说实话，心里挺美。 这半年以来，除了小狐狸不让我动手动脚以外，你别说还挺暖心的。昨天晚上，我犹豫了半天，还是没憋住，低眉顺眼的对小狐狸说：宝贝，我们认识这么久了，你到底对我有没有感觉啊？小狐狸一边舔着哈根达斯，一边很不耐烦的说：有！有！天呐，我的心都快跳出来了，激动的我磕磕巴巴的对她说：亲爱的，快告诉我什么感觉？小狐狸一脸不屑的样子，都没拿眼皮瞄我一下，没好气的说：就像怀了孕一样"
    },
    authorname: {
      type: String,
      default: "苏安_56",
    },
    diamond: {
      type: Number,
      default: 10,
    },
    messagenum: {
      type: Number,
      default: 26,
    },
    heartnum: {
      type: Number,
      default: 360,
    }
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
}
</script>

<style>
</style>