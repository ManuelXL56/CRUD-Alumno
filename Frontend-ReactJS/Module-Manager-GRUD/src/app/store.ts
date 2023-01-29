// @ts-ignore
import { configureStore } from "@reduxjs/toolkit";
import ActionsMenuReducer from "../features/actions/ActionMenuSlice";
import ActionsAuthReducer from "../features/actions/ActionAuthSlice";

import { Authenticate_API_Slide, Authenticate_API_Slide_CRUD } from "../features/AuthQuerys/AuthSession"

export const store = configureStore({
  reducer: {
    [Authenticate_API_Slide.reducerPath]: Authenticate_API_Slide.reducer,
    [Authenticate_API_Slide_CRUD.reducerPath]: Authenticate_API_Slide_CRUD.reducer,
    actionsMenu: ActionsMenuReducer,
    actionsAuth: ActionsAuthReducer,
  },
  middleware: (getDefaultMiddleware: () => any[]) =>
    getDefaultMiddleware().concat(Authenticate_API_Slide.middleware, Authenticate_API_Slide_CRUD.middleware),
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;