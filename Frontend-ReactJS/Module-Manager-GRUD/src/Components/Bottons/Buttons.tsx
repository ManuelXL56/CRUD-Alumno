import "./Buttons.css";
// @ts-ignore
import React from "react";

interface Context {
  type?: any;
  className: string;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
  content: any;
  disabled?: boolean;
}

function Buttons(props: Context) {
  return (
    <button
      type={props.type}
      className={props.className}
      onClick={props.onClick}
      disabled={props.disabled}
    >
      {props.content}
    </button>
  );
}

export default React.memo(Buttons);
