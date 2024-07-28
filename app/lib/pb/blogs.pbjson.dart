//
//  Generated code. Do not modify.
//  source: blogs.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use commentDescriptor instead')
const Comment$json = {
  '1': 'Comment',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
  ],
};

/// Descriptor for `Comment`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List commentDescriptor = $convert.base64Decode(
    'CgdDb21tZW50EhoKCHVzZXJuYW1lGAEgASgJUgh1c2VybmFtZRIgCgtkZXNjcmlwdGlvbhgCIA'
    'EoCVILZGVzY3JpcHRpb24=');

@$core.Deprecated('Use blogDescriptor instead')
const Blog$json = {
  '1': 'Blog',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
    {'1': 'tag', '3': 3, '4': 1, '5': 9, '10': 'tag'},
    {'1': 'description', '3': 4, '4': 1, '5': 9, '10': 'description'},
    {'1': 'likes', '3': 5, '4': 1, '5': 13, '10': 'likes'},
    {'1': 'likedUsername', '3': 6, '4': 3, '5': 9, '10': 'likedUsername'},
    {'1': 'comments', '3': 7, '4': 3, '5': 11, '6': '.proto.Comment', '10': 'comments'},
    {'1': 'imagePath', '3': 8, '4': 1, '5': 9, '10': 'imagePath'},
  ],
};

/// Descriptor for `Blog`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List blogDescriptor = $convert.base64Decode(
    'CgRCbG9nEhoKCHVzZXJuYW1lGAEgASgJUgh1c2VybmFtZRIUCgV0aXRsZRgCIAEoCVIFdGl0bG'
    'USEAoDdGFnGAMgASgJUgN0YWcSIAoLZGVzY3JpcHRpb24YBCABKAlSC2Rlc2NyaXB0aW9uEhQK'
    'BWxpa2VzGAUgASgNUgVsaWtlcxIkCg1saWtlZFVzZXJuYW1lGAYgAygJUg1saWtlZFVzZXJuYW'
    '1lEioKCGNvbW1lbnRzGAcgAygLMg4ucHJvdG8uQ29tbWVudFIIY29tbWVudHMSHAoJaW1hZ2VQ'
    'YXRoGAggASgJUglpbWFnZVBhdGg=');

@$core.Deprecated('Use getBlogsRequestDescriptor instead')
const GetBlogsRequest$json = {
  '1': 'GetBlogsRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'offset', '3': 2, '4': 1, '5': 5, '10': 'offset'},
    {'1': 'tag', '3': 3, '4': 1, '5': 9, '10': 'tag'},
  ],
};

/// Descriptor for `GetBlogsRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getBlogsRequestDescriptor = $convert.base64Decode(
    'Cg9HZXRCbG9nc1JlcXVlc3QSIgoMc2Vzc2lvblRva2VuGAEgASgJUgxzZXNzaW9uVG9rZW4SFg'
    'oGb2Zmc2V0GAIgASgFUgZvZmZzZXQSEAoDdGFnGAMgASgJUgN0YWc=');

@$core.Deprecated('Use getBlogsResponseDescriptor instead')
const GetBlogsResponse$json = {
  '1': 'GetBlogsResponse',
  '2': [
    {'1': 'blogs', '3': 1, '4': 3, '5': 11, '6': '.proto.Blog', '10': 'blogs'},
    {'1': 'nextOffset', '3': 2, '4': 1, '5': 5, '10': 'nextOffset'},
  ],
};

/// Descriptor for `GetBlogsResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getBlogsResponseDescriptor = $convert.base64Decode(
    'ChBHZXRCbG9nc1Jlc3BvbnNlEiEKBWJsb2dzGAEgAygLMgsucHJvdG8uQmxvZ1IFYmxvZ3MSHg'
    'oKbmV4dE9mZnNldBgCIAEoBVIKbmV4dE9mZnNldA==');

@$core.Deprecated('Use postBlogRequestDescriptor instead')
const PostBlogRequest$json = {
  '1': 'PostBlogRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
    {'1': 'tag', '3': 3, '4': 1, '5': 9, '10': 'tag'},
    {'1': 'description', '3': 4, '4': 1, '5': 9, '10': 'description'},
    {'1': 'imageData', '3': 5, '4': 1, '5': 12, '10': 'imageData'},
  ],
};

/// Descriptor for `PostBlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List postBlogRequestDescriptor = $convert.base64Decode(
    'Cg9Qb3N0QmxvZ1JlcXVlc3QSIgoMc2Vzc2lvblRva2VuGAEgASgJUgxzZXNzaW9uVG9rZW4SFA'
    'oFdGl0bGUYAiABKAlSBXRpdGxlEhAKA3RhZxgDIAEoCVIDdGFnEiAKC2Rlc2NyaXB0aW9uGAQg'
    'ASgJUgtkZXNjcmlwdGlvbhIcCglpbWFnZURhdGEYBSABKAxSCWltYWdlRGF0YQ==');

@$core.Deprecated('Use postBlogResponseDescriptor instead')
const PostBlogResponse$json = {
  '1': 'PostBlogResponse',
  '2': [
    {'1': 'isBlogAdded', '3': 1, '4': 1, '5': 8, '10': 'isBlogAdded'},
  ],
};

/// Descriptor for `PostBlogResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List postBlogResponseDescriptor = $convert.base64Decode(
    'ChBQb3N0QmxvZ1Jlc3BvbnNlEiAKC2lzQmxvZ0FkZGVkGAEgASgIUgtpc0Jsb2dBZGRlZA==');

@$core.Deprecated('Use updateBlogRequestDescriptor instead')
const UpdateBlogRequest$json = {
  '1': 'UpdateBlogRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'oldTitle', '3': 2, '4': 1, '5': 9, '10': 'oldTitle'},
    {'1': 'newTitle', '3': 3, '4': 1, '5': 9, '10': 'newTitle'},
    {'1': 'newTag', '3': 4, '4': 1, '5': 9, '10': 'newTag'},
    {'1': 'newDescription', '3': 5, '4': 1, '5': 9, '10': 'newDescription'},
    {'1': 'newImageData', '3': 6, '4': 1, '5': 12, '10': 'newImageData'},
  ],
};

/// Descriptor for `UpdateBlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateBlogRequestDescriptor = $convert.base64Decode(
    'ChFVcGRhdGVCbG9nUmVxdWVzdBIiCgxzZXNzaW9uVG9rZW4YASABKAlSDHNlc3Npb25Ub2tlbh'
    'IaCghvbGRUaXRsZRgCIAEoCVIIb2xkVGl0bGUSGgoIbmV3VGl0bGUYAyABKAlSCG5ld1RpdGxl'
    'EhYKBm5ld1RhZxgEIAEoCVIGbmV3VGFnEiYKDm5ld0Rlc2NyaXB0aW9uGAUgASgJUg5uZXdEZX'
    'NjcmlwdGlvbhIiCgxuZXdJbWFnZURhdGEYBiABKAxSDG5ld0ltYWdlRGF0YQ==');

@$core.Deprecated('Use updateBlogResponseDescriptor instead')
const UpdateBlogResponse$json = {
  '1': 'UpdateBlogResponse',
  '2': [
    {'1': 'isBlogUpdated', '3': 1, '4': 1, '5': 8, '10': 'isBlogUpdated'},
  ],
};

/// Descriptor for `UpdateBlogResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateBlogResponseDescriptor = $convert.base64Decode(
    'ChJVcGRhdGVCbG9nUmVzcG9uc2USJAoNaXNCbG9nVXBkYXRlZBgBIAEoCFINaXNCbG9nVXBkYX'
    'RlZA==');

@$core.Deprecated('Use deleteBlogRequestDescriptor instead')
const DeleteBlogRequest$json = {
  '1': 'DeleteBlogRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
  ],
};

/// Descriptor for `DeleteBlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteBlogRequestDescriptor = $convert.base64Decode(
    'ChFEZWxldGVCbG9nUmVxdWVzdBIiCgxzZXNzaW9uVG9rZW4YASABKAlSDHNlc3Npb25Ub2tlbh'
    'IUCgV0aXRsZRgCIAEoCVIFdGl0bGU=');

@$core.Deprecated('Use deleteBlogResponseDescriptor instead')
const DeleteBlogResponse$json = {
  '1': 'DeleteBlogResponse',
  '2': [
    {'1': 'isBlogDeleted', '3': 1, '4': 1, '5': 8, '10': 'isBlogDeleted'},
  ],
};

/// Descriptor for `DeleteBlogResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteBlogResponseDescriptor = $convert.base64Decode(
    'ChJEZWxldGVCbG9nUmVzcG9uc2USJAoNaXNCbG9nRGVsZXRlZBgBIAEoCFINaXNCbG9nRGVsZX'
    'RlZA==');

@$core.Deprecated('Use getBlogRequestDescriptor instead')
const GetBlogRequest$json = {
  '1': 'GetBlogRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
  ],
};

/// Descriptor for `GetBlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getBlogRequestDescriptor = $convert.base64Decode(
    'Cg5HZXRCbG9nUmVxdWVzdBIiCgxzZXNzaW9uVG9rZW4YASABKAlSDHNlc3Npb25Ub2tlbhIUCg'
    'V0aXRsZRgCIAEoCVIFdGl0bGU=');

@$core.Deprecated('Use getBlogResponseDescriptor instead')
const GetBlogResponse$json = {
  '1': 'GetBlogResponse',
  '2': [
    {'1': 'blog', '3': 1, '4': 1, '5': 11, '6': '.proto.Blog', '10': 'blog'},
    {'1': 'isBlogLiked', '3': 2, '4': 1, '5': 8, '10': 'isBlogLiked'},
    {'1': 'isBlogUpdatableDeletable', '3': 3, '4': 1, '5': 8, '10': 'isBlogUpdatableDeletable'},
  ],
};

/// Descriptor for `GetBlogResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getBlogResponseDescriptor = $convert.base64Decode(
    'Cg9HZXRCbG9nUmVzcG9uc2USHwoEYmxvZxgBIAEoCzILLnByb3RvLkJsb2dSBGJsb2cSIAoLaX'
    'NCbG9nTGlrZWQYAiABKAhSC2lzQmxvZ0xpa2VkEjoKGGlzQmxvZ1VwZGF0YWJsZURlbGV0YWJs'
    'ZRgDIAEoCFIYaXNCbG9nVXBkYXRhYmxlRGVsZXRhYmxl');

@$core.Deprecated('Use likeBlogRequestDescriptor instead')
const LikeBlogRequest$json = {
  '1': 'LikeBlogRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
  ],
};

/// Descriptor for `LikeBlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List likeBlogRequestDescriptor = $convert.base64Decode(
    'Cg9MaWtlQmxvZ1JlcXVlc3QSIgoMc2Vzc2lvblRva2VuGAEgASgJUgxzZXNzaW9uVG9rZW4SFA'
    'oFdGl0bGUYAiABKAlSBXRpdGxl');

@$core.Deprecated('Use likeBlogResponseDescriptor instead')
const LikeBlogResponse$json = {
  '1': 'LikeBlogResponse',
  '2': [
    {'1': 'isBlogLiked', '3': 1, '4': 1, '5': 8, '10': 'isBlogLiked'},
  ],
};

/// Descriptor for `LikeBlogResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List likeBlogResponseDescriptor = $convert.base64Decode(
    'ChBMaWtlQmxvZ1Jlc3BvbnNlEiAKC2lzQmxvZ0xpa2VkGAEgASgIUgtpc0Jsb2dMaWtlZA==');

@$core.Deprecated('Use dislikeBlogRequestDescriptor instead')
const DislikeBlogRequest$json = {
  '1': 'DislikeBlogRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
  ],
};

/// Descriptor for `DislikeBlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List dislikeBlogRequestDescriptor = $convert.base64Decode(
    'ChJEaXNsaWtlQmxvZ1JlcXVlc3QSIgoMc2Vzc2lvblRva2VuGAEgASgJUgxzZXNzaW9uVG9rZW'
    '4SFAoFdGl0bGUYAiABKAlSBXRpdGxl');

@$core.Deprecated('Use dislikeBlogResponseDescriptor instead')
const DislikeBlogResponse$json = {
  '1': 'DislikeBlogResponse',
  '2': [
    {'1': 'isBlogDisliked', '3': 1, '4': 1, '5': 8, '10': 'isBlogDisliked'},
  ],
};

/// Descriptor for `DislikeBlogResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List dislikeBlogResponseDescriptor = $convert.base64Decode(
    'ChNEaXNsaWtlQmxvZ1Jlc3BvbnNlEiYKDmlzQmxvZ0Rpc2xpa2VkGAEgASgIUg5pc0Jsb2dEaX'
    'NsaWtlZA==');

@$core.Deprecated('Use addCommentRequestDescriptor instead')
const AddCommentRequest$json = {
  '1': 'AddCommentRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
    {'1': 'commentDescription', '3': 3, '4': 1, '5': 9, '10': 'commentDescription'},
  ],
};

/// Descriptor for `AddCommentRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addCommentRequestDescriptor = $convert.base64Decode(
    'ChFBZGRDb21tZW50UmVxdWVzdBIiCgxzZXNzaW9uVG9rZW4YASABKAlSDHNlc3Npb25Ub2tlbh'
    'IUCgV0aXRsZRgCIAEoCVIFdGl0bGUSLgoSY29tbWVudERlc2NyaXB0aW9uGAMgASgJUhJjb21t'
    'ZW50RGVzY3JpcHRpb24=');

@$core.Deprecated('Use addCommentResponseDescriptor instead')
const AddCommentResponse$json = {
  '1': 'AddCommentResponse',
  '2': [
    {'1': 'isCommentAdded', '3': 1, '4': 1, '5': 8, '10': 'isCommentAdded'},
  ],
};

/// Descriptor for `AddCommentResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addCommentResponseDescriptor = $convert.base64Decode(
    'ChJBZGRDb21tZW50UmVzcG9uc2USJgoOaXNDb21tZW50QWRkZWQYASABKAhSDmlzQ29tbWVudE'
    'FkZGVk');

@$core.Deprecated('Use searchBlogRequestDescriptor instead')
const SearchBlogRequest$json = {
  '1': 'SearchBlogRequest',
  '2': [
    {'1': 'sessionToken', '3': 1, '4': 1, '5': 9, '10': 'sessionToken'},
    {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
  ],
};

/// Descriptor for `SearchBlogRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List searchBlogRequestDescriptor = $convert.base64Decode(
    'ChFTZWFyY2hCbG9nUmVxdWVzdBIiCgxzZXNzaW9uVG9rZW4YASABKAlSDHNlc3Npb25Ub2tlbh'
    'IUCgV0aXRsZRgCIAEoCVIFdGl0bGU=');

@$core.Deprecated('Use searchBlogResponseDescriptor instead')
const SearchBlogResponse$json = {
  '1': 'SearchBlogResponse',
  '2': [
    {'1': 'doesBlogExists', '3': 1, '4': 1, '5': 8, '10': 'doesBlogExists'},
  ],
};

/// Descriptor for `SearchBlogResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List searchBlogResponseDescriptor = $convert.base64Decode(
    'ChJTZWFyY2hCbG9nUmVzcG9uc2USJgoOZG9lc0Jsb2dFeGlzdHMYASABKAhSDmRvZXNCbG9nRX'
    'hpc3Rz');

