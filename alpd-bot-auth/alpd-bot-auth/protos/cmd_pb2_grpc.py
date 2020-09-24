# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import cmd_pb2 as cmd__pb2


class CommandServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ExecuteCommandV1 = channel.unary_unary(
                '/alpd.CommandService/ExecuteCommandV1',
                request_serializer=cmd__pb2.CmdRequest.SerializeToString,
                response_deserializer=cmd__pb2.CmdReply.FromString,
                )


class CommandServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ExecuteCommandV1(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CommandServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ExecuteCommandV1': grpc.unary_unary_rpc_method_handler(
                    servicer.ExecuteCommandV1,
                    request_deserializer=cmd__pb2.CmdRequest.FromString,
                    response_serializer=cmd__pb2.CmdReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'alpd.CommandService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CommandService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ExecuteCommandV1(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/alpd.CommandService/ExecuteCommandV1',
            cmd__pb2.CmdRequest.SerializeToString,
            cmd__pb2.CmdReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
