import apiClient from "./client";
import { IManga } from "@/types/manga";
import { IDetail } from "@/types/detail";

async function searchManga(query: string): Promise<IManga[]> {
  const res = await apiClient.get<IManga[]>("/search", {
    params: { query: query.split(" ").join("-") },
  });
  return res.data;
}

async function getManga(manga: string | undefined): Promise<IDetail> {
  const res = await apiClient.get<IDetail>(`/manga/${manga}`);
  console.log(res.data);
  return res.data;
}

export default { searchManga, getManga };
