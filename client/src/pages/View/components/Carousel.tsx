import Loading from "@/components/common/Loading";
import { IChapter } from "@/types/chapter";
import { Container, Box, Image, IconButton } from "@chakra-ui/react";
import React, { useState } from "react";
import { FaArrowCircleLeft, FaArrowCircleRight } from "react-icons/fa";

const Carousel: React.FC<IChapter> = ({ page_urls }) => {
  const [page, setPage] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(false);

  const showNext = () => {
    if (page + 1 < page_urls.length) {
      setLoading(true);
      setPage(page + 1);
    }
  }

  const showPrev = () => {
    if (page - 1 >= 0) {
      setLoading(true);
      setPage(page - 1);
    }
  }

  return (
    <Container
      display={"flex"}
      justifyContent={"center"}
      alignItems={"center"}
    >
      <Box width={"100%"} height={"100%"} pos={"relative"}>
        {loading && <Loading />}
        <Image src={page_urls[page]} display={loading ? "none" : "block"} onLoad={() => setLoading(false)} objectFit={"cover"} />
        <Box pt={"10px"} display={"flex"} justifyContent={"space-between"}>
          <IconButton onClick={showPrev}>
            <FaArrowCircleLeft />
          </IconButton>
          <IconButton onClick={showNext}>
            <FaArrowCircleRight />
          </IconButton>
        </Box>
      </Box>
    </ Container>
  );
}

export default Carousel;