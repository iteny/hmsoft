import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import './assets/css/global.css'
import axios from 'axios'

// axios.defaults.baseURL = "http://localhost:9090/";
// // 请求超时
// axios.defaults.timeout =30000;
// //  post 请求头
// axios.defaults.headers.post['Content-Type'] ='application/json';
Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
