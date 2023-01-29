import "./InputForm.css";
// @ts-ignore
import React, { useCallback } from "react";

interface Context {
  setValues?: any;
  values?: any;
  name: string;
  type?: string;
  placeholder?: string;
  list?: string;
  readOnly?: boolean;
  width?: string;
  height?: string;
  minLength?: number;
  pattern?: string;
  title?: string;
  OptionsList?: string[];
}

function InputForm(props: Context) {
  
  const setData = useCallback((event: any) => {
    event.preventDefault();
    if (props.values !== undefined && props.setValues !== undefined) {
      props.setValues({
        ...props.values,
        [event.target.name]: event.target.value,
      });
    }
  }, [props.values]);

  return (
    <>
      <input
        style={{ width: props.width, height: props.height }}
        className="InputForm"
        type={props.type}
        name={props.name}
        onChange={setData}
        placeholder={props.placeholder}
        list={props.list}
        readOnly={props.readOnly}
        minLength={props.minLength}
        pattern={props.pattern}
        title={props.title}
        //required
      />
      {props.list !== undefined && props.OptionsList !== undefined ? (
        <datalist id="DataList">
          {props.OptionsList.map((Option: string, index: number) => (
            <option key={index}>{Option}</option>
          ))}
        </datalist>
      ) : null}
    </>
  );
}

export default React.memo(InputForm);
