import ApiClient from "@/app/api/client";
import { IChapter } from "@/types/chapter";
import { IDetail } from "@/types/detail";
import { FetchState } from "@/types/fetch";
import { IHistory } from "@/types/history";
import { IManga } from "@/types/manga";
import { useEffect, useState } from "react";

export class ApiClientHooks {
  apiClient: ApiClient;
  constructor(apiClient: ApiClient) {
    this.apiClient = apiClient;
  }

  getLibrary = () => {
    const [state, setState] = useState<FetchState<IManga[]>>({ state: "pending"});

    useEffect(() => {
      this.apiClient.getLibrary()
        .then((res) => setState({ state: "resolved", value: res }))
        .catch((err) => setState({ state: "rejected", error: err }));
    }, [this.apiClient]);

    return state;
  }

  addLibrary = (manga: string, coverImage: string) => {
    const [state, setState] = useState<FetchState<null>>({ state: "pending" });
    
    useEffect(() => {
      this.apiClient.addLibrary(manga, coverImage)
        .then(() => setState({ state: "resolved", value: null }))
        .catch((err) => setState({ state: "rejected", error: err }))
    }, [this.apiClient, manga, coverImage])

    return state;
  }

  removeLibrary = (manga: string) => {
    const [state, setState] = useState<FetchState<null>>({ state: "pending" });
    
    useEffect(() => {
      this.apiClient.removeLibrary(manga)
        .then(() => setState({ state: "resolved", value: null }))
        .catch((err) => setState({ state: "rejected", error: err }))
    }, [this.apiClient, manga])

    return state;
  }

  searchManga = (query: string) => {
    const [state, setState] = useState<FetchState<IManga[]>>({ state: "pending" });

    useEffect(() => {
      this.apiClient.searchManga(query)
        .then((res) => setState({ state: "resolved", value: res }))
        .catch((err) => setState({ state: "rejected", error: err }))
    }, [this.apiClient, query])

    return state;
  }
  
  getManga = (manga: string) => {
    const [state, setState] = useState<FetchState<IDetail>>({ state: "pending" });

    useEffect(() => {
      this.apiClient.getManga(manga)
        .then((res) => setState({ state: "resolved", value: res }))
        .catch((err) => setState({ state: "rejected", error: err}))
    }, [this.apiClient, manga])

    return state;
  }

  getHistory = () => {
    const [state, setState] = useState<FetchState<IHistory[]>>({ state: "pending" });

    useEffect(() => {
      this.apiClient.getHistory()
        .then((res) => setState({ state: "resolved", value: res }))
        .catch((err) => setState({ state: "rejected", error: err }))
    }, [this.apiClient])

    return state;
  }

  removeHistory = (id: number) => {
    const [state, setState] = useState<FetchState<null>>({ state: "pending" });

    useEffect(() => {
      this.apiClient.removeHistory(id)
        .then(() => setState({ state: "resolved", value: null }))
        .catch((err) => setState({ state: "rejected", error: err}))
    }, [this.apiClient, id])

    return state;
  }

  getChapter = (manga: string, chapter: string) => {
    const [state, setState] = useState<FetchState<IChapter>>({ state: "pending" });

    useEffect(() => {
      this.apiClient.getChapter(manga, chapter)
        .then((res) => setState({ state: "resolved", value: res }))
        .catch((err) => setState({ state: "rejected", error: err}))
    }, [this.apiClient, manga, chapter])

    return state;
  }
}
