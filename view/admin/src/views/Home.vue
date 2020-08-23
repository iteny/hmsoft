<template>
  <div class="home">
    <el-container>
      <el-header>
        <div>
          <img src="../assets/img/avatar.jpg" />
          <span>系统管理后台</span>
        </div>
        <el-button type="success" @click="logout()" size="mini">退出按钮</el-button>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <el-menu
            default-active="2"
            class="el-menu-vertical-demo"
            @open="handleOpen"
            @close="handleClose"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
          >
            <!-- 一级菜单 -->
            <el-submenu index="1">
              <!-- 一级菜单模板 -->
              <template slot="title">
                <!-- 菜单图标 -->
                <i class="el-icon-location"></i>
                <!-- 文本 -->
                <span>导航一</span>
              </template>
              <!-- 一级菜单下二级菜单 -->
              <el-submenu index="1-4">
                <template slot="title">
                  <!-- 菜单图标 -->
                  <i class="el-icon-location"></i>
                  <!-- 文本 -->
                  <span>二级菜单</span>
                </template>

                <!-- <el-menu-item index="1-4-1">选项1</el-menu-item> -->
              </el-submenu>
            </el-submenu>
          </el-menu>
        </el-aside>
        <el-main>Main</el-main>
      </el-container>
    </el-container>
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
    async login() {
      const res = await this.$http.post("/admin/loginstatus", {
        // username: this.formdata.username,
        // password: this.formdata.password,
      });
      if (res.data.status == "success") {
        this.$message.success("您的用户名是" + res.data.info);
      } else if (res.data.status == "faild") {
        this.$message.error("您还未登录！");
        this.$router.push("login");
      }
    },
  },
  created() {
    this.login();
  },
};
</script>
<style scoped>
.el-container {
  height: 100%;
}
.el-header {
  background: chocolate;
  line-height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.el-header img {
  widows: 40px;
  height: 40px;
  border: 1px solid #eee;
  border-radius: 50%;
  padding: 1px;
  box-shadow: 0 0 10px #ddd;
  margin-right: 25px;
}
.el-header div {
  display: flex;
  align-items: center;
}
.el-aside {
  background: cornflowerblue;
  line-height: 200px;
}
.el-main {
  background: darkgreen;
  line-height: 200px;
}
</style>
