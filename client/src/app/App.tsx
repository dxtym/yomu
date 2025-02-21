import "./App.css";

import Browse from "@/pages/Browse/Browse";
import Library from "@/pages/Library/Library";
import History from "@/pages/History/History";
import View from "@/pages/View/View";
import Manga from "@/pages/Manga/Manga";
import Loading from "@/components/common/Loading";

import { Suspense, createContext, useMemo } from "react";
import { LocaleProvider } from "@chakra-ui/react";
import {
  Route,
  RouterProvider,
  createBrowserRouter,
  createRoutesFromElements,
} from "react-router-dom";
import ApiClient from "@/app/api/client";
import { ApiClientHooks } from "@/hooks/client";

const apiClient = new ApiClient(); 
const ApiClientHooksContext = createContext<ApiClientHooks>(new ApiClientHooks(apiClient));

export default function App() {
  const apiClientHook = useMemo(() => new ApiClientHooks(apiClient), []);
  
  const router = createBrowserRouter(
    createRoutesFromElements(
      <Route path="/">
        <Route index element={<Library />} />
        <Route path="library" element={<Library />} />
        <Route path="browse" element={<Browse />} />
        <Route path="history" element={<History />} />
        <Route path="browse/:manga" element={<Manga />} />
        <Route path="chapter/:manga/:chapter" element={<View />} />
      </Route>
    )
  );
  
  return (
    <Suspense fallback={<Loading />}>
      <LocaleProvider locale="en-US">
        <ApiClientHooksContext.Provider value={apiClientHook}>
          <RouterProvider router={router} />
        </ApiClientHooksContext.Provider>
      </LocaleProvider>
    </Suspense>
  );
}

export { ApiClientHooksContext };