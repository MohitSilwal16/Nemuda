//
//  Generated code. Do not modify.
//  source: auth.proto
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

import 'auth.pb.dart' as $0;

export 'auth.pb.dart';

@$pb.GrpcServiceName('proto.AuthService')
class AuthServiceClient extends $grpc.Client {
  static final _$login = $grpc.ClientMethod<$0.AuthRequest, $0.AuthResponse>(
      '/proto.AuthService/Login',
      ($0.AuthRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.AuthResponse.fromBuffer(value));
  static final _$logout = $grpc.ClientMethod<$0.LogoutRequest, $0.LogoutResponse>(
      '/proto.AuthService/Logout',
      ($0.LogoutRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.LogoutResponse.fromBuffer(value));
  static final _$register = $grpc.ClientMethod<$0.AuthRequest, $0.AuthResponse>(
      '/proto.AuthService/Register',
      ($0.AuthRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.AuthResponse.fromBuffer(value));
  static final _$verifySessionToken = $grpc.ClientMethod<$0.ValidationRequest, $0.ValidationResponse>(
      '/proto.AuthService/VerifySessionToken',
      ($0.ValidationRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ValidationResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.AuthResponse> login($0.AuthRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$login, request, options: options);
  }

  $grpc.ResponseFuture<$0.LogoutResponse> logout($0.LogoutRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$logout, request, options: options);
  }

  $grpc.ResponseFuture<$0.AuthResponse> register($0.AuthRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$register, request, options: options);
  }

  $grpc.ResponseFuture<$0.ValidationResponse> verifySessionToken($0.ValidationRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$verifySessionToken, request, options: options);
  }
}

@$pb.GrpcServiceName('proto.AuthService')
abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'proto.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.AuthRequest, $0.AuthResponse>(
        'Login',
        login_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.AuthRequest.fromBuffer(value),
        ($0.AuthResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.LogoutRequest, $0.LogoutResponse>(
        'Logout',
        logout_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LogoutRequest.fromBuffer(value),
        ($0.LogoutResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.AuthRequest, $0.AuthResponse>(
        'Register',
        register_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.AuthRequest.fromBuffer(value),
        ($0.AuthResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ValidationRequest, $0.ValidationResponse>(
        'VerifySessionToken',
        verifySessionToken_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ValidationRequest.fromBuffer(value),
        ($0.ValidationResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.AuthResponse> login_Pre($grpc.ServiceCall call, $async.Future<$0.AuthRequest> request) async {
    return login(call, await request);
  }

  $async.Future<$0.LogoutResponse> logout_Pre($grpc.ServiceCall call, $async.Future<$0.LogoutRequest> request) async {
    return logout(call, await request);
  }

  $async.Future<$0.AuthResponse> register_Pre($grpc.ServiceCall call, $async.Future<$0.AuthRequest> request) async {
    return register(call, await request);
  }

  $async.Future<$0.ValidationResponse> verifySessionToken_Pre($grpc.ServiceCall call, $async.Future<$0.ValidationRequest> request) async {
    return verifySessionToken(call, await request);
  }

  $async.Future<$0.AuthResponse> login($grpc.ServiceCall call, $0.AuthRequest request);
  $async.Future<$0.LogoutResponse> logout($grpc.ServiceCall call, $0.LogoutRequest request);
  $async.Future<$0.AuthResponse> register($grpc.ServiceCall call, $0.AuthRequest request);
  $async.Future<$0.ValidationResponse> verifySessionToken($grpc.ServiceCall call, $0.ValidationRequest request);
}
