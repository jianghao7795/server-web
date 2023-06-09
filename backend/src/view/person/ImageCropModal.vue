<template>
  <el-dialog v-model="status" title="裁剪图片">
    <div style="width: 100%; height: 400px; margin-bottom: 10px">
      <vue-cropper
        ref="cropper"
        :img="props.imgUrl"
        :output-size="options.size"
        :output-type="options.outputType"
        :info="true"
        :full="options.full"
        :fixed="false"
        :fixed-number="[75, 34]"
        :can-move="options.canMove"
        :can-move-box="options.canMoveBox"
        :fixed-box="options.fixedBox"
        :original="options.original"
        :auto-crop="options.autoCrop"
        :auto-crop-width="options.autoCropWidth"
        :auto-crop-height="options.autoCropHeight"
        :center-box="options.centerBox"
        @real-time="realTime"
        :high="options.high"
        mode="contain"
        :max-img-size="options.max"
      ></vue-cropper>
    </div>
    <el-card header="预览">
      <div style="min-height: 200px">
        <div :style="{ width: previews.w + 'px', height: previews.h + 'px', overflow: 'hidden', margin: '5px' }">
          <div :style="previews.div">
            <img :src="previews.url" :style="previews.img" />
          </div>
        </div>
      </div>
    </el-card>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="oncancel">取消</el-button>
        <el-button type="primary" @click="onOk">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
export default {
  name: "ImageCropModal",
};
</script>

<script setup>
import { uploadFile } from "@/api/fileUploadAndDownload";
import { reactive, ref, computed } from "vue";
import { VueCropper } from "vue-cropper";

const props = defineProps({
  imgUrl: String,
  dialogFormVisible: Boolean,
  cropData: Object,
});

const emits = defineEmits(["changeImageStatus", "changeAvatar"]);

const options = reactive({
  img: props.imgUrl,
  size: 1,
  full: false,
  outputType: "png",
  canMove: true,
  fixedBox: false,
  original: false,
  canMoveBox: true,
  autoCrop: true,
  // 只有自动截图开启 宽度高度才生效
  autoCropWidth: 750,
  autoCropHeight: 340,
  centerBox: true,
  high: true,
  max: 99999,
});

const status = computed({
  get() {
    return props.dialogFormVisible;
  },
  set(val) {
    emits("changeImageStatus", val, "");
  },
});

const previews = ref({});

const cropper = ref(null);

const realTime = (data) => {
  previews.value = data;
  //   emits("changeImage", { url: data, name: "预览" });
};

const oncancel = () => {
  status.value = false;
};
const onOk = () => {
  cropper.value.getCropBlob(async (data) => {
    const file = new window.File([data], "screenshot_" + props.cropData.name, { type: data.type });
    const forms = new FormData();
    forms.append("file", file);
    const resp = await uploadFile(forms, 2);
    if (resp.code === 0) {
      ElMessage({
        type: "success",
        message: "截图成功 上传成功！",
      });

      status.value = false;
      emits("changeAvatar", resp.data.file.url);
    }
  });
};
</script>

<style scoped>
.model-show {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
