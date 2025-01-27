import { Container, Flex, Heading, Input } from "@chakra-ui/react";
import { FC, ReactElement } from "react";
import { IoMdArrowRoundBack } from "react-icons/io";

interface HeaderProps {
  name: string;
  hasSearch?: boolean;
  onClick?: () => void;
  setQuery?: (query: string) => void;
}

const Header: FC<HeaderProps> = ({ name, hasSearch, onClick, setQuery }): ReactElement => {
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
        <Heading textStyle={"2xl"} onClick={onClick}>
          {name === "Back" ? <IoMdArrowRoundBack /> : name}
        </Heading>
        <Input
          display={hasSearch ? "block" : "none"}
          placeholder={"Search for manga..."}
          variant={"subtle"}
          size={"lg"}
          maxW={{ base: "100%", md: "400px" }}
          onChange={(e) => setQuery && setQuery(e.target.value)}
        />
      </Flex>
    </Container>
  );
}

export default Header;