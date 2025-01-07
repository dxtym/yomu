import { Flex, Container, Text } from "@chakra-ui/react";

const Empty = () => {
  return (
    <Container
      height={"100vh"}
      display={"flex"}
      justifyContent={"center"}
      alignItems={"center"}
    >
      <Flex justifyContent={"center"} alignItems={"center"} flexDir={"column"}>
        <Text fontSize={"3xl"}>^•⩊•^</Text>
        <Text>Nothing Here Yet</Text>
      </Flex>
    </Container>
  );
};

export default Empty;
