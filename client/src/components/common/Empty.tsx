import { Flex, Container, Text } from "@chakra-ui/react";

export default function Empty() {
  return (
    <Container
      height={"100%"}
      display={"flex"}
      justifyContent={"center"}
      alignItems={"center"}
    >
      <Flex justifyContent={"center"} alignItems={"center"} flexDir={"column"}>
        <Text fontSize={"3xl"}>^•⩊•^</Text>
        <Text>Gomenasai!</Text>
      </Flex>
    </Container>
  );
}
