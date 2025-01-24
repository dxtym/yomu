import { Box, Text } from "@chakra-ui/react";
import { Link } from "react-router-dom";

interface ChapterProps {
  url: string;
  name: string;
}

const Chapter: React.FC<ChapterProps> = ({ url, name }) => {
  const chapter = url.split("/").slice(4).join("/");

  return (
    <Link to={`/chapter/${chapter}`}>
      <Box borderWidth={"1px"} borderRadius={"5px"} padding={2}>
        <Text padding={"2px"} truncate>{name}</Text>
      </Box>
    </Link>
  );
}

export default Chapter;