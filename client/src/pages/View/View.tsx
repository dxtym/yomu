import ChapterService from "@/api/chapter";
import ProgressService from "@/api/progress";
import Loading from "@/components/common/Loading";

import { IChapter } from "@/types/chapter";
import { Container } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { PhotoProvider, PhotoSlider } from "react-photo-view";
import { useNavigate, useParams } from "react-router-dom";

export default function View() {
  const params = useParams();
  const navigate = useNavigate();
  const [page, setPage] = useState<number>(0);
  const [data, setData] = useState<IChapter>();
  const [loading, setLoading] = useState<boolean>(true);

  const handleProgress = async () => {
    try {
      await ProgressService.updateProgress(params.manga, params.chapter, page);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    ChapterService.getChapter(params.manga, params.chapter)
      .then((res) => setData(res))
      .catch((err) => console.error(err))
      .finally(() => setLoading(false));
    ProgressService.getProgress(params.manga, params.chapter)
      .then((res) => setPage(res))
      .catch((err) => console.error(err));
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
          loadingElement={<Loading />}
          onClose={() => {
            navigate(-1);
            handleProgress();
          }}
          onIndexChange={(index: number) => setPage(index)}
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
