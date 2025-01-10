import LibraryService from "@/api/library";
import Toaster from "@/components/common/Toaster";

import { IManga } from "@/types/manga";
import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import { Container, Grid, GridItem, Image } from "@chakra-ui/react";

interface GalleryProps {
  data: IManga[];
  hasSearch?: boolean;
}

export default function Gallery(props: GalleryProps) {
  const [manga, setManga] = useState<string>("");
  const [toast, setToast] = useState<boolean>(false);
  const [coverImage, setCoverImage] = useState<string>("");
  const [longPress, setLongPress] = useState<boolean>(false);
  const action = props.hasSearch
    ? LibraryService.addLibrary(manga, coverImage)
    : LibraryService.removeLibrary(manga);

  const handlePress = (item: IManga) => {
    setLongPress(true);
    setManga(item.manga);
    setCoverImage(item.cover_image);
  };
  
  useEffect(() => {
    let timer: any;
    if (longPress) {
      timer = setTimeout(() => {
        action
          .then(() => {
            setToast(true);
            setTimeout(() => {
              setToast(false);
            }, 3000);
          })
          .catch((err) => console.error(err));
      }, 1000);
    } else {
      clearTimeout(timer);
    }

    return () => clearTimeout(timer);
  }, [manga, longPress, coverImage]);

  return (
    <Container
      mb={"80px"}
      px={"25px"}
      position={"relative"}
      mt={props.hasSearch ? "150px" : "80px"}
    >
      {toast && <Toaster />}
      <Grid templateColumns={"repeat(2, 1fr)"} gap={5}>
        {props.data?.map((item: IManga, index: number) => {
          return (
            <GridItem
              key={index}
              onMouseDown={() => handlePress(item)}
              onMouseUp={() => setLongPress(false)}
              onMouseLeave={() => setLongPress(false)}
              onTouchStart={() => handlePress(item)}
              onTouchEnd={() => setLongPress(false)}
            >
              <Link to={`/browse/${item.manga}`}>
                <Image
                  src={item.cover_image}
                  height={"220px"}
                  width={"220px"}
                />
              </Link>
            </GridItem>
          );
        })}
      </Grid>
    </Container>
  );
}
