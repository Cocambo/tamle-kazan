import axios from 'axios';
import { useAuthStore } from '@/stores/authStore';

const API_URL = 'http://localhost:8080/api';

export const http = axios.create({  
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

http.interceptors.request.use((config) => {
  const authStore = useAuthStore();
  if (authStore.accessToken) {
    config.headers.Authorization = `Bearer ${authStore.accessToken}`;
  }
  return config;
});

http.interceptors.response.use(
  (response) => response,
  async (error) => {
    const authStore = useAuthStore();
    const originalRequest = error.config;

    if (
      error.response?.status === 401 &&
      !originalRequest._retry &&
      authStore.refreshToken
    ) {
      originalRequest._retry = true;

      try {
        const { data } = await axios.post(
          `${API_URL}/user/refresh`,
          { refresh_token: authStore.refreshToken }
        );

        authStore.setAccessToken(data.access_token);
        originalRequest.headers.Authorization = `Bearer ${data.access_token}`;

        return http(originalRequest);
      } catch {
        authStore.logout();
      }
    }

    return Promise.reject(error);
  }
);
