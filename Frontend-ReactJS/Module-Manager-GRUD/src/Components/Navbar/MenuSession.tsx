import React, { useEffect, useRef, useCallback } from "react";
import { NavLink } from "react-router-dom";
import {
  faUser,
  faPowerOff,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import "./MenuSession.css";
import { setMenuInterface } from "./IntefaceMenu";
import { useSelector, useDispatch } from "react-redux";
import { desactive_menu_session } from "../../features/actions/ActionMenuSlice";
import { close_session } from "../../features/actions/ActionAuthSlice";
import type { RootState } from "../../app/store";

function MenuSession(props: setMenuInterface) {

  const stateIsActive = useSelector(
    (state: RootState) => state.actionsMenu.isActiveMenu_Session
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
    dispatch(desactive_menu_session());
  };

  const MenuActivateRef = useOutsideClick(handleClickOutside);
  
  useEffect(() => {
    props.setMenuActivateRef(MenuActivateRef);
  }, [MenuActivateRef]);

  const ActionBottons = useCallback(() => {
    dispatch(desactive_menu_session());
  }, [stateIsActive]);

  const stateUsername = useSelector(
    (state: RootState) => state.actionsAuth.matricule
  );

  const CloseSession = () => {
    dispatch(close_session())
    window.location.reload();
  };

  const MenuContent = () => {
    return (
      <>
        <NavLink to={"#"} onClick={ActionBottons}>
          <li>
            <FontAwesomeIcon icon={faUser} />
            &nbsp; {stateUsername}
          </li>
        </NavLink>
        <NavLink to={"#"} onClick={CloseSession}>
          <li>
            <FontAwesomeIcon icon={faPowerOff} />
          </li>
        </NavLink>
      </>
    );
  };
  
  return (
    <nav>
      <div
        id="MenuSession"
        className={
          stateIsActive ? "Menu-Session-Activate" : "Menu-Session-Desactive"
        }
      >
        <ul>
          <MenuContent />
        </ul>
      </div>
    </nav>
  );
}

export default React.memo(MenuSession);
