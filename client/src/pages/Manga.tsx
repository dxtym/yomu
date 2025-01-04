import axios from "axios";
import Detail from "@/components/Detail";
import Header from "@/components/Header";
import WebApp from "@twa-dev/sdk";
import { Flex, Spinner } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { IoMdArrowRoundBack } from "react-icons/io";
import { useParams, useNavigate } from "react-router-dom";
import { IDetail } from "@/types";

export default function Manga() {
  const params = useParams();
  const navigate = useNavigate();

  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<IDetail>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    axios
      .get(`${url}/manga/${params.url}`, {
        headers: { 
          authorization: `tma ${WebApp.initData}`,
          "ngrok-skip-browser-warning": "true",
        },
      })
      .then((res) => setData(res.data))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, [url, params]);

  if (loading) {
    return (
      <Flex justifyContent={"center"} alignItems={"center"} height={"100vh"}>
        <Spinner size={"xl"} />
      </Flex>
    );
  }

  return (
    <>
      <Header name={<IoMdArrowRoundBack />} onClick={() => navigate(-1)} />
      <Detail data={data} />
    </>
  );
}
