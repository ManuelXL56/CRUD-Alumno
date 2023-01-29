import "./index.css"

import React from "react"
import ReactDOM from "react-dom/client"
import { LoadWasm } from "./LoadWasm/LoadWasm"

import AppRouter from "./Routers/AppRouter"

import { store } from "./app/store"
import { Provider } from "react-redux"
import AuthProvider from "./Auth/AuthProvider"

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <LoadWasm>
      <Provider store={store}>
        <AuthProvider>
          <AppRouter />
        </AuthProvider>
      </Provider>
    </LoadWasm>
  </React.StrictMode>
);