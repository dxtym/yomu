import MangaService from "@/api/manga";
import Empty from "@/components/common/Empty";
import Gallery from "@/components/common/Gallery";

import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";

import { IManga } from "@/types/manga";
import { useEffect, useState } from "react";

export default function Browse() {
  const [data, setData] = useState<IManga[]>();
  const [query, setQuery] = useState<string>("");

  useEffect(() => {
    const fetch = setTimeout(() => {
      MangaService.searchManga(query)
        .then((res) => {
          if (res) {
            setData(res);
            document.body.style.height = "auto";
            document.body.style.overflow = "auto";
          }
        })
        .catch((err) => console.error(err));
    }, 1000);

    return () => clearTimeout(fetch);
  }, [query]);

  return (
    <>
      <Header name={"Browse"} setQuery={setQuery} hasSearch />
      {data ? <Gallery data={data} hasSearch /> : <Empty />}
      <Navbar />
    </>
  );
}
