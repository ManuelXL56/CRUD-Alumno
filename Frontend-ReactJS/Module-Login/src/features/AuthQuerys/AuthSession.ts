import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const Authenticate_API_Slide = createApi({
  reducerPath: "Authenticate_API_Slide",
  baseQuery: fetchBaseQuery({ baseUrl: "https://localhost/graphql/auth" }),
  tagTypes: ["Post"],
  endpoints: (build: any) => ({
    getRSAPublicKeysBackend: build.query({
      query: () => ({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          query: `query RSAPublicKeyBackend { 
            rSAPublicKeyBackend {
              publicKey
            } 
          }`,
        }),
        providesTags: ["GetRSAPublicKeysBackend"],
      }),
    }),
    login: build.mutation({
      query: (input: any) => ({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          query: `mutation Login($input: DataChiper!) {
            login(input: $input) {
              token
              result
            }
          }`,
          variables: { input },
        }),
        invalidatesTags: ["GetRSAPublicKeysBackend"],
      }),
    }),
  }),
});

export const {
  useGetRSAPublicKeysBackendQuery,
  useLoginMutation,
} = Authenticate_API_Slide;
