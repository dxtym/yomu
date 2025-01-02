export interface IChapter {
  name: string;
  url: string;
}

export interface IDetail {
  title: string;
  cover_image: string;
  description: string;
  chapters: IChapter[];
}
