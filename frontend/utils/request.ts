import axios from 'axios';

const service = axios.create({
  baseURL: 'http://127.0.0.1',
  timeout: 9999,
});

service.interceptors.request.use(
  (config) => {
    config.headers = {
      'Content-Type': 'application/json',
      ...config.headers,
    };

    return config;
  },
  (err) => {
    return err;
  }
);

export default service;
