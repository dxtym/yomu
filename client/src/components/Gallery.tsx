import { Manga } from "@/pages/Library";
import { Container, Grid, GridItem, Image } from "@chakra-ui/react";
import terminal from "virtual:terminal";

interface GalleryProps {
  data: Array<Manga>;
  hasSearch?: boolean;
}

export default function Gallery({ data, hasSearch = false }: GalleryProps) {
  terminal.log(hasSearch);
  return (
    <Container 
      py={"25px"} 
      px={hasSearch ? "0" : "25px"} 
      position={"relative"} 
      mt={hasSearch ? "0" : "75px"}
      mb={"75px"}
    >
      <Grid templateColumns={"repeat(2, 1fr)"} gap={"10"}>
        {data.map((item: Manga) => {
          return (
            <GridItem key={item.manga_id}>
              <Image src={item.cover_image} height={"220px"} width={"200px"} />
            </GridItem>
          );
        })}
      </Grid>
    </Container>
  );
}
