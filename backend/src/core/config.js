/**
 * 网站配置文件
 */

const config = {
  appName: "后台管理",
  appLogo: "/Thesquid.ink-Free-Flat-Sample-Space-rocket.ico",
  showViteLogo: true,
};

export const viteLogo = (env) => {
  if (config.showViteLogo) {
  }
};

export const fileSize = 2; // 上传文件大小限制，单位MB

export default config;
