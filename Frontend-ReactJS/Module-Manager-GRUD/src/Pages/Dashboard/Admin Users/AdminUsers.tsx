import InputForm from '../../../Components/InputForm/InputForm';
import Buttons from '../../../Components/Bottons/Buttons';
import React, { useState, useCallback } from 'react';
import { useSearchUserMutation, useRegisterUserMutation, useUpdateUserMutation, useDeletUserMutation} from "../../../features/AuthQuerys/AuthSession";
import { EncryptRSA, DecryptRSA } from "../../../RSA Encrypt";
import "../Tabs.css";
import Loading from "../../../Components/Loading/Loading";

interface props {
  RSAPublicKeyBackend: string;
}

interface ResultTableInt {
  matricule: string,
  full_name: string,
  mail: string,
  direction: string,
  phone: string,
  role: string,
}

function AdminUsers(props: props) {
  const [isLoadingPages, setIsLoadingPages] = useState(false);
  const [SearchUser] = useSearchUserMutation();
  const [RegisterUser] = useRegisterUserMutation();
  const [UpdateUser] = useUpdateUserMutation();
  const [DeleteUser] = useDeletUserMutation();
  let [DataUser, setDataUser] = useState<any>();
  let [ResultTable, setResultTable] = useState<ResultTableInt[]>([
    {
      matricule: '',
      full_name: '',
      mail: '',
      direction: '',
      phone: '',
      role: '',
    },
  ])

  const _SerachUser = useCallback(async () => {
    if (DataUser !== undefined) {
      const values = {
        dataChiper: DataUser.matricule,
        publicKey: import.meta.env.VITE_PublicKey_RSA,
      };
      try {
        const data: any = await SearchUser(values)
        const obj = {
          ...data.data.data.searchUser
        }
        await Object.keys(obj).forEach(async (key) => {
          if (key !== 'publicKey') {
            obj[key] = await DecryptRSA(obj.publicKey, obj[key]);
          }
        });
        setResultTable([{ ...obj }]);
      } catch (error) {
        console.log(error);
      }
    }
  }, [DataUser]);

  const _RegisterUser = useCallback(async () => {
    if (DataUser !== undefined) {
      const values = {
        ...DataUser,
        publicKey: import.meta.env.VITE_PublicKey_RSA,
      };
      try {
        const data: any = await RegisterUser(values)
        const obj = {
          ...data.data.data.registerUser
        }
        alert(obj.result);
      } catch (error) {
        console.log(error);
      }
    }
  }, [DataUser]);

  const _UpdateUser = useCallback(async () => {
    if (DataUser !== undefined) {
      const values = {
        ...DataUser,
        publicKey: import.meta.env.VITE_PublicKey_RSA,
      };
      try {
        const data: any = await UpdateUser(values)
        const obj = {
          ...data.data.data.updateUser
        }
        alert(obj.result);
      } catch (error) {
        console.log(error);
      }
    }
  }, [DataUser]);

  const _DeleteUser = useCallback(async () => {
    if (DataUser !== undefined) {
      const values = {
        dataChiper: DataUser.matricule,
        publicKey: import.meta.env.VITE_PublicKey_RSA,
      };
      try {
        const data: any = await DeleteUser(values)
        alert(data.data.data.deleteUser.result)
      } catch (error) {
        console.log(error);
      }
    }
  }, [DataUser]);

  // Execute mutations from GraophQL server Manager GRUD
  const handleSubmit = useCallback(async (e: any, func: any) => {
    e.preventDefault();
    setIsLoadingPages(true);

    //console.log(DataUser);
    if (DataUser !== undefined) {
      await Object.keys(DataUser).forEach(async (key) => {
        if(DataUser[key] !== "") {
          DataUser[key] = await EncryptRSA(props.RSAPublicKeyBackend, DataUser[key] || "");
        }
      });
      await func();
      setDataUser(DataUser = undefined)
    }

    setIsLoadingPages(false);
  }, [DataUser, props.RSAPublicKeyBackend]);

  // Recopile data from forms
  const handleChange = useCallback((event: any) => {
    event.preventDefault();
    setDataUser(DataUser = {
      ...DataUser,
      [event.target.name]: event.target.value,
    })
    //console.log(DataUser);
  }, [DataUser]);

  if (isLoadingPages) return <Loading open={isLoadingPages} />;
  return (
    <>
      <div className={"tabs__menu grid-cols-4"} >
        {/* Titulos */}
        <input
          type="radio"
          name="tabs-menu"
          className="hidden peer/tab-1"
          id="tab-1"
          defaultChecked={true}
          readOnly
          onClick={() => setDataUser(DataUser = undefined)}
        />
        <label
          className={
            "tabs__label peer-checked/tab-1:text-blue-500 peer-checked/tab-1:border-blue-500"
          }
          htmlFor="tab-1"
        >
          Buscar Usuario
        </label>

        <input
          type="radio"
          name="tabs-menu"
          className="hidden peer/tab-2"
          id="tab-2"
          readOnly
          onClick={() => setDataUser(DataUser = undefined)}
        />
        <label
          className={
            "tabs__label peer-checked/tab-2:text-blue-500 peer-checked/tab-2:border-blue-500"
          }
          htmlFor="tab-2"
        >
          Registar Usuario
        </label>

        <input
          type="radio"
          name="tabs-menu"
          className="hidden peer/tab-3"
          id="tab-3"
          readOnly
          onClick={() => setDataUser(DataUser = undefined)}
        />
        <label
          className={
            "tabs__label peer-checked/tab-3:text-blue-500 peer-checked/tab-3:border-blue-500"
          }
          htmlFor="tab-3"
        >
          Actualziar Usuario
        </label>

        <input
          type="radio"
          name="tabs-menu"
          className="hidden peer/tab-4"
          id="tab-4"
          readOnly
          onClick={() => setDataUser(DataUser = undefined)}
        />
        <label
          className={
            "tabs__label peer-checked/tab-4:text-blue-500 peer-checked/tab-4:border-blue-500"
          }
          htmlFor="tab-4"
        >
          Eliminar Usuario
        </label>

        {/* Contenido */}
        <form className={"tabs__content peer-checked/tab-1:block col-span-full"} onChange={handleChange} onSubmit={(e) => handleSubmit(e, _SerachUser)}>
          <InputForm type="search" name="matricule" placeholder='Mantricula' required={true} />
          <Buttons type="submit" className={'BottonBlue'} content={'Buscar'} disabled={DataUser === undefined} />
          <div className=' overflow-x-auto'>
            <table className="table-auto m-auto border-collapse border border-slate-300 w-11/12">
              <thead>
                <tr className='border border-slate-400 bg-zinc-300'>
                  <th>Matricula</th>
                  <th>Nombre Completo</th>
                  <th>Email</th>
                  <th>Dirección</th>
                  <th>Telefono</th>
                  <th>Rol</th>
                </tr>
              </thead>
              <tbody>
              {ResultTable.map(item => (
                <tr key={item.matricule} className='border border-slate-500 bg-zinc-100'>
                  <td>{item.matricule}</td>
                  <td>{item.full_name}</td>
                  <td>{item.mail}</td>
                  <td>{item.direction}</td>
                  <td>{item.phone}</td>
                  <td>{item.role}</td>
                </tr>
              ))}
              </tbody>
            </table>
          </div>
        </form>
        <form className={"tabs__content peer-checked/tab-2:block col-span-full"} onChange={handleChange} onSubmit={(e) => handleSubmit(e, _RegisterUser)}>
          <InputForm type="text" name="matricule" placeholder='Matricula' required={true} /><br />
          <InputForm type="password" name="password" placeholder='Contraseña' pattern='^(?=.*\W+)(?=.*[\d].*[\d])(?=.*[A-Z].*[A-Z])(?=.*[a-z].*[a-z]).{8,}$' title="La contraseña debe tener al menos un carácter especial, dos dígitos, dos letras mayúsculas y dos minúsculas y una longitud mínima de 8 caracteres" required={true} /><br />
          <InputForm type="text" name="full_name" placeholder='Nombre Completo' required={true} /><br />
          <InputForm type="email" name="mail" placeholder='Email' required={true} /><br />
          <InputForm type="text" name="direction" placeholder='Dirección' required={true} /><br />
          <InputForm type="number" name="phone" placeholder='Telefono' required={true} /><br />
          <InputForm type="search" list="RoleList" OptionsList={['Admin', 'Alumno']} name="role" placeholder='Rol' required={true} /><br />
          <Buttons type="submit" className={'BottonBlue'} content={'Registrar'} disabled={DataUser === undefined} /><br />
        </form>
        <form className={"tabs__content peer-checked/tab-3:block col-span-full"} onChange={handleChange} onSubmit={(e) => handleSubmit(e, _UpdateUser)}>
          <InputForm type="text" name="old_matricule" placeholder='Matricula Anterior' required={true}/><br />
          <InputForm type="text" name="new_matricule" placeholder='Matricula nueva' required={true} /><br />
          <InputForm type="password" name="password" placeholder='Contraseña'  pattern='^(?=.*\W+)(?=.*[\d].*[\d])(?=.*[A-Z].*[A-Z])(?=.*[a-z].*[a-z]).{8,}$' title="La contraseña debe tener al menos un carácter especial, dos dígitos, dos letras mayúsculas y dos minúsculas y una longitud mínima de 8 caracteres" required={true} /><br />
          <InputForm type="text" name="full_name" placeholder='Nombre Completo' /><br />
          <InputForm type="email" name="mail" placeholder='Email' required={true} /><br />
          <InputForm type="text" name="direction" placeholder='Dirección' required={true} /><br />
          <InputForm type="number" name="phone" placeholder='Telefono' required={true} /><br />
          <InputForm type="search" list="RoleList" OptionsList={['Admin', 'Alumno']} name="role" placeholder='Rol' required={true} /><br />
          <Buttons type="submit" className={'BottonBlue'} content={'Actualizar'} /><br />
        </form>
        <form className={"tabs__content peer-checked/tab-4:block col-span-full"} onChange={handleChange} onSubmit={(e) => handleSubmit(e, _DeleteUser)}>
          <InputForm type="search" name="matricule" placeholder='Matricula' required={true} /><br/>
          <Buttons type="submit" className={'BottonRed'} content={'Eliminar Usuario'} disabled={DataUser === undefined} /><br/>
        </form>
      </div>
    </>
  );
}


export default React.memo(AdminUsers);