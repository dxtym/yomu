import ApiClient from "./client";
import { IManga } from "@/types/manga";
import { IDetail } from "@/types/detail";

export async function searchManga(this: ApiClient, query: string): Promise<IManga[]> {
  const res = await this.client.get<IManga[]>("/search", {
    params: { query: query.split(" ").join("-") },
  });
  return res.data;
}

export async function getManga(this: ApiClient, manga: string): Promise<IDetail> {
  const res = await this.client.get<IDetail>(`/manga/${manga}`);
  return res.data;
}
