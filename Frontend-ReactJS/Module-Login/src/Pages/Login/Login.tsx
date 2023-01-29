import "./Login.css";
// @ts-ignore
import React, { useState, useCallback } from "react";
import Buttons from "../../Components/Bottons/Buttons";
import Loading from "../../Components/Loading/Loading";
import InputForm from "../../Components/InputForm/InputForm";
// @ts-ignore
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
// @ts-ignore
import { faUser } from "@fortawesome/free-solid-svg-icons";
import { useGetRSAPublicKeysBackendQuery } from "../../features/AuthQuerys/AuthSession";
import { useLoginMutation } from "../../features/AuthQuerys/AuthSession";
import { EncryptRSA } from "../../RSA Encrypt";

interface InterfaceValues {
  Matricule: string;
  Password: string;
}

function Login() {
  const [Values, setValues] = useState<InterfaceValues>({
    Matricule: "",
    Password: "",
  });

  const [isLoadingPages, setIsLoading] = useState(false);
  const { data, isError, error, isLoading } = useGetRSAPublicKeysBackendQuery<any>(null);
  const [Login] = useLoginMutation();

  const Auth = async (RSAPublicKeyBackend: string) => {
    let DataCipher = await EncryptRSA(
      RSAPublicKeyBackend,
      JSON.stringify(Values)
    );
    let LoginData: any;
    try {
      const values = {
        publicKey: import.meta.env.VITE_PublicKey_RSA,
        dataChiper: DataCipher,
      };
      LoginData = await Login(values);
      LoginData = LoginData.data.data.login;
      if (LoginData.result !== "session auth failed") {
        return { ...LoginData, isSession: true };
      } else {
        return { result: LoginData.result, isSession: false };
      }
    } catch (error) {
      alert("¡Error Login! " + error);
      console.error(error);
      return { result: LoginData.result, isSession: false };
    }
  };

  const handleSubmit = useCallback(
    async (event: any) => {
      event.preventDefault();
      setIsLoading(true);
      try {
        let Session: any = await Auth(data.data.rSAPublicKeyBackend.publicKey);
        if (Session.isSession) {
          window.parent.postMessage({message: Session.token}, "*");
        } else {
          alert("¡Error Login! " + Session.result);
        }
      } catch (error) {
        console.log(error);
      }

      setIsLoading(false);
    },
    [Values, isLoadingPages, data, Auth]
  );

  const Inputs = () => {
    let Objects: string[][] = [
      ["Matricule", "text", "Nombre de Usuario", "90%", "2.5rem"],
      ["Password", "password", "Contraseña", "90%", "2.5rem"],
    ];
    return Objects.map((data: string[], Index: number) => (
      <InputForm
        key={Index}
        values={Values}
        setValues={setValues}
        name={data[0]}
        type={data[1]}
        placeholder={data[2]}
        width={data[3]}
        height={data[4]}
      />
    ));
  };

  if (isError) return <label>Error: {error} </label>;
  if (isLoading) return <Loading open={isLoading} />;
  return (
    <>
      {!isLoadingPages ? (
        <div className="box-login">
          <form className="form" onSubmit={handleSubmit} method="post">
            <FontAwesomeIcon icon={faUser} size={"4x"} />
            {Inputs()}
            <Buttons
              type={"submit"}
              className={"BottonGreen"}
              content={"iniciar"}
            />
          </form>
        </div>
      ) : (
        <Loading open={isLoadingPages} />
      )}
    </>
  );
}

export default React.memo(Login);
