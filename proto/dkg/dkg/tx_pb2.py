# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/dkg/dkg/tx.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from third_party.proto.gogoproto import gogo_pb2 as third__party_dot_proto_dot_gogoproto_dot_gogo__pb2
from third_party.proto.cosmos_proto import cosmos_pb2 as third__party_dot_proto_dot_cosmos__proto_dot_cosmos__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/dkg/dkg/tx.proto',
  package='dkg.dkg',
  syntax='proto3',
  serialized_options=b'Z\017dkg/x/dkg/types',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x16proto/dkg/dkg/tx.proto\x12\x07\x64kg.dkg\x1a\x19google/protobuf/any.proto\x1a&third_party/proto/gogoproto/gogo.proto\x1a+third_party/proto/cosmos_proto/cosmos.proto\"\xa6\x01\n\x13MsgRefundMsgRequest\x12\x0f\n\x07\x63reator\x18\x01 \x01(\t\x12\x41\n\x06sender\x18\x02 \x01(\x0c\x42\x31\xfa\xde\x1f-github.com/cosmos/cosmos-sdk/types.AccAddress\x12;\n\rinner_message\x18\x03 \x01(\x0b\x32\x14.google.protobuf.AnyB\x0e\xca\xb4-\nRefundable\"\x1d\n\x1bMsgRefundMsgRequestResponse2]\n\x03Msg\x12V\n\x10RefundMsgRequest\x12\x1c.dkg.dkg.MsgRefundMsgRequest\x1a$.dkg.dkg.MsgRefundMsgRequestResponseB\x11Z\x0f\x64kg/x/dkg/typesb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_any__pb2.DESCRIPTOR,third__party_dot_proto_dot_gogoproto_dot_gogo__pb2.DESCRIPTOR,third__party_dot_proto_dot_cosmos__proto_dot_cosmos__pb2.DESCRIPTOR,])




_MSGREFUNDMSGREQUEST = _descriptor.Descriptor(
  name='MsgRefundMsgRequest',
  full_name='dkg.dkg.MsgRefundMsgRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='creator', full_name='dkg.dkg.MsgRefundMsgRequest.creator', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sender', full_name='dkg.dkg.MsgRefundMsgRequest.sender', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372\336\037-github.com/cosmos/cosmos-sdk/types.AccAddress', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='inner_message', full_name='dkg.dkg.MsgRefundMsgRequest.inner_message', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\312\264-\nRefundable', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=148,
  serialized_end=314,
)


_MSGREFUNDMSGREQUESTRESPONSE = _descriptor.Descriptor(
  name='MsgRefundMsgRequestResponse',
  full_name='dkg.dkg.MsgRefundMsgRequestResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=316,
  serialized_end=345,
)

_MSGREFUNDMSGREQUEST.fields_by_name['inner_message'].message_type = google_dot_protobuf_dot_any__pb2._ANY
DESCRIPTOR.message_types_by_name['MsgRefundMsgRequest'] = _MSGREFUNDMSGREQUEST
DESCRIPTOR.message_types_by_name['MsgRefundMsgRequestResponse'] = _MSGREFUNDMSGREQUESTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MsgRefundMsgRequest = _reflection.GeneratedProtocolMessageType('MsgRefundMsgRequest', (_message.Message,), {
  'DESCRIPTOR' : _MSGREFUNDMSGREQUEST,
  '__module__' : 'proto.dkg.dkg.tx_pb2'
  # @@protoc_insertion_point(class_scope:dkg.dkg.MsgRefundMsgRequest)
  })
_sym_db.RegisterMessage(MsgRefundMsgRequest)

MsgRefundMsgRequestResponse = _reflection.GeneratedProtocolMessageType('MsgRefundMsgRequestResponse', (_message.Message,), {
  'DESCRIPTOR' : _MSGREFUNDMSGREQUESTRESPONSE,
  '__module__' : 'proto.dkg.dkg.tx_pb2'
  # @@protoc_insertion_point(class_scope:dkg.dkg.MsgRefundMsgRequestResponse)
  })
_sym_db.RegisterMessage(MsgRefundMsgRequestResponse)


DESCRIPTOR._options = None
_MSGREFUNDMSGREQUEST.fields_by_name['sender']._options = None
_MSGREFUNDMSGREQUEST.fields_by_name['inner_message']._options = None

_MSG = _descriptor.ServiceDescriptor(
  name='Msg',
  full_name='dkg.dkg.Msg',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=347,
  serialized_end=440,
  methods=[
  _descriptor.MethodDescriptor(
    name='RefundMsgRequest',
    full_name='dkg.dkg.Msg.RefundMsgRequest',
    index=0,
    containing_service=None,
    input_type=_MSGREFUNDMSGREQUEST,
    output_type=_MSGREFUNDMSGREQUESTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MSG)

DESCRIPTOR.services_by_name['Msg'] = _MSG

# @@protoc_insertion_point(module_scope)
