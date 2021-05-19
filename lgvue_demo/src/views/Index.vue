<template>
  <v-container class="pa-0 d-flex justify-center fill-height"
               fluid>
    <v-card class="float-none"
            flat>
      <v-card-text>
        <v-card class=" d-flex justify-center "
                flat>
          <p class="text-h5">
            {{$t('Welcometo')+":"}}
            <strong class="red--text text--lighten-1">{{userinfo.NickName}}</strong>
          </p>
        </v-card>
        <v-card class=" d-flex justify-center "
                flat>
          <text-scroll :unitime="10000"
                       :dataList="dataList"></text-scroll>
        </v-card>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import TextScroll from '@/components/textccroll/'
import { mapGetters } from 'vuex'
export default {
  name: "Index",
  components: {
    TextScroll
  },
  computed: {
    ...mapGetters([
      'userinfo',
    ]),
  },
  data: () => {
    return {
      dataList: [
        "生活是一面鏡子。你對它笑，它就對你笑；你對它哭，它也對你哭",
        "活著一天，就是有福氣，就該珍惜。當我哭泣我沒有鞋子穿的時候，我發現有人卻沒有腳",
        "人生是個圓，有的人走了一輩子也沒有走出命運畫出的圓圈，其實，圓上的每一個點都有一條騰飛的切線",
        "千萬別迷戀網路遊戲，要玩就玩好人生這場大遊戲.",
        "命運負責洗牌，但是玩牌的是我們自己!",
        "我們心中的恐懼，永遠比真正的危險巨大的多",
        "命運掌握在自己手中。要麼你駕馭生命，要麼生命駕馭你，你的心態決定你是坐騎還是騎手",
        "寧願做過了後悔，也不要錯過了後悔",
        "不要拿小人的錯誤來懲罰自己，不要在這些微不足道的事情上折磨浪費自己的寶貴時間",
        "如果我們都去做自己能力做得到的事，我們會讓自己大吃一驚",
        "出路出路，走出去了，總是會有路的。困難苦難，困在家裡就是難",
        "學的到東西的事情是鍛煉，學不到的是磨練",
        "過錯是暫時的遺憾，而錯過則是永遠的遺憾",
        "環境不會改變，解決之道在於改變自己",
        "勇氣是控制恐懼心理，而不是心裡毫無恐懼",
        "天道酬勤。也許你付出了不一定得到回報，但不付出一定得不到回報",
        "如果你能像看別人缺點一樣，如此準確般的發現自己的缺點，那麼你的生命將會不平凡",
      ],
    }
  },
}
</script>

<style>
</style>