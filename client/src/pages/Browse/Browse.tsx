import Header from "@/components/common/Header";
import Navbar from "@/components/common/Navbar";
import Search from "@/pages/Browse/components/Search";
import { useEffect, useState } from "react";

const Browse = () => {
  const [search, setSearch] = useState<string>("");

  useEffect(() => {
    document.body.style.height = "100vh";
    document.body.style.overflow = "hidden";
  }, []);

  return (
    <>
      <Header name={"Browse"} setSearch={setSearch} hasSearch />
      <Search search={search} />
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
};

export default Browse;
