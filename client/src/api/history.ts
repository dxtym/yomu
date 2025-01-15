import apiClient from "./client";
import { IHistory } from "@/types/history";

async function getHistory(): Promise<IHistory[]> {
  const res = await apiClient.get<IHistory[]>("/history");
  return res.data;
}

async function removeHistory(id: number): Promise<void> {
  await apiClient.delete("/history", {
    params: { id: id }
  });
}

export default { getHistory, removeHistory };
