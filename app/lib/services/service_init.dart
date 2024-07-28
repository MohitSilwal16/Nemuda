import 'dart:async';

import 'package:app/main.dart';
import 'package:app/pb/blogs.pbgrpc.dart';
import 'package:app/pb/user.pbgrpc.dart';
import 'package:grpc/grpc.dart';
import 'package:app/pb/auth.pbgrpc.dart';
import 'package:hive/hive.dart';

const shortContextTimeout = Duration(seconds: 3);
const contextTimeout = Duration(seconds: 5);
const longContextTimeout = Duration(seconds: 10);

class Clients {
  static final Clients _instance = Clients._internal();

  factory Clients() {
    return _instance;
  }

  Clients._internal();

  late ClientChannel _channel;
  late AuthServiceClient authClient;
  late BlogsServiceClient blogClient;
  late UserServiceClient userClient;
  late Box<dynamic> hiveBox;

  Future<void> init({
    String host = serviceURL,
    int port = servicePort,
  }) async {
    _channel = ClientChannel(
      host,
      port: port,
      options: const ChannelOptions(
        credentials: ChannelCredentials.insecure(),
      ),
    );

    authClient = AuthServiceClient(_channel);
    userClient = UserServiceClient(_channel);
    blogClient = BlogsServiceClient(_channel);

    hiveBox = Hive.box("session");
  }
}
