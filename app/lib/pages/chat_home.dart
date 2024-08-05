import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'dart:convert';

import 'package:app/main.dart';
import 'package:app/models/message.dart';
import 'package:app/pb/user.pb.dart';
import 'package:app/pages/login.dart';
import 'package:app/pages/chat_user.dart';
import 'package:app/services/user.dart';
import 'package:app/services/auth.dart';
import 'package:app/services/service_init.dart';
import 'package:app/utils/components/loading.dart';
import 'package:app/utils/components/user_card.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/utils/utils.dart';

class ChatHomePage extends StatefulWidget {
  const ChatHomePage({super.key});

  @override
  State<ChatHomePage> createState() => _ChatHomePageState();
}

class _ChatHomePageState extends State<ChatHomePage> {
  final controllerSearch = TextEditingController();
  bool isSearchBarClosed = true;
  late List<UserAndLastMessage> usersAndLastMessage;
  late Future<void> finalFutureFunc;

  late WebSocketChannel channel;

  navigateToChatPage(String user) {
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => ChatUserPage(
          user: user,
          channel: channel,
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
    finalFutureFunc = futureFunction();

    final sessionToken = Clients().hiveBox.get("sessionToken");
    channel = WebSocketChannel.connect(
        Uri.parse("ws://$serviceURL:$wsPort/ws/chat/$sessionToken"));
    super.initState();
  }

  @override
  void dispose() {
    // TODO: Ensure that Websocket is closed & check when dispose() is called
    controllerSearch.dispose();
    channel.sink.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: finalFutureFunc,
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const CustomCircularProgressIndicator();
        }

        if (snapshot.hasError) {
          // Navigator.of(context)
          //   ..pop()
          //   ..pop()
          //   ..pop();
          return LoginPage();
        }
        return stream();
      },
    );
  }

  // Websocket for sending & receiving messages
  StreamBuilder stream() {
    return StreamBuilder(
      stream: channel.stream,
      builder: (context, snapshot) {
        if (snapshot.hasError) {
          return LoginPage();
        }
        if (snapshot.hasData) {
          print(snapshot.data);

          Map<String, dynamic> jsonMap = jsonDecode(snapshot.data.toString());
          ModelMessage message = ModelMessage.fromJson(jsonMap);

          if (message.messageType != "Message") {
            return chatHomePage();
          }

          return chatHomePageWithNotification(
              context, message.sender, message.messageContent);
        }
        return chatHomePage();
      },
    );
  }

  Scaffold chatHomePage() {
    final size = MediaQuery.of(context).size;
    return Scaffold(
      appBar: AppBar(
        title: isSearchBarClosed
            ? const Text(
                "Nemu Chat",
                style: TextStyle(fontWeight: FontWeight.w700, fontSize: 24),
              )
            : searchTextField(context),
        // Back Button
        leading: IconButton(
          onPressed: () {
            if (isSearchBarClosed) {
              Navigator.pop(context);
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
                      Clients().hiveBox.delete("sessionToken");
                      Navigator.pushReplacementNamed(context, "login");
                    }).catchError((err) {
                      final trimmedGrpcError =
                          trimGrpcErrorMessage(err.toString());

                      ScaffoldMessenger.of(context)
                          .showSnackBar(returnSnackbar(trimmedGrpcError));
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
    );
  }

  Scaffold chatHomePageWithNotification(
      BuildContext context, String notificationTitle, String notificationBody) {
    final size = MediaQuery.of(context).size;
    bool showNotification = true;

    return Scaffold(
      appBar: AppBar(
        title: isSearchBarClosed
            ? const Text(
                "Nemu Chat",
                style: TextStyle(fontWeight: FontWeight.w700, fontSize: 24),
              )
            : searchTextField(context),
        // Back Button
        leading: IconButton(
          onPressed: () {
            if (isSearchBarClosed) {
              Navigator.pop(context);
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
                      Clients().hiveBox.delete("sessionToken");
                      Navigator.pushReplacementNamed(context, "login");
                    }).catchError((err) {
                      final trimmedGrpcError =
                          trimGrpcErrorMessage(err.toString());

                      ScaffoldMessenger.of(context)
                          .showSnackBar(returnSnackbar(trimmedGrpcError));
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
      body: Stack(
        children: [
          Container(
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
          Positioned(
            bottom: 26,
            right: 10,
            child: AnimatedOpacity(
              opacity: 1.0,
              duration: const Duration(milliseconds: 300),
              child: Visibility(
                visible: showNotification,
                child: GestureDetector(
                  onTap: () {
                    setState(() {
                      showNotification = false;
                    });
                  },
                  child: Container(
                    width: size.width * .6,
                    padding: const EdgeInsets.symmetric(
                        vertical: 15, horizontal: 20),
                    decoration: BoxDecoration(
                      color: Colors.blue,
                      borderRadius: BorderRadius.circular(10),
                      boxShadow: const [
                        BoxShadow(
                          color: Colors.black26,
                          blurRadius: 10,
                          offset: Offset(0, 4),
                        ),
                      ],
                    ),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          notificationTitle,
                          style: const TextStyle(
                            fontSize: 20,
                            fontWeight: FontWeight.bold,
                            color: Colors.white,
                          ),
                        ),
                        const SizedBox(height: 10),
                        Text(
                          notificationBody,
                          style: const TextStyle(
                              color: Colors.white, fontSize: 16),
                        ),
                        const SizedBox(height: 10),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            ElevatedButton(
                              onPressed: () {
                                showNotification = false;
                                Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => ChatUserPage(
                                      user: notificationTitle,
                                      channel: channel,
                                    ),
                                  ),
                                );
                              },
                              style: ElevatedButton.styleFrom(
                                backgroundColor: Colors.black,
                                shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(10),
                                ),
                              ),
                              child: const Text('Open'),
                            ),
                            ElevatedButton(
                              onPressed: () {
                                setState(() {
                                  showNotification = false;
                                });
                              },
                              style: ElevatedButton.styleFrom(
                                backgroundColor: Colors.black,
                                shape: RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(10),
                                ),
                              ),
                              child: const Text('Close'),
                            ),
                          ],
                        ),
                      ],
                    ),
                  ),
                ),
              ),
            ),
          ),

          // END
        ],
      ),
    );
  }

  TextField searchTextField(BuildContext context) {
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
          final trimmedGrpcError = trimGrpcErrorMessage(err.toString());

          ScaffoldMessenger.of(context)
              .showSnackBar(returnSnackbar(trimmedGrpcError));

          if (trimmedGrpcError == "INVALID SESSION TOKEN") {
            Navigator.pushReplacementNamed(context, "login");
          }
        });
      },
    );
  }
}
