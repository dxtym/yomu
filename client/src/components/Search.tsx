import { Container, Input } from "@chakra-ui/react";
import axios from "axios";
import { useEffect, useState } from "react";
import Empty from "./Empty";
import Gallery from "./Gallery";
import WebApp from "@twa-dev/sdk";

export default function Search() {
  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<any>();
  const [search, setSearch] = useState<string>("");

  useEffect(() => {
    const fetchManga = setTimeout(() => {
      if (!search) return;
      axios
        .get(`${url}/search`, {
          headers: { authorization: `tma ${WebApp.initData}` },
          params: { title: search.split(" ").join("-") },
        })
        .then((res) => {
          console.log(res.data);
          setData(res.data);
          document.body.style.height = "auto";
          document.body.style.overflow = "auto";
        })
        .catch((err) => console.error(err));
    }, 1000);

    return () => clearTimeout(fetchManga);
  }, [search]);

  return (
    <Container py={"20px"} px={"25px"} position={"relative"} mt={"75px"}>
      <Input
        placeholder={"Search for manga..."}
        variant={"subtle"}
        size={"lg"}
        onChange={(e) => setSearch(e.target.value)}
      />
      {data && data.length ? (
        <Gallery data={data} hasSearch />
      ) : (
        <Empty hasSearch />
      )}
    </Container>
  );
}
