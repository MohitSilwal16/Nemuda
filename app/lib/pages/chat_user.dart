import 'package:app/utils/components/show_notification.dart';
import 'package:flutter/material.dart';
import 'package:visibility_detector/visibility_detector.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'dart:convert';

import 'package:app/pb/user.pb.dart';
import 'package:app/models/message.dart';
import 'package:app/pages/chat_home.dart';
import 'package:app/pages/login.dart';
import 'package:app/services/service_init.dart';
import 'package:app/services/user.dart';
import 'package:app/utils/components/alert_dialogue.dart';
import 'package:app/utils/colors.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/utils/utils.dart';
import 'package:app/utils/components/loading.dart';
import 'package:app/utils/components/message_textfield.dart';
import 'package:app/utils/components/message_card.dart';

class ChatUserPage extends StatefulWidget {
  const ChatUserPage({
    super.key,
    required this.user,
    required this.channel,
    required this.broadcastStream,
  });

  final String user;
  final WebSocketChannel channel;
  final Stream broadcastStream;

  @override
  State<ChatUserPage> createState() => _ChatUserPageState();
}

class _ChatUserPageState extends State<ChatUserPage> {
  final controllerMessage = TextEditingController();
  final ScrollController controllerScroll = ScrollController();
  late final Future<void> finalFutureFunc;
  late final String sessionToken;

  late List<Message> messages;
  late int offset;

  futureFunction() async {
    final res = await getMessages(widget.user, 0);
    messages = res.messages;
    offset = res.nextOffset;
  }

  loadMoreMessages(int index, Size size) {
    if (index == 0) {
      return VisibilityDetector(
        key: const Key("Load-More-Messages"),
        child: noMoreMessagesContainer(),
        onVisibilityChanged: (VisibilityInfo info) {
          if (info.visibleFraction > 0) {
            if (offset == -1) {
              return;
            }

            getMessages(widget.user, offset).then((res) {
              setState(() {
                final temp = messages;
                messages = res.messages;
                messages += temp;
                offset = res.nextOffset;
              });
            }).catchError((err) {
              if (err.toString() ==
                  "TimeoutException after 0:00:05.000000: Future not completed") {
                ScaffoldMessenger.of(context)
                    .showSnackBar(returnSnackbar("Request Timed Out"));
                return;
              }

              final trimmedGrpcError = trimGrpcErrorMessage(err.toString());

              ScaffoldMessenger.of(context)
                  .showSnackBar(returnSnackbar(trimmedGrpcError));

              if (trimmedGrpcError == "INVALID SESSION TOKEN") {
                Navigator.pushReplacementNamed(context, "login");
              }
            });
          }
        },
      );
    }

    return MessageCard(
      message: messages[index - 1],
      user: widget.user,
    );
  }

  sendMessage(String message) {
    if (message == "") {
      return;
    }
    controllerMessage.text = "";

    widget.channel.sink.add(
      jsonEncode(
        ModelWSMessage(
          message: message.trim(),
          receiver: widget.user,
          sessionToken: sessionToken,
          messageType: "Message",
        ).toJson(),
      ),
    );
  }

  navigateToChatPage(String user) {
    Navigator.pushReplacement(
      context,
      MaterialPageRoute(
        builder: (context) => ChatUserPage(
          user: user,
          channel: widget.channel,
          broadcastStream: widget.broadcastStream,
        ),
      ),
    );
  }

  @override
  void initState() {
    finalFutureFunc = futureFunction();

    sessionToken = ServiceManager().hiveBox.get("sessionToken");
    widget.channel.sink.add(
      jsonEncode(
        ModelWSMessage(
          message: "",
          receiver: widget.user,
          sessionToken: sessionToken,
          messageType: "Read",
        ).toJson(),
      ),
    );

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: finalFutureFunc,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          // TODO: Skeleton Page
          return const CustomCircularProgressIndicator();
        }

        if (snapshot.hasError) {
          return LoginPage();
        }

        if (messages.length <= 9) {
          // Initial message loading
          // Scroll to the bottom
          WidgetsBinding.instance.addPostFrameCallback((_) {
            controllerScroll.jumpTo(controllerScroll.position.maxScrollExtent);
          });
        } else {
          // Scroll a bit downwards
          WidgetsBinding.instance.addPostFrameCallback((_) {
            controllerScroll.jumpTo(70);
          });
        }
        return chatUserPage();
      },
    );
  }

  StreamBuilder stream() {
    return StreamBuilder(
      stream: widget.broadcastStream,
      builder: (context, snapshot) {
        print("Build");
        if (snapshot.hasData) {
          print("Build 1");
          Map<String, dynamic> jsonMap = jsonDecode(snapshot.data.toString());
          ModelMessage message = ModelMessage.fromJson(jsonMap);

          if (message.error != "") {
            showErrorDialog(context, message.error);
          }

          if (message.sender == widget.user ||
              (message.selfMessage && widget.user == message.receiver)) {
            if (message.messageType == "Message") {
              // Scroll to the bottom
              WidgetsBinding.instance.addPostFrameCallback((_) {
                controllerScroll
                    .jumpTo(controllerScroll.position.maxScrollExtent);
              });

              messages.add(
                Message(
                  sender: message.sender,
                  receiver: message.receiver,
                  messageContent: message.messageContent,
                  dateTime: message.dateTime,
                  status: message.status,
                ),
              );
              print(messages.length);

              // Retain only the last 18 messages
              int keepLastN = 18;
              if (messages.length > keepLastN) {
                messages = messages.sublist(messages.length - keepLastN);
              }

              // Acknowledge user that I read his/her message
              if (!message.selfMessage) {
                widget.channel.sink.add(
                  jsonEncode(
                    ModelWSMessage(
                            message: "",
                            receiver: message.sender,
                            sessionToken: sessionToken,
                            messageType: "Read")
                        .toJson(),
                  ),
                );
              }
            } else if (message.messageType == "Read") {
              // Update message status if available in list
              for (int i = messages.length - 1; i >= 0; --i) {
                if (messages[i].status == "Read") {
                  break;
                }
                messages[i].status = "Read";
              }
            } else if (message.messageType == "Delivered") {
              for (int i = messages.length - 1; i >= 0; --i) {
                if (messages[i].status == "Read" ||
                    messages[i].status == "Delivered") {
                  break;
                }
                messages[i].status = "Delivered";
              }
            }
          } else {
            // The message is from some another user
            if (message.messageType == "Message") {
              showNotificationDialog(
                  context,
                  message.sender,
                  message.messageContent,
                  () => navigateToChatPage(message.sender));
            }
          }
        }

        if (snapshot.hasError) {
          showErrorDialog(context, "Error in WebSocket");
        }

        return chatUserPage();
      },
    );
  }

  PopScope chatUserPage() {
    final size = MediaQuery.of(context).size;

    return PopScope(
      canPop: false,
      onPopInvoked: (didPop) {
        if (didPop) {
          return;
        }
        Navigator.pushReplacement(
          context,
          MaterialPageRoute(
            builder: (context) => ChatHomePage(
              channel: widget.channel,
              broadcastStream: widget.broadcastStream,
            ),
          ),
        );
      },
      child: Scaffold(
        appBar: AppBar(
          leading: CircleAvatar(
            backgroundColor: Colors.transparent,
            child: IconButton(
              onPressed: () => Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                  builder: (context) => ChatHomePage(
                    channel: widget.channel,
                    broadcastStream: widget.broadcastStream,
                  ),
                ),
              ),
              icon: const Icon(Icons.arrow_back_ios_new),
            ),
          ),
          title: Text(
            widget.user,
            style: const TextStyle(fontWeight: FontWeight.w700, fontSize: 24),
          ),
          centerTitle: true,
        ),
        body: Container(
          width: size.width,
          height: size.height,
          padding: const EdgeInsets.symmetric(vertical: 20),
          decoration: const BoxDecoration(
            image: DecorationImage(
              image: AssetImage("assets/chat_bg.jpg"),
              fit: BoxFit.cover,
            ),
          ),
          child: ListView(
            children: [
              SizedBox(
                height: size.height * .7,
                child: messages.isNotEmpty
                    ? ListView(
                        controller: controllerScroll,
                        children: List.generate(
                          messages.length + 1,
                          (index) => loadMoreMessages(index, size),
                        ),
                      )
                    : noMoreMessagesContainer(),
              ),

              SizedBox(height: size.height * .06),

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

  Padding noMoreMessagesContainer() {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Center(
        child: Container(
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 12),
          decoration: BoxDecoration(
            color: MyColors.primaryColor,
            borderRadius: BorderRadius.circular(10),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.2),
                spreadRadius: 1,
                blurRadius: 4,
              ),
            ],
          ),
          child: Text(
            offset == -1 ? "No Messages" : "Loading ...",
            style: TextStyle(
              fontSize: 15,
              fontWeight: FontWeight.bold,
              color: Colors.grey.shade100,
            ),
          ),
        ),
      ),
    );
  }
}
