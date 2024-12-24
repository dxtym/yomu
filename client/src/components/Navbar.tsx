import { Container, Flex } from "@chakra-ui/react";
import NavButton from "./NavButton";

export default function Navbar(props: any) {
    return (
    <Container className='navbar'>
      <Flex>
        {props.navs.map((nav: string, index: number) => (
            <NavButton text={nav} index={index} />
        ))}
      </Flex>
    </Container>
  );
}