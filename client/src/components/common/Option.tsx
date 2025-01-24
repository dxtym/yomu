import { Button, Container, Text } from "@chakra-ui/react";
import { FaSearch, FaBook, FaHistory } from "react-icons/fa";

interface OptionProps {
  text: string;
  index: number;
}

const Option: React.FC<OptionProps> = ({ text, index }) => {
  const renderIcons = () => {
    switch (index) {
      case 0:
        return <FaBook />;
      case 1:
        return <FaSearch />;
      case 2:
        return <FaHistory />;
    }
  };

  return (
    <Container>
      <Button
        variant={"plain"}
        display={"flex"}
        flexDir={"column"}
        justifyContent={"center"}
        alignItems={"center"}
        gap={"5px"}
      >
        {renderIcons()}
        <Text textStyle={"md"}>{text}</Text>
      </Button>
    </Container>
  );
}

export default Option;