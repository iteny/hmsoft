<template>
  <div class="login">
    <div class="loginBox">
      <div class="avatarBox">
        <img src="../assets/img/avatar.jpg" />
      </div>
      <el-form ref="loginFormRef" :model="formdata" :rules="rules" class="loginForm">
        <el-form-item prop="username">
          <el-input v-model="formdata.username" placeholder="请输入您的用户名" suffix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="formdata.password"
            placeholder="请输入您的密码"
            suffix-icon="el-icon-lock"
            type="password"
          ></el-input>
        </el-form-item>
        <el-form-item class="buttonBox">
          <el-button type="success" @click="loginButton()" style="width:48%">登录</el-button>
          <el-button type="info" @click="resetButton()" style="width:48%">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      formdata: {
        username: "",
        password: ""
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" }, //trigger鼠标失去焦点触发事件
          { min: 5, max: 11, message: "长度在 5 到 11 个字符", trigger: "blur" }
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" }, //trigger鼠标失去焦点触发事件
          { min: 5, max: 11, message: "长度在 5 到 11 个字符", trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    resetButton() {
      this.$refs.loginFormRef.resetFields();
    },
    loginButton() {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return;
        const res = await this.$http.get("/api/admin/login", {
          params: this.formdata
        });
        console.log(res.data);
      });
    }
  }
};
</script>
<style scoped>
.login {
  background: #000000;
  background: -moz-linear-gradient(top, #000000 0%, #ffffff 100%);
  background: -webkit-gradient(
    linear,
    left top,
    left bottom,
    color-stop(0%, #000000),
    color-stop(100%, #ffffff)
  );
  background: -webkit-linear-gradient(top, #000000 0%, #ffffff 100%);
  background: -o-linear-gradient(top, #000000 0%, #ffffff 100%);
  background: -ms-linear-gradient(top, #000000 0%, #ffffff 100%);
  background: linear-gradient(to bottom, #000000 0%, #ffffff 100%);
  filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#000000', endColorstr='#ffffff',GradientType=0 );
  height: 100%;
}
.loginBox {
  width: 360px;
  height: 310px;
  /* background-color: white; */
  top: 50%;
  left: 50%;
  position: absolute;
  transform: translate(-50%, -50%);
  border-radius: 5px;
}
.avatarBox {
  width: 60px;
  height: 60px;
  border: 1px solid #eee;
  border-radius: 50%;
  padding: 5px;
  box-shadow: 0 0 10px #ddd;
  position: absolute;
  transform: translate(200%, -15%);
  background-color: #fff;
}
.avatarBox img {
  width: 60px;
  height: 60px;
  border: 1px solid #eee;
  border-radius: 50%;
}
.buttonBox {
  /* display: flex; */
  /* justify-content: end; */
}
.loginForm {
  position: absolute;
  width: 100%;
  padding: 0 20px;
  bottom: 30px;
  box-sizing: border-box;
}
</style>
