import { API_URL } from "@/utils/api";

export function getMainPhoto(restaurant) {
  if (!restaurant?.photos?.length) return null;

  const main = restaurant.photos.find(p => p.is_main);
  if (!main?.url) return null;

  if (main.url.startsWith("http")) {
    return main.url;
  }

  return `${API_URL}${main.url}`;
}
