import Gallery from "@/components/Gallery";
import Header from "@/components/Header";
import Navbar from "@/components/Navbar";
import axios from "axios";
import { useEffect, useState } from "react";
import { Text } from "@chakra-ui/react";

export interface Manga {
  id: number;
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
        console.log(res.data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchLibrary();
  }, []);

  return (
    <>
      <Header name="Library" />
      {data ? <Gallery data={data} /> : <Text>Loading...</Text>}
      <Navbar navs={["Library", "Search", "History"]} />
    </>
  );
}
