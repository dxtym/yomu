export interface IDetail {
  title: string;
  cover_image: string;
  description: string;
  chapters: {
    name: string;
    url: string;
  }[];
}