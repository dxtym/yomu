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

  const handleDelete = async (id: number) => {
    try {
      await HistoryService.removeHistory(id);
      window.location.reload();
    } catch (err) {
      console.error(err);
    }
  };

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
            {data?.map((item: IHistory) => {
              return (
                <Box
                  borderWidth={"1px"}
                  borderRadius={"5px"}
                  padding={3}
                  key={item.id}
                >
                  <Flex justifyContent={"space-between"} alignItems={"center"}>
                    <Stack>
                      <Text textStyle={"lg"} fontWeight={"semibold"}>
                        {item.manga}
                      </Text>
                      <Text>{item.read_at}</Text>
                    </Stack>
                    <IconButton
                      variant={"ghost"}
                      onClick={() => handleDelete(item.id)}
                    >
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
