export default {
  install: (app, options) => {
    app.config.globalProperties.$translate = (key) => {
      // 获取 `options` 对象的深层属性
      // 使用 `key` 作为索引
      return key.split(".").reduce((o, i) => {
        if (o) return o[i];
      }, options);
    };
  },
};
