// @ts-ignore
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const Authenticate_API_Slide = createApi({
  reducerPath: "Authenticate_API_Slide",
  baseQuery: fetchBaseQuery({ baseUrl: "https://localhost/graphql/auth" }),
  tagTypes: ["Post"],
  endpoints: (builder: any) => ({
    getRSAPublicKeysBackend: builder.query({
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
    ValidToken: builder.mutation({
      query: (input: any) => ({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          query: `mutation ValidateToken($input: DataChiper!) {
            validateToken(input: $input) {
              matricule
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

export const Authenticate_API_Slide_CRUD = createApi({
  reducerPath: "Authenticate_API_Slide_CRUD",
  baseQuery: fetchBaseQuery({ baseUrl: "https://localhost/graphql/Manager-GRUD" }),
  tagTypes: ["Post"],
  endpoints: (builder: any) => ({
    SearchUser: builder.mutation({
      query: (input: any) => ({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          'Authorization': `Bearer ${localStorage.getItem('Session')}`
        },
        body: JSON.stringify({
          query: `mutation SearchUser($input: DataChiper!) {
            searchUser(input: $input) {
              matricule
              full_name
              mail
              direction
              phone
              role
              publicKey
            }
          }`,
          variables: { input },
        }),
        invalidatesTags: ["GetRSAPublicKeysBackend"],
      }),
    }),
    RegisterUser: builder.mutation({
      query: (input: any) => ({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          'Authorization': `Bearer ${localStorage.getItem('Session')}`
        },
        body: JSON.stringify({
          query: `mutation RegisterUser($input: DataUser_Input!) {
            registerUser(input: $input) {
              result
            }
          }`,
          variables: { input },
        }),
        invalidatesTags: ["GetRSAPublicKeysBackend"],
      }),
    }),
    UpdateUser: builder.mutation({
      query: (input: any) => ({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          'Authorization': `Bearer ${localStorage.getItem('Session')}`
        },
        body: JSON.stringify({
          query: `mutation UpdateUser($input: DataUser_Update_Input!) {
            updateUser(input: $input) {
              result
            }
          }`,
          variables: { input },
        }),
        invalidatesTags: ["GetRSAPublicKeysBackend"],
      }),
    }),
    DeletUser: builder.mutation({
      query: (input: any) => ({
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          'Authorization': `Bearer ${localStorage.getItem('Session')}`
        },
        body: JSON.stringify({
          query: `mutation DeleteUser($input: DataChiper!) {
            deleteUser(input: $input) {
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
  useValidTokenMutation,
} = Authenticate_API_Slide;

export const {
  useSearchUserMutation,
  useRegisterUserMutation,
  useUpdateUserMutation,
  useDeletUserMutation,
} = Authenticate_API_Slide_CRUD;
