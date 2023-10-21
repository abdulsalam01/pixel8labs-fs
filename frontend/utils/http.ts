// utils/http.js
import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: process.env.API_URL,
  timeout: +process.env.API_TIMEOUT ?? 5000,
  headers: {
    'Content-Type': 'application/json',
  },
});

export default axiosInstance;
