import { Container, Flex, Heading, Input } from "@chakra-ui/react";

const Header = (props: any) => {
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
          {props.name}
        </Heading>
        <Input
          display={props.hasSearch ? "block" : "none"}
          placeholder={"Search for manga..."}
          variant={"subtle"}
          size={"lg"}
          maxW={{ base: "100%", md: "400px" }}
          onChange={(e) => props.setSearch(e.target.value)}
        />
      </Flex>
    </Container>
  );
};

export default Header;
