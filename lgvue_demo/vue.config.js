const path = require('path')
const name = 'LgVue' // page title

function resolve (dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    port: 8080,
    open: true,
    overlay: {
      warnings: false,
      errors: true
    },
    proxy: {
      // 把key的路径代理到target位置
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      ["/lego"]: { //需要代理的路径   例如 '/lego'
        target: `http://127.0.0.1:9567`, //代理到 目标路径
        changeOrigin: true,
        secure: false,
        logLevel: 'debug',
        // pathRewrite: { // 修改路径数据
        //   ['^/lego']: '' // 举例 '^/lego:""' 把路径中的/lego字符串删除
        // }
      }
    },
  },
  configureWebpack: {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    devtool: 'source-map',
    name: name,
    resolve: {
      alias: {
        '@': resolve('src')
      }
    }
  },
}
