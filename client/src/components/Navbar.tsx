import { Container, Flex } from "@chakra-ui/react";
import Option from "./Option";

export default function Navbar(props: any) {
  return (
    <Container py="25px" position="fixed" bottom={0} zIndex={1} bgColor="black">
      <Flex>
        {props.navs.map((nav: string, index: number) => (
          <Option text={nav} index={index} />
        ))}
      </Flex>
    </Container>
  );
}
