// @ts-ignore
import { createSlice } from "@reduxjs/toolkit";
// @ts-ignore
//import type { PayloadAction } from "@reduxjs/toolkit";

export interface ActionState {
  isActiveMenu_Main: boolean;
  isActiveMenu_Session: boolean;
}
const initialState: ActionState = {
  isActiveMenu_Main: false,
  isActiveMenu_Session: false,
};

export const ActionsSlice = createSlice({
  name: "actions",
  initialState,
  reducers: {
    active_menu_main: (state: any) => {
      state.isActiveMenu_Main = true;
    },
    desactive_menu_main: (state: any) => {
      state.isActiveMenu_Main = false;
    },
    active_menu_session: (state: any) => {
      state.isActiveMenu_Session = true;
    },
    desactive_menu_session: (state: any) => {
      state.isActiveMenu_Session = false;
    },
  },
});

export const {
  active_menu_main,
  desactive_menu_main,
  active_menu_session,
  desactive_menu_session,
} = ActionsSlice.actions;
export default ActionsSlice.reducer;
