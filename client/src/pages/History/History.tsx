import HistoryService from "@/api/history";
import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";

import { useEffect, useState } from "react";
import { FaTrash } from "react-icons/fa";
import { IHistory } from "@/types/history";
import {
  Box,
  Container,
  Flex,
  IconButton,
  Stack,
  Text,
} from "@chakra-ui/react";
import Empty from "@/components/common/Empty";

export default function History() {
  const [data, setData] = useState<IHistory[]>();

  useEffect(() => {
    HistoryService.getHistory()
      .then((res) => setData(res))
      .catch((err) => console.error(err));
  }, []);

  return (
    <>
      <Header name="History" />
      {data ? (
        <Container mb={"80px"} px={"25px"} position={"relative"} mt={"80px"}>
          <Stack>
            {data?.map((item: IHistory, index: number) => {
              console.log(item.read_at);
              return (
                <Box
                  borderWidth={"1px"}
                  borderRadius={"5px"}
                  padding={3}
                  key={index}
                >
                  <Flex justifyContent={"space-between"} alignItems={"center"}>
                    <Stack>
                      <Text textStyle={"lg"} fontWeight={"semibold"}>
                        {item.manga}
                      </Text>
                      <Text>{item.read_at}</Text>
                    </Stack>
                    <IconButton variant={"ghost"}>
                      <FaTrash />
                    </IconButton>
                  </Flex>
                </Box>
              );
            })}
          </Stack>
        </Container>
      ) : (
        <Empty />
      )}
      <Navbar />
    </>
  );
}
