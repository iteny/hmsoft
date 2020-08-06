<template>
  <div class="home">
    <el-button type="success" plain>成功按钮</el-button>
    <el-button type="success" plain>成功按钮</el-button>
    <el-button type="success" plain>成功按钮</el-button>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";

export default {
  data() {
    return {
      loginObj: {
        username: "admindsaf",
        password: "123456",
      },
    };
  },
  name: "Home",
  components: {
    // HelloWorld
  },
  methods: {
    // async/ await来发送异步请求，从服务端获取数据，代码很简洁，同时async/await 已经被标准化。
    // async login () {
    //   const res = await this.$http.get('/api/login', { params: this.loginObj })
    //   console.log(res.data)
    // }
  },
  created() {
    this.login();
  },
};
console.log("asdfsadf");
</script>
