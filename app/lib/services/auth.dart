import 'dart:async';

import 'package:app/services/service_init.dart';
import 'package:app/pb/auth.pbgrpc.dart';

Future<AuthResponse> login(String username, String password) async {
  final request = AuthRequest(username: username, password: password);
  final response =
      await ServiceManager().authClient.login(request).timeout(contextTimeout);
  return response;
}

Future<AuthResponse> register(String username, String password) async {
  final request = AuthRequest(username: username, password: password);
  final response = await ServiceManager()
      .authClient
      .register(request)
      .timeout(contextTimeout);
  return response;
}

Future<ValidationResponse> validateSessionToken() async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  final request = ValidationRequest(sessionToken: sessionToken);
  final response = await ServiceManager()
      .authClient
      .verifySessionToken(request)
      .timeout(contextTimeout);
  return response;
}

Future<LogoutResponse> logout() async {
  final sessionToken = ServiceManager().hiveBox.get("sessionToken");
  final request = LogoutRequest(sessionToken: sessionToken);
  final response =
      await ServiceManager().authClient.logout(request).timeout(contextTimeout);
  return response;
}
