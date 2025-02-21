import Toaster from "@/components/common/Toaster";

import { IManga } from "@/types/manga";
import { Link } from "react-router-dom";
import { FC, ReactElement, useContext, useState } from "react";
import { Center, Container, Grid, GridItem, Image } from "@chakra-ui/react";
import { ApiClientHooksContext } from "@/app/App";

interface GalleryProps {
  data?: IManga[];
  hasSearch?: boolean;
}

const Gallery: FC<GalleryProps> = ({ data, hasSearch }): ReactElement => {
  const [toast, setToast] = useState<boolean>(false);
  const [timer, setTimer] = useState<any | null>(null);
  const apiClientHooks = useContext(ApiClientHooksContext);

  const handlePress = (item: IManga) => {
    const newTimer = setTimeout(() => {
      const action = hasSearch 
        ? apiClientHooks.addLibrary(item.manga, item.cover_image) 
        : apiClientHooks.removeLibrary(item.manga);
      
      if (action.state === "rejected") console.error(action.error); // TODO: error handling
      if (action.state === "resolved") {
        setToast(true);
        setTimeout(() => setToast(false), 3000);
        if (!hasSearch) {
          window.location.reload();
        }
      }

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
      mt={hasSearch ? "150px" : "80px"}
    >
      {toast && <Toaster />}
      <Grid
        templateColumns={{ base: "repeat(2, 1fr)", md: "repeat(4, 1fr)" }}
        gap={5}
      >
        {data?.map((item: IManga, index: number) => {
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

export default Gallery;