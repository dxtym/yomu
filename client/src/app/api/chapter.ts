import ApiClient from "./client";
import { IChapter } from "@/types/chapter";

export async function getChapter(
  this: ApiClient,
  manga: string,
  chapter: string,
): Promise<IChapter> {
  const res = await this.client.get<IChapter>(`/chapter/${manga}/${chapter}`);
  return res.data;
}
