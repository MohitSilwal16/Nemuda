import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'dart:convert';

import 'package:app/main.dart';
import 'package:app/models/message.dart';
import 'package:app/pb/user.pb.dart';
import 'package:app/pages/chat_user.dart';
import 'package:app/services/user.dart';
import 'package:app/services/auth.dart';
import 'package:app/services/service_init.dart';
import 'package:app/utils/components/alert_dialogue.dart';
import 'package:app/utils/components/error.dart';
import 'package:app/utils/components/show_notification.dart';
import 'package:app/utils/components/loading.dart';
import 'package:app/utils/components/user_card.dart';
import 'package:app/utils/components/button.dart';

class ChatHomePage extends StatefulWidget {
  const ChatHomePage({
    super.key,
    this.channel,
    this.broadcastStream,
  });

  final WebSocketChannel? channel;
  final Stream? broadcastStream;

  @override
  State<ChatHomePage> createState() => _ChatHomePageState();
}

class _ChatHomePageState extends State<ChatHomePage> {
  final controllerSearch = TextEditingController();
  bool isSearchBarClosed = true;
  late List<UserAndLastMessage> usersAndLastMessage;
  late final Future<void> finalFutureFunc;
  late final Stream broadcastStream;

  late WebSocketChannel channel;

  navigateToChatPage(String user) {
    Navigator.pushReplacement(
      context,
      MaterialPageRoute(
        builder: (context) => ChatUserPage(
          user: user,
          channel: channel,
          broadcastStream: broadcastStream,
        ),
      ),
    );
  }

  futureFunction() async {
    final res = await searchUsersByStartingPattern("");
    usersAndLastMessage = res.usersAndLastMessage;
  }

  @override
  void initState() {
    if (widget.channel == null) {
      final sessionToken = ServiceManager().hiveBox.get("sessionToken");
      channel = WebSocketChannel.connect(
          Uri.parse("ws://$serviceURL:$wsPort/ws/chat/$sessionToken"));
      broadcastStream = channel.stream.asBroadcastStream();
    } else {
      channel = widget.channel!;
      broadcastStream = widget.broadcastStream!;
    }

    finalFutureFunc = futureFunction();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: finalFutureFunc,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.done) {
          return stream();
        }

        if (snapshot.hasError) {
          handleErrorsFutureBuilder(context, snapshot.error!);
        }
        // TODO: Skeleton Page
        return const CustomCircularProgressIndicator();

        // TODO: How to refresh build while using Future & Stream Builder
      },
    );
  }

  // Websocket for sending & receiving messages
  StreamBuilder stream() {
    return StreamBuilder(
      stream: broadcastStream,
      builder: (context, snapshot) {
        if (snapshot.hasError) {
          showErrorDialog(context, "Error in WebSocket");
        }
        if (snapshot.hasData) {
          print(snapshot.data);
          Map<String, dynamic> jsonMap = jsonDecode(snapshot.data.toString());
          ModelMessage message = ModelMessage.fromJson(jsonMap);

          if (message.error != "") {
            showErrorDialog(context, message.error);
            return chatHomePage();
          }

          if (message.selfMessage) {
            return chatHomePage();
          }

          if (message.messageType == "Message") {
            // TODO: Edit the design of Notification Box
            showNotificationDialog(
              context,
              message.sender,
              message.messageContent,
              () => navigateToChatPage(message.sender),
            );

            // Update last message if available in list
            for (int i = 0; i < usersAndLastMessage.length; ++i) {
              if (usersAndLastMessage[i].username == message.sender) {
                usersAndLastMessage[i].lastMessage = Message(
                  sender: message.sender,
                  receiver: message.receiver,
                  status: message.status,
                  dateTime: message.dateTime,
                  messageContent: message.messageContent,
                );
                break;
              }
            }
          } else {
            // Update last message status if available in list
            for (int i = 0; i < usersAndLastMessage.length; ++i) {
              if (usersAndLastMessage[i].username == message.sender) {
                usersAndLastMessage[i].lastMessage.status = message.messageType;
                break;
              }
            }
          }
        }

        return chatHomePage();
      },
    );
  }

  PopScope chatHomePage() {
    final size = MediaQuery.of(context).size;
    return PopScope(
      canPop: false,
      onPopInvoked: (didPop) {
        if (didPop) {
          return;
        }
        Navigator.pushReplacementNamed(context, "home");
      },
      child: Scaffold(
        appBar: AppBar(
          title: isSearchBarClosed
              ? const Text(
                  "Nemu Chat",
                  style: TextStyle(fontWeight: FontWeight.w700, fontSize: 24),
                )
              : searchTextField(),
          // Back Button
          leading: IconButton(
            onPressed: () {
              if (isSearchBarClosed) {
                Navigator.pushReplacementNamed(context, "home");
              } else {
                setState(() {
                  isSearchBarClosed = true;
                });
              }
            },
            icon: const Icon(Icons.arrow_back_ios_new),
          ),
          actions: [
            // Search Users
            isSearchBarClosed
                ? IconButton(
                    onPressed: () {
                      setState(() {
                        isSearchBarClosed = false;
                      });
                    },
                    icon: const Icon(Icons.search),
                  )
                : Padding(
                    padding: const EdgeInsets.only(bottom: 7),
                    child: IconButton(
                        icon: const Icon(Icons.close),
                        onPressed: () {
                          setState(() {
                            isSearchBarClosed = true;
                          });
                        }),
                  ),
            const SizedBox(width: 10),

            // Log out Button
            isSearchBarClosed
                ? MyButton(
                    size: size,
                    text: "Logout",
                    onPressed: () {
                      logout().then((res) {
                        ServiceManager().hiveBox.delete("sessionToken");
                        Navigator.pushReplacementNamed(context, "login");
                      });
                    },
                    widthWRTScreen: .26,
                    heightWRTScreen: .05,
                    fontSize: 16,
                  )
                : const SizedBox(),

            const SizedBox(width: 15),
          ],
        ),
        body: Container(
          width: size.width,
          height: size.height,
          padding: const EdgeInsets.only(bottom: 10),
          decoration: const BoxDecoration(
            image: DecorationImage(
              image: AssetImage("assets/home_bg.jpg"),
              fit: BoxFit.cover,
            ),
          ),
          child: ListView(
            children: List.generate(
              usersAndLastMessage.length,
              (index) => UserMessageCard(
                usersAndLastMessage: usersAndLastMessage[index],
                navigateToChatPage: navigateToChatPage,
              ),
            ),
          ),
        ),
      ),
    );
  }

  TextField searchTextField() {
    return TextField(
      style: const TextStyle(
        fontSize: 22,
        fontWeight: FontWeight.bold,
      ),
      autofocus: true,
      decoration: const InputDecoration(
        hintStyle: TextStyle(
          fontSize: 22,
          fontWeight: FontWeight.bold,
        ),
        counter: SizedBox(),
        hintText: "Search Users",
        border: InputBorder.none,
      ),
      maxLength: 20,
      controller: controllerSearch,
      onChanged: (val) {
        searchUsersByStartingPattern(val).then((res) {
          setState(() {
            usersAndLastMessage = res.usersAndLastMessage;
          });
        }).catchError((err) {
          handleErrors(context, err);
        });
      },
    );
  }
}
