import axios from "axios";
import { useStore } from "./store";

export const instance = axios.create({
  baseURL: import.meta.env.VITE_API_URL as string,
  withCredentials: true,
});

// Add a request interceptor
instance.interceptors.request.use(
  (config) => {
    // Get sessionId from Zustand store
    const sessionId = useStore.getState().sessionId;

    // If sessionId exists, set it in the headers
    if (sessionId) {
      config.headers["session_id"] = sessionId;
    }

    return config;
  },
  (error) => {
    // Do something with request error
    return Promise.reject(error);
  }
);
