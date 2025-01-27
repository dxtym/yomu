import { ApiClientHooksContext } from "@/app/App";
import Empty from "@/components/common/Empty";
import Gallery from "@/components/common/Gallery";

import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";

import { IManga } from "@/types/manga";
import { FC, ReactElement, useContext, useEffect, useState } from "react";

const Browse: FC = (): ReactElement => {
  const [manga, setManga] = useState<IManga[]>();
  const [query, setQuery] = useState<string>("");
  const apiClientHooks = useContext(ApiClientHooksContext);

  useEffect(() => {
    const fetch = setTimeout(() => {
      const manga = apiClientHooks.searchManga(query);
      if (manga.state === "resolved") {
        setManga(manga.value);
        document.body.style.height = "auto";
        document.body.style.overflow = "auto";
      } else if (manga.state === "rejected") {
        console.error(manga.error); // TODO: error handling
      }
    }, 1000);

    return () => clearTimeout(fetch);
  }, [query]);

  return (
    <>
      <Header name={"Browse"} setQuery={setQuery} hasSearch />
      {manga ? <Gallery data={manga} hasSearch /> : <Empty />}
      <Navbar />
    </>
  );
}

export default Browse;