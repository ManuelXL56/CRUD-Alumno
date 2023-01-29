export function EncryptRSA(PublicKey: string, data: string) {
  return new Promise((resolve, reject) => {
    const DataCipher: string = `${window.EncryptRSA(
      import.meta.env.VITE_PrivateKey_RSA,
      PublicKey,
      data
    )}`;
    if (DataCipher === "Invalid Data") {
      //alert("¡Error! " + DataCipher);
      //return { isError: false };
      reject(new Error(DataCipher));
    }
    resolve(DataCipher);
    //return { isError: false, DataCipher: DataCipher };
  });
}

export function DecryptRSA(PublicKey: string, data: string) {
    return new Promise((resolve, reject) => {
      const Data: string = `${window.DecryptRSA(
          import.meta.env.VITE_PrivateKey_RSA,
          PublicKey,
          data
      )}`;
      if (Data === "Invalid Data") {
        //alert("¡Error! " + Data);
        //return { isError: false };
        reject(new Error(Data));
      }
      resolve(Data);
      //return { isError: false, data: Data };
  });
}