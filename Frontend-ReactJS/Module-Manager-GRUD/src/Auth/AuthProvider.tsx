import React, { createContext, useState, useEffect } from "react";
import { useDispatch } from "react-redux";
import {
  close_session,
  add_matricule,
} from "../features/actions/ActionAuthSlice";
import {
  useValidTokenMutation,
  useGetRSAPublicKeysBackendQuery,
} from "../features/AuthQuerys/AuthSession";
import Autheticathe_Token, { AuthTokenInterface } from "./Autheticathe_Token";

export const AuthContext = createContext<string | null>(null);

interface Props {
  // @ts-ignore
  children?: React.JSX.Element;
}

export default function AuthProvider(elements: Props) {
  const Token = localStorage.getItem("Session") || "";
  const [contextValue, setContextValue] = useState<string | null>(
    Token !== "" ? Token : null
  );

  const [ValidateToken] = useValidTokenMutation();
  const dispatch = useDispatch();

  const { data, isError, error, isLoading } = useGetRSAPublicKeysBackendQuery<any>(null);

  const CloseSession = () => {
    setContextValue(null);
    dispatch(close_session());
  };

  useEffect(() => {
    // @ts-ignore
    //console.log(data);
    if (contextValue === null || contextValue === "" || isError) {
      CloseSession();
    } else if (!isLoading && !isError) {
      const AuntToken_Data: AuthTokenInterface = {
        Token: Token,
        RSAPublicKeyBackend: data.data.rSAPublicKeyBackend.publicKey,
        ValidateToken: ValidateToken,
      };
      new Autheticathe_Token(AuntToken_Data)
        .GetSession()
        .then((Data: any) => {
          if (Data.isAuth) {
            dispatch(add_matricule(Data.matricule));
          } else {
            CloseSession();
            window.location.reload();
          }
        })
        .catch((err: any) => {
          console.log(err);
        });
    }
  }, [contextValue, data, isError, isLoading, error]);

  return (
    <AuthContext.Provider value={contextValue}>
      {elements.children}
    </AuthContext.Provider>
  );
}
