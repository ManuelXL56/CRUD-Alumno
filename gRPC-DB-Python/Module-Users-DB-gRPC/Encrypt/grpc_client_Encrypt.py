import grpc
from Config_Encrypt_pb2_grpc import RouteGuideStub
from Config_Encrypt_pb2 import Empty, KeysAES_Data, KeysAES


def GetKeysAES():
    with grpc.insecure_channel('svc-backend-grpc:50051') as channel:
        stub = RouteGuideStub(channel)
        response = stub.GetKeysAES(Empty())
    return (response.key, response.nonce)

def EncryptionAES(Key, Nonce, Data):
    with grpc.insecure_channel('svc-backend-grpc:50051') as channel:
        stub = RouteGuideStub(channel)
        response = stub.EncryptionAES(KeysAES_Data(keys=KeysAES(key=Key, nonce=Nonce), data=Data))
    return response.data

def DecryptionAES(Key, Nonce, Data):
    with grpc.insecure_channel('svc-backend-grpc:50051') as channel:
        stub = RouteGuideStub(channel)
        response = stub.DecryptionAES(KeysAES_Data(keys=KeysAES(key=Key, nonce=Nonce), data=Data))
    return response.data