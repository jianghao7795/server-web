/*
 * @Author: jianghao
 * @Date: 2022-07-29 09:48:24
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-27 10:27:59
 */
/*
 * web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from "./global";

export default {
  install: (app) => {
    register(app);
  },
};
