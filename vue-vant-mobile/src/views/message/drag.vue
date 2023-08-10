<template>
  <TransitionGroup name="list" tag="div" class="container">
    <div
      class="item"
      v-for="(item, i) in drag.list"
      :key="item.id"
      draggable="true"
      @dragstart="dragstart($event, i)"
      @dragenter="dragenter($event, i)"
      @dragend="dragend"
      @dragover="dragover"
    >
      {{ item.name }}
    </div>
  </TransitionGroup>
</template>
<script setup lang="ts">
  import { reactive } from 'vue';
  const drag = reactive({
    list: [
      { name: 'a', id: 1 },
      { name: 'b', id: 2 },
      { name: 'c', id: 3 },
      { name: 'd', id: 4 },
      { name: 'e', id: 5 },
    ],
  });

  let dragIndex = 0;

  function dragstart(e: DragEvent, index: number) {
    e.stopPropagation();
    dragIndex = index;
    // setTimeout(() => {
    //   e?.target?.class?.add('moveing');
    // }, 0);
  }
  function dragenter(e: DragEvent, index: number) {
    console.log(e, index);
    e.preventDefault();
    // 拖拽到原位置时不触发
    if (dragIndex !== index) {
      const source = drag.list[dragIndex];
      drag.list.splice(dragIndex, 1);
      drag.list.splice(index, 0, source);

      // 更新节点位置
      dragIndex = index;
    }
  }
  function dragover(e: DragEvent) {
    e.preventDefault();
    if (e?.dataTransfer?.dropEffect) {
      e.dataTransfer.dropEffect = 'move';
    }
  }
  function dragend(e: DragEvent) {
    // e.target.classList.remove('moveing');
  }
</script>

<style lang="less" scoped>
  .item {
    width: 200px;
    height: 40px;
    line-height: 40px;
    // background-color: #f5f6f8;
    background-color: skyblue;
    text-align: center;
    margin: 10px;
    color: #fff;
    font-size: 18px;
  }

  .container {
    position: relative;
    padding: 0;
  }

  .moveing {
    opacity: 0;
  }

  .list-move, /* 对移动中的元素应用的过渡 */
    .list-enter-active,
    .list-leave-active {
    transition: all 0.2s ease;
  }
</style>
