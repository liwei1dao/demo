<template>
  <v-app>
    <v-app-bar class="white"
               extension-height="2px"
               app
               flat>
      <template v-slot:default="{}">
        <v-container>
          <v-row justify="center">
            <v-col cols="2">
              <v-btn class="text-h4 red--text font-weight-black"
                     @click="goindex"
                     height="100%"
                     text>简书</v-btn>
            </v-col>
            <v-col>
              <v-btn class="text-h5 red--text"
                     height="100%"
                     @click="goindex"
                     text>
                <v-icon large
                        color="red">
                  mdi-atom-variant
                </v-icon>
                首页
              </v-btn>
            </v-col>
            <v-spacer></v-spacer>
            <v-col>
              <template v-if="userinfo">
                <v-btn class="text-h5  ml-6 red--text"
                       rounded
                       to='/components/message'
                       outlined
                       height="100%"
                       text>
                  后台
                </v-btn>
              </template>
              <template v-else>
                <v-btn color="error"
                       fab
                       small
                       dark>
                  <v-icon>mdi-cog-outline</v-icon>
                </v-btn>
                <v-btn class="text-h5  ml-6"
                       height="100%"
                       to='/login'
                       text>
                  登录
                </v-btn>
                <v-btn class="text-h5  ml-6 red--text"
                       rounded
                       to='/register'
                       outlined
                       height="100%"
                       text>
                  注册
                </v-btn>
              </template>
              <v-btn class="text-h5  ml-6 white--text"
                     color="error"
                     rounded
                     to='/writing'
                     height="100%"
                     dark>
                <v-icon>mdi-lead-pencil</v-icon>
                写文章
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </template>
      <template v-slot:extension="{}">
        <v-divider></v-divider>
      </template>
    </v-app-bar>
    <v-main>
      <router-view />
    </v-main>
    <v-footer class="d-flex flex-column justify-center"
              padless>
      <v-col class="text-center text-subtitle-2 pt-12 pb-1"
             cols="6">
        本站不提供任何视听上传服务，所有内容均来自互联网其他分享站所提供的公开引用资源。所有视频及图文均归其来源站所有。我们只提供WEB页面服务，本站不存储、不制作不上传任何视频与图文。本站不承担任何由于内容所引起的争议和法律责任
      </v-col>
      <v-col class="text-center text-subtitle-2 pa-1"
             cols="6">
        任何单位或个人，认为本站收录的内容，可能涉嫌侵犯其信息网络传播权，应该向我们通过邮件的方式，提出书面权利通知并提供版权所属证明。我们在收到上述证明文件后，将会尽量尽快断开删除相关链接内容。
      </v-col>
      <v-col class="text-center text-subtitle-2 pa-1"
             cols="6">
        E-Mail邮箱： tangrenjie88@gmail.com
      </v-col>
      <v-col class="text-center text-subtitle-2 pa-1 pt-6 pb-6"
             cols="6">
        <v-btn text>RSS订阅</v-btn>
        <v-btn text>百度蜘蛛</v-btn>
        <v-btn text>神马爬虫</v-btn>
        <v-btn text>搜狗蜘蛛</v-btn>
        <v-btn text>奇虎地图</v-btn>
      </v-col>
    </v-footer>
  </v-app>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: "LayoutA",
  computed: {
    ...mapGetters([
      'userinfo',
    ]),
  },
  methods: {
    goindex () {
      this.$router.push({
        path: `/index`,
      })
    }
  }
}
</script>

<style>
</style>