import { Container, Grid, GridItem, Image } from "@chakra-ui/react";
import WebApp from "@twa-dev/sdk";
import axios from "axios";
import { useEffect, useRef, useState } from "react";
import { Link } from "react-router-dom";

interface GalleryProps {
  data: Array<any>;
  hasSearch?: boolean;
}

export default function Gallery({ data, hasSearch = false }: GalleryProps) {
  const url = import.meta.env.VITE_API_URL;
  const [manga, setManga] = useState<string>("");
  const [coverImage, setCoverImage] = useState<string>("");
  const [longPress, setLongPress] = useState<boolean>(false);

  useEffect(() => {
    let timer: any;
    if (longPress) {
      timer = setTimeout(() => {
        axios
          .post(`${url}/library`,
            { manga: manga, cover_image: coverImage },
            {
              headers: {
                "authorization": `tma ${WebApp.initData}`,
                "ngrok-skip-browser-warning": "true",
              },
            })
          .then((res) => console.log(res))
          .catch((err) => console.error(err));
        
        // toaster.create({
        //   description: "Added to the Library",
        //   status: "info",
        // })
      }, 1500);
    } else {
      clearTimeout(timer);
    }

    return () => clearTimeout(timer);
  }, [longPress, manga, coverImage]);

  return (
    <Container
      py={"25px"}
      px={hasSearch ? "0" : "25px"}
      position={"relative"}
      mt={hasSearch ? "0" : "75px"}
      mb={"75px"}
    >
      <Grid templateColumns={"repeat(2, 1fr)"} gap={"10"}>
        {data && data.map((item: any, index: number) => {
          return (
            <GridItem 
              key={index}
              onMouseDown={() => {
                setLongPress(true);
                setManga(item.manga_url);
                setCoverImage(item.cover_image);
              }}
              onMouseUp={() => setLongPress(false)}
              onMouseLeave={() => setLongPress(false)}
              onTouchStart={() => {
                setLongPress(true);
                setManga(item.manga_url);
                setCoverImage(item.cover_image);
              }}
              onTouchEnd={() => setLongPress(false)}
            >
              <Link to={`/browse/${item.manga_url}`}>
                <Image
                  src={item.cover_image}
                  height={"220px"}
                  width={"200px"}
                />
              </Link>
            </GridItem>
          );
        })}
      </Grid>
    </Container>
  );
}
