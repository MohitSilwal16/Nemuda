import 'dart:async';

import 'package:app/pb/user.pb.dart';
import 'package:app/pb/user.pbgrpc.dart';
import 'package:app/services/service_init.dart';

Future<UserExistsResponse> doesUserExists(String username) async {
  final request = UserExistsRequest(username: username);
  final response = await Clients()
      .userClient
      .doesUserExists(request);
      // .timeout(contextTimeout);
  return response;
}

Future<SearchUsersByStartingPatternResponse> searchUsersByStartingPattern(
    String searchPattern) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = SearchUsersByStartingPatternRequest(
      sessionToken: sessionToken, searchPattern: searchPattern);

  final response = Clients().userClient.searchUsersByStartingPattern(request);
  return response;
}

Future<GetMessagesResponse> getMessages(
    String user1, int offset) async {
  final sessionToken = Clients().hiveBox.get("sessionToken");

  final request = GetMessagesRequest(
      sessionToken: sessionToken, user1: user1, offset: offset);

  final response = Clients().userClient.getMessagesWithPagination(request);
  return response;
}

