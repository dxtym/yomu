import Header from "@/components/Header";
import Navbar from "@/components/Navbar";
import Empty from "@/components/Empty";
import { useEffect, useState } from "react";
import WebApp from "@twa-dev/sdk";
import axios from "axios";
import Gallery from "@/components/Gallery";

export default function Library() {
  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<any>([]);

  useEffect(() => {
    axios
      .get(`${url}/library`, {
        headers: { 
          authorization: `tma ${WebApp.initData}`,
          "ngrok-skip-browser-warning": "true",
        },
      })
      .then((res) => {
        setData(res.data);
        console.log(res.data);
        if (!res.data || res.data.length == 0) {
          document.body.style.height = "100vh";
          document.body.style.overflow = "hidden";
        }
      })
      .catch((err) => console.error(err));
  }, []);

  return (
    <>
      <Header name={"Library"} />
      {data && data.length ? <Gallery data={data} /> : <Empty />}
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
}
