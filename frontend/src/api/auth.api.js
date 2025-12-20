import { http } from '@/utils/http';

export const authApi = {
  register(data) {
    return http.post('/user/register', data);
  },

  login(data) {
    return http.post('/user/login', data);
  },

  getProfile() {
    return http.get('/user/profile');
  },

  resendConfirmation(email) {
    return http.post('/user/resend-confirmation', { email });
  },

  logout(refreshToken) {
    return http.post('/user/logout', {
      refresh_token: refreshToken,
    });
  },
};
