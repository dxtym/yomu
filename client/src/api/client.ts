import axios from "axios";

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    Authorization: `tma ${window.Telegram.WebApp.initData}`,
  },
});

export default apiClient;
