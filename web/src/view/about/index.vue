<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-divider>About Me</el-divider>
          </template>
          <div>
            <el-row>
              <el-col>{{ x }} {{ y }}</el-col>
              <el-col>{{ isDark }}</el-col>
              <el-col>{{ store.name }}</el-col>
              <el-col>
                <!-- <useMouse v-slot="{ x, y }"> x: {{ x }} y: {{ y }} </useMouse> -->
                <Button :lists="lists">
                  <template v-slot:two="{ item }">
                    <div>{{ item.name }}</div>
                  </template>
                  <template v-slot:active>
                    <div>123123123</div>
                  </template>
                </Button>
              </el-col>
              <!-- <el-col :span="8" :offset="0">
                <ButtonSlot :action="['default']">bilibili</ButtonSlot>
              </el-col> -->
              <el-col>
                <draggable-vue />
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "About",
};
</script>

<script setup>
// import ButtonSlot from "./slot";
import { version } from "vue";
import draggableVue from "./draggable.vue";
import { useMouse, usePreferredDark, useLocalStorage } from "@vueuse/core";
import Button from "./Tabs/Button.vue";

console.log(version);
const { x, y } = useMouse();
const isDark = usePreferredDark();
const store = useLocalStorage("appColor", {
  name: "Apple",
  color: "red",
});

const lists = [
  { name: "aaa", age: 12 },
  { name: "bbb", age: 78 },
  { name: "ccc", age: 45 },
];
</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
