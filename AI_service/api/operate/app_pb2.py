# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: app.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'app.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tapp.proto\x12\x0b\x61pi.operate\"\x83\x01\n\x07\x43ontent\x12\n\n\x02id\x18\x01 \x01(\x03\x12\r\n\x05title\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x13\n\x0bpicture_url\x18\x04 \x01(\t\x12\r\n\x05price\x18\x05 \x01(\x02\x12\x10\n\x08quantity\x18\x06 \x01(\r\x12\x12\n\ncategories\x18\x07 \x03(\t\"@\n\x0e\x46indContentReq\x12\r\n\x05query\x18\x01 \x01(\t\x12\x0c\n\x04page\x18\x02 \x01(\x05\x12\x11\n\tpage_size\x18\x03 \x01(\x05\"b\n\x0e\x46indContentRsp\x12\r\n\x05total\x18\x01 \x01(\x03\x12&\n\x08\x63ontents\x18\x02 \x03(\x0b\x32\x14.api.operate.Content\x12\x0c\n\x04\x63ode\x18\x03 \x01(\x05\x12\x0b\n\x03msg\x18\x04 \x01(\t\"\x1b\n\rGetContentReq\x12\n\n\x02id\x18\x01 \x03(\x03\"R\n\rGetContentRsp\x12&\n\x08\x63ontents\x18\x01 \x03(\x0b\x32\x14.api.operate.Content\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t2\x94\x01\n\x03\x41pp\x12G\n\x0b\x46indContent\x12\x1b.api.operate.FindContentReq\x1a\x1b.api.operate.FindContentRsp\x12\x44\n\nGetContent\x12\x1a.api.operate.GetContentReq\x1a\x1a.api.operate.GetContentRspb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'app_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_CONTENT']._serialized_start=27
  _globals['_CONTENT']._serialized_end=158
  _globals['_FINDCONTENTREQ']._serialized_start=160
  _globals['_FINDCONTENTREQ']._serialized_end=224
  _globals['_FINDCONTENTRSP']._serialized_start=226
  _globals['_FINDCONTENTRSP']._serialized_end=324
  _globals['_GETCONTENTREQ']._serialized_start=326
  _globals['_GETCONTENTREQ']._serialized_end=353
  _globals['_GETCONTENTRSP']._serialized_start=355
  _globals['_GETCONTENTRSP']._serialized_end=437
  _globals['_APP']._serialized_start=440
  _globals['_APP']._serialized_end=588
# @@protoc_insertion_point(module_scope)
