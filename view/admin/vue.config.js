// module.exports = {
//   lintOnSave: false,
//   // devServer: {
//   //   proxy: {
//   //     '/admin': {
//   //       target: "http://localhost:9090/", //目标服务器
//   //       ws: true,
//   //       changeOrigin: true,//是否改变请求源
//   //       pathRewrite: { //路径重写
//   //         "^/admin": ''
//   //       }
//   //     },
//   //     // "/log": {
//   //     //   target: "http://192.168.0.201:12368",     // 另一个要请求的地址
//   //     //   ws: true,
//   //     //   changeOrigin: true,
//   //     //   pathRewrite: {
//   //     //     "^/log": "/"
//   //     //   }
//   //     // }
//   //   }
//   // }
// }
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
  }

};
