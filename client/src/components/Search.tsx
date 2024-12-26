import { Manga } from "@/pages/Library";
import { Container, Input } from "@chakra-ui/react";
import axios from "axios";
import { useState } from "react";
import terminal from "virtual:terminal";
import Empty from "./Empty";
import Gallery from "./Gallery";

export default function Search() {
  const [data, setData] = useState<Array<Manga>>([]);
  const [search, setSearch] = useState<string>("");

  const handleSearch = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const url = import.meta.env.VITE_API_URL as string;
    setSearch(e.target.value);
    setTimeout(async () => {
      try {
        const res = await axios.get(`${url}/search`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
          params: {
            title: search,
          }
        })
        setData(res.data);
        if (res.data  && res.data.length !== 0) {
          document.body.style.height = "";
          document.body.style.overflow = "";
        }
        terminal.log(res.data);
      } catch (error) {
        terminal.error(error);
      }
    }, 2000);
  }

  return (
    <Container py={"20px"} px={"25px"} position={"relative"} mt={"75px"}>
      <Input placeholder={"Search for manga..."} variant={"subtle"} size={"lg"} onChange={handleSearch}/>
      {data && data.length ? <Gallery data={data} hasSearch /> : <Empty hasSearch />}
    </Container>
  );
}
