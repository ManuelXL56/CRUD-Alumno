import logging
from grpc_server_User import serve_users
from ConnectionDB import getConnectionDB
import sys
sys.path.insert(1, "/app/mysql")
from DriverDB import configure_db
import time
#from grpc_client_Encrypt import EncryptionAES, GetKeysAES

if __name__ == "__main__":
    #KeysAES = GetKeysAES()
    #print(KeysAES)
    #_EncryptionAES = EncryptionAES("xMhFJY4hdpsNsZnF2IiCriRFPyqdOm-Q8rl9qMqDAQo=", "WMPGx4Ky2s6iXvvC", "Admin")
    #print(_EncryptionAES)
    #_EncryptionAES = EncryptionAES("xMhFJY4hdpsNsZnF2IiCriRFPyqdOm-Q8rl9qMqDAQo=", "WMPGx4Ky2s6iXvvC", "Alumno")
    #print(_EncryptionAES)

    time.sleep(120)

    connection = getConnectionDB()
    configure_db(connection)

    logging.basicConfig(level=logging.INFO)
    serve_users("[::]:50052")