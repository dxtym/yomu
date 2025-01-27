import axios, { AxiosInstance } from "axios";
import { addLibrary, getLibrary, removeLibrary } from "./library";
import { getHistory, removeHistory } from "./history";
import { getManga, searchManga } from "./manga";
import { getProgress, updateProgress } from "./progress";
import { getChapter } from "./chapter";

class ApiClient {
  public client: AxiosInstance;
  constructor() {
    this.client = axios.create({
      baseURL: import.meta.env.VITE_API_URL,
      headers: { Authorization: `tma ${window.Telegram.WebApp.initData}`}
    })
  }
  
  public getLibrary = getLibrary.bind(this);
  public addLibrary = addLibrary.bind(this);
  public removeLibrary = removeLibrary.bind(this);
  
  public getHistory = getHistory.bind(this);
  public removeHistory = removeHistory.bind(this);

  public getManga = getManga.bind(this);
  public searchManga = searchManga.bind(this);

  public getProgress = getProgress.bind(this);
  public updateProgress = updateProgress.bind(this);

  public getChapter = getChapter.bind(this);
}

export default ApiClient;
