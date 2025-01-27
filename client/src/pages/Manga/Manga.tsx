import { ApiClientHooksContext } from "@/app/App";
import Loading from "@/components/common/Loading";
import Detail from "@/pages/Manga/components/Detail";

import { IDetail } from "@/types/detail";
import { FC, ReactElement, useContext, useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

const Manga: FC = (): ReactElement => {
  const params = useParams();
  const navigate = useNavigate();
  const [details, setDetails] = useState<IDetail>();
  const [loading, setLoading] = useState<boolean>(true);
  const apiClientHooks = useContext(ApiClientHooksContext);

  const backBtn = window.Telegram.WebApp.BackButton;
  backBtn.show();
  backBtn.onClick(() => navigate(-1));

  useEffect(() => {
    const details = apiClientHooks.getManga(params.manga ?? "");
    if (details.state === "resolved") {
      setLoading(false);
      setDetails(details.value);
    } else if (details.state === "rejected") {
      console.error(details.error); // TODO: error handling
    }
  }, []);

  if (loading) {
    return <Loading />;
  }

  return <Detail data={details} />
}

export default Manga;