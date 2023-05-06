<template>
  <footer class="footerend">
    <el-backtop></el-backtop>
    <div class="bottom-info">
      <div>
        <span>Powered by chengdu</span>
        <el-divider direction="vertical" />
        <span>Copyright @</span>
      </div>
    </div>
  </footer>
</template>

<script>
// 此文件内容为版权信息，如需改动，请联系wx:shouzi_1994购买授权 未授权状态只需保留此代码 不影响任何正常使用
// 项目为apatch协议 请遵守版权协议内容
export default {
  name: "BottomInfo",
};
</script>

<style lang="scss">
.footerend {
  flex: auto;
  min-height: 0;
}
.bottom-info {
  color: #888;
  flex: auto;
  min-height: 0;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span {
      margin: 0 3px;
    }
  }
}
</style>
