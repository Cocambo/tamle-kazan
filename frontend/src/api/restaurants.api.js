import { http } from "@/utils/http";

export const restaurantsApi = {
  getAll(params = {}) {
    return http.get("/restaurants", { params });
  },

  getById(id) {
    return http.get(`/restaurants/${id}`);
  },
};
