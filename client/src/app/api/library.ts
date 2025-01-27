import ApiClient from "./client";
import { IManga } from "@/types/manga";

export async function getLibrary(this: ApiClient): Promise<IManga[]> {
  const res = await this.client.get<IManga[]>("/library");
  return res.data;
}

export async function addLibrary(
  this: ApiClient, 
  manga: string, 
  coverImage: string
): Promise<void> {
  await this.client.post("/library", {
    manga: manga,
    cover_image: coverImage,
  });
}

export async function removeLibrary(this: ApiClient, manga: string): Promise<void> {
  await this.client.delete("/library", {
    data: { manga: manga },
  });
}
