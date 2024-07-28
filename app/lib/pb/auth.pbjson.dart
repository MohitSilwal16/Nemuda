//
//  Generated code. Do not modify.
//  source: auth.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use authRequestDescriptor instead')
const AuthRequest$json = {
  '1': 'AuthRequest',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `AuthRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List authRequestDescriptor = $convert.base64Decode(
    'CgtBdXRoUmVxdWVzdBIaCgh1c2VybmFtZRgBIAEoCVIIdXNlcm5hbWUSGgoIcGFzc3dvcmQYAi'
    'ABKAlSCHBhc3N3b3Jk');

@$core.Deprecated('Use authResponseDescriptor instead')
const AuthResponse$json = {
  '1': 'AuthResponse',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
  ],
};

/// Descriptor for `AuthResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List authResponseDescriptor = $convert.base64Decode(
    'CgxBdXRoUmVzcG9uc2USIgoMc2Vzc2lvblRva2VuGAEgASgJUgxzZXNzaW9uVG9rZW4=');

@$core.Deprecated('Use validationRequestDescriptor instead')
const ValidationRequest$json = {
  '1': 'ValidationRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
  ],
};

/// Descriptor for `ValidationRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List validationRequestDescriptor = $convert.base64Decode(
    'ChFWYWxpZGF0aW9uUmVxdWVzdBIiCgxzZXNzaW9uVG9rZW4YASABKAlSDHNlc3Npb25Ub2tlbg'
    '==');

@$core.Deprecated('Use validationResponseDescriptor instead')
const ValidationResponse$json = {
  '1': 'ValidationResponse',
  '2': [
    {'1': 'isUserVerified', '3': 1, '4': 1, '5': 8, '10': 'isUserVerified'},
  ],
};

/// Descriptor for `ValidationResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List validationResponseDescriptor = $convert.base64Decode(
    'ChJWYWxpZGF0aW9uUmVzcG9uc2USJgoOaXNVc2VyVmVyaWZpZWQYASABKAhSDmlzVXNlclZlcm'
    'lmaWVk');

@$core.Deprecated('Use logoutRequestDescriptor instead')
const LogoutRequest$json = {
  '1': 'LogoutRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
  ],
};

/// Descriptor for `LogoutRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List logoutRequestDescriptor = $convert.base64Decode(
    'Cg1Mb2dvdXRSZXF1ZXN0EiIKDHNlc3Npb25Ub2tlbhgBIAEoCVIMc2Vzc2lvblRva2Vu');

@$core.Deprecated('Use logoutResponseDescriptor instead')
const LogoutResponse$json = {
  '1': 'LogoutResponse',
  '2': [
    {'1': 'isUserLoggedOut', '3': 1, '4': 1, '5': 8, '10': 'isUserLoggedOut'},
  ],
};

/// Descriptor for `LogoutResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List logoutResponseDescriptor = $convert.base64Decode(
    'Cg5Mb2dvdXRSZXNwb25zZRIoCg9pc1VzZXJMb2dnZWRPdXQYASABKAhSD2lzVXNlckxvZ2dlZE'
    '91dA==');

