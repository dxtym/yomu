import Carousel from "./components/Carousel";

import { IChapter } from "@/types/chapter";
import { useNavigate, useParams } from "react-router-dom";
import { FC, ReactElement, useContext, useEffect, useState } from "react";
import { ApiClientHooksContext } from "@/app/App";

const View: FC = (): ReactElement => {
  const params = useParams();
  const navigate = useNavigate();
  const [chapter, setChapter] = useState<IChapter>();
  const apiClientHooks = useContext(ApiClientHooksContext);

  const backBtn = window.Telegram.WebApp.BackButton;
  backBtn.show();
  backBtn.onClick(() => navigate(-1));

  useEffect(() => {
    const chapter = apiClientHooks.getChapter(params.manga ?? "", params.chapter ?? "");
    if (chapter.state === "resolved") setChapter(chapter.value);
    else if (chapter.state === "rejected") console.error(chapter.error); // TODO: error handling
  }, []);

  return <Carousel page_urls={chapter?.page_urls ?? []} />;
}

export default View;