import 'package:flutter/material.dart';

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

  navigateToChatPage(String user) {
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => ChatUserPage(user: user),
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
    super.initState();
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
          return LoginPage();
        }
        return chatHomePage(context);
      },
    );
  }

  Scaffold chatHomePage(BuildContext context) {
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
