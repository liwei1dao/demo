<template>
  <v-list dense
          v-model="selectedItem">
    <template v-for="(router,i) in routers">
      <template v-if="router.children && router.children.length > 0">
        <v-list-group :key="i"
                      :prepend-icon="router.meta.icon">
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title v-text="$t(router.meta.title)"></v-list-item-title>
            </v-list-item-content>
          </template>
          <v-list-item v-for="(child,n) in router.children"
                       :key="n"
                       :to="router.path +'/'+ child.path"
                       link
                       class="pl-8">
            <v-list-item-icon>
              <v-icon v-text="child.meta.icon"></v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title v-text="$t(child.meta.title)"></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-group>
      </template>
      <template v-else>
        <v-list-item :key="i"
                     :to="router.path"
                     link>
          <v-list-item-icon>
            <v-icon v-text="router.meta.icon"></v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title v-text="router.meta.title"></v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
    </template>
  </v-list>
</template>

<script>
export default {
  name: "Navigation",
  props: {
    routers: {
      type: Array,
      default: function () {
        return [
          {
            meta: { title: 'Console', icon: 'mdi-console' },
            children: [
              {
                meta: { title: 'ProjjectInfo', icon: 'mdi-card-text' },
              },
              {
                meta: { title: 'HostInfo', icon: 'mdi-laptop' },
              },
              {
                meta: { title: 'ClusterInfo', icon: 'mdi-server-network' },
              },
            ]
          },
          {
            meta: { title: 'Components', icon: 'mdi-video-input-component' },
          }
        ]
      }
    }
  },
  data () {
    return {
      selectedItem: 1,
    }
  },
}
</script>

<style>
</style>