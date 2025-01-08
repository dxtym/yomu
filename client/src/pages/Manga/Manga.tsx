import MangaService from "@/api/manga";

import Header from "@/components/common/Header";
import Loading from "@/components/common/Loading";
import Detail from "@/pages/Manga/components/Detail";

import { IDetail } from "@/types/detail";
import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

function Manga() {
  const params = useParams();
  const navigate = useNavigate();
  const [data, setData] = useState<IDetail>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    MangaService.getManga(params.manga)
      .then((res) => setData(res))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return <Loading />;
  }

  return (
    <>
      <Header name={"Back"} onClick={() => navigate(-1)} />
      <Detail data={data} />
    </>
  );
}

export default Manga;
