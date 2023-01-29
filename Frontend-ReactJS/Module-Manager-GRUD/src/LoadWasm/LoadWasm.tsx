import './wasm_exec.js';
// @ts-ignore
import React, { useEffect } from 'react';
import Loading from '../Components/Loading/Loading'

async function loadWasm(): Promise<void> {
  const goWasm = new window.Go();
  let pathWasm: string = "https://localhost/main.wasm";
  //let pathWasm: string = "http://localhost:3001/main.wasm";
  const result = await WebAssembly.instantiateStreaming(
    fetch(pathWasm),
    goWasm.importObject
  );
  goWasm.run(result.instance);
}

export const LoadWasm: React.FC<React.PropsWithChildren<{}>> = (props) => {
  const [isLoading, setIsLoading] = React.useState(true);

  useEffect(() => {
    loadWasm().then(() => {
      setIsLoading(false);
    });
  }, []);

  if (isLoading) {
    return (
      <Loading open={isLoading} />
    );
  } else {
    return <React.Fragment>{props.children}</React.Fragment>;
  }
};
