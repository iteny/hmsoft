const webpack = require('webpack');
module.exports = {
  publicPath: '/',
  outputDir: 'dist',
  lintOnSave: false,
  devServer: {
    open: true,
    host: 'localhost',
    port: '8080',
    proxy: {
      '': {
        target: 'http://localhost:9090', // 要请求的地址
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^': ''
        }
      }
    }
  },
  configureWebpack: {//引入jquery
    plugins: [
      new webpack.ProvidePlugin({
        $: "jquery",
        jQuery: "jquery",
        "windows.jQuery": "jquery"
      })
    ]
  },

};
