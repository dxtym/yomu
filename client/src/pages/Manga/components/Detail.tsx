import {
  Center,
  Container,
  Flex,
  Image,
  Heading,
  HStack,
  Text,
} from "@chakra-ui/react";
import Chapter from "./Chapter";

const Detail = (props: any) => {
  return (
    <Container
      py={"25px"}
      px={"25px"}
      position={"relative"}
      mt={"75px"}
      mb={"75px"}
    >
      <Flex direction={"column"} gap={"5"}>
        <HStack gapX={5}>
          <Image
            src={props.data.cover_image}
            height={"220px"}
            width={"150px"}
          />
          <Heading size={"lg"}>
            <Center>{props.data.title}</Center>
          </Heading>
        </HStack>
        <Text textStyle={"md"} fontWeight={"medium"}>
          Description
        </Text>
        <Text lineClamp={2}>{props.data.description}</Text>
        {props.data.chapters &&
          props.data.chapters.map((chapter: any, index: number) => {
            return (
              <Chapter key={index} name={chapter.name} url={chapter.url} />
            );
          })}
      </Flex>
    </Container>
  );
};

export default Detail;
