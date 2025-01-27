import ApiClient from "./client";
import { IHistory } from "@/types/history";

export async function getHistory(this: ApiClient): Promise<IHistory[]> {
  const res = await this.client.get<IHistory[]>("/history");
  return res.data;
}

export async function removeHistory(this: ApiClient, id: number): Promise<void> {
  await this.client.delete("/history", {
    params: { id: id },
  });
}
