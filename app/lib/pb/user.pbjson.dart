//
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use userExistsRequestDescriptor instead')
const UserExistsRequest$json = {
  '1': 'UserExistsRequest',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
  ],
};

/// Descriptor for `UserExistsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userExistsRequestDescriptor = $convert.base64Decode(
    'ChFVc2VyRXhpc3RzUmVxdWVzdBIaCgh1c2VybmFtZRgBIAEoCVIIdXNlcm5hbWU=');

@$core.Deprecated('Use userExistsResponseDescriptor instead')
const UserExistsResponse$json = {
  '1': 'UserExistsResponse',
  '2': [
    {'1': 'doesUserExists', '3': 1, '4': 1, '5': 8, '10': 'doesUserExists'},
  ],
};

/// Descriptor for `UserExistsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userExistsResponseDescriptor = $convert.base64Decode(
    'ChJVc2VyRXhpc3RzUmVzcG9uc2USJgoOZG9lc1VzZXJFeGlzdHMYASABKAhSDmRvZXNVc2VyRX'
    'hpc3Rz');

@$core.Deprecated('Use messageDescriptor instead')
const Message$json = {
  '1': 'Message',
  '2': [
    {'1': 'sender', '3': 1, '4': 1, '5': 9, '10': 'sender'},
    {'1': 'receiver', '3': 2, '4': 1, '5': 9, '10': 'receiver'},
    {'1': 'messageContent', '3': 3, '4': 1, '5': 9, '10': 'messageContent'},
    {'1': 'status', '3': 4, '4': 1, '5': 9, '10': 'status'},
    {'1': 'dateTime', '3': 5, '4': 1, '5': 9, '10': 'dateTime'},
  ],
};

/// Descriptor for `Message`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List messageDescriptor = $convert.base64Decode(
    'CgdNZXNzYWdlEhYKBnNlbmRlchgBIAEoCVIGc2VuZGVyEhoKCHJlY2VpdmVyGAIgASgJUghyZW'
    'NlaXZlchImCg5tZXNzYWdlQ29udGVudBgDIAEoCVIObWVzc2FnZUNvbnRlbnQSFgoGc3RhdHVz'
    'GAQgASgJUgZzdGF0dXMSGgoIZGF0ZVRpbWUYBSABKAlSCGRhdGVUaW1l');

@$core.Deprecated('Use userAndLastMessageDescriptor instead')
const UserAndLastMessage$json = {
  '1': 'UserAndLastMessage',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'lastMessage', '3': 2, '4': 1, '5': 11, '6': '.proto.Message', '10': 'lastMessage'},
  ],
};

/// Descriptor for `UserAndLastMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userAndLastMessageDescriptor = $convert.base64Decode(
    'ChJVc2VyQW5kTGFzdE1lc3NhZ2USGgoIdXNlcm5hbWUYASABKAlSCHVzZXJuYW1lEjAKC2xhc3'
    'RNZXNzYWdlGAIgASgLMg4ucHJvdG8uTWVzc2FnZVILbGFzdE1lc3NhZ2U=');

@$core.Deprecated('Use searchUsersByStartingPatternRequestDescriptor instead')
const SearchUsersByStartingPatternRequest$json = {
  '1': 'SearchUsersByStartingPatternRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'searchPattern', '3': 2, '4': 1, '5': 9, '10': 'searchPattern'},
  ],
};

/// Descriptor for `SearchUsersByStartingPatternRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List searchUsersByStartingPatternRequestDescriptor = $convert.base64Decode(
    'CiNTZWFyY2hVc2Vyc0J5U3RhcnRpbmdQYXR0ZXJuUmVxdWVzdBIiCgxzZXNzaW9uVG9rZW4YAS'
    'ABKAlSDHNlc3Npb25Ub2tlbhIkCg1zZWFyY2hQYXR0ZXJuGAIgASgJUg1zZWFyY2hQYXR0ZXJu');

@$core.Deprecated('Use searchUsersByStartingPatternResponseDescriptor instead')
const SearchUsersByStartingPatternResponse$json = {
  '1': 'SearchUsersByStartingPatternResponse',
  '2': [
    {'1': 'usersAndLastMessage', '3': 1, '4': 3, '5': 11, '6': '.proto.UserAndLastMessage', '10': 'usersAndLastMessage'},
  ],
};

/// Descriptor for `SearchUsersByStartingPatternResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List searchUsersByStartingPatternResponseDescriptor = $convert.base64Decode(
    'CiRTZWFyY2hVc2Vyc0J5U3RhcnRpbmdQYXR0ZXJuUmVzcG9uc2USSwoTdXNlcnNBbmRMYXN0TW'
    'Vzc2FnZRgBIAMoCzIZLnByb3RvLlVzZXJBbmRMYXN0TWVzc2FnZVITdXNlcnNBbmRMYXN0TWVz'
    'c2FnZQ==');

@$core.Deprecated('Use getMessagesRequestDescriptor instead')
const GetMessagesRequest$json = {
  '1': 'GetMessagesRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'offset', '3': 2, '4': 1, '5': 5, '10': 'offset'},
    {'1': 'user1', '3': 3, '4': 1, '5': 9, '10': 'user1'},
  ],
};

/// Descriptor for `GetMessagesRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMessagesRequestDescriptor = $convert.base64Decode(
    'ChJHZXRNZXNzYWdlc1JlcXVlc3QSIgoMc2Vzc2lvblRva2VuGAEgASgJUgxzZXNzaW9uVG9rZW'
    '4SFgoGb2Zmc2V0GAIgASgFUgZvZmZzZXQSFAoFdXNlcjEYAyABKAlSBXVzZXIx');

@$core.Deprecated('Use getMessagesResponseDescriptor instead')
const GetMessagesResponse$json = {
  '1': 'GetMessagesResponse',
  '2': [
    {'1': 'messages', '3': 1, '4': 3, '5': 11, '6': '.proto.Message', '10': 'messages'},
    {'1': 'nextOffset', '3': 2, '4': 1, '5': 5, '10': 'nextOffset'},
  ],
};

/// Descriptor for `GetMessagesResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getMessagesResponseDescriptor = $convert.base64Decode(
    'ChNHZXRNZXNzYWdlc1Jlc3BvbnNlEioKCG1lc3NhZ2VzGAEgAygLMg4ucHJvdG8uTWVzc2FnZV'
    'IIbWVzc2FnZXMSHgoKbmV4dE9mZnNldBgCIAEoBVIKbmV4dE9mZnNldA==');

@$core.Deprecated('Use userDescriptor instead')
const User$json = {
  '1': 'User',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
    {'1': 'token', '3': 3, '4': 1, '5': 9, '10': 'token'},
  ],
};

/// Descriptor for `User`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userDescriptor = $convert.base64Decode(
    'CgRVc2VyEhoKCHVzZXJuYW1lGAEgASgJUgh1c2VybmFtZRIaCghwYXNzd29yZBgCIAEoCVIIcG'
    'Fzc3dvcmQSFAoFdG9rZW4YAyABKAlSBXRva2Vu');

