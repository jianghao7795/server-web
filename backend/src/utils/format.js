import { formatTimeToStr } from "@/utils/date";
import { getDict } from "@/utils/dictionary";

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? "开启" : "停用";
  } else {
    return "";
  }
};
export const formatDate = (time) => {
  if (time !== null && time !== "") {
    var date = new Date(time);
    return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
  } else {
    return "";
  }
};

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter((item) => item.value === value);
  return rowLabel && rowLabel[0] && rowLabel[0].label;
};

export const getDictFunc = async (type) => {
  const dicts = await getDict(type);
  return dicts;
};

export const fileSizeChange = (value = 0, type = "B") => {
  if (value === 0) return 0;
  let size = 0;
  switch (type) {
    case "MB":
      size = value / 1024 / 1024;
      break;
    case "GB":
      size = value / 1024 / 1024 / 1024;
      break;
    case "KB":
      size = value / 1024;
      break;
    default:
      size = value;
      break;
  }

  return size.toFixed(2);
};
