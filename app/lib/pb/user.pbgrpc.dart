//
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'user.pb.dart' as $2;

export 'user.pb.dart';

@$pb.GrpcServiceName('proto.UserService')
class UserServiceClient extends $grpc.Client {
  static final _$doesUserExists = $grpc.ClientMethod<$2.UserExistsRequest, $2.UserExistsResponse>(
      '/proto.UserService/DoesUserExists',
      ($2.UserExistsRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.UserExistsResponse.fromBuffer(value));
  static final _$searchUsersByStartingPattern = $grpc.ClientMethod<$2.SearchUsersByStartingPatternRequest, $2.SearchUsersByStartingPatternResponse>(
      '/proto.UserService/SearchUsersByStartingPattern',
      ($2.SearchUsersByStartingPatternRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.SearchUsersByStartingPatternResponse.fromBuffer(value));
  static final _$getMessagesWithPagination = $grpc.ClientMethod<$2.GetMessagesRequest, $2.GetMessagesResponse>(
      '/proto.UserService/GetMessagesWithPagination',
      ($2.GetMessagesRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.GetMessagesResponse.fromBuffer(value));

  UserServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$2.UserExistsResponse> doesUserExists($2.UserExistsRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$doesUserExists, request, options: options);
  }

  $grpc.ResponseFuture<$2.SearchUsersByStartingPatternResponse> searchUsersByStartingPattern($2.SearchUsersByStartingPatternRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$searchUsersByStartingPattern, request, options: options);
  }

  $grpc.ResponseFuture<$2.GetMessagesResponse> getMessagesWithPagination($2.GetMessagesRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getMessagesWithPagination, request, options: options);
  }
}

@$pb.GrpcServiceName('proto.UserService')
abstract class UserServiceBase extends $grpc.Service {
  $core.String get $name => 'proto.UserService';

  UserServiceBase() {
    $addMethod($grpc.ServiceMethod<$2.UserExistsRequest, $2.UserExistsResponse>(
        'DoesUserExists',
        doesUserExists_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.UserExistsRequest.fromBuffer(value),
        ($2.UserExistsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.SearchUsersByStartingPatternRequest, $2.SearchUsersByStartingPatternResponse>(
        'SearchUsersByStartingPattern',
        searchUsersByStartingPattern_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.SearchUsersByStartingPatternRequest.fromBuffer(value),
        ($2.SearchUsersByStartingPatternResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.GetMessagesRequest, $2.GetMessagesResponse>(
        'GetMessagesWithPagination',
        getMessagesWithPagination_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.GetMessagesRequest.fromBuffer(value),
        ($2.GetMessagesResponse value) => value.writeToBuffer()));
  }

  $async.Future<$2.UserExistsResponse> doesUserExists_Pre($grpc.ServiceCall call, $async.Future<$2.UserExistsRequest> request) async {
    return doesUserExists(call, await request);
  }

  $async.Future<$2.SearchUsersByStartingPatternResponse> searchUsersByStartingPattern_Pre($grpc.ServiceCall call, $async.Future<$2.SearchUsersByStartingPatternRequest> request) async {
    return searchUsersByStartingPattern(call, await request);
  }

  $async.Future<$2.GetMessagesResponse> getMessagesWithPagination_Pre($grpc.ServiceCall call, $async.Future<$2.GetMessagesRequest> request) async {
    return getMessagesWithPagination(call, await request);
  }

  $async.Future<$2.UserExistsResponse> doesUserExists($grpc.ServiceCall call, $2.UserExistsRequest request);
  $async.Future<$2.SearchUsersByStartingPatternResponse> searchUsersByStartingPattern($grpc.ServiceCall call, $2.SearchUsersByStartingPatternRequest request);
  $async.Future<$2.GetMessagesResponse> getMessagesWithPagination($grpc.ServiceCall call, $2.GetMessagesRequest request);
}
