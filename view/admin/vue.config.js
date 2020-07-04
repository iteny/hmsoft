// module.exports = {
//     configureWebpack: {
//         devServer: {
//             //MOCK接口编写的地方
//             //每次做更改这个配置文件的时候，都必须重启项目才会生效
//             before(app) {
//                 //模拟登陆的一个接口
//                 let tockenStr = 'dhclass'
//                 //(req,res)回调方法，req：请求；res:返回的结果
//                 app.get('/api/login', (req, res) => {
//                     const { username, password } = req.query
//                     if (username == 'admin' && password == '123456' ||
//                         username == 'dh' && password == '123456') {
//                         res.json({
//                             code: 0,
//                             message: '登陆成功',
//                             tocken: tockenStr + '-' +
//                                 username + '-' +
//                                 (new Date().getTime() + 60 * 60 * 1000)
//                         })
//                     } else {
//                         res.json({
//                             code: 1,
//                             message: '账号或密码错误'
//                         })
//                     }
//                 })
//             }
//         }
//     }
// }
