import { DecryptRSA } from "../RSA Encrypt";

export interface AuthTokenInterface {
  Token: string;
  RSAPublicKeyBackend: any;
  ValidateToken: any;
}

export default class Autheticathe_Token {
  #Token: string;
  #RSAPublicKeyBackend: any;
  #ValidateToken: any;
  #result_auth_Token: any;

  constructor(TokenReq: AuthTokenInterface) {
    this.#RSAPublicKeyBackend = TokenReq.RSAPublicKeyBackend;
    this.#ValidateToken = TokenReq.ValidateToken;
    this.#Token = TokenReq.Token;
  }

  async #AuthToken(): Promise<boolean> {
    try {
      this.#result_auth_Token = await this.#ValidateToken({
        publicKey: import.meta.env.VITE_PublicKey_RSA,
        dataChiper: this.#Token,
      });
      this.#result_auth_Token = this.#result_auth_Token.data.data.validateToken;
      return true;
    } catch (err) {
      console.log(err);
      alert(err);
      return false;
    }
  }

  async GetSession(): Promise<any> {
    const AuthToken = await this.#AuthToken();
    if (!AuthToken) return { isAuth: false };
    if (this.#result_auth_Token.result === "valid token failed")
      return { isAuth: false };
    const data = await DecryptRSA(
      this.#RSAPublicKeyBackend,
      this.#result_auth_Token.matricule
    );
    //if (!isError) return { matricule: data, isAuth: true };
    return { matricule: data, isAuth: true };
  }
}
