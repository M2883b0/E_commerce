# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import api.order.order_service_pb2 as order__service__pb2

GRPC_GENERATED_VERSION = '1.70.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in order_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class OrderServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.PlaceOrder = channel.unary_unary(
                '/api.order.OrderService/PlaceOrder',
                request_serializer=order__service__pb2.PlaceOrderReq.SerializeToString,
                response_deserializer=order__service__pb2.PlaceOrderResp.FromString,
                _registered_method=True)
        self.ListOrder = channel.unary_unary(
                '/api.order.OrderService/ListOrder',
                request_serializer=order__service__pb2.ListOrderReq.SerializeToString,
                response_deserializer=order__service__pb2.ListOrderResp.FromString,
                _registered_method=True)
        self.GetOrderById = channel.unary_unary(
                '/api.order.OrderService/GetOrderById',
                request_serializer=order__service__pb2.GetOrderByIdReq.SerializeToString,
                response_deserializer=order__service__pb2.GetOrderByIdResp.FromString,
                _registered_method=True)
        self.DelOrderById = channel.unary_unary(
                '/api.order.OrderService/DelOrderById',
                request_serializer=order__service__pb2.DelOrderByIdReq.SerializeToString,
                response_deserializer=order__service__pb2.DelOrderByIdResp.FromString,
                _registered_method=True)
        self.MarkOrderPaid = channel.unary_unary(
                '/api.order.OrderService/MarkOrderPaid',
                request_serializer=order__service__pb2.MarkOrderPaidReq.SerializeToString,
                response_deserializer=order__service__pb2.MarkOrderPaidResp.FromString,
                _registered_method=True)
        self.MarkOrderCancel = channel.unary_unary(
                '/api.order.OrderService/MarkOrderCancel',
                request_serializer=order__service__pb2.MarkOrderCancelReq.SerializeToString,
                response_deserializer=order__service__pb2.MarkOrderCancelResp.FromString,
                _registered_method=True)


class OrderServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def PlaceOrder(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListOrder(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOrderById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DelOrderById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def MarkOrderPaid(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def MarkOrderCancel(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OrderServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'PlaceOrder': grpc.unary_unary_rpc_method_handler(
                    servicer.PlaceOrder,
                    request_deserializer=order__service__pb2.PlaceOrderReq.FromString,
                    response_serializer=order__service__pb2.PlaceOrderResp.SerializeToString,
            ),
            'ListOrder': grpc.unary_unary_rpc_method_handler(
                    servicer.ListOrder,
                    request_deserializer=order__service__pb2.ListOrderReq.FromString,
                    response_serializer=order__service__pb2.ListOrderResp.SerializeToString,
            ),
            'GetOrderById': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOrderById,
                    request_deserializer=order__service__pb2.GetOrderByIdReq.FromString,
                    response_serializer=order__service__pb2.GetOrderByIdResp.SerializeToString,
            ),
            'DelOrderById': grpc.unary_unary_rpc_method_handler(
                    servicer.DelOrderById,
                    request_deserializer=order__service__pb2.DelOrderByIdReq.FromString,
                    response_serializer=order__service__pb2.DelOrderByIdResp.SerializeToString,
            ),
            'MarkOrderPaid': grpc.unary_unary_rpc_method_handler(
                    servicer.MarkOrderPaid,
                    request_deserializer=order__service__pb2.MarkOrderPaidReq.FromString,
                    response_serializer=order__service__pb2.MarkOrderPaidResp.SerializeToString,
            ),
            'MarkOrderCancel': grpc.unary_unary_rpc_method_handler(
                    servicer.MarkOrderCancel,
                    request_deserializer=order__service__pb2.MarkOrderCancelReq.FromString,
                    response_serializer=order__service__pb2.MarkOrderCancelResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.order.OrderService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('api.order.OrderService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class OrderService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def PlaceOrder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/api.order.OrderService/PlaceOrder',
            order__service__pb2.PlaceOrderReq.SerializeToString,
            order__service__pb2.PlaceOrderResp.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ListOrder(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/api.order.OrderService/ListOrder',
            order__service__pb2.ListOrderReq.SerializeToString,
            order__service__pb2.ListOrderResp.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetOrderById(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/api.order.OrderService/GetOrderById',
            order__service__pb2.GetOrderByIdReq.SerializeToString,
            order__service__pb2.GetOrderByIdResp.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def DelOrderById(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/api.order.OrderService/DelOrderById',
            order__service__pb2.DelOrderByIdReq.SerializeToString,
            order__service__pb2.DelOrderByIdResp.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def MarkOrderPaid(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/api.order.OrderService/MarkOrderPaid',
            order__service__pb2.MarkOrderPaidReq.SerializeToString,
            order__service__pb2.MarkOrderPaidResp.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def MarkOrderCancel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/api.order.OrderService/MarkOrderCancel',
            order__service__pb2.MarkOrderCancelReq.SerializeToString,
            order__service__pb2.MarkOrderCancelResp.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
