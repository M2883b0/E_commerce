# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: ai_agent.proto
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
    'ai_agent.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0e\x61i_agent.proto\x12\x0c\x61pi.ai_agent\"7\n\x0eUserRequestReq\x12\x0f\n\x07user_id\x18\x01 \x01(\x03\x12\x14\n\x0cuser_message\x18\x02 \x01(\t\")\n\x0fUserRequestResp\x12\x16\n\x0e\x61gent_response\x18\x01 \x01(\t2W\n\x07\x41iAgent\x12L\n\x0bUserRequest\x12\x1c.api.ai_agent.UserRequestReq\x1a\x1d.api.ai_agent.UserRequestResp\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ai_agent_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_USERREQUESTREQ']._serialized_start=32
  _globals['_USERREQUESTREQ']._serialized_end=87
  _globals['_USERREQUESTRESP']._serialized_start=89
  _globals['_USERREQUESTRESP']._serialized_end=130
  _globals['_AIAGENT']._serialized_start=132
  _globals['_AIAGENT']._serialized_end=219
# @@protoc_insertion_point(module_scope)
