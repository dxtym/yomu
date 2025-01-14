import apiClient from "./client";

async function updateProgress(
  manga: string | undefined,
  chapter: string | undefined,
  page: number,
): Promise<void> {
  await apiClient.put("/progress", {
    manga: manga,
    chapter: chapter,
    page: page,
  });
}

async function getProgress(
  manga: string | undefined,
  chapter: string | undefined,
): Promise<number> {
  const res = await apiClient.get<number>("/progress", {
    params: { manga: manga, chapter: chapter },
  });

  return res.data;
}

export default { updateProgress, getProgress };
