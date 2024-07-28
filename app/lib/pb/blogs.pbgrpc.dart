//
//  Generated code. Do not modify.
//  source: blogs.proto
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

import 'blogs.pb.dart' as $1;

export 'blogs.pb.dart';

@$pb.GrpcServiceName('proto.BlogsService')
class BlogsServiceClient extends $grpc.Client {
  static final _$getBlogsByTagWithPagination = $grpc.ClientMethod<$1.GetBlogsRequest, $1.GetBlogsResponse>(
      '/proto.BlogsService/GetBlogsByTagWithPagination',
      ($1.GetBlogsRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.GetBlogsResponse.fromBuffer(value));
  static final _$postBlog = $grpc.ClientMethod<$1.PostBlogRequest, $1.PostBlogResponse>(
      '/proto.BlogsService/PostBlog',
      ($1.PostBlogRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.PostBlogResponse.fromBuffer(value));
  static final _$updateBlog = $grpc.ClientMethod<$1.UpdateBlogRequest, $1.UpdateBlogResponse>(
      '/proto.BlogsService/UpdateBlog',
      ($1.UpdateBlogRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.UpdateBlogResponse.fromBuffer(value));
  static final _$deleteBlog = $grpc.ClientMethod<$1.DeleteBlogRequest, $1.DeleteBlogResponse>(
      '/proto.BlogsService/DeleteBlog',
      ($1.DeleteBlogRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.DeleteBlogResponse.fromBuffer(value));
  static final _$getBlogByTitle = $grpc.ClientMethod<$1.GetBlogRequest, $1.GetBlogResponse>(
      '/proto.BlogsService/GetBlogByTitle',
      ($1.GetBlogRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.GetBlogResponse.fromBuffer(value));
  static final _$likeBlog = $grpc.ClientMethod<$1.LikeBlogRequest, $1.LikeBlogResponse>(
      '/proto.BlogsService/LikeBlog',
      ($1.LikeBlogRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.LikeBlogResponse.fromBuffer(value));
  static final _$dislikeBlog = $grpc.ClientMethod<$1.DislikeBlogRequest, $1.DislikeBlogResponse>(
      '/proto.BlogsService/DislikeBlog',
      ($1.DislikeBlogRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.DislikeBlogResponse.fromBuffer(value));
  static final _$addComment = $grpc.ClientMethod<$1.AddCommentRequest, $1.AddCommentResponse>(
      '/proto.BlogsService/AddComment',
      ($1.AddCommentRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.AddCommentResponse.fromBuffer(value));
  static final _$searchBlogByTitle = $grpc.ClientMethod<$1.SearchBlogRequest, $1.SearchBlogResponse>(
      '/proto.BlogsService/SearchBlogByTitle',
      ($1.SearchBlogRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $1.SearchBlogResponse.fromBuffer(value));

  BlogsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$1.GetBlogsResponse> getBlogsByTagWithPagination($1.GetBlogsRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getBlogsByTagWithPagination, request, options: options);
  }

  $grpc.ResponseFuture<$1.PostBlogResponse> postBlog($1.PostBlogRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$postBlog, request, options: options);
  }

  $grpc.ResponseFuture<$1.UpdateBlogResponse> updateBlog($1.UpdateBlogRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateBlog, request, options: options);
  }

  $grpc.ResponseFuture<$1.DeleteBlogResponse> deleteBlog($1.DeleteBlogRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteBlog, request, options: options);
  }

  $grpc.ResponseFuture<$1.GetBlogResponse> getBlogByTitle($1.GetBlogRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getBlogByTitle, request, options: options);
  }

  $grpc.ResponseFuture<$1.LikeBlogResponse> likeBlog($1.LikeBlogRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$likeBlog, request, options: options);
  }

  $grpc.ResponseFuture<$1.DislikeBlogResponse> dislikeBlog($1.DislikeBlogRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$dislikeBlog, request, options: options);
  }

  $grpc.ResponseFuture<$1.AddCommentResponse> addComment($1.AddCommentRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$addComment, request, options: options);
  }

  $grpc.ResponseFuture<$1.SearchBlogResponse> searchBlogByTitle($1.SearchBlogRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$searchBlogByTitle, request, options: options);
  }
}

@$pb.GrpcServiceName('proto.BlogsService')
abstract class BlogsServiceBase extends $grpc.Service {
  $core.String get $name => 'proto.BlogsService';

  BlogsServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.GetBlogsRequest, $1.GetBlogsResponse>(
        'GetBlogsByTagWithPagination',
        getBlogsByTagWithPagination_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.GetBlogsRequest.fromBuffer(value),
        ($1.GetBlogsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.PostBlogRequest, $1.PostBlogResponse>(
        'PostBlog',
        postBlog_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.PostBlogRequest.fromBuffer(value),
        ($1.PostBlogResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.UpdateBlogRequest, $1.UpdateBlogResponse>(
        'UpdateBlog',
        updateBlog_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.UpdateBlogRequest.fromBuffer(value),
        ($1.UpdateBlogResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.DeleteBlogRequest, $1.DeleteBlogResponse>(
        'DeleteBlog',
        deleteBlog_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.DeleteBlogRequest.fromBuffer(value),
        ($1.DeleteBlogResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.GetBlogRequest, $1.GetBlogResponse>(
        'GetBlogByTitle',
        getBlogByTitle_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.GetBlogRequest.fromBuffer(value),
        ($1.GetBlogResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.LikeBlogRequest, $1.LikeBlogResponse>(
        'LikeBlog',
        likeBlog_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.LikeBlogRequest.fromBuffer(value),
        ($1.LikeBlogResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.DislikeBlogRequest, $1.DislikeBlogResponse>(
        'DislikeBlog',
        dislikeBlog_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.DislikeBlogRequest.fromBuffer(value),
        ($1.DislikeBlogResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.AddCommentRequest, $1.AddCommentResponse>(
        'AddComment',
        addComment_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.AddCommentRequest.fromBuffer(value),
        ($1.AddCommentResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.SearchBlogRequest, $1.SearchBlogResponse>(
        'SearchBlogByTitle',
        searchBlogByTitle_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.SearchBlogRequest.fromBuffer(value),
        ($1.SearchBlogResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.GetBlogsResponse> getBlogsByTagWithPagination_Pre($grpc.ServiceCall call, $async.Future<$1.GetBlogsRequest> request) async {
    return getBlogsByTagWithPagination(call, await request);
  }

  $async.Future<$1.PostBlogResponse> postBlog_Pre($grpc.ServiceCall call, $async.Future<$1.PostBlogRequest> request) async {
    return postBlog(call, await request);
  }

  $async.Future<$1.UpdateBlogResponse> updateBlog_Pre($grpc.ServiceCall call, $async.Future<$1.UpdateBlogRequest> request) async {
    return updateBlog(call, await request);
  }

  $async.Future<$1.DeleteBlogResponse> deleteBlog_Pre($grpc.ServiceCall call, $async.Future<$1.DeleteBlogRequest> request) async {
    return deleteBlog(call, await request);
  }

  $async.Future<$1.GetBlogResponse> getBlogByTitle_Pre($grpc.ServiceCall call, $async.Future<$1.GetBlogRequest> request) async {
    return getBlogByTitle(call, await request);
  }

  $async.Future<$1.LikeBlogResponse> likeBlog_Pre($grpc.ServiceCall call, $async.Future<$1.LikeBlogRequest> request) async {
    return likeBlog(call, await request);
  }

  $async.Future<$1.DislikeBlogResponse> dislikeBlog_Pre($grpc.ServiceCall call, $async.Future<$1.DislikeBlogRequest> request) async {
    return dislikeBlog(call, await request);
  }

  $async.Future<$1.AddCommentResponse> addComment_Pre($grpc.ServiceCall call, $async.Future<$1.AddCommentRequest> request) async {
    return addComment(call, await request);
  }

  $async.Future<$1.SearchBlogResponse> searchBlogByTitle_Pre($grpc.ServiceCall call, $async.Future<$1.SearchBlogRequest> request) async {
    return searchBlogByTitle(call, await request);
  }

  $async.Future<$1.GetBlogsResponse> getBlogsByTagWithPagination($grpc.ServiceCall call, $1.GetBlogsRequest request);
  $async.Future<$1.PostBlogResponse> postBlog($grpc.ServiceCall call, $1.PostBlogRequest request);
  $async.Future<$1.UpdateBlogResponse> updateBlog($grpc.ServiceCall call, $1.UpdateBlogRequest request);
  $async.Future<$1.DeleteBlogResponse> deleteBlog($grpc.ServiceCall call, $1.DeleteBlogRequest request);
  $async.Future<$1.GetBlogResponse> getBlogByTitle($grpc.ServiceCall call, $1.GetBlogRequest request);
  $async.Future<$1.LikeBlogResponse> likeBlog($grpc.ServiceCall call, $1.LikeBlogRequest request);
  $async.Future<$1.DislikeBlogResponse> dislikeBlog($grpc.ServiceCall call, $1.DislikeBlogRequest request);
  $async.Future<$1.AddCommentResponse> addComment($grpc.ServiceCall call, $1.AddCommentRequest request);
  $async.Future<$1.SearchBlogResponse> searchBlogByTitle($grpc.ServiceCall call, $1.SearchBlogRequest request);
}
