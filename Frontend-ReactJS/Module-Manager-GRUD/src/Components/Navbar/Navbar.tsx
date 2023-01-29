import "./Navbar.css";
import Logo from "../../assets/logo512.png";
//import React from "react";
import useAuth from "../../Auth/useAuth";
import { NavLink } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faBars,
  faUser,
} from "@fortawesome/free-solid-svg-icons";
import type { RootState } from "../../app/store";
import { useSelector, useDispatch } from "react-redux";
import {
  active_menu_main,
  desactive_menu_main,
  active_menu_session,
  desactive_menu_session,
} from "../../features/actions/ActionMenuSlice";

import { MenuInterface } from "./IntefaceMenu";

//------------------------ Menu navigation ------------------------//

export default function Navbar(props: MenuInterface) {
  function isActiveBottons(isActive: boolean): string {
    return isActive ? "isActive" : "";
  }

  //const isAuth = useAuth();
  const Session: boolean = useAuth() !== null;

  const stateIsActive_main = useSelector(
    (state: RootState) => state.actionsMenu.isActiveMenu_Main
  );
  const stateIsActive_session = useSelector(
    (state: RootState) => state.actionsMenu.isActiveMenu_Session
  );
  const dispatch = useDispatch();

  const ActionMenu_Main = (e: any) => {
    e.preventDefault();
    if (stateIsActive_main !== true) {
      dispatch(active_menu_main());
    } else {
      dispatch(desactive_menu_main());
    }
  };

  const ActionMenu_Session = (e: any) => {
    e.preventDefault();
    if (stateIsActive_session !== true) {
      dispatch(active_menu_session());
    } else {
      dispatch(desactive_menu_session());
    }
  };

  return (
    <>
      <header id="Navbar">
        <NavLink className={"Logo"} to="/Login/">
          <img src={Logo} width="40px" height="40px" alt="Logo" />
        </NavLink>
        <nav>
          {Session === true ? (
            <div className="Bars">
              <NavLink
                to="#"
                title="Menu"
                ref={props.MenuActivateRef}
                onClick={ActionMenu_Main}
              >
                <FontAwesomeIcon icon={faBars} />
              </NavLink>
            </div>
          ) : null}
        </nav>
        <nav className="User">
          <ul>
            <li>
              {Session !== true ? (
                <NavLink
                  className={({ isActive }): string =>
                    isActiveBottons(isActive)
                  }
                  to="/Login/"
                  title="Menu"
                >
                  Sing in
                </NavLink>
              ) : (
                <NavLink
                  to="#"
                  title="User Profile"
                  ref={props.MenuActivateRef_Session}
                  onClick={ActionMenu_Session}
                >
                  <FontAwesomeIcon icon={faUser} />
                </NavLink>
              )}
            </li>
          </ul>
        </nav>
      </header>
    </>
  );
}
