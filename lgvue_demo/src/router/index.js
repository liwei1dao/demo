import Vue from 'vue'
import VueRouter from 'vue-router'
import LayoutV from '@/layout/LayoutV.vue'
import store from '@/store'
import Message from '@/components/message/'

Vue.use(VueRouter)
/*权限*/
const Role = {
  Guester: 0,                    //游客
  Operator: 1,                   //运营
  Developer: 2,                  //开发
  Master: 3,                     //超级管理员
}

/*静态路由*/
const staticRoutes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/Register.vue')
  },
  {
    path: '/',
    name: 'Index',
    component: LayoutV,
    redirect: '/index',
    meta: { title: 'Index', icon: 'mdi-home-floor-b', role: Role.Guester },
    children: [
      {
        path: 'index',
        component: () => import('@/views/Index.vue'),
        meta: { title: 'Index', icon: 'mdi-home-floor-b', role: Role.Guester },
      }
    ]
  },
]

/*动态路由*/
const Dynamicroutes = [

  {//控制台
    path: '/console',
    component: LayoutV,
    meta: { title: 'Console', icon: 'mdi-home-floor-b', role: Role.Operator },
    children: [
      {
        path: 'index',
        component: () => import('@/views/console/Index.vue'),
        meta: { title: 'ProjjectInfo', icon: 'mdi-card-text', role: Role.Operator },
      },
      {
        path: 'hostinfo',
        component: () => import('@/views/console/HostInfo.vue'),
        meta: { title: 'HostInfo', icon: 'mdi-laptop', role: Role.Operator },
      },
      {
        path: 'clusterinfo',
        component: () => import('@/views/console/ClusterInfo.vue'),
        meta: { title: 'ClusterInfo', icon: 'mdi-server-network', role: Role.Developer },
      },
    ]
  },
  {//组件演示
    path: '/components',
    component: LayoutV,
    meta: { title: 'Components', icon: 'mdi-video-input-component', role: Role.Guester },
    children: [
      {
        path: 'message',
        component: () => import('@/views/components/Message.vue'),
        meta: { title: 'Message', icon: 'mdi-message-processing-outline', role: Role.Guester },
      },
      {
        path: 'echarts',
        component: () => import('@/views/components/Echarts.vue'),
        meta: { title: 'Echarts', icon: 'mdi-chart-sankey', role: Role.Guester },
      },
      {
        path: 'textccroll',
        component: () => import('@/views/components/TextCcroll.vue'),
        meta: { title: 'TextCcroll', icon: 'mdi-sign-text', role: Role.Guester },
      },
    ]
  }
]

// 定义登录页面名称（为了方便理解才定义的）
const BasePages = ['Login', 'Register']
const router = createRouter()

function createRouter () {
  return new VueRouter({
    routes: staticRoutes,
  })
}
/*重置路由*/
export function resetRouter () {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

/*添加动态路由*/
export function addRoutes (role) {
  var routes = []
  for (const route of Dynamicroutes) {
    if (route.meta.role == 0 || route.meta.role & role > 0) {
      var temproute = { path: route.path, component: route.component, meta: route.meta }
      temproute.children = null
      if (route.children && route.children.length > 0) {
        temproute.children = []
        for (const child of route.children) {
          var tempchild = { path: child.path, component: child.component, meta: child.meta }
          var isopn = child.meta.role & role
          if (child.meta.role == 0 || isopn > 0) {
            temproute.children.push(tempchild)
          }
        }
      }
      routes.push(temproute)
    }
  }
  return routes
}

// /*定义全局前置守卫（里面有两个坑要注意）*/
router.beforeEach(async (to, from, next) => {
  const token = store.getters.token
  const userinfo = store.getters.userinfo
  // if (to.name == LOGIN_PAGE_NAME) {
  if (BasePages.includes(to.name)) {
    next()
  } else {
    if (token && userinfo) {
      next()
    } else {
      if (token) {
        try {
          await store.dispatch('user/getuserinfo')
          await store.dispatch('app/setrouters', store.getters.userinfo.UserRole ? store.getters.userinfo.UserRole : 0)
          router.addRoutes(store.getters.routers)
          next({ ...to, replace: true });
        } catch (error) {
          Message.error(error.message)
          await store.dispatch('user/resetToken')
          // 如果没有登录而且前往的页面不是登录页面，跳转到登录页
          next({ name: BasePages[0] })
        }
      } else {
        next({ name: BasePages[0] })
      }
    }
  }
})

export default router
