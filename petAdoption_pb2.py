# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: petAdoption.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'petAdoption.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11petAdoption.proto\x12\x0bpetAdoption\"\\\n\x03Pet\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06gender\x18\x03 \x01(\t\x12\x0b\n\x03\x61ge\x18\x04 \x01(\x05\x12\r\n\x05\x62reed\x18\x05 \x01(\t\x12\x0f\n\x07picture\x18\x06 \x01(\x0c\"]\n\x10InsertPetRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06gender\x18\x02 \x01(\t\x12\x0b\n\x03\x61ge\x18\x03 \x01(\x05\x12\r\n\x05\x62reed\x18\x04 \x01(\t\x12\x0f\n\x07picture\x18\x05 \x01(\x0c\"$\n\x11InsertPetResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"&\n\x10SearchPetRequest\x12\x12\n\nsearchTerm\x18\x01 \x01(\t\"3\n\x11SearchPetResponse\x12\x1e\n\x04pets\x18\x01 \x03(\x0b\x32\x10.petAdoption.Pet2\xa5\x01\n\x0bPetAdoption\x12J\n\tInsertPet\x12\x1d.petAdoption.InsertPetRequest\x1a\x1e.petAdoption.InsertPetResponse\x12J\n\tSearchPet\x12\x1d.petAdoption.SearchPetRequest\x1a\x1e.petAdoption.SearchPetResponseB\x10Z\x0ego/petAdoptionb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'petAdoption_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\016go/petAdoption'
  _globals['_PET']._serialized_start=34
  _globals['_PET']._serialized_end=126
  _globals['_INSERTPETREQUEST']._serialized_start=128
  _globals['_INSERTPETREQUEST']._serialized_end=221
  _globals['_INSERTPETRESPONSE']._serialized_start=223
  _globals['_INSERTPETRESPONSE']._serialized_end=259
  _globals['_SEARCHPETREQUEST']._serialized_start=261
  _globals['_SEARCHPETREQUEST']._serialized_end=299
  _globals['_SEARCHPETRESPONSE']._serialized_start=301
  _globals['_SEARCHPETRESPONSE']._serialized_end=352
  _globals['_PETADOPTION']._serialized_start=355
  _globals['_PETADOPTION']._serialized_end=520
# @@protoc_insertion_point(module_scope)
