import apiClient from "./client";
import { IChapter } from "@/types/chapter";

async function getChapter(
  manga: string | undefined,
  chapter: string | undefined,
): Promise<IChapter> {
  const res = await apiClient.get<IChapter>(`/chapter/${manga}/${chapter}`);
  return res.data;
}

export default { getChapter };
