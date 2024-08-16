import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:visibility_detector/visibility_detector.dart';

import 'package:app/main.dart';
import 'package:app/pb/user.pbgrpc.dart';
import 'package:app/bloc/chat_bloc.dart';
import 'package:app/bloc/chat_event.dart';
import 'package:app/bloc/chat_state.dart';
import 'package:app/utils/colors.dart';
import 'package:app/utils/components/error.dart';
import 'package:app/utils/components/loading.dart';
import 'package:app/utils/components/message_card.dart';
import 'package:app/utils/components/message_textfield.dart';
import 'package:app/utils/components/show_notification.dart';

class ChatUserPage extends StatefulWidget {
  final String user;

  const ChatUserPage({
    super.key,
    required this.user,
  });

  @override
  State<ChatUserPage> createState() => _ChatUserPageState();
}

class _ChatUserPageState extends State<ChatUserPage> {
  final TextEditingController controllerMessage = TextEditingController();

  void sendMessage(String message) {
    context
        .read<ChatBloc>()
        .add(EventSendMessage(message: message, receiver: widget.user));
    controllerMessage.clear();
  }

  @override
  void dispose() {
    controllerMessage.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset:
          true, // To Scroll Properly when keyboard is open
      appBar: AppBar(
        leading: CircleAvatar(
          backgroundColor: Colors.transparent,
          child: IconButton(
            onPressed: () =>
                Navigator.pushReplacementNamed(context, "chat_home"),
            icon: const Icon(Icons.arrow_back_ios_new),
          ),
        ),
        title: Text(
          widget.user,
          style: const TextStyle(fontWeight: FontWeight.w700, fontSize: 24),
        ),
        centerTitle: true,
      ),
      body: PopScope(
        canPop: false,
        onPopInvoked: (didPop) {
          if (didPop) {
            return;
          }
          Navigator.pushReplacementNamed(context, "chat_home");
        },
        child: Container(
          width: size.width,
          height: size.height,
          padding: const EdgeInsets.symmetric(vertical: 20),
          decoration: const BoxDecoration(
            image: DecorationImage(
              image: AssetImage("assets/chat_bg.jpg"),
              fit: BoxFit.cover,
            ),
          ),
          child: Column(
            children: [
              Expanded(
                child: SingleChildScrollView(
                  reverse: true,
                  child: _BuildMessages(user: widget.user),
                ),
              ),

              SizedBox(height: size.height * .035),

              MyMessageTextField(
                controller: controllerMessage,
                sendMessage: sendMessage,
              ),

              // END
            ],
          ),
        ),
      ),
    );
  }
}

class _BuildMessages extends StatefulWidget {
  const _BuildMessages({required this.user});

  final String user;

  @override
  State<_BuildMessages> createState() => _BuildMessagesState();
}

class _BuildMessagesState extends State<_BuildMessages> {
  int offset = 0;
  List<Message> messages = [];
  final ScrollController controllerScroll = ScrollController();

  void navigateToChatPage(String user) {
    Navigator.pushReplacement(
      context,
      MaterialPageRoute(
        builder: (context) => ChatUserPage(user: user),
      ),
    );
  }

  @override
  void initState() {
    context.read<ChatBloc>().add(EventMarkMsgAsRead(receiver: widget.user));
    context
        .read<ChatBloc>()
        .add(EventFetchPrevMessages(offset: 0, receiver: widget.user));
    super.initState();
  }

  @override
  void dispose() {
    controllerScroll.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: size.height * .7,
      child: BlocBuilder<ChatBloc, ChatState>(
        buildWhen: (previous, current) {
          return current is! StateNothing && previous != current;
        },
        builder: (context, state) {
          final currentState = state;
          if (currentState is StateChatLoading) {
            return const CustomCircularProgressIndicator();
          }

          if (currentState is StateChatError) {
            handleErrorsBlocBuilder(context, currentState.errorMessage);
          }

          if (currentState is StateChatLoaded) {
            // Set to avoid duplicacy
            Set<Message> messageSet = currentState.messages.toSet();
            messageSet.addAll(messages);
            messages = messageSet.toList()
              ..sort((a, b) => a.dateTime.compareTo(b.dateTime));
            offset = currentState.nextOffset;

            // Scroll to the bottom
            if (messages.length <= 9) {
              WidgetsBinding.instance.addPostFrameCallback((_) {
                controllerScroll
                    .jumpTo(controllerScroll.position.maxScrollExtent);
              });
            } else {
              // Scroll lil bit downwards when old message is loaded
              // Else it's gonna request for older messages again & again
              WidgetsBinding.instance.addPostFrameCallback((_) {
                controllerScroll.jumpTo(70);
              });
            }
          }

          if (currentState is StateNewMsgReceived) {
            if (currentState.message.error != "") {
              handleErrorsBlocBuilder(context, currentState.message.error);
            } else if (currentState.message.messageType == "Message" &&
                messages.isNotEmpty &&
                messages.last.messageContent !=
                    currentState.message.messageContent &&
                messages.last.dateTime != currentState.message.dateTime) {
              if (currentState.message.sender == widget.user ||
                  currentState.message.selfMessage) {
                // Retain only the last 18 messages
                int keepLastN = 18;
                if (messages.length > 27) {
                  // Remove 18 messages if msg list size is more than 27
                  messages = messages.sublist(messages.length - keepLastN);
                  offset = 18;
                }

                messages.add(
                  Message(
                    sender: currentState.message.sender,
                    receiver: currentState.message.receiver,
                    messageContent: currentState.message.messageContent,
                    dateTime: currentState.message.dateTime,
                    status: currentState.message.status,
                  ),
                );

                // Acknowledge user that I've read your msg
                context
                    .read<ChatBloc>()
                    .add(EventMarkMsgAsRead(receiver: widget.user));
                WidgetsBinding.instance.addPostFrameCallback((_) {
                  controllerScroll
                      .jumpTo(controllerScroll.position.maxScrollExtent);
                });
              } else {
                showNotificationDialog(
                  context,
                  currentState.message.sender,
                  currentState.message.messageContent,
                  () => navigateToChatPage(currentState.message.sender),
                );
              }
            } else if (currentState.message.messageType == "Read") {
              for (int i = messages.length - 1; i >= 0; --i) {
                if (messages[i].status == "Read") break;
                messages[i].status = "Read";
              }
            } else if (currentState.message.messageType == "Delivered") {
              for (int i = messages.length - 1; i >= 0; --i) {
                if (messages[i].status == "Read" ||
                    messages[i].status == "Delivered") break;
                messages[i].status = "Delivered";
              }
            }
          }

          return ListView.builder(
            physics: const BouncingScrollPhysics(),
            itemCount: messages.length + 1,
            controller: controllerScroll,
            itemBuilder: (context, index) {
              if (index == 0) {
                return _LoadMoreMessages(
                  offset: offset,
                  user: widget.user,
                );
              }
              return MessageCard(
                user: widget.user,
                message: messages[index - 1],
              );
            },
          );
        },
      ),
    );
  }
}

class _LoadMoreMessages extends StatelessWidget {
  const _LoadMoreMessages({
    required this.offset,
    required this.user,
  });

  final int offset;
  final String user;

  @override
  Widget build(BuildContext context) {
    return VisibilityDetector(
      key: const Key("Load-More-Msg"),
      child: _NoMessagesContainer(offset: offset),
      onVisibilityChanged: (info) {
        if (info.visibleFraction > 0) {
          if (offset == -1) {
            return;
          }
          context.read<ChatBloc>().add(
                EventFetchPrevMessages(
                  offset: offset,
                  receiver: user,
                ),
              );
        }
      },
    );
  }
}

class _NoMessagesContainer extends StatelessWidget {
  const _NoMessagesContainer({
    required this.offset,
  });

  final int offset;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Center(
        child: Container(
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 12),
          decoration: BoxDecoration(
            color: offset == -1 ? MyColors.primaryColor : Colors.transparent,
            borderRadius: BorderRadius.circular(10),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.2),
                spreadRadius: 1,
                blurRadius: 4,
              ),
            ],
          ),
          child: offset == -1
              ? Text(
                  "No Messages",
                  style: TextStyle(
                    fontSize: 15,
                    fontWeight: FontWeight.bold,
                    color: Colors.grey.shade100,
                  ),
                )
              : const CustomCircularProgressIndicator(),
        ),
      ),
    );
  }
}
