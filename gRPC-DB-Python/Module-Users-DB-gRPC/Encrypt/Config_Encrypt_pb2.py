# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Config_Encrypt.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14\x43onfig_Encrypt.proto\x12\x0fgrpc_encryption\"\x07\n\x05\x45mpty\"%\n\x07KeysAES\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05nonce\x18\x02 \x01(\t\"D\n\x0cKeysAES_Data\x12&\n\x04keys\x18\x01 \x01(\x0b\x32\x18.grpc_encryption.KeysAES\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\t\"\x14\n\x04\x44\x61ta\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t2\xe0\x01\n\nRouteGuide\x12@\n\nGetKeysAES\x12\x16.grpc_encryption.Empty\x1a\x18.grpc_encryption.KeysAES\"\x00\x12G\n\rEncryptionAES\x12\x1d.grpc_encryption.KeysAES_Data\x1a\x15.grpc_encryption.Data\"\x00\x12G\n\rDecryptionAES\x12\x1d.grpc_encryption.KeysAES_Data\x1a\x15.grpc_encryption.Data\"\x00\x42\x31Z/Module.com/GraphQL-Server/graph/grpc-encryptionb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'Config_Encrypt_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z/Module.com/GraphQL-Server/graph/grpc-encryption'
  _EMPTY._serialized_start=41
  _EMPTY._serialized_end=48
  _KEYSAES._serialized_start=50
  _KEYSAES._serialized_end=87
  _KEYSAES_DATA._serialized_start=89
  _KEYSAES_DATA._serialized_end=157
  _DATA._serialized_start=159
  _DATA._serialized_end=179
  _ROUTEGUIDE._serialized_start=182
  _ROUTEGUIDE._serialized_end=406
# @@protoc_insertion_point(module_scope)