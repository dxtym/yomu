import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";
import Gallery from "@/components/common/Gallery";

import { ApiClientHooksContext } from "@/app/App";
import { FC, ReactElement, useContext } from "react";
import Loading from "@/components/common/Loading";

const Library: FC = (): ReactElement => {
  const apiClientHooks = useContext(ApiClientHooksContext);
  const manga = apiClientHooks.getLibrary();

  if (manga.state === "resolved") {
    document.body.style.height = "auto";
    document.body.style.overflow = "auto";
  } else if (manga.state === "rejected") {
    console.error(manga.error); // TODO: error handling
  }

  return (
    <>
      <Header name={"Library"} />
      {manga.state === "resolved" ? <Gallery data={manga.value} /> : <Loading /> }
      <Navbar />
    </>
  );
}

export default Library;