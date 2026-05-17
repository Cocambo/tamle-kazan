import { defineStore } from "pinia";
import { http } from "@/utils/http";

export const useUsersStore = defineStore("users", {
  state: () => ({
    users: {},
  }),

  actions: {
    async fetchUser(id) {
      if (this.users[id]) return;

      try {
        const res = await http.get(`/user/users/${id}`);
        this.users[id] = res.data;
      } catch (e) {
        console.error("Ошибка загрузки пользователя:", e);
      }
    },
  },

  getters: {
    getUserById: (state) => (id) => state.users[id],
  },
});