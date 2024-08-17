import 'dart:async';

import 'package:app/pb/user.pb.dart';
import 'package:app/pb/user.pbgrpc.dart';
import 'package:app/services/service_init.dart';

Future<UserExistsResponse> doesUserExists(String username) async {
  final request = UserExistsRequest(username: username);

  final response = await ServiceManager()
      .userClient
      .doesUserExists(request)
      .timeout(shortContextTimeout);
  return response;
}
