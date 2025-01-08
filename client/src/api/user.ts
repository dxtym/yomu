import apiClient from "./client";
import { IUser } from "@/types/user";

async function createUser(user: IUser): Promise<IUser> {
  const res = await apiClient.post<IUser>("/user", user);
  return res.data;
}

export default { createUser };
