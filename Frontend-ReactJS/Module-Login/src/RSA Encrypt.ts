export function EncryptRSA(PublicKey: string, data: string) {
  return new Promise((resolve, reject) => {
    const DataCipher: string = `${window.EncryptRSA(
      import.meta.env.VITE_PrivateKey_RSA,
      PublicKey,
      data
    )}`;
    if (DataCipher === "Invalid Data") {
      reject(new Error(DataCipher));
    }
    resolve(DataCipher);
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
      reject(new Error(Data));
    }
    resolve(Data);
  });
}
