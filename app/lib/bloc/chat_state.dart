import 'package:app/pb/user.pb.dart';
import 'package:app/models/message.dart';

abstract class ChatState {}

class StateChatInitial extends ChatState {}

class StateChatLoading extends ChatState {}

class StateChatLoaded extends ChatState {
  final List<Message> messages;
  final int nextOffset;

  StateChatLoaded({required this.messages, required this.nextOffset});
}

class StateNewMsgReceived extends ChatState {
  final ModelMessage message;

  StateNewMsgReceived({
    required this.message,
  });
}

class StateUserLoaded extends ChatState {
  final List<UserAndLastMessage> usersAndLastMsg;

  StateUserLoaded({required this.usersAndLastMsg});
}

class StateChatError extends ChatState {
  final String errorMessage;

  StateChatError({
    required this.errorMessage,
  });
}

class StateMsgMarkedAsRead extends ChatState {}

class StateNothing extends ChatState {}