import { Navigate } from "react-router-dom";
import UseAuth from "../Auth/useAuth";

interface Element {
  // @ts-ignore
  element: JSX.Element;
}

export default function PrivateRouters(props: Element) {
  const session = UseAuth();
  
  if (session === null) return <Navigate replace to="/Login/" />;
  return props.element;
}
