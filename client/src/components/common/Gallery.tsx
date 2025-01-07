import { Container, Grid, GridItem, Image, Alert } from "@chakra-ui/react";
import WebApp from "@twa-dev/sdk";
import axios from "axios";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Gallery = (props: any) => {
  const url = import.meta.env.VITE_API_URL;
  const [manga, setManga] = useState<string>("");
  const [alert, setAlert] = useState<boolean>(false);
  const [coverImage, setCoverImage] = useState<string>("");
  const [longPress, setLongPress] = useState<boolean>(false);

  useEffect(() => {
    let timer: any;
    if (longPress) {
      timer = setTimeout(() => {
        axios
          .post(
            `${url}/library`,
            { manga: manga, cover_image: coverImage },
            {
              headers: {
                authorization: `tma ${WebApp.initData}`,
                "ngrok-skip-browser-warning": "true",
              },
            },
          )
          .then((res) => {
            console.log(res);
            setAlert(true);
            setTimeout(() => {
              setAlert(false);
            }, 3000);
          })
          .catch((err) => console.error(err));
      }, 1500);
    } else {
      clearTimeout(timer);
    }

    return () => clearTimeout(timer);
  }, [longPress, manga, coverImage]);

  return (
    <Container
      position={"relative"}
      mt={props.hasSearch ? "150px" : "80px"}
      mb={"80px"}
      px={"25px"}
    >
      {alert && (
        <Alert.Root
          status={"success"}
          pos={"fixed"}
          top={"15%"}
          left={"50%"}
          transform={"translate(-50%, -50%)"}
          zIndex={"1000"}
          width={"90%"}
        >
          <Alert.Indicator />
          <Alert.Title>Added to the Library</Alert.Title>
        </Alert.Root>
      )}
      <Grid templateColumns={"repeat(2, 1fr)"} gap={5}>
        {props.data &&
          props.data.map((item: any, index: number) => {
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
                    width={"220px"}
                  />
                </Link>
              </GridItem>
            );
          })}
      </Grid>
    </Container>
  );
};

export default Gallery;
