from concurrent.futures import ThreadPoolExecutor

import grpc
import sys
sys.path.insert(1, "/app/Users")
from Config_Users_pb2 import LoginRequest, LoginReply, Data, DataUser, DataUser_Update
from Config_Users_pb2_grpc import add_RouteGuideServicer_to_server, RouteGuide
from ConnectionDB import getConnectionDB
sys.path.insert(1, "/app/Encrypt")
from grpc_client_Encrypt import EncryptionAES, DecryptionAES
import re
import logging

class RouteGuide(RouteGuide):
    def __init__(self):
        self._Result_reply = Data()
        self.KeysAES = 'xMhFJY4hdpsNsZnF2IiCriRFPyqdOm-Q8rl9qMqDAQo='
        self.NonceAES = 'WMPGx4Ky2s6iXvvC'
        
    def Login(self, request: LoginRequest, context: grpc.aio.ServicerContext) -> LoginReply: 
        _Login_reply = LoginReply()
        if request.matricule == "" or request.password == "":
            _Login_reply.result = "false"
            return _Login_reply
        matricule = EncryptionAES(self.KeysAES, self.NonceAES, request.matricule)
        password = EncryptionAES(self.KeysAES, self.NonceAES, request.password)
        sql = "CALL Login(\"{}\", \"{}\")".format(
            matricule, 
            password
        )
        try: 
            connection = getConnectionDB()
            with connection.cursor() as cursor:
                cursor.execute(sql)
                temp = cursor.fetchone()
                if temp is not None:
                    _Login_reply.result = temp['result']
                    if _Login_reply.result == "true":
                        _Login_reply.role = DecryptionAES(self.KeysAES, self.NonceAES, temp['role'])
            connection.commit()
        except Exception as ex:
            print(ex)
        finally:
            connection.close()
        return _Login_reply

    def SearchUser(self, request: Data, context: grpc.aio.ServicerContext) -> DataUser:
        _SearchUser_reply = DataUser()
        sql = "CALL `Search_User`(\"{}\");".format(EncryptionAES(self.KeysAES, self.NonceAES, request.data))
        IsEmply = lambda data : data if data == '' else DecryptionAES(self.KeysAES, self.NonceAES, data)
        try: 
            connection = getConnectionDB()
            with connection.cursor() as cursor:
                cursor.execute(sql)
                row = cursor.fetchone()
                if row is not None:
                    _SearchUser_reply.matricule = IsEmply(row['Matricule'])
                    _SearchUser_reply.password = IsEmply(row['Password'])
                    _SearchUser_reply.full_Name = IsEmply(row['Full_Name'])
                    _SearchUser_reply.mail = IsEmply(row['Mail'])
                    _SearchUser_reply.direction = IsEmply(row['Direction']) 
                    _SearchUser_reply.phone = IsEmply(row['Phone'])
                    _SearchUser_reply.role = IsEmply(row['Role'])
            connection.commit()
        except Exception as ex:
            print(ex)
        finally:
            connection.close()
        return _SearchUser_reply

    def RegisterUser(self, request: DataUser, context:  grpc.aio.ServicerContext) -> Data:
        if len(request.matricule) < 8:
            self._Result_reply.data = "Matricule Invalid"
            return self._Result_reply
        if not(re.match("^(?=.*\W+)(?=.*[\d].*[\d])(?=.*[A-Z].*[A-Z])(?=.*[a-z].*[a-z]).{8,}$", request.password)):
            self._Result_reply.data = "Password Invalid"
            return self._Result_reply
        sql = "CALL `Register_User`(\"{}\", \"{}\", \"{}\", \"{}\", \"{}\", \"{}\", \"{}\");".format(
            EncryptionAES(self.KeysAES, self.NonceAES, request.matricule), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.password), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.full_Name), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.mail), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.direction),
            EncryptionAES(self.KeysAES, self.NonceAES, request.phone), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.role)
        )
        print(sql)
        try: 
            connection = getConnectionDB()
            with connection.cursor() as cursor:
                cursor.execute(sql)
                temp = cursor.fetchone()['result']
                if temp is not None:
                    self._Result_reply.data = temp
            connection.commit()
        except Exception as ex:
            print(ex)
        finally:
            connection.close()
        return self._Result_reply

    def UpdateUser(self, request: DataUser_Update, context:  grpc.aio.ServicerContext) -> Data:
        if len(request.new_matricule) < 8:
            self._Result_reply.data = "Matricule Invalid"
            return self._Result_reply
        if re.match("^(?=.*\W+)(?=.*[\d].*[\d])(?=.*[A-Z].*[A-Z])(?=.*[a-z].*[a-z]).{8,}$", request.password) is None:
            self._Result_reply.data = "Password Invalid"
            return self._Result_reply
        if request.old_matricule == "" or request.new_matricule == "" or request.password == "" or request.full_Name == "" or request.mail == "" or request.direction == "" or request.phone == "" or request.role == "":
            self._Result_reply.data = "Data Invalid, you must not leave empty spaces"
            return self._Result_reply
        sql = "CALL Update_User(\"{}\", \"{}\", \"{}\", \"{}\", \"{}\", \"{}\", \"{}\", \"{}\")".format(
            EncryptionAES(self.KeysAES, self.NonceAES, request.old_matricule), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.new_matricule), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.password), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.full_Name), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.mail), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.direction), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.phone), 
            EncryptionAES(self.KeysAES, self.NonceAES, request.role)
        )
        try: 
            connection = getConnectionDB()
            with connection.cursor() as cursor:
                cursor.execute(sql)
                temp = cursor.fetchone()['result']
                if temp is not None:
                    self._Result_reply.data = temp
            connection.commit()
        except Exception as ex:
            print(ex)
        finally:
            connection.close()
        return self._Result_reply

    def DeleteUser(self, request: Data, context:  grpc.aio.ServicerContext) -> Data:
        sql = "CALL Delete_User(\"{}\")".format(EncryptionAES(self.KeysAES, self.NonceAES, request.data))
        try: 
            connection = getConnectionDB()
            with connection.cursor() as cursor:
                cursor.execute(sql)
                temp = cursor.fetchone()['result']
                if temp is not None:
                    self._Result_reply.data = temp
            connection.commit()
        except Exception as ex:
            print(ex)
        finally:
            connection.close()
        return self._Result_reply

def serve_users(address: str):
    server = grpc.server(ThreadPoolExecutor())
    add_RouteGuideServicer_to_server(RouteGuide(), server)
    server.add_insecure_port(address)
    logging.info("Starting server on %s", address)
    server.start()
    server.wait_for_termination()