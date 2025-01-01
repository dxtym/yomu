import Detail from "@/components/Detail";
import Empty from "@/components/Empty";
import Header from "@/components/Header";
import WebApp from "@twa-dev/sdk";
import axios from "axios";
import { useEffect, useState } from "react";
import { IoMdArrowRoundBack } from "react-icons/io";
import { useParams, useNavigate } from "react-router-dom";

interface Detail {
  title: string;
  cover_image: string;
  description: string;
}

export default function Manga() {
  const params = useParams();
  const navigate = useNavigate();

  const url = import.meta.env.VITE_API_URL;
  const [data, setData] = useState<Detail>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    axios
      .get(`${url}/manga/${params.url}`, {
        headers: { authorization: `tma ${WebApp.initData}` },
      })
      .then((res) => setData(res.data))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return <Empty />;
  }

  return (
    <>
      <Header name={<IoMdArrowRoundBack />} onClick={() => navigate(-1)} />
      <Detail data={data} />
    </>
  );
}
