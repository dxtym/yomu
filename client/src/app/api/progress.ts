import ApiClient from "./client";

export async function updateProgress(
  this: ApiClient,
  manga: string | undefined,
  chapter: string | undefined,
  page: number,
): Promise<void> {
  await this.client.put("/progress", {
    manga: manga,
    chapter: chapter,
    page: page,
  });
}

export async function getProgress(
  this: ApiClient,
  manga: string | undefined,
  chapter: string | undefined,
): Promise<number> {
  const res = await this.client.get<number>("/progress", {
    params: { manga: manga, chapter: chapter },
  });

  return res.data;
}
