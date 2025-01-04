import { Container, Grid, GridItem, Image } from "@chakra-ui/react";
import { Link } from "react-router-dom";

interface GalleryProps {
  data: Array<any>;
  hasSearch?: boolean;
}

export default function Gallery({ data, hasSearch = false }: GalleryProps) {
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
            <GridItem key={index}>
              <Link to={`${item.manga_url}`}>
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
