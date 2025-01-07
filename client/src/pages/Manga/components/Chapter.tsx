import { Box, Text } from "@chakra-ui/react";
import { Link } from "react-router-dom";

const Chapter = (props: any) => {
  return (
    <Link to={`/chapter/${props.url.split("/").slice(4).join("/")}`}>
      <Box borderWidth={"1px"} borderRadius={"5px"}>
        <Text padding={"2px"} truncate>
          {props.name}
        </Text>
      </Box>
    </Link>
  );
};

export default Chapter;
