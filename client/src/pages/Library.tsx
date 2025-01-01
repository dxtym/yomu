import Header from "@/components/Header";
import Navbar from "@/components/Navbar";
import axios from "axios";
import { useEffect, useState } from "react";
import Empty from "@/components/Empty";
import WebApp from "@twa-dev/sdk";

interface Manga {
  manga_url: number;
  cover_image: string;
}

export default function Library() {
  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<Array<Manga>>([]);

  useEffect(() => {
    const fetchLibrary = async () => {
      axios
        .get(`${url}/library`, {
          headers: { authorization: `tma ${WebApp.initData}` },
        })
        .then((res) => {
          setData(res.data);
          if (!res.data || res.data.length == 0) {
            document.body.style.height = "100vh";
            document.body.style.overflow = "hidden";
          }
        })
        .catch((err) => console.error(err));
    };

    fetchLibrary();
  }, []);

  return (
    <>
      <Header name={"Library"} />
      {/* TODO: fetch data from library */}
      <Empty />
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
}
