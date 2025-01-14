import apiClient from "./client";
import { IManga } from "@/types/manga";

async function getLibrary(): Promise<IManga[]> {
  const res = await apiClient.get<IManga[]>("/library");
  return res.data;
}

async function addLibrary(manga: string, coverImage: string): Promise<void> {
  await apiClient.post("/library", {
    manga: manga,
    cover_image: coverImage,
  });
}

async function removeLibrary(manga: string): Promise<void> {
  await apiClient.delete("/library", {
    data: { manga: manga },
  });
}

export default { getLibrary, addLibrary, removeLibrary };
