// @ts-ignore
import { createSlice } from "@reduxjs/toolkit";
// @ts-ignore
import type { PayloadAction } from "@reduxjs/toolkit";

export interface ActionState {
  matricule: string;
}
const initialState: ActionState = {
  matricule: "",
};

export const ActionsSlice = createSlice({
  name: "actions",
  initialState,
  reducers: {
    add_matricule: (state: any, action: PayloadAction) => {
      state.matricule = action.payload;
    },
    close_session: (state: any) => {
      state.username = "";
      localStorage.clear();
    },
  },
});

export const {
  add_matricule,
  close_session,
} = ActionsSlice.actions;
export default ActionsSlice.reducer;
