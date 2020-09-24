# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='auth.proto',
  package='alpd',
  syntax='proto3',
  serialized_options=b'Z\rprotos;protos',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\nauth.proto\x12\x04\x61lpd\"/\n\rpasswdRequest\x12\x0c\n\x04user\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"/\n\tauthReply\x12\x15\n\rauthenticated\x18\x02 \x01(\x08\x12\x0b\n\x03msg\x18\x03 \x01(\t2K\n\x0b\x41uthService\x12<\n\x12\x61uthWithPasswordV1\x12\x13.alpd.passwdRequest\x1a\x0f.alpd.authReply\"\x00\x42\x0fZ\rprotos;protosb\x06proto3'
)




_PASSWDREQUEST = _descriptor.Descriptor(
  name='passwdRequest',
  full_name='alpd.passwdRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='alpd.passwdRequest.user', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='password', full_name='alpd.passwdRequest.password', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=20,
  serialized_end=67,
)


_AUTHREPLY = _descriptor.Descriptor(
  name='authReply',
  full_name='alpd.authReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='authenticated', full_name='alpd.authReply.authenticated', index=0,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='alpd.authReply.msg', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=69,
  serialized_end=116,
)

DESCRIPTOR.message_types_by_name['passwdRequest'] = _PASSWDREQUEST
DESCRIPTOR.message_types_by_name['authReply'] = _AUTHREPLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

passwdRequest = _reflection.GeneratedProtocolMessageType('passwdRequest', (_message.Message,), {
  'DESCRIPTOR' : _PASSWDREQUEST,
  '__module__' : 'auth_pb2'
  # @@protoc_insertion_point(class_scope:alpd.passwdRequest)
  })
_sym_db.RegisterMessage(passwdRequest)

authReply = _reflection.GeneratedProtocolMessageType('authReply', (_message.Message,), {
  'DESCRIPTOR' : _AUTHREPLY,
  '__module__' : 'auth_pb2'
  # @@protoc_insertion_point(class_scope:alpd.authReply)
  })
_sym_db.RegisterMessage(authReply)


DESCRIPTOR._options = None

_AUTHSERVICE = _descriptor.ServiceDescriptor(
  name='AuthService',
  full_name='alpd.AuthService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=118,
  serialized_end=193,
  methods=[
  _descriptor.MethodDescriptor(
    name='authWithPasswordV1',
    full_name='alpd.AuthService.authWithPasswordV1',
    index=0,
    containing_service=None,
    input_type=_PASSWDREQUEST,
    output_type=_AUTHREPLY,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_AUTHSERVICE)

DESCRIPTOR.services_by_name['AuthService'] = _AUTHSERVICE

# @@protoc_insertion_point(module_scope)