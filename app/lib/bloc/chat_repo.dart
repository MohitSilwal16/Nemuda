import 'dart:convert';
import 'package:app/main.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

import 'package:app/pb/user.pb.dart';
import 'package:app/models/message.dart';
import 'package:app/services/service_init.dart';

class ChatRepo {
  late WebSocketChannel _channel;
  late String? _sessionToken;

  static final ChatRepo _instance = ChatRepo._internal();

  factory ChatRepo() {
    return _instance;
  }

  ChatRepo._internal();

  void init() {
    _sessionToken = ServiceManager().hiveBox.get("sessionToken");
    if (_sessionToken == null) {
      throw Exception("INVALID SESSION TOKEN");
    }
    _channel = WebSocketChannel.connect(
      Uri.parse("ws://$serviceURL:$wsPort/ws/chat/$_sessionToken"),
    );
  }

  Stream<ModelMessage> get messagesStream {
    return _channel.stream.map((data) {
      final jsonData = jsonDecode(data as String) as Map<String, dynamic>;

      final ModelMessage modelMessage = ModelMessage.fromJson(jsonData);
      if (modelMessage.error != "") {
        throw Exception(modelMessage.error);
      }

      return modelMessage;
    });
  }

  Future<GetMessagesResponse> fetchPrevMessages(
      String user1, int offset) async {
    final request = GetMessagesRequest(
        sessionToken: _sessionToken, user1: user1, offset: offset);

    final res = await ServiceManager()
        .userClient
        .getMessagesWithPagination(request)
        .timeout(contextTimeout);

    return res;
  }

  void sendMessage(String message, String receiver) {
    _channel.sink.add(
      jsonEncode(ModelWSMessage(
        message: message,
        receiver: receiver,
        sessionToken: _sessionToken!,
        messageType: "Message",
      ).toJson()),
    );
  }

  void markMsgAsRead(String receiver) {
    _channel.sink.add(
      jsonEncode(ModelWSMessage(
        message: "",
        receiver: receiver, // Receiver = User whose message I'm read
        sessionToken: _sessionToken!,
        messageType: "Read",
      ).toJson()),
    );
  }

  Future<List<UserAndLastMessage>> getUsersByStartingPattern(
      String searchPattern) async {
    final request = SearchUsersByStartingPatternRequest(
        sessionToken: _sessionToken, searchPattern: searchPattern);

    final res = await ServiceManager()
        .userClient
        .searchUsersByStartingPattern(request)
        .timeout(shortContextTimeout);

    return res.usersAndLastMessage;
  }

  void dispose() {
    _channel.sink.close();
  }
}
