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
  final ScrollController controllerScroll = ScrollController();
  FocusNode focusNode = FocusNode();

  void sendMessage(String message) {
    context
        .read<ChatBloc>()
        .add(EventSendMessage(message: message, receiver: widget.user));
    controllerMessage.clear();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      controllerScroll.jumpTo(controllerScroll.position.maxScrollExtent);
    });
  }

  @override
  void initState() {
    super.initState();
    focusNode.addListener(() {
      if (focusNode.hasFocus) {
        Future.delayed(
          const Duration(milliseconds: 500),
          () => controllerScroll.animateTo(
            controllerScroll.position.maxScrollExtent,
            duration: const Duration(seconds: 1),
            curve: Curves.fastOutSlowIn,
          ),
        );
      }
    });
  }

  @override
  void dispose() {
    controllerMessage.dispose();
    controllerScroll.dispose();
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
                  child: _BuildMessages(
                    user: widget.user,
                    controllerScroll: controllerScroll,
                  ),
                ),
              ),

              SizedBox(height: size.height * .035),

              MyMessageTextField(
                controller: controllerMessage,
                sendMessage: sendMessage,
                focusNode: focusNode,
                user: widget.user,
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
  const _BuildMessages({
    required this.user,
    required this.controllerScroll,
  });

  final String user;
  final ScrollController controllerScroll;

  @override
  State<_BuildMessages> createState() => _BuildMessagesState();
}

class _BuildMessagesState extends State<_BuildMessages> {
  int offset = 0;
  List<Message> messages = [];

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
    super.initState();
    context.read<ChatBloc>().add(EventMarkMsgAsRead(receiver: widget.user));
    context
        .read<ChatBloc>()
        .add(EventFetchPrevMessages(offset: 0, receiver: widget.user));
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: size.height * .7,
      child: BlocBuilder<ChatBloc, ChatState>(
        builder: (context, state) {
          final currentState = state;
          if (currentState is StateNothing) {
            return messageBody();
          }
          if (currentState is StateChatLoading ||
              currentState is StateUserLoaded) {
            return const CustomCircularProgressIndicator();
          }

          if (currentState is StateChatError) {
            handleErrorsBlocBuilder(context, currentState.exception);
            context.read<ChatBloc>().add(EventNothing());
            return messageBody();
          }

          if (currentState is StateChatLoaded) {
            messages = currentState.messages + messages;
            offset = currentState.nextOffset;

            // Scroll to the bottom
            if (messages.length <= 9) {
              WidgetsBinding.instance.addPostFrameCallback((_) {
                widget.controllerScroll
                    .jumpTo(widget.controllerScroll.position.maxScrollExtent);
              });
            } else {
              // Scroll lil bit downwards when old message is loaded
              // Else it's gonna request for older messages again & again
              WidgetsBinding.instance.addPostFrameCallback((_) {
                widget.controllerScroll.jumpTo(70);
              });
            }
            context.read<ChatBloc>().add(EventNothing());
            return messageBody();
          }

          // New Msg
          // TODO: Handle bug when 1st msg is sent
          // TODO: Self Messages're not displayed
          if (currentState is StateNewMsgReceived) {
            context.read<ChatBloc>().add(EventNothing());
            if (currentState.message.error != "") {
              handleErrorsBlocBuilder(context, currentState.message.error);
              return messageBody();
            }

            if (currentState.message.sender != widget.user &&
                !currentState.message.selfMessage) {
              if (currentState.message.messageType == "Message") {
                showNotificationDialog(
                    context,
                    currentState.message.sender,
                    currentState.message.messageContent,
                    () => navigateToChatPage(currentState.message.sender));
              }
              return messageBody();
            }

            if (messages.isEmpty) {
              messages.add(
                Message(
                  sender: currentState.message.sender,
                  receiver: currentState.message.receiver,
                  messageContent: currentState.message.messageContent,
                  dateTime: currentState.message.dateTime,
                  status: currentState.message.status,
                ),
              );
              return messageBody();
            }

            if (currentState.message.messageType == "Read") {
              for (int i = messages.length - 1; i >= 0; --i) {
                if (messages[i].status == "Read") break;
                messages[i].status = "Read";
              }
              return messageBody();
            }

            if (currentState.message.messageType == "Delivered") {
              for (int i = messages.length - 1; i >= 0; --i) {
                if (messages[i].status == "Read" ||
                    messages[i].status == "Delivered") break;
                messages[i].status = "Delivered";
              }
              return messageBody();
            }

            // Remove 17 messages if msg list size is more than 27
            if (messages.length > 27) {
              messages = messages.sublist(messages.length - 17);
              offset = 17;
            }
            offset += 1;

            messages.add(
              Message(
                sender: currentState.message.sender,
                receiver: currentState.message.receiver,
                messageContent: currentState.message.messageContent,
                dateTime: currentState.message.dateTime,
                status: currentState.message.status,
              ),
            );
            WidgetsBinding.instance.addPostFrameCallback((_) {
              widget.controllerScroll
                  .jumpTo(widget.controllerScroll.position.maxScrollExtent);
            });

            // It's not Self Message so we need to send Ack
            if (!currentState.message.selfMessage) {
              context
                  .read<ChatBloc>()
                  .add(EventMarkMsgAsRead(receiver: widget.user));
            }

            return messageBody();
          }
          return const Text("Nemu Chat");
        },
      ),
    );
  }

  ListView messageBody() {
    return ListView.builder(
      physics: const BouncingScrollPhysics(),
      itemCount: messages.length + 1,
      controller: widget.controllerScroll,
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
