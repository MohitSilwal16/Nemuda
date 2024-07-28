import 'dart:async';

import 'package:app/pb/user.pb.dart';
import 'package:app/services/service_init.dart';

Future<UserExistsResponse> doesUserExists(String username) async {
  final request = UserExistsRequest(username: username);
  final response = await Clients()
      .userClient
      .doesUserExists(request)
      .timeout(contextTimeout);
  return response;
}
