import { FC, ReactElement } from "react";
import { Box, Flex, IconButton } from "@chakra-ui/react";
import { FaArrowCircleLeft, FaArrowCircleRight } from "react-icons/fa";

interface PagerProps {
  showNext: () => void;
  showPrev: () => void;
}

const Pager: FC<PagerProps> = ({showNext, showPrev}): ReactElement => {
  return (
    <Box 
      py={"20px"} 
      bottom={0} 
      left={0} 
      pos={"fixed"} 
      width={"100%"} 
      display={"block"} 
      bgColor={"black"}
    >
      <Flex justifyContent={"space-around"}>
        <IconButton onClick={showPrev}>
          <FaArrowCircleLeft />
        </IconButton>
        <IconButton onClick={showNext}>
          <FaArrowCircleRight />
        </IconButton>
      </Flex>
    </Box>
  )
}

export default Pager;