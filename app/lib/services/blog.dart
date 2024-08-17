import 'dart:async';

import 'package:app/pb/blogs.pb.dart';
import 'package:app/pb/blogs.pbgrpc.dart';
import 'package:app/services/service_init.dart';

Future<GetBlogsResponse> getBlogsWithPagination(String tag, int offset) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request =
      GetBlogsRequest(sessionToken: sessionToken, tag: tag, offset: offset);
  final response = await ServiceManager()
      .blogClient
      .getBlogsByTagWithPagination(request)
      .timeout(contextTimeout);
  return response;
}

Future<UpdateBlogResponse> updateBlog(String oldTitle, String newTitle,
    String newDescription, String newTag, List<int> imageData) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = UpdateBlogRequest(
    sessionToken: sessionToken,
    newTag: newTag,
    newDescription: newDescription,
    newTitle: newTitle,
    oldTitle: oldTitle,
    newImageData: imageData,
  );

  final response = await ServiceManager()
      .blogClient
      .updateBlog(request)
      .timeout(longContextTimeout);
  return response;
}

Future<PostBlogResponse> postBlog(
    String title, String description, String tag, List<int> imageData) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = PostBlogRequest(
    sessionToken: sessionToken,
    tag: tag,
    description: description,
    title: title,
    imageData: imageData,
  );

  final response = await ServiceManager()
      .blogClient
      .postBlog(request)
      .timeout(longContextTimeout);
  return response;
}

Future<DeleteBlogResponse> deleteBlog(String title) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = DeleteBlogRequest(sessionToken: sessionToken, title: title);
  final response = ServiceManager()
      .blogClient
      .deleteBlog(request)
      .timeout(longContextTimeout);
  return response;
}

// Other Blog Operations
Future<LikeBlogResponse> likeBlog(String title) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = LikeBlogRequest(sessionToken: sessionToken, title: title);
  final response = await ServiceManager()
      .blogClient
      .likeBlog(request)
      .timeout(contextTimeout);
  return response;
}

Future<DislikeBlogResponse> dislikeBlog(String title) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = DislikeBlogRequest(sessionToken: sessionToken, title: title);
  final response = await ServiceManager()
      .blogClient
      .dislikeBlog(request)
      .timeout(contextTimeout);
  return response;
}

Future<AddCommentResponse> addComment(String title, String comment) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = AddCommentRequest(
      sessionToken: sessionToken, title: title, commentDescription: comment);

  final respone = await ServiceManager()
      .blogClient
      .addComment(request)
      .timeout(contextTimeout);
  return respone;
}

Future<SearchBlogResponse> searchBlog(String title) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = SearchBlogRequest(sessionToken: sessionToken, title: title);

  final response = await ServiceManager()
      .blogClient
      .searchBlogByTitle(request)
      .timeout(shortContextTimeout);
  return response;
}

Future<GetBlogResponse> getBlogByTitle(String title) async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  if (sessionToken == null) {
    throw Exception("INVALID SESSION TOKEN");
  }

  final request = GetBlogRequest(sessionToken: sessionToken, title: title);

  final response = await ServiceManager()
      .blogClient
      .getBlogByTitle(request)
      .timeout(contextTimeout);
  return response;
}
