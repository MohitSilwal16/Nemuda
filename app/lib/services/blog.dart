import 'dart:async';

import 'package:app/pb/blogs.pbgrpc.dart';
import 'package:app/services/service_init.dart';
import 'package:hive/hive.dart';

Future<GetBlogsResponse> getBlogsWithPagination(String tag, int offset) async {
  final sessionToken = Hive.box("session").get("sessionToken");

  final request =
      GetBlogsRequest(sessionToken: sessionToken, tag: tag, offset: offset);
  final response = await Clients()
      .blogClient
      .getBlogsByTagWithPagination(request)
      .timeout(contextTimeout);
  return response;
}
