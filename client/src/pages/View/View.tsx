import ChapterService from "@/api/chapter";
import Carousel from "./components/Carousel";

import { IChapter } from "@/types/chapter";
import { Container } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

export default function View() {
  const params = useParams();
  const [data, setData] = useState<IChapter>();

  useEffect(() => {
    ChapterService.getChapter(params.manga, params.chapter)
      .then((res) => setData(res))
      .catch((err) => console.error(err));
  }, []);

  // TODO: rewrite this to custom one
  return (
    <Container>
      <Carousel />
    </Container>
  );
}
