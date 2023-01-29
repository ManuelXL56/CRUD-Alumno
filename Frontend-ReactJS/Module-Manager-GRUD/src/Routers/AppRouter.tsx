import {
  BrowserRouter,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import PublicRouters from "./PublicRouters";
import PrivateRouters from "./PrivateRouters";
import { useState } from "react";
import Login from "../Pages/Login/Login";
import Navbar from "../Components/Navbar/Navbar";
import Dashboard from "../Pages/Dashboard/Dashboard";
import NotFoundPages from "../Pages/NotFoundPages/NotFoundPages";
import MenuMain from "../Components/Navbar/MenuMain";
import MenuSession from "../Components/Navbar/MenuSession";
import useAuth from "../Auth/useAuth";
import AdminUsers from "../Pages/Dashboard/Admin Users/AdminUsers";

import Loading from "../Components/Loading/Loading";
import { useGetRSAPublicKeysBackendQuery } from "../features/AuthQuerys/AuthSession";

export default function AppRouters() {
  const [MenuActivateRef, setMenuActivateRef] = useState();
  const [MenuActivateRef_Session, setMenuActivateRef_Session] = useState();

  const Session: boolean = useAuth() !== null;

  const { data, isError, error, isLoading } = useGetRSAPublicKeysBackendQuery<any>(null);
  
  if (isError) return <label>Error: {error} </label>;
  if (isLoading) return <Loading open={isLoading} />;
  return (
    <BrowserRouter>
      <Navbar
        MenuActivateRef={MenuActivateRef}
        MenuActivateRef_Session={MenuActivateRef_Session}
      />
      {Session ? (
        <>
          <MenuMain setMenuActivateRef={setMenuActivateRef} />
          <MenuSession setMenuActivateRef={setMenuActivateRef_Session} />
        </>
      ) : null}
      <main>
        <Routes>
          <Route
            path="/"
            element={<Navigate replace to="Login" />} 
          />
          <Route
            path="Login"
            element={<PublicRouters element={<Login />} />}
          />
          <Route
            path="Dashboard"
            element={<PrivateRouters element={<Dashboard />} />}
          >
            <Route
              path="AdminUsers"
              element={<PrivateRouters element={<AdminUsers RSAPublicKeyBackend={data.data.rSAPublicKeyBackend.publicKey} />} />}
            />
          </Route>
          <Route path="*" element={<NotFoundPages />} />
        </Routes>
      </main>
    </BrowserRouter>
  );
}