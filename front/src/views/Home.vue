<template>
  <div class="login-box">
    <el-form :model="form" label-width="90px">
      <el-form-item label="用户名">
        <el-input v-model="form.loginName" placeholder="请输入用户名" />
      </el-form-item>
      <el-form-item label="密  码">
        <el-input v-model="form.password" placeholder="请输入密码" />
      </el-form-item>

      <el-form-item class="btn">
        <el-button type="primary" @click="onSubmit" class="btn-login"
          >登录</el-button
        >
        <el-button @click="onReset" class="btn-login"> 重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script setup>
import { reactive } from "vue";
import { useRouter } from "vue-router"
const router = useRouter()
const form = reactive({
  loginName: "",
  password: "",
});

const onSubmit = () => {
  console.log("submit!");
  router.push("/center")
};
const onReset = () => {
  console.log("onReset!");
};
</script>
<style lang="scss">
 .login-box {
   top: 40%;
   position: absolute;
   left: 50%;
   transform: translate(-50%,-50%);
   max-width: 600px;
   min-width: 350px;
   .el-form-item__label{
    text-align:justify;
    text-justify:inter-ideograph;
   }
   .btn{

   }
 }
.btn-login{
 width: 80px;
 margin: 0 20px;
}
</style>
