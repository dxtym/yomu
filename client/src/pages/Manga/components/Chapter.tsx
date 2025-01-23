import { Box, Text } from "@chakra-ui/react";
import { Link } from "react-router-dom";

export default function Chapter(props: { url: string; name: string }) {
  const chapter = props.url.split("/").slice(4).join("/");

  return (
    <Link to={`/chapter/${chapter}`}>
      <Box borderWidth={"1px"} borderRadius={"5px"} padding={2}>
        <Text padding={"2px"} truncate>
          {props.name}
        </Text>
      </Box>
    </Link>
  );
}
