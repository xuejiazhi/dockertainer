<template>
  <div class="center-wrap">
    <div class="title">用户中心</div>
    <div class="content">
      <div class="left"><img src="../assets/logo.png" alt="user-logo"></div>
      <div class="right">admin</div>
    </div>
  </div>
</template>

<style lang="scss">
.center-wrap{
  margin: auto;
  width: 500px;
  margin-top: 100px;
  text-align: left;
  border: 1px solid #ddd;
  .title{
    height: 50px;
    line-height: 50px;
    background: #dfdfdf;
    color: #333;
    font-weight: bold;
    padding-left: 20px;
    font-size: 14px;
    border-bottom: 1px solid #ddd;
  }
  .content{
    display: flex;
    flex-direction: row;
    padding: 23px 0;
    .left{
      width: 100px;
      text-align: center;
      img{
        display: inline-block;
        width: 48px;
        height: 48px;
        background: #ccc;
        border-radius: 50%;
        border: 1px solid #ddd;
      }
    }
    .right{
      height: 48px;
      line-height: 48px;
      font-size: 14px;
      color: #666;
    }

  }

}
</style>