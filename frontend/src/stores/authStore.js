import { defineStore } from 'pinia';
import { authApi } from '@/api/auth.api';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    accessToken: localStorage.getItem('accessToken'),
    refreshToken: localStorage.getItem('refreshToken'),
  }),

  getters: {
    isAuthenticated: (state) => !!state.user,
  },

  actions: {
    setTokens(access, refresh) {
      this.accessToken = access;
      this.refreshToken = refresh;
      localStorage.setItem('accessToken', access);
      localStorage.setItem('refreshToken', refresh);
    },

    setAccessToken(access) {
      this.accessToken = access;
      localStorage.setItem('accessToken', access);
    },

    async login(credentials) {
      const { data } = await authApi.login(credentials);
      this.setTokens(data.access_token, data.refresh_token);
      await this.fetchProfile();
    },

    async register(data) {
      return authApi.register(data);
    },

    async fetchProfile() {
      const { data } = await authApi.getProfile();
      this.user = data;
    },

    async resendEmailConfirmation() {
      if (!this.user?.email) return;
      await authApi.resendConfirmation(this.user.email);
    },

    async logout() {
      if (this.refreshToken) {
        await authApi.logout(this.refreshToken);
      }
      this.user = null;
      this.accessToken = null;
      this.refreshToken = null;
      localStorage.clear();
    },
  },
});
