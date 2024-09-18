import 'dart:async';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'package:app/utils/utils.dart';
import 'package:app/bloc/chat_event.dart';
import 'package:app/bloc/chat_repo.dart';
import 'package:app/bloc/chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final ChatRepo repo;
  late StreamSubscription messageSubscription;

  static final ChatBloc _instance = ChatBloc._internal(repo: ChatRepo());

  factory ChatBloc() {
    return _instance;
  }

  void updateStreamSubscription(ChatRepo repo) {
    messageSubscription = repo.messagesStream.listen(
      (message) {
        add(EventNewMsgReceived(message: message));
      },
      onError: (err) => add(EventError(message: err.toString())),
      onDone: () => add(EventError(message: "WebSocket Disconnected")),
      cancelOnError: true,
    );
  }

  ChatBloc._internal({required this.repo}) : super(StateChatInitial()) {
    on<EventSendMessage>(_sendMessage);
    on<EventFetchPrevMessages>(_fetchPrevMessages);
    on<EventNewMsgReceived>(_newMsgReceived);
    on<EventSearchUser>(_searchUser);
    on<EventMarkMsgAsRead>(_markMsgAsRead);
    on<EventError>(_onerror);
    on<EventNothing>(_onNothing);
  }

  Future<void> _fetchPrevMessages(
      EventFetchPrevMessages event, Emitter<ChatState> emit) async {
    if (state is StateUserLoaded) {
      emit(StateChatLoading());
    }
    try {
      final res = await repo.fetchPrevMessages(event.receiver, event.offset);
      emit(StateChatLoaded(messages: res.messages, nextOffset: res.nextOffset));
    } catch (err) {
      final errMsg = trimGrpcErrorMessage(err.toString());
      emit(StateChatError(exception: errMsg));
    }
  }

  Future<void> _sendMessage(
      EventSendMessage event, Emitter<ChatState> emit) async {
    try {
      repo.sendMessage(event.message, event.receiver);
    } catch (err) {
      emit(StateChatError(exception: err.toString()));
    }
  }

  Future<void> _newMsgReceived(
      EventNewMsgReceived event, Emitter<ChatState> emit) async {
    try {
      emit(StateNewMsgReceived(message: event.message));
    } catch (err) {
      emit(StateChatError(exception: err.toString()));
    }
  }

  Future<void> _searchUser(
      EventSearchUser event, Emitter<ChatState> emit) async {
    emit(StateChatLoading());
    try {
      final res = await repo.getUsersByStartingPattern(event.searchPattern);
      emit(StateUserLoaded(usersAndLastMsg: res));
    } catch (err) {
      final errMsg = trimGrpcErrorMessage(err.toString());
      emit(StateChatError(exception: errMsg));
    }
  }

  Future<void> _markMsgAsRead(
      EventMarkMsgAsRead event, Emitter<ChatState> emit) async {
    try {
      repo.markMsgAsRead(event.receiver);
    } catch (err) {
      emit(StateChatError(exception: err.toString()));
    }
  }

  Future<void> _onerror(EventError event, Emitter<ChatState> emit) async {
    emit(StateChatError(exception: event.message));
  }

  Future<void> _onNothing(EventNothing event, Emitter<ChatState> emit) async {
    emit(StateNothing());
  }

  @override
  Future<void> close() {
    messageSubscription.cancel();
    return super.close();
  }
}
