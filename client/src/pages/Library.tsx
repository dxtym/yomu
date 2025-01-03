import Header from "@/components/Header";
import Navbar from "@/components/Navbar";
import Empty from "@/components/Empty";

export default function Library() {
  // const url = import.meta.env.VITE_API_URL;
  // const [data, setData] = useState<Array<IManga>>([]);

  // useEffect(() => {
  //   axios
  //     .get(`${url}/library`, {
  //       headers: { authorization: `tma ${WebApp.initData}` },
  //     })
  //     .then((res) => {
  //       setData(res.data);
  //       if (!res.data || res.data.length == 0) {
  //         document.body.style.height = "100vh";
  //         document.body.style.overflow = "hidden";
  //       }
  //     })
  //     .catch((err) => console.error(err));
  // }, []);

  return (
    <>
      <Header name={"Library"} />
      <Empty />
      <Navbar navs={["Library", "Browse", "History"]} />
    </>
  );
}
