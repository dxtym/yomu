import axios from "axios";
import { useEffect, useState } from "react";
import WebApp from "@twa-dev/sdk";
import Gallery from "@/components/common/Gallery";
import Empty from "@/components/common/Empty";

const Search = (props: any) => {
  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<any>([]);

  useEffect(() => {
    const fetchManga = setTimeout(() => {
      axios
        .get(`${url}/search`, {
          headers: {
            authorization: `tma ${WebApp.initData}`,
            "ngrok-skip-browser-warning": "true",
          },
          params: { title: props.search.split(" ").join("-") },
        })
        .then((res) => {
          setData(res.data);
          console.log(res.data);
          document.body.style.height = "auto";
          document.body.style.overflow = "auto";
        })
        .catch((err) => console.error(err));
    }, 1000);

    return () => clearTimeout(fetchManga);
  }, [props.search]);

  return <>{data ? <Gallery data={data} hasSearch /> : <Empty />}</>;
};

export default Search;
