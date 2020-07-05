<template>
  <div class="login">
    <div class="loginBox">
      <div class="avatarBox">
        <img src="../assets/img/avatar.jpg" />
      </div>
      <el-form label-width="80px">
        <el-form-item label="活动名称">
          <el-input></el-input>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<style scoped>
.login {
  background-color: cadetblue;
  height: 100%;
}
.loginBox {
  width: 450px;
  height: 350px;
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
</style>
