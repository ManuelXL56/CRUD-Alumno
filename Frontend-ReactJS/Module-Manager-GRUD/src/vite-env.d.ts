/// <reference types="vite/client" />

declare global {
  export interface Window {
    Go: any;
    EncryptRSA: (PrivateKey: string, PublicKey: string, Data: string) => string;
    DecryptRSA: (PrivateKey: string, PublicKey: string, ciphertext: string) => string;
    EncryptAES: (Key: string, Nonce: string, Data: string) => string;
    DecryptAES: (Key: string, Nonce: string, Ciphertext: string) => string;
  }
}

declare module "*.jpg" {
  const path: string;
  export default path;
}

declare module "*.png" {
  const path: string;
  export default path;
}

export {};
