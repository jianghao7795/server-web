<template>
  <div class="chatIcon">
    <n-popover popper-class="biaoqing" :width="400" trigger="click" placement="bottom">
      <template #trigger>
        <n-button quaternary>
          <grinning-face theme="outline" size="18" fill="#333" :strokeWidth="3" @click="handleEmoji" />
        </n-button>
      </template>
      <div class="emotionList">
        <a href="javascript:void(0);" @click="getEmo(item)" v-for="item in emoji" :key="item.codes" class="emotionItem">
          {{ item.char }}
        </a>
      </div>
    </n-popover>
  </div>
</template>
<script lang="ts" setup>
import { defineEmits } from "vue";
import emoji from "@/common/emoji";
import { GrinningFace } from "@icon-park/vue-next";

// defineProps 和 defineEmits 用法
// const props = defineProps<{ view: number }>();
const emits = defineEmits<{ (e: "changeActive", num: EmojiType.emoji): void; (e: "handleFouces"): void }>();

const getEmo = (emoji: EmojiType.emoji) => {
  emits("changeActive", emoji);
};

const handleEmoji = () => {
  emits("handleFouces");
};
// onMounted(() => {
//   faceList.value = emoji.map((i) => i.char);
// });
// export default {
//   mounted() {
//     for (let i in appData) {
//       this.faceList.push(appData[i].char);
//     }
//   },
//   data() {
//     return {
//       faceList: [],
//       emojiItem: null,
//     };
//   },
//   methods: {
//     getEmo(index) {
//       this.emojiItem = index;
//       this.$emit("changeText");
//     },
//   },
// };
</script>
<style scoped>
.emotionSelect {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.emotionSelect img {
  width: 14px;
  height: 14px;
}
.chatIcon {
  padding: 0 10px;
  font-size: 25px;
}
.emotionList {
  display: flex;
  flex-wrap: wrap;
  padding: 5px;
}
.emotionItem {
  width: 10%;
  font-size: 20px;
  text-align: center;
}
/*包含以下四种的链接*/
.emotionItem {
  text-decoration: none;
}
/*正常的未被访问过的链接*/
.emotionItem:link {
  text-decoration: none;
}
/*已经访问过的链接*/
.emotionItem:visited {
  text-decoration: none;
}
/*鼠标划过(停留)的链接*/
.emotionItem:hover {
  text-decoration: none;
}
/* 正在点击的链接*/
.emotionItem:active {
  text-decoration: none;
}
</style>
