import service from "@/utils/request";
// @Tags FileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取文件户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
export const getFileList = (data) => {
  return service({
    url: "/fileUploadAndDownload/getFileList",
    method: "get",
    params: data,
  });
};

// @Tags FileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.FileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
export const deleteFile = (data) => {
  return service({
    url: `/fileUploadAndDownload/deleteFile/${data.ID}`,
    method: "delete",
  });
};

/**
 * 编辑文件名或者备注
 * @param data
 * @returns {*}
 */
export const editFileName = (data) => {
  return service({
    url: "/fileUploadAndDownload/editFileName",
    method: "put",
    data,
  });
};

export const uploadFile = (data, is_cropper) => {
  if (is_cropper) {
    return service({
      url: `/fileUploadAndDownload/upload?is_cropper=${is_cropper}`,
      method: "post",
      data,
    });
  }
  return service({
    url: "/fileUploadAndDownload/upload",
    method: "post",
    data,
  });
};
