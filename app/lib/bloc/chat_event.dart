import 'package:app/models/message.dart';

abstract class ChatEvent {}

class EventSendMessage extends ChatEvent {
  final String message;
  final String receiver;

  EventSendMessage({
    required this.message,
    required this.receiver,
  });
}

class EventFetchPrevMessages extends ChatEvent {
  final String receiver;
  final int offset;

  EventFetchPrevMessages({required this.offset, required this.receiver});
}

class EventMarkMsgAsRead extends ChatEvent {
  final String receiver;

  EventMarkMsgAsRead({required this.receiver});
}

class EventNewMsgReceived extends ChatEvent {
  final ModelMessage message;

  EventNewMsgReceived({required this.message});
}

class EventSearchUser extends ChatEvent {
  final String searchPattern;

  EventSearchUser({required this.searchPattern});
}

class EventError extends ChatEvent {
  final String message;

  EventError({required this.message});
}