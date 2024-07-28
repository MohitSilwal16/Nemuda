//
//  Generated code. Do not modify.
//  source: blogs.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Comment extends $pb.GeneratedMessage {
  factory Comment({
    $core.String? username,
    $core.String? description,
  }) {
    final $result = create();
    if (username != null) {
      $result.username = username;
    }
    if (description != null) {
      $result.description = description;
    }
    return $result;
  }
  Comment._() : super();
  factory Comment.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Comment.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Comment', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'username')
    ..aOS(2, _omitFieldNames ? '' : 'description')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Comment clone() => Comment()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Comment copyWith(void Function(Comment) updates) => super.copyWith((message) => updates(message as Comment)) as Comment;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Comment create() => Comment._();
  Comment createEmptyInstance() => create();
  static $pb.PbList<Comment> createRepeated() => $pb.PbList<Comment>();
  @$core.pragma('dart2js:noInline')
  static Comment getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Comment>(create);
  static Comment? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get username => $_getSZ(0);
  @$pb.TagNumber(1)
  set username($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUsername() => $_has(0);
  @$pb.TagNumber(1)
  void clearUsername() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get description => $_getSZ(1);
  @$pb.TagNumber(2)
  set description($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDescription() => $_has(1);
  @$pb.TagNumber(2)
  void clearDescription() => clearField(2);
}

class Blog extends $pb.GeneratedMessage {
  factory Blog({
    $core.String? username,
    $core.String? title,
    $core.String? tag,
    $core.String? description,
    $core.int? likes,
    $core.Iterable<$core.String>? likedUsername,
    $core.Iterable<Comment>? comments,
    $core.String? imagePath,
  }) {
    final $result = create();
    if (username != null) {
      $result.username = username;
    }
    if (title != null) {
      $result.title = title;
    }
    if (tag != null) {
      $result.tag = tag;
    }
    if (description != null) {
      $result.description = description;
    }
    if (likes != null) {
      $result.likes = likes;
    }
    if (likedUsername != null) {
      $result.likedUsername.addAll(likedUsername);
    }
    if (comments != null) {
      $result.comments.addAll(comments);
    }
    if (imagePath != null) {
      $result.imagePath = imagePath;
    }
    return $result;
  }
  Blog._() : super();
  factory Blog.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Blog.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Blog', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'username')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..aOS(3, _omitFieldNames ? '' : 'tag')
    ..aOS(4, _omitFieldNames ? '' : 'description')
    ..a<$core.int>(5, _omitFieldNames ? '' : 'likes', $pb.PbFieldType.OU3)
    ..pPS(6, _omitFieldNames ? '' : 'likedUsername', protoName: 'likedUsername')
    ..pc<Comment>(7, _omitFieldNames ? '' : 'comments', $pb.PbFieldType.PM, subBuilder: Comment.create)
    ..aOS(8, _omitFieldNames ? '' : 'imagePath', protoName: 'imagePath')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Blog clone() => Blog()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Blog copyWith(void Function(Blog) updates) => super.copyWith((message) => updates(message as Blog)) as Blog;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Blog create() => Blog._();
  Blog createEmptyInstance() => create();
  static $pb.PbList<Blog> createRepeated() => $pb.PbList<Blog>();
  @$core.pragma('dart2js:noInline')
  static Blog getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Blog>(create);
  static Blog? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get username => $_getSZ(0);
  @$pb.TagNumber(1)
  set username($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUsername() => $_has(0);
  @$pb.TagNumber(1)
  void clearUsername() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get tag => $_getSZ(2);
  @$pb.TagNumber(3)
  set tag($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTag() => $_has(2);
  @$pb.TagNumber(3)
  void clearTag() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get description => $_getSZ(3);
  @$pb.TagNumber(4)
  set description($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDescription() => $_has(3);
  @$pb.TagNumber(4)
  void clearDescription() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get likes => $_getIZ(4);
  @$pb.TagNumber(5)
  set likes($core.int v) { $_setUnsignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLikes() => $_has(4);
  @$pb.TagNumber(5)
  void clearLikes() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.String> get likedUsername => $_getList(5);

  @$pb.TagNumber(7)
  $core.List<Comment> get comments => $_getList(6);

  @$pb.TagNumber(8)
  $core.String get imagePath => $_getSZ(7);
  @$pb.TagNumber(8)
  set imagePath($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasImagePath() => $_has(7);
  @$pb.TagNumber(8)
  void clearImagePath() => clearField(8);
}

class GetBlogsRequest extends $pb.GeneratedMessage {
  factory GetBlogsRequest({
    $core.String? sessionToken,
    $core.int? offset,
    $core.String? tag,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (offset != null) {
      $result.offset = offset;
    }
    if (tag != null) {
      $result.tag = tag;
    }
    return $result;
  }
  GetBlogsRequest._() : super();
  factory GetBlogsRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetBlogsRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetBlogsRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..a<$core.int>(2, _omitFieldNames ? '' : 'offset', $pb.PbFieldType.O3)
    ..aOS(3, _omitFieldNames ? '' : 'tag')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetBlogsRequest clone() => GetBlogsRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetBlogsRequest copyWith(void Function(GetBlogsRequest) updates) => super.copyWith((message) => updates(message as GetBlogsRequest)) as GetBlogsRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetBlogsRequest create() => GetBlogsRequest._();
  GetBlogsRequest createEmptyInstance() => create();
  static $pb.PbList<GetBlogsRequest> createRepeated() => $pb.PbList<GetBlogsRequest>();
  @$core.pragma('dart2js:noInline')
  static GetBlogsRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetBlogsRequest>(create);
  static GetBlogsRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get offset => $_getIZ(1);
  @$pb.TagNumber(2)
  set offset($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasOffset() => $_has(1);
  @$pb.TagNumber(2)
  void clearOffset() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get tag => $_getSZ(2);
  @$pb.TagNumber(3)
  set tag($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTag() => $_has(2);
  @$pb.TagNumber(3)
  void clearTag() => clearField(3);
}

class GetBlogsResponse extends $pb.GeneratedMessage {
  factory GetBlogsResponse({
    $core.Iterable<Blog>? blogs,
    $core.int? nextOffset,
  }) {
    final $result = create();
    if (blogs != null) {
      $result.blogs.addAll(blogs);
    }
    if (nextOffset != null) {
      $result.nextOffset = nextOffset;
    }
    return $result;
  }
  GetBlogsResponse._() : super();
  factory GetBlogsResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetBlogsResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetBlogsResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..pc<Blog>(1, _omitFieldNames ? '' : 'blogs', $pb.PbFieldType.PM, subBuilder: Blog.create)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'nextOffset', $pb.PbFieldType.O3, protoName: 'nextOffset')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetBlogsResponse clone() => GetBlogsResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetBlogsResponse copyWith(void Function(GetBlogsResponse) updates) => super.copyWith((message) => updates(message as GetBlogsResponse)) as GetBlogsResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetBlogsResponse create() => GetBlogsResponse._();
  GetBlogsResponse createEmptyInstance() => create();
  static $pb.PbList<GetBlogsResponse> createRepeated() => $pb.PbList<GetBlogsResponse>();
  @$core.pragma('dart2js:noInline')
  static GetBlogsResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetBlogsResponse>(create);
  static GetBlogsResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Blog> get blogs => $_getList(0);

  @$pb.TagNumber(2)
  $core.int get nextOffset => $_getIZ(1);
  @$pb.TagNumber(2)
  set nextOffset($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasNextOffset() => $_has(1);
  @$pb.TagNumber(2)
  void clearNextOffset() => clearField(2);
}

class PostBlogRequest extends $pb.GeneratedMessage {
  factory PostBlogRequest({
    $core.String? sessionToken,
    $core.String? title,
    $core.String? tag,
    $core.String? description,
    $core.List<$core.int>? imageData,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (title != null) {
      $result.title = title;
    }
    if (tag != null) {
      $result.tag = tag;
    }
    if (description != null) {
      $result.description = description;
    }
    if (imageData != null) {
      $result.imageData = imageData;
    }
    return $result;
  }
  PostBlogRequest._() : super();
  factory PostBlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PostBlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PostBlogRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..aOS(3, _omitFieldNames ? '' : 'tag')
    ..aOS(4, _omitFieldNames ? '' : 'description')
    ..a<$core.List<$core.int>>(5, _omitFieldNames ? '' : 'imageData', $pb.PbFieldType.OY, protoName: 'imageData')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PostBlogRequest clone() => PostBlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PostBlogRequest copyWith(void Function(PostBlogRequest) updates) => super.copyWith((message) => updates(message as PostBlogRequest)) as PostBlogRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PostBlogRequest create() => PostBlogRequest._();
  PostBlogRequest createEmptyInstance() => create();
  static $pb.PbList<PostBlogRequest> createRepeated() => $pb.PbList<PostBlogRequest>();
  @$core.pragma('dart2js:noInline')
  static PostBlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PostBlogRequest>(create);
  static PostBlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get tag => $_getSZ(2);
  @$pb.TagNumber(3)
  set tag($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTag() => $_has(2);
  @$pb.TagNumber(3)
  void clearTag() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get description => $_getSZ(3);
  @$pb.TagNumber(4)
  set description($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasDescription() => $_has(3);
  @$pb.TagNumber(4)
  void clearDescription() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get imageData => $_getN(4);
  @$pb.TagNumber(5)
  set imageData($core.List<$core.int> v) { $_setBytes(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasImageData() => $_has(4);
  @$pb.TagNumber(5)
  void clearImageData() => clearField(5);
}

class PostBlogResponse extends $pb.GeneratedMessage {
  factory PostBlogResponse({
    $core.bool? isBlogAdded,
  }) {
    final $result = create();
    if (isBlogAdded != null) {
      $result.isBlogAdded = isBlogAdded;
    }
    return $result;
  }
  PostBlogResponse._() : super();
  factory PostBlogResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PostBlogResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PostBlogResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'isBlogAdded', protoName: 'isBlogAdded')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PostBlogResponse clone() => PostBlogResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PostBlogResponse copyWith(void Function(PostBlogResponse) updates) => super.copyWith((message) => updates(message as PostBlogResponse)) as PostBlogResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PostBlogResponse create() => PostBlogResponse._();
  PostBlogResponse createEmptyInstance() => create();
  static $pb.PbList<PostBlogResponse> createRepeated() => $pb.PbList<PostBlogResponse>();
  @$core.pragma('dart2js:noInline')
  static PostBlogResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PostBlogResponse>(create);
  static PostBlogResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isBlogAdded => $_getBF(0);
  @$pb.TagNumber(1)
  set isBlogAdded($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsBlogAdded() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsBlogAdded() => clearField(1);
}

class UpdateBlogRequest extends $pb.GeneratedMessage {
  factory UpdateBlogRequest({
    $core.String? sessionToken,
    $core.String? oldTitle,
    $core.String? newTitle,
    $core.String? newTag,
    $core.String? newDescription,
    $core.List<$core.int>? newImageData,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (oldTitle != null) {
      $result.oldTitle = oldTitle;
    }
    if (newTitle != null) {
      $result.newTitle = newTitle;
    }
    if (newTag != null) {
      $result.newTag = newTag;
    }
    if (newDescription != null) {
      $result.newDescription = newDescription;
    }
    if (newImageData != null) {
      $result.newImageData = newImageData;
    }
    return $result;
  }
  UpdateBlogRequest._() : super();
  factory UpdateBlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateBlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UpdateBlogRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'oldTitle', protoName: 'oldTitle')
    ..aOS(3, _omitFieldNames ? '' : 'newTitle', protoName: 'newTitle')
    ..aOS(4, _omitFieldNames ? '' : 'newTag', protoName: 'newTag')
    ..aOS(5, _omitFieldNames ? '' : 'newDescription', protoName: 'newDescription')
    ..a<$core.List<$core.int>>(6, _omitFieldNames ? '' : 'newImageData', $pb.PbFieldType.OY, protoName: 'newImageData')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateBlogRequest clone() => UpdateBlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateBlogRequest copyWith(void Function(UpdateBlogRequest) updates) => super.copyWith((message) => updates(message as UpdateBlogRequest)) as UpdateBlogRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UpdateBlogRequest create() => UpdateBlogRequest._();
  UpdateBlogRequest createEmptyInstance() => create();
  static $pb.PbList<UpdateBlogRequest> createRepeated() => $pb.PbList<UpdateBlogRequest>();
  @$core.pragma('dart2js:noInline')
  static UpdateBlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateBlogRequest>(create);
  static UpdateBlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get oldTitle => $_getSZ(1);
  @$pb.TagNumber(2)
  set oldTitle($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasOldTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearOldTitle() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get newTitle => $_getSZ(2);
  @$pb.TagNumber(3)
  set newTitle($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNewTitle() => $_has(2);
  @$pb.TagNumber(3)
  void clearNewTitle() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get newTag => $_getSZ(3);
  @$pb.TagNumber(4)
  set newTag($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasNewTag() => $_has(3);
  @$pb.TagNumber(4)
  void clearNewTag() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get newDescription => $_getSZ(4);
  @$pb.TagNumber(5)
  set newDescription($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasNewDescription() => $_has(4);
  @$pb.TagNumber(5)
  void clearNewDescription() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.int> get newImageData => $_getN(5);
  @$pb.TagNumber(6)
  set newImageData($core.List<$core.int> v) { $_setBytes(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasNewImageData() => $_has(5);
  @$pb.TagNumber(6)
  void clearNewImageData() => clearField(6);
}

class UpdateBlogResponse extends $pb.GeneratedMessage {
  factory UpdateBlogResponse({
    $core.bool? isBlogUpdated,
  }) {
    final $result = create();
    if (isBlogUpdated != null) {
      $result.isBlogUpdated = isBlogUpdated;
    }
    return $result;
  }
  UpdateBlogResponse._() : super();
  factory UpdateBlogResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UpdateBlogResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'UpdateBlogResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'isBlogUpdated', protoName: 'isBlogUpdated')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UpdateBlogResponse clone() => UpdateBlogResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UpdateBlogResponse copyWith(void Function(UpdateBlogResponse) updates) => super.copyWith((message) => updates(message as UpdateBlogResponse)) as UpdateBlogResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UpdateBlogResponse create() => UpdateBlogResponse._();
  UpdateBlogResponse createEmptyInstance() => create();
  static $pb.PbList<UpdateBlogResponse> createRepeated() => $pb.PbList<UpdateBlogResponse>();
  @$core.pragma('dart2js:noInline')
  static UpdateBlogResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UpdateBlogResponse>(create);
  static UpdateBlogResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isBlogUpdated => $_getBF(0);
  @$pb.TagNumber(1)
  set isBlogUpdated($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsBlogUpdated() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsBlogUpdated() => clearField(1);
}

class DeleteBlogRequest extends $pb.GeneratedMessage {
  factory DeleteBlogRequest({
    $core.String? sessionToken,
    $core.String? title,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (title != null) {
      $result.title = title;
    }
    return $result;
  }
  DeleteBlogRequest._() : super();
  factory DeleteBlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteBlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeleteBlogRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteBlogRequest clone() => DeleteBlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteBlogRequest copyWith(void Function(DeleteBlogRequest) updates) => super.copyWith((message) => updates(message as DeleteBlogRequest)) as DeleteBlogRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeleteBlogRequest create() => DeleteBlogRequest._();
  DeleteBlogRequest createEmptyInstance() => create();
  static $pb.PbList<DeleteBlogRequest> createRepeated() => $pb.PbList<DeleteBlogRequest>();
  @$core.pragma('dart2js:noInline')
  static DeleteBlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteBlogRequest>(create);
  static DeleteBlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);
}

class DeleteBlogResponse extends $pb.GeneratedMessage {
  factory DeleteBlogResponse({
    $core.bool? isBlogDeleted,
  }) {
    final $result = create();
    if (isBlogDeleted != null) {
      $result.isBlogDeleted = isBlogDeleted;
    }
    return $result;
  }
  DeleteBlogResponse._() : super();
  factory DeleteBlogResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteBlogResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeleteBlogResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'isBlogDeleted', protoName: 'isBlogDeleted')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteBlogResponse clone() => DeleteBlogResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteBlogResponse copyWith(void Function(DeleteBlogResponse) updates) => super.copyWith((message) => updates(message as DeleteBlogResponse)) as DeleteBlogResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeleteBlogResponse create() => DeleteBlogResponse._();
  DeleteBlogResponse createEmptyInstance() => create();
  static $pb.PbList<DeleteBlogResponse> createRepeated() => $pb.PbList<DeleteBlogResponse>();
  @$core.pragma('dart2js:noInline')
  static DeleteBlogResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteBlogResponse>(create);
  static DeleteBlogResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isBlogDeleted => $_getBF(0);
  @$pb.TagNumber(1)
  set isBlogDeleted($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsBlogDeleted() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsBlogDeleted() => clearField(1);
}

/// Other Operations on Blogs
class GetBlogRequest extends $pb.GeneratedMessage {
  factory GetBlogRequest({
    $core.String? sessionToken,
    $core.String? title,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (title != null) {
      $result.title = title;
    }
    return $result;
  }
  GetBlogRequest._() : super();
  factory GetBlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetBlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetBlogRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetBlogRequest clone() => GetBlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetBlogRequest copyWith(void Function(GetBlogRequest) updates) => super.copyWith((message) => updates(message as GetBlogRequest)) as GetBlogRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetBlogRequest create() => GetBlogRequest._();
  GetBlogRequest createEmptyInstance() => create();
  static $pb.PbList<GetBlogRequest> createRepeated() => $pb.PbList<GetBlogRequest>();
  @$core.pragma('dart2js:noInline')
  static GetBlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetBlogRequest>(create);
  static GetBlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);
}

class GetBlogResponse extends $pb.GeneratedMessage {
  factory GetBlogResponse({
    Blog? blog,
    $core.bool? isBlogLiked,
    $core.bool? isBlogUpdatableDeletable,
  }) {
    final $result = create();
    if (blog != null) {
      $result.blog = blog;
    }
    if (isBlogLiked != null) {
      $result.isBlogLiked = isBlogLiked;
    }
    if (isBlogUpdatableDeletable != null) {
      $result.isBlogUpdatableDeletable = isBlogUpdatableDeletable;
    }
    return $result;
  }
  GetBlogResponse._() : super();
  factory GetBlogResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetBlogResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetBlogResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOM<Blog>(1, _omitFieldNames ? '' : 'blog', subBuilder: Blog.create)
    ..aOB(2, _omitFieldNames ? '' : 'isBlogLiked', protoName: 'isBlogLiked')
    ..aOB(3, _omitFieldNames ? '' : 'isBlogUpdatableDeletable', protoName: 'isBlogUpdatableDeletable')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetBlogResponse clone() => GetBlogResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetBlogResponse copyWith(void Function(GetBlogResponse) updates) => super.copyWith((message) => updates(message as GetBlogResponse)) as GetBlogResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetBlogResponse create() => GetBlogResponse._();
  GetBlogResponse createEmptyInstance() => create();
  static $pb.PbList<GetBlogResponse> createRepeated() => $pb.PbList<GetBlogResponse>();
  @$core.pragma('dart2js:noInline')
  static GetBlogResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetBlogResponse>(create);
  static GetBlogResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Blog get blog => $_getN(0);
  @$pb.TagNumber(1)
  set blog(Blog v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasBlog() => $_has(0);
  @$pb.TagNumber(1)
  void clearBlog() => clearField(1);
  @$pb.TagNumber(1)
  Blog ensureBlog() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.bool get isBlogLiked => $_getBF(1);
  @$pb.TagNumber(2)
  set isBlogLiked($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasIsBlogLiked() => $_has(1);
  @$pb.TagNumber(2)
  void clearIsBlogLiked() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get isBlogUpdatableDeletable => $_getBF(2);
  @$pb.TagNumber(3)
  set isBlogUpdatableDeletable($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasIsBlogUpdatableDeletable() => $_has(2);
  @$pb.TagNumber(3)
  void clearIsBlogUpdatableDeletable() => clearField(3);
}

class LikeBlogRequest extends $pb.GeneratedMessage {
  factory LikeBlogRequest({
    $core.String? sessionToken,
    $core.String? title,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (title != null) {
      $result.title = title;
    }
    return $result;
  }
  LikeBlogRequest._() : super();
  factory LikeBlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LikeBlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'LikeBlogRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  LikeBlogRequest clone() => LikeBlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  LikeBlogRequest copyWith(void Function(LikeBlogRequest) updates) => super.copyWith((message) => updates(message as LikeBlogRequest)) as LikeBlogRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static LikeBlogRequest create() => LikeBlogRequest._();
  LikeBlogRequest createEmptyInstance() => create();
  static $pb.PbList<LikeBlogRequest> createRepeated() => $pb.PbList<LikeBlogRequest>();
  @$core.pragma('dart2js:noInline')
  static LikeBlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LikeBlogRequest>(create);
  static LikeBlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);
}

class LikeBlogResponse extends $pb.GeneratedMessage {
  factory LikeBlogResponse({
    $core.bool? isBlogLiked,
  }) {
    final $result = create();
    if (isBlogLiked != null) {
      $result.isBlogLiked = isBlogLiked;
    }
    return $result;
  }
  LikeBlogResponse._() : super();
  factory LikeBlogResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LikeBlogResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'LikeBlogResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'isBlogLiked', protoName: 'isBlogLiked')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  LikeBlogResponse clone() => LikeBlogResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  LikeBlogResponse copyWith(void Function(LikeBlogResponse) updates) => super.copyWith((message) => updates(message as LikeBlogResponse)) as LikeBlogResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static LikeBlogResponse create() => LikeBlogResponse._();
  LikeBlogResponse createEmptyInstance() => create();
  static $pb.PbList<LikeBlogResponse> createRepeated() => $pb.PbList<LikeBlogResponse>();
  @$core.pragma('dart2js:noInline')
  static LikeBlogResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LikeBlogResponse>(create);
  static LikeBlogResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isBlogLiked => $_getBF(0);
  @$pb.TagNumber(1)
  set isBlogLiked($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsBlogLiked() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsBlogLiked() => clearField(1);
}

class DislikeBlogRequest extends $pb.GeneratedMessage {
  factory DislikeBlogRequest({
    $core.String? sessionToken,
    $core.String? title,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (title != null) {
      $result.title = title;
    }
    return $result;
  }
  DislikeBlogRequest._() : super();
  factory DislikeBlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DislikeBlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DislikeBlogRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DislikeBlogRequest clone() => DislikeBlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DislikeBlogRequest copyWith(void Function(DislikeBlogRequest) updates) => super.copyWith((message) => updates(message as DislikeBlogRequest)) as DislikeBlogRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DislikeBlogRequest create() => DislikeBlogRequest._();
  DislikeBlogRequest createEmptyInstance() => create();
  static $pb.PbList<DislikeBlogRequest> createRepeated() => $pb.PbList<DislikeBlogRequest>();
  @$core.pragma('dart2js:noInline')
  static DislikeBlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DislikeBlogRequest>(create);
  static DislikeBlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);
}

class DislikeBlogResponse extends $pb.GeneratedMessage {
  factory DislikeBlogResponse({
    $core.bool? isBlogDisliked,
  }) {
    final $result = create();
    if (isBlogDisliked != null) {
      $result.isBlogDisliked = isBlogDisliked;
    }
    return $result;
  }
  DislikeBlogResponse._() : super();
  factory DislikeBlogResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DislikeBlogResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DislikeBlogResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'isBlogDisliked', protoName: 'isBlogDisliked')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DislikeBlogResponse clone() => DislikeBlogResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DislikeBlogResponse copyWith(void Function(DislikeBlogResponse) updates) => super.copyWith((message) => updates(message as DislikeBlogResponse)) as DislikeBlogResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DislikeBlogResponse create() => DislikeBlogResponse._();
  DislikeBlogResponse createEmptyInstance() => create();
  static $pb.PbList<DislikeBlogResponse> createRepeated() => $pb.PbList<DislikeBlogResponse>();
  @$core.pragma('dart2js:noInline')
  static DislikeBlogResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DislikeBlogResponse>(create);
  static DislikeBlogResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isBlogDisliked => $_getBF(0);
  @$pb.TagNumber(1)
  set isBlogDisliked($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsBlogDisliked() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsBlogDisliked() => clearField(1);
}

class AddCommentRequest extends $pb.GeneratedMessage {
  factory AddCommentRequest({
    $core.String? sessionToken,
    $core.String? title,
    $core.String? commentDescription,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (title != null) {
      $result.title = title;
    }
    if (commentDescription != null) {
      $result.commentDescription = commentDescription;
    }
    return $result;
  }
  AddCommentRequest._() : super();
  factory AddCommentRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddCommentRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AddCommentRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..aOS(3, _omitFieldNames ? '' : 'commentDescription', protoName: 'commentDescription')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddCommentRequest clone() => AddCommentRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddCommentRequest copyWith(void Function(AddCommentRequest) updates) => super.copyWith((message) => updates(message as AddCommentRequest)) as AddCommentRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AddCommentRequest create() => AddCommentRequest._();
  AddCommentRequest createEmptyInstance() => create();
  static $pb.PbList<AddCommentRequest> createRepeated() => $pb.PbList<AddCommentRequest>();
  @$core.pragma('dart2js:noInline')
  static AddCommentRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddCommentRequest>(create);
  static AddCommentRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get commentDescription => $_getSZ(2);
  @$pb.TagNumber(3)
  set commentDescription($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasCommentDescription() => $_has(2);
  @$pb.TagNumber(3)
  void clearCommentDescription() => clearField(3);
}

class AddCommentResponse extends $pb.GeneratedMessage {
  factory AddCommentResponse({
    $core.bool? isCommentAdded,
  }) {
    final $result = create();
    if (isCommentAdded != null) {
      $result.isCommentAdded = isCommentAdded;
    }
    return $result;
  }
  AddCommentResponse._() : super();
  factory AddCommentResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AddCommentResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'AddCommentResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'isCommentAdded', protoName: 'isCommentAdded')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  AddCommentResponse clone() => AddCommentResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  AddCommentResponse copyWith(void Function(AddCommentResponse) updates) => super.copyWith((message) => updates(message as AddCommentResponse)) as AddCommentResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static AddCommentResponse create() => AddCommentResponse._();
  AddCommentResponse createEmptyInstance() => create();
  static $pb.PbList<AddCommentResponse> createRepeated() => $pb.PbList<AddCommentResponse>();
  @$core.pragma('dart2js:noInline')
  static AddCommentResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AddCommentResponse>(create);
  static AddCommentResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get isCommentAdded => $_getBF(0);
  @$pb.TagNumber(1)
  set isCommentAdded($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIsCommentAdded() => $_has(0);
  @$pb.TagNumber(1)
  void clearIsCommentAdded() => clearField(1);
}

class SearchBlogRequest extends $pb.GeneratedMessage {
  factory SearchBlogRequest({
    $core.String? sessionToken,
    $core.String? title,
  }) {
    final $result = create();
    if (sessionToken != null) {
      $result.sessionToken = sessionToken;
    }
    if (title != null) {
      $result.title = title;
    }
    return $result;
  }
  SearchBlogRequest._() : super();
  factory SearchBlogRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SearchBlogRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SearchBlogRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'sessionToken', protoName: 'sessionToken')
    ..aOS(2, _omitFieldNames ? '' : 'title')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SearchBlogRequest clone() => SearchBlogRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SearchBlogRequest copyWith(void Function(SearchBlogRequest) updates) => super.copyWith((message) => updates(message as SearchBlogRequest)) as SearchBlogRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SearchBlogRequest create() => SearchBlogRequest._();
  SearchBlogRequest createEmptyInstance() => create();
  static $pb.PbList<SearchBlogRequest> createRepeated() => $pb.PbList<SearchBlogRequest>();
  @$core.pragma('dart2js:noInline')
  static SearchBlogRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SearchBlogRequest>(create);
  static SearchBlogRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get sessionToken => $_getSZ(0);
  @$pb.TagNumber(1)
  set sessionToken($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSessionToken() => $_has(0);
  @$pb.TagNumber(1)
  void clearSessionToken() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get title => $_getSZ(1);
  @$pb.TagNumber(2)
  set title($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTitle() => $_has(1);
  @$pb.TagNumber(2)
  void clearTitle() => clearField(2);
}

class SearchBlogResponse extends $pb.GeneratedMessage {
  factory SearchBlogResponse({
    $core.bool? doesBlogExists,
  }) {
    final $result = create();
    if (doesBlogExists != null) {
      $result.doesBlogExists = doesBlogExists;
    }
    return $result;
  }
  SearchBlogResponse._() : super();
  factory SearchBlogResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SearchBlogResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SearchBlogResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'proto'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'doesBlogExists', protoName: 'doesBlogExists')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SearchBlogResponse clone() => SearchBlogResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SearchBlogResponse copyWith(void Function(SearchBlogResponse) updates) => super.copyWith((message) => updates(message as SearchBlogResponse)) as SearchBlogResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SearchBlogResponse create() => SearchBlogResponse._();
  SearchBlogResponse createEmptyInstance() => create();
  static $pb.PbList<SearchBlogResponse> createRepeated() => $pb.PbList<SearchBlogResponse>();
  @$core.pragma('dart2js:noInline')
  static SearchBlogResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SearchBlogResponse>(create);
  static SearchBlogResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get doesBlogExists => $_getBF(0);
  @$pb.TagNumber(1)
  set doesBlogExists($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDoesBlogExists() => $_has(0);
  @$pb.TagNumber(1)
  void clearDoesBlogExists() => clearField(1);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
