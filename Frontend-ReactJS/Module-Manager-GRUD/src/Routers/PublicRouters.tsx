// @ts-ignore
import React from "react";
// @ts-ignore
import { Navigate } from "react-router-dom";
import UseAuth from "../Auth/useAuth";

interface Element {
  // @ts-ignore
  element: JSX.Element;
}

export default function PublicRouters(props: Element) {
  const session: any = UseAuth();
  
  if (session === null) return props.element;
  return <Navigate replace to="/Dashboard/" />;
}
