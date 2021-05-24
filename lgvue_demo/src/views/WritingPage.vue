<template>
  <v-container class="d-flex justify-center mt-6">
    <v-card width="100%"
            height="100%"
            flat>
      <v-card-text>
        <v-row justify="center">
          <v-col cols="3">
            <v-card height="830px">
              <v-card-title>文章列表</v-card-title>
              <v-divider></v-divider>
              <v-card-text>
                <v-responsive v-if="articles.length > 0"
                              class="overflow-y-auto"
                              max-height="720">
                  <template v-for="(item,i) in articles">
                    <v-hover v-slot="{ hover }"
                             :key="i">
                      <v-card class="pa-2 mb-2"
                              @click="getarticlehandle(i)"
                              v-ripple>
                        <v-card class="text-h5 mb-2"
                                flat>
                          {{item.Title}}
                        </v-card>
                        <v-card class="mt-1 text-justify text--disabled"
                                width="100%"
                                flat>
                          {{item.Content | ellipsis}}
                        </v-card>
                        <v-divider class="mt-2"></v-divider>
                        <v-card class="d-flex justify-end mt-1"
                                flat>
                          <v-card class="text--disabled"
                                  flat>{{item.CreationTime|formatDate}}</v-card>
                        </v-card>
                      </v-card>
                    </v-hover>
                  </template>
                </v-responsive>
                <v-card v-else
                        height="720px"
                        class="text-center"
                        flat>
                  <v-card-text class="text--disabled"> 您当前还没有文章</v-card-text>

                </v-card>
              </v-card-text>
            </v-card>
            <v-btn class="mt-2 red"
                   @click="restarteditor"
                   dark
                   block>添加新的文章</v-btn>
            <v-btn class="mt-2 green"
                   @click="restarteditor"
                   :disabled="currarticle.Id == 0 ? true:false"
                   block>删除文章</v-btn>
          </v-col>
          <v-col cols="9">
            <v-card>
              <v-card-title>编辑文章</v-card-title>
              <v-divider></v-divider>
              <v-card-text>
                <v-text-field label="Title"
                              v-model="currarticle.Title"></v-text-field>
                <mavon-editor v-model="currarticle.Content"
                              ref=md
                              @imgAdd="$imgAdd"
                              @imgDel="$imgDel"></mavon-editor>
              </v-card-text>
              <v-card-actions class="d-flex justify-end">
                <v-btn @click="saveearticle"
                       rounded>保存</v-btn>
                <v-btn class="text-h6  ml-6 white--text"
                       @click="releasearticle"
                       color="error"
                       rounded>发布</v-btn>
              </v-card-actions>
            </v-card>

          </v-col>
        </v-row>
      </v-card-text>

    </v-card>

  </v-container>
</template>

<script>
import { createarticle, deletearticle, getauthoIrdarticles } from '@/api/article.js'
import moment from 'moment'
export default {
  props: {
    articleid: {
      type: Number,
      default: 0,
    }
  },
  data: () => {
    return {
      currindex: -1,
      articles: [],
      currarticle: {
        Id: 0,
        Title: "",
        Content: "",
        CreationTime: 0,
      },
      img_file: {}
    }
  },
  filters: {
    ellipsis (value) {
      if (!value) return "";
      if (value.length > 60) {
        return value.slice(0, 60) + "...";
      }
      return value;
    },
    formatDate (time) {
      var date = new Date(time * 1000);
      return moment(date).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  created () {
    getauthoIrdarticles(null).then(response => {
      this.articles = response.data
    })
  },
  methods: {
    // 绑定@imgAdd event
    $imgAdd (pos, $file) {
      // 缓存图片信息
      this.img_file[pos] = $file;
    },
    $imgDel (pos) {
      delete this.img_file[pos];
    },
    uploadimg ($e) {
      // 第一步.将图片上传到服务器.
      var formdata = new FormData();
      for (var _img in this.img_file) {
        formdata.append(_img, this.img_file[_img]);
      }
      axios({
        url: 'server url',
        method: 'post',
        data: formdata,
        headers: { 'Content-Type': 'multipart/form-data' },
      }).then((res) => {
        /**
         * 例如：返回数据为 res = [[pos, url], [pos, url]...]
         * pos 为原图片标志（0）
         * url 为上传后图片的url地址
         */
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        for (var img in res) {
          // $vm.$img2Url 详情见本页末尾
          $vm.$img2Url(img[0], img[1]);
        }
      })
    },
    getarticlehandle (i) {
      // getarticle({ ArticleId: aid }).then(response => {
      //   this.article = response.data
      // })
      this.currindex = i
      this.currarticle = this.articles[i]
    },
    saveearticle () {
      createarticle({ ArticleId: this.currarticle.Id, Title: this.currarticle.Title, Content: this.currarticle.Content }).then(response => {
        if (response.data.Id != this.currarticle.Id) {
          this.articles.push(response.data)
        }
      })
    },
    releasearticle () {

    },
    deleteArticle () {
      deletearticle({ ArticleId: this.currarticle.Id }).then(response => {
        this.articles.splice(currindex, 1);
        this.currarticle = {
          Id: 0,
          Title: "",
          Content: "",
          CreationTime: 0,
        }
      })
    },
    restarteditor () {
      this.currarticle = {
        Id: 0,
        Title: "",
        Content: "",
        CreationTime: 0,
      }
    }
  }
}
</script>

<style>
</style>