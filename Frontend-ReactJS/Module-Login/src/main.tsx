import "./index.css";
import React from "react";
import ReactDOM from "react-dom/client";
import { LoadWasm } from "./LoadWasm/LoadWasm";
// @ts-ignore
import { ApiProvider } from "@reduxjs/toolkit/query/react";
import { Authenticate_API_Slide } from "./features/AuthQuerys/AuthSession";
import Login from "./Pages/Login/Login";

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <LoadWasm>
      <ApiProvider api={Authenticate_API_Slide}>
        <Login />
      </ApiProvider>
    </LoadWasm>
  </React.StrictMode>,
)
