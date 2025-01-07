import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";
import Empty from "@/components/common/Empty";
import { useEffect, useState } from "react";
import WebApp from "@twa-dev/sdk";
import axios from "axios";
import Gallery from "@/components/common/Gallery";

const Library = () => {
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
      {data ? <Gallery data={data} /> : <Empty />}
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
};

export default Library;
