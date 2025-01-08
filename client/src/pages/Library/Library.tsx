import LibraryService from "@/api/library";

import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";
import Empty from "@/components/common/Empty";
import Gallery from "@/components/common/Gallery";

import { IManga } from "@/types/manga";
import { useEffect, useState } from "react";

function Library() {
  const [data, setData] = useState<IManga[]>([]);

  useEffect(() => {
    LibraryService.getLibrary()
      .then((res) => {
        setData(res);
        document.body.style.height = "auto";
        document.body.style.overflow = "auto";
      })
      .catch((err) => console.error(err));
  }, []);

  return (
    <>
      <Header name={"Library"} />
      {data ? <Gallery data={data} /> : <Empty />}
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
}

export default Library;
