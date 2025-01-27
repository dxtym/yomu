import { Flex, Spinner } from "@chakra-ui/react";
import { FC, ReactElement } from "react";

const Loading: FC = (): ReactElement => {
  return (
    <Flex justifyContent={"center"} alignItems={"center"} height={"100vh"}>
      <Spinner size={"xl"} />
    </Flex>
  );
}

export default Loading;