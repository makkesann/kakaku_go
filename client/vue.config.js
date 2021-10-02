module.exports = {
  //   css: {
  //   loaderOptions: {
  //     scss: {
  //       prependData: '@import "./src/assets/style/global.scss";'
  //     }
  //   }
  // },
  pluginOptions: {
    autoRouting: {
      chunkNamePrefix: 'page-'
    }
  },
  devServer: {
    port: 443,
    disableHostCheck: true,
    https: true
},
}
