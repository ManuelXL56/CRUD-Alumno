import "./Loading.css";
// @ts-ignore
import React, { useRef, useEffect } from "react";

interface context {
    style?: any;
    open: boolean;
}

function Loading(props: context) {
    const _Dialog = useRef<any>();
    
    useEffect(() => {
        _Dialog.current.close()
        if (props.open === true) {
            _Dialog.current.showModal();
        } else if (props.open === false) {
            _Dialog.current.close();
        }
      }, [props.open]);

    return (
        <dialog id="Loading" style={props.style} ref={_Dialog}>
            <div className="Loader" />
        </dialog>
    );
}

export default Loading