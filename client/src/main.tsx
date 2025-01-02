import "./index.css";
import App from "./App";
import WebApp from "@twa-dev/sdk";
import { createRoot } from "react-dom/client";
import { Provider } from "@/components/ui/provider";
import "react-photo-view/dist/react-photo-view.css";

WebApp.ready();

createRoot(document.getElementById("root")!).render(
  <Provider>
    <App />
  </Provider>,
);
