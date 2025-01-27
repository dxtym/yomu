import { Flex, Container, Text } from "@chakra-ui/react";
import { FC, ReactElement } from "react";

const Empty: FC = (): ReactElement => {
  return (
    <Container
      height={"100vh"}
      display={"flex"}
      justifyContent={"center"}
      alignItems={"center"}
    >
      <Flex justifyContent={"center"} alignItems={"center"} flexDir={"column"}>
        <Text fontSize={"3xl"}>^•⩊•^</Text>
        <Text>Empty</Text>
      </Flex>
    </Container>
  );
}

export default Empty;