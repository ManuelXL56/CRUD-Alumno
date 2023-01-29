import "./DialogAlert.css";
// @ts-ignore
import React, { useRef, useEffect } from "react";
// @ts-ignore
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
// @ts-ignore
import { faTriangleExclamation } from "@fortawesome/free-solid-svg-icons";
import Buttons from "../../Components/Bottons/Buttons";

interface Context {
  open: boolean;
  setOpen: any;
  setSubmit: any;
}

function DialogAlert(props: Context) {
  const _Dialog = useRef<any>();

  useEffect(() => {
    _Dialog.current.close();
    if (props.open === true) {
      _Dialog.current.showModal();
    } else if (props.open === false) {
      _Dialog.current.close();
    }
  }, [props.open]);

  return (
    <>
      <dialog className="DialogAlert" ref={_Dialog}>
        <FontAwesomeIcon icon={faTriangleExclamation} size={"2x"} />
        <h1>Está Seguro de esta acción.</h1>
        <h2>Sera irreversible.</h2>
        <Buttons
          type={"submit"}
          className={"BottonRed"}
          content={"Yes"}
          onClick={(event) => {
            props.setSubmit(event);
            props.setOpen(false);
          }}
        />
        <Buttons
          type={"submit"}
          className={"BottonGreen"}
          content={"No"}
          onClick={(event) => {
            props.setOpen(false);
          }}
        />
      </dialog>
    </>
  );
}

export default React.memo(DialogAlert);
