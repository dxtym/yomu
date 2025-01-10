import apiClient from "./client";
import { IManga } from "@/types/manga";

async function getLibrary(): Promise<IManga[]> {
  const res = await apiClient.get<IManga[]>("/library");
  return res.data;
}

async function addLibrary(manga: string, coverImage: string) {
  const res = await apiClient.post("/library", {
    manga: manga,
    cover_image: coverImage,
  });
  return res.data;
}

async function removeLibrary(manga: string) {
  const res = await apiClient.delete("/library", {
    data: {manga: manga}
  })
  return res.data;
}

export default { getLibrary, addLibrary, removeLibrary };
