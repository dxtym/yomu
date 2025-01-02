import Header from "@/components/Header";
import Navbar from "@/components/Navbar";
import Search from "@/components/Search";
import { useEffect } from "react";

export default function Browse() {
  useEffect(() => {
    document.body.style.height = "100vh";
    document.body.style.overflow = "hidden";
  }, []);

  return (
    <>
      <Header name={"Browse"} />
      <Search />
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
}
