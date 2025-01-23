import LibraryService from "@/api/library";
import Toaster from "@/components/common/Toaster";

import { IManga } from "@/types/manga";
import { Link } from "react-router-dom";
import { useState } from "react";
import { Center, Container, Grid, GridItem, Image } from "@chakra-ui/react";

export default function Gallery(props: {
  data?: IManga[];
  hasSearch?: boolean;
}) {
  const [toast, setToast] = useState<boolean>(false);
  const [timer, setTimer] = useState<NodeJS.Timeout | null>(null);

  const handlePress = (item: IManga) => {
    const newTimer = setTimeout(() => {
      const action = props.hasSearch
        ? LibraryService.addLibrary(item.manga, item.cover_image)
        : LibraryService.removeLibrary(item.manga);

      action
        .then(() => {
          setToast(true);
          setTimeout(() => setToast(false), 3000);
          if (!props.hasSearch) {
            window.location.reload();
          }
        })
        .catch((err) => console.error(err));
    }, 1000);

    setTimer(newTimer);
  };

  const cancelPress = () => {
    if (timer) {
      clearTimeout(timer);
      setTimer(null);
    }
  };

  return (
    <Container
      mb={"80px"}
      px={"25px"}
      position={"relative"}
      mt={props.hasSearch ? "150px" : "80px"}
    >
      {toast && <Toaster />}
      <Grid
        templateColumns={{ base: "repeat(2, 1fr)", md: "repeat(4, 1fr)" }}
        gap={5}
      >
        {props.data?.map((item: IManga, index: number) => {
          return (
            <GridItem
              key={index}
              onMouseDown={() => handlePress(item)}
              onMouseUp={() => cancelPress()}
              onMouseLeave={() => cancelPress()}
              onTouchStart={() => handlePress(item)}
              onTouchEnd={() => cancelPress()}
            >
              <Link to={`/browse/${item.manga}`}>
                <Center>
                  <Image src={item.cover_image} height="250px" width="100%" />
                </Center>
              </Link>
            </GridItem>
          );
        })}
      </Grid>
    </Container>
  );
}
