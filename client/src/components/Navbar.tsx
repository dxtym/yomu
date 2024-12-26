import { Container, Flex, Link } from "@chakra-ui/react";
import Option from "./Option";

export default function Navbar(props: any) {
  return (
    <Container
      py={"25px"}
      position={"fixed"}
      bottom={0}
      zIndex={1}
      bgColor={"black"}
    >
      <Flex justifyContent={"center"}>
        {props.navs.map((nav: string, index: number) => (
          <Link href={`/${nav.toLowerCase()}`}>
            <Option text={nav} index={index} />
          </Link>
        ))}
      </Flex>
    </Container>
  );
}
