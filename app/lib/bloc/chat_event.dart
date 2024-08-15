import 'package:app/models/message.dart';
import 'package:equatable/equatable.dart';

abstract class ChatEvent extends Equatable {}

class EventSendMessage extends ChatEvent {
  final String message;
  final String receiver;

  EventSendMessage({
    required this.message,
    required this.receiver,
  });

  @override
  List<Object?> get props => [message, receiver];
}

class EventFetchPrevMessages extends ChatEvent {
  final String receiver;
  final int offset;

  EventFetchPrevMessages({required this.offset, required this.receiver});

  @override
  List<Object?> get props => [receiver, offset];
}

class EventMarkMsgAsRead extends ChatEvent {
  final String receiver;

  EventMarkMsgAsRead({required this.receiver});

  @override
  List<Object?> get props => [receiver];
}

class EventNewMsgReceived extends ChatEvent {
  final ModelMessage message;

  EventNewMsgReceived({required this.message});

  @override
  List<Object?> get props => [message];
}

class EventSearchUser extends ChatEvent {
  final String searchPattern;

  EventSearchUser({required this.searchPattern});

  @override
  List<Object?> get props => [searchPattern];
}

class EventError extends ChatEvent {
  final String message;

  EventError({required this.message});

  @override
  List<Object?> get props => [message];
}
