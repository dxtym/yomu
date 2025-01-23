import { Container, Flex, Heading, Input } from "@chakra-ui/react";
import { IoMdArrowRoundBack } from "react-icons/io";

export default function Header(props: {
  name: string;
  hasSearch?: boolean;
  onClick?: () => void;
  setQuery?: (query: string) => void;
}) {
  return (
    <Container
      top={0}
      left={0}
      py={"20px"}
      px={"25px"}
      zIndex={1000}
      position={"fixed"}
      bgColor={"black"}
    >
      <Flex
        direction={{ base: "column", md: "row" }}
        justify={{ md: "space-between" }}
        gap={5}
      >
        <Heading textStyle={"2xl"} onClick={props.onClick}>
          {props.name === "Back" ? <IoMdArrowRoundBack /> : props.name}
        </Heading>
        <Input
          display={props.hasSearch ? "block" : "none"}
          placeholder={"Search for manga..."}
          variant={"subtle"}
          size={"lg"}
          maxW={{ base: "100%", md: "400px" }}
          onChange={(e) => props.setQuery && props.setQuery(e.target.value)}
        />
      </Flex>
    </Container>
  );
}
