import ChapterService from "@/api/chapter";
import Loading from "@/components/common/Loading";

import { IChapter } from "@/types/chapter";
import { Container } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { PhotoProvider, PhotoSlider } from "react-photo-view";
import { useNavigate, useParams } from "react-router-dom";

function View() {
  const params = useParams();
  const navigate = useNavigate();
  const [page, setPage] = useState<number>(1);
  const [data, setData] = useState<IChapter>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    ChapterService.getChapter(params.manga, params.chapter)
      .then((res) => setData(res))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return <Loading />;
  }

  return (
    <PhotoProvider>
      <Container>
        <PhotoSlider
          loop={false}
          index={page}
          visible={true}
          onClose={() => navigate(-1)}
          onIndexChange={(index: number) => setPage(index + 1)}
          images={
            data?.page_urls.map((img: string, index: number) => ({
              src: img,
              key: index,
            })) ?? []
          }
        />
      </Container>
    </PhotoProvider>
  );
}

export default View;
