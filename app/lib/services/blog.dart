import 'dart:async';

import 'package:app/pb/blogs.pb.dart';
import 'package:app/pb/blogs.pbgrpc.dart';
import 'package:app/services/service_init.dart';

Future<GetBlogsResponse> getBlogsWithPagination(String tag, int offset) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request =
      GetBlogsRequest(sessionToken: sessionToken, tag: tag, offset: offset);
  final response = await Clients()
      .blogClient
      .getBlogsByTagWithPagination(request)
      .timeout(contextTimeout);
  return response;
}

Future<UpdateBlogResponse> updateBlog(String oldTitle, String newTitle,
    String newDescription, String newTag, List<int> imageData) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = UpdateBlogRequest(
    sessionToken: sessionToken,
    newTag: newTag,
    newDescription: newDescription,
    newTitle: newTitle,
    oldTitle: oldTitle,
    newImageData: imageData,
  );

  final response = await Clients().blogClient.updateBlog(request);
  return response;
}

Future<PostBlogResponse> postBlog(
    String title, String description, String tag, List<int> imageData) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = PostBlogRequest(
    sessionToken: sessionToken,
    tag: tag,
    description: description,
    title: title,
    imageData: imageData,
  );

  final response = await Clients().blogClient.postBlog(request);
  return response;
}

Future<DeleteBlogResponse> deleteBlog(String title) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = DeleteBlogRequest(sessionToken: sessionToken, title: title);

  final response = Clients().blogClient.deleteBlog(request);
  return response;
}

// Other Blog Operations
Future<LikeBlogResponse> likeBlog(String title) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = LikeBlogRequest(sessionToken: sessionToken, title: title);
  final response =
      await Clients().blogClient.likeBlog(request).timeout(contextTimeout);
  return response;
}

Future<DislikeBlogResponse> dislikeBlog(String title) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = DislikeBlogRequest(sessionToken: sessionToken, title: title);
  final response =
      await Clients().blogClient.dislikeBlog(request).timeout(contextTimeout);
  return response;
}

Future<AddCommentResponse> addComment(String title, String comment) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = AddCommentRequest(
      sessionToken: sessionToken, title: title, commentDescription: comment);

  final respone = await Clients().blogClient.addComment(request);
  return respone;
}

Future<SearchBlogResponse> searchBlog(String title) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = SearchBlogRequest(sessionToken: sessionToken, title: title);

  final response = await Clients().blogClient.searchBlogByTitle(request);
  return response;
}

Future<GetBlogResponse> getBlogByTitle(String title) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = GetBlogRequest(sessionToken: sessionToken, title: title);

  final response = await Clients().blogClient.getBlogByTitle(request);
  return response;
}
