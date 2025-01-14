import apiClient from "./client";
import { IHistory } from "@/types/history";

async function getHistory(): Promise<IHistory[]> {
  const res = await apiClient.get<IHistory[]>("/history");
  return res.data;
}

export default { getHistory };
