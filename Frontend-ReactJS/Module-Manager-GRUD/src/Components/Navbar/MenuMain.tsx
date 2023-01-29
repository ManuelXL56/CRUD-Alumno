import React, { useEffect, useRef, useCallback } from "react";
import { NavLink } from "react-router-dom";
import {
  faUsers,
  faBars,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import "./MenuMain.css";
import { setMenuInterface } from "./IntefaceMenu";
import { useSelector, useDispatch } from "react-redux";
import { desactive_menu_main } from "../../features/actions/ActionMenuSlice";
import type { RootState } from "../../app/store";

function MenuMain(props: setMenuInterface) {
  
  function isActiveBottons(isActive: boolean): string {
    return isActive ? "isActive" : "";
  }

  const stateIsActive = useSelector(
    (state: RootState) => state.actionsMenu.isActiveMenu_Main 
  );
  const dispatch = useDispatch();

  const useOutsideClick = (callback: () => void) => {
    const ref = useRef();
  
    useEffect(() => {
      const handleClick = (event: any) => {
        // @ts-ignore
        if (ref.current && !ref.current.contains(event.target)) {
          //window.addEventListener('click', function(e) {
          if (event.target !== document.getElementById("MenuBar") && stateIsActive === true) {
            callback();
          }
          //})
        }
      };
      document.addEventListener("click", handleClick);
      return () => {
        document.removeEventListener("click", handleClick);
      };
    }, [callback, ref]);
  
    return ref;
  };

  const handleClickOutside = () => {
    dispatch(desactive_menu_main());
  };

  const MenuActivateRef = useOutsideClick(handleClickOutside);

  const ActionBottons = useCallback(() => {
    dispatch(desactive_menu_main());
  }, [stateIsActive]);

  useEffect(() => {
    props.setMenuActivateRef(MenuActivateRef);
  }, [MenuActivateRef]);

  const MenuContent = () => {
    return (
      <>
        <NavLink to={"#"} onClick={ActionBottons}>
          <li className="Bars">
            <FontAwesomeIcon icon={faBars} />
          </li>
        </NavLink>
        <NavLink
          to={"/Dashboard/AdminUsers/"}
          className={({ isActive }): string => isActiveBottons(isActive)}
          onClick={ActionBottons}
        >
          <li className="Buttons">
            <FontAwesomeIcon icon={faUsers} />
            &nbsp; Administrar Usuarios
          </li>
        </NavLink>
      </>
    );
  };
  
  return (
    <nav>
      <div
        id="MenuBar"
        className={stateIsActive ? "Menu-Activate" : "Menu-Desactive"}
      >
        <ul>
          <MenuContent />
        </ul>
      </div>
    </nav>
  );
}

export default React.memo(MenuMain);
