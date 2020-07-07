<template>
  <div class="login">
    <div class="loginBox">
      <div class="avatarBox">
        <img src="../assets/img/avatar.jpg" />
      </div>
      <el-form label-width="60px" :model="formdata" class="loginForm">
        <el-form-item label="用户名">
          <el-input v-model="formdata.username" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item label="密 码">
          <el-input v-model="formdata.password" prefix-icon="el-icon-lock"></el-input>
        </el-form-item>
        <el-form-item class="buttonBox">
          <el-button type="success">登录</el-button>
          <el-button type="info">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  data(){
    return{
      formdata:{
        username:'asdfdsa',
        password:'asdfasdf',
      }
    };
  }
};
</script>
<style scoped>
.login {
  background-color: cadetblue;
  height: 100%;
}
.loginBox {
  width: 450px;
  height: 310px;
  background-color: white;
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
  transform: translate(-25%, -15%);
  background-color: #fff;
}
.avatarBox img {
  width: 60px;
  height: 60px;
  border: 1px solid #eee;
  border-radius: 50%;
}
.buttonBox{
  display: flex;
  /* justify-content: end; */
}
.loginForm{
  position: absolute;
  width: 100%;
  padding: 0 20px;
  bottom: 30px;
  box-sizing: border-box;
}
</style>
