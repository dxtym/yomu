import { Container, Heading } from "@chakra-ui/react";

export default function Header(props: any) {
  return (
    <Container
      top={0}
      py="25px"
      px="25px"
      zIndex={1}
      position="fixed"
      bgColor="black"
    >
      <Heading textStyle="2xl">{props.name}</Heading>
    </Container>
  );
}
