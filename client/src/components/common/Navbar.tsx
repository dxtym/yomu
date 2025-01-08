import Option from "./Option";
import { Container, Flex, Link } from "@chakra-ui/react";

interface NavbarProps {
  navs: string[];
}

export default function Navbar({ navs }: NavbarProps) {
  return (
    <Container
      py={"20px"}
      position={"fixed"}
      bottom={0}
      zIndex={1}
      bgColor={"black"}
    >
      <Flex justifyContent={"center"} alignItems={"center"}>
        {navs.map((nav: string, index: number) => (
          <Link key={index} href={`/${nav.toLowerCase()}`}>
            <Option text={nav} index={index} />
          </Link>
        ))}
      </Flex>
    </Container>
  );
}
