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
                    <v-card class=" pa-2"
                            :key="i"
                            @click="getarticlehandle(i)"
                            outlined>
                      <v-card class="text-h5 mb-2 transparent"
                              flat>
                        {{item.Title}}
                      </v-card>
                      <v-card class="mt-1 text-justify text--disabled transparent"
                              width="100%"
                              flat>
                        {{item.Content | ellipsis}}
                      </v-card>
                      <v-divider class="mt-2"></v-divider>
                      <v-card class="d-flex justify-end mt-1 transparent"
                              flat>
                        <v-card class="text--disabled transparent"
                                flat>{{item.CreationTime|formatDate}}</v-card>
                      </v-card>
                    </v-card>
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
                   @click="deleteArticle"
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
      articles: [],
      currarticle: {
        Id: 0,
        Title: "",
        Content: "",
        ShortContent: "",
        Images: [],
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
      console.log("添加图片" + $file)
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
      this.currarticle = this.articles[i]
    },
    saveearticle () {
      this.getmdimage()
      this.getshortContent()
      createarticle({ ArticleId: this.currarticle.Id, Title: this.currarticle.Title, Content: this.currarticle.Content, ShortContent: this.currarticle.ShortContent, Images: this.currarticle.Images }).then(response => {
        if (response.data.Id != this.currarticle.Id) {
          this.articles.push(response.data)
        }
        const { message } = response
        this.$message.success(message)
      })
    },
    releasearticle () {

    },
    deleteArticle () {
      deletearticle({ ArticleId: this.currarticle.Id }).then(response => {
        for (let i = 0, len = this.articles.length; i < len; i++) {
          if (this.articles[i].Id == response.data.Id) {
            this.articles.splice(i, 1);
          }
        }
        this.currarticle = {
          Id: 0,
          Title: "",
          Content: "",
          CreationTime: 0,
        }
        const { message } = response
        this.$message.success(message)
      })
    },
    restarteditor () {
      this.currarticle = {
        Id: 0,
        Title: "",
        Content: "",
        CreationTime: 0,
      }
    },
    getmdimage () {
      const pattern = /!\[(.*?)\]\((.*?)\)/mg;
      let matcher;
      this.currarticle.Images = []
      while ((matcher = pattern.exec(this.currarticle.Content)) !== null) {
        this.currarticle.Images.push(matcher[2])
      }
      console.log(this.currarticle.Images)
    },
    getshortContent () {
      if (this.currarticle.Content != "") {
        var str = this.currarticle.Content.replace(/(\*\*|__)(.*?)(\*\*|__)/g, '')          //全局匹配内粗体
          .replace(/\!\[[\s\S]*?\]\([\s\S]*?\)/g, '')                  //全局匹配图片
          .replace(/\[[\s\S]*?\]\([\s\S]*?\)/g, '')                    //全局匹配连接
          .replace(/<\/?.+?\/?>/g, '')                                 //全局匹配内html标签
          .replace(/(\*)(.*?)(\*)/g, '')                               //全局匹配内联代码块
          .replace(/`{1,2}[^`](.*?)`{1,2}/g, '')                       //全局匹配内联代码块
          .replace(/```([\s\S]*?)```[\s]*/g, '')                       //全局匹配代码块
          .replace(/\~\~(.*?)\~\~/g, '')                               //全局匹配删除线
          .replace(/[\s]*[-\*\+]+(.*)/g, '')                           //全局匹配无序列表
          .replace(/[\s]*[0-9]+\.(.*)/g, '')                           //全局匹配有序列表
          .replace(/(#+)(.*)/g, '')                                    //全局匹配标题
          .replace(/(>+)(.*)/g, '')                                    //全局匹配摘要
          .replace(/\r\n/g, "")                                        //全局匹配换行
          .replace(/\n/g, "")                                          //全局匹配换行
          .replace(/\s/g, "")                                          //全局匹配空字符;
        this.currarticle.ShortContent = str.slice(0, 155);
      }
    }
  },

}
</script>

<style>
</style>