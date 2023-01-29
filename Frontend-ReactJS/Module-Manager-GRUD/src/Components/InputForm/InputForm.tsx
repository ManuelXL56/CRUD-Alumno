import "./InputForm.css";
// @ts-ignore
import React, { useCallback } from "react";

interface Context {
  onChange?: any;
  name: string;
  type?: string;
  placeholder?: string;
  list?: string;
  readOnly?: boolean;
  minLength?: number;
  pattern?: string;
  title?: string;
  OptionsList?: string[];
  required?: boolean;
}

function InputForm(props: Context) {
  return (
    <>
      <input
        className="InputForm"
        type={props.type}
        name={props.name}
        onChange={props.onChange}
        placeholder={props.placeholder}
        list={props.OptionsList !== undefined && props.list !== undefined ? props.list : undefined}
        readOnly={props.readOnly}
        minLength={props.minLength}
        pattern={props.pattern}
        title={props.title}
        required={props.required}
      />
      {props.OptionsList !== undefined && props.list !== undefined ? (
        <datalist id={props.list}>
          {props.OptionsList.map((Option: string, index: number) => (
            <option key={index}>{Option}</option>
          ))}
        </datalist>
      ) : null}
    </>
  );
}

export default React.memo(InputForm);
