import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";

import { IHistory } from "@/types/history";
import { FaTrash } from "react-icons/fa";
import { FC, ReactElement, useContext, useEffect, useState } from "react";
import {
  Box,
  Container,
  Flex,
  IconButton,
  Stack,
  Text,
} from "@chakra-ui/react";
import Empty from "@/components/common/Empty";
import { ApiClientHooksContext } from "@/app/App";

const History: FC = (): ReactElement => {
  const [history, setHistory] = useState<IHistory[]>();
  const apiClientHooks = useContext(ApiClientHooksContext);

  const handleDelete = (id: number) => {
    const history = apiClientHooks.removeHistory(id);
    if (history.state === "rejected") console.error(history.error);
  };

  useEffect(() => {
    const history = apiClientHooks.getHistory();
    if (history.state === "resolved") setHistory(history.value);
    else if (history.state === "rejected") console.error(history.error); // TODO: error handling
  }, []);

  return (
    <>
      <Header name="History" />
      {history ? (
        <Container mb={"80px"} px={"25px"} position={"relative"} mt={"80px"}>
          <Stack>
            {history?.map((item: IHistory) => {
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
                      <Text>{item.updated_at}</Text>
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

export default History;