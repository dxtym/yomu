import { Container, Flex, Link } from "@chakra-ui/react";
import Option from "./Option";

const Navbar = (props: any) => {
  return (
    <Container
      py={"20px"}
      position={"fixed"}
      bottom={0}
      zIndex={1}
      bgColor={"black"}
    >
      <Flex justifyContent={"center"} alignItems={"center"}>
        {props.navs.map((nav: string, index: number) => (
          <Link key={index} href={`/${nav.toLowerCase()}`}>
            <Option text={nav} index={index} />
          </Link>
        ))}
      </Flex>
    </Container>
  );
};

export default Navbar;
