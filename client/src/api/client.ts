import WebApp from "@twa-dev/sdk";
import axios from "axios";

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    Authorization: `tma ${WebApp.initData}`,
    "ngrok-skip-browser-warning": true,
  },
});

export default apiClient;
