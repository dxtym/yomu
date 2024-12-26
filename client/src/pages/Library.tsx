import Gallery from "@/components/Gallery";
import Header from "@/components/Header";
import Navbar from "@/components/Navbar";
import axios from "axios";
import { useEffect, useState } from "react";
import Empty from "@/components/Empty";
import terminal from "virtual:terminal";

export interface Manga {
  manga_id: number;
  cover_image: string;
}

export default function Library() {
  const [data, setData] = useState<Array<Manga>>([]);

  useEffect(() => {
    const fetchLibrary = async () => {
      const url = import.meta.env.VITE_API_URL as string;

      try {
        const res = await axios.get(`${url}/library`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        });
        setData(res.data);
        if (!res.data || res.data.length === 0) {
          document.body.style.height = "100vh";
          document.body.style.overflow = "hidden";
        }
        terminal.log(res.data);
      } catch (error) {
        terminal.error(error);
      }
    };

    fetchLibrary();
  }, []);

  return (
    <>
      <Header name={"Library"} />
      {data && data.length ? <Gallery data={data} /> : <Empty />}
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
}
