import Chapter from "./Chapter";

import {
  Center,
  Container,
  Flex,
  Image,
  Heading,
  HStack,
  Text,
} from "@chakra-ui/react";
import { IDetail } from "@/types/detail";

interface DetailProps {
  data?: IDetail;
}

function Detail(props: DetailProps) {
  const title = props.data?.title ?? "Untitled";
  const coverImage = props.data?.cover_image ?? "/default.jpg";
  const description = props.data?.description ?? "No description available.";
  const chapters = props.data?.chapters ?? [];

  return (
    <Container px={"25px"} my={"80px"} position={"relative"}>
      <Flex direction={"column"} gap={5}>
        <HStack gapX={5}>
          <Image src={coverImage} height={"220px"} width={"150px"} />
          <Heading size={"lg"}>
            <Center>{title}</Center>
          </Heading>
        </HStack>
        <Text textStyle={"md"} fontWeight={"medium"}>
          Description
        </Text>
        <Text lineClamp={2}>{description}</Text>
        {chapters?.map((chapter: any, index: number) => {
          return <Chapter key={index} name={chapter.name} url={chapter.url} />;
        })}
      </Flex>
    </Container>
  );
}

export default Detail;
