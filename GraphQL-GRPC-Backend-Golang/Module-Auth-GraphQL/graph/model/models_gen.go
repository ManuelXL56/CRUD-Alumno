// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DataChiper struct {
	PublicKey  string `json:"publicKey"`
	DataChiper string `json:"dataChiper"`
}

type DataToken struct {
	Matricule string `json:"matricule"`
	Result    string `json:"result"`
}

type RSAPublicKey struct {
	PublicKey string `json:"publicKey"`
}

type Session struct {
	Token  string `json:"token"`
	Result string `json:"result"`
}