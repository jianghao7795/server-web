export const colorItem = ["primary", "info", "success", "warning", "default", "error"];

export const colorIndex = (index: number): "default" | "primary" | "info" | "success" | "warning" | "error" => {
  return colorItem[index % 6] as "default" | "primary" | "info" | "success" | "warning" | "error";
};
