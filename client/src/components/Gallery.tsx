import { Manga } from "@/pages/Library";
import { Container, Grid, GridItem, Image } from "@chakra-ui/react";

export default function Gallery(props: any) {
  return (
    <Container py="25px" px="25px" position="relative" mt="75px">
      <Grid templateColumns="repeat(2, 1fr)" gap="10">
        {props.data.map((item: Manga) => {
          return (
            <GridItem key={item.id}>
              <Image src={item.cover_image} height="200px" width="180px" />
            </GridItem>
          );
        })}
      </Grid>
    </Container>
  );
}
