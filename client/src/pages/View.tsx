import { Container, Flex, Spinner } from "@chakra-ui/react";
import WebApp from "@twa-dev/sdk";
import axios from "axios";
import { useEffect, useState } from "react";
import { PhotoProvider, PhotoSlider } from "react-photo-view";
import { useNavigate, useParams } from "react-router-dom";

export default function View() {
  const params = useParams();
  const navigate = useNavigate();

  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<any>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    axios
      .get(`${url}/chapter/${params.url}/${params.id}`, {
        headers: { 
          "authorization": `tma ${WebApp.initData}`,
          "ngrok-skip-browser-warning": "true",
        },
      })
      .then((res) => setData(res.data))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return (
      <Flex justifyContent={"center"} alignItems={"center"} height={"100vh"}>
        <Spinner size={"xl"} />
      </Flex>
    );
  }

  return (
    <PhotoProvider>
      <Container>
        <PhotoSlider
          images={data?.page_urls.map((img: string, index: number) => ({
            src: img,
            key: index,
          }))}
          visible={true}
          onClose={() => navigate(-1)}
          loop={false}
        />
      </Container>
    </PhotoProvider>
  );
}
