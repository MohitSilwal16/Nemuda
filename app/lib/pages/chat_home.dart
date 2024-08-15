import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'package:app/main.dart';
import 'package:app/bloc/chat_repo.dart';
import 'package:app/pb/user.pb.dart';
import 'package:app/utils/components/show_notification.dart';
import 'package:app/bloc/chat_bloc.dart';
import 'package:app/bloc/chat_event.dart';
import 'package:app/bloc/chat_state.dart';
import 'package:app/services/auth.dart';
import 'package:app/services/service_init.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/components/error.dart';
import 'package:app/utils/components/user_card.dart';
import 'package:app/pages/chat_user.dart';
import 'package:app/pages/static/chat_home_skeleton.dart';

class ChatHomePage extends StatelessWidget {
  const ChatHomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const _BuildAppBar(),
      body: PopScope(
        canPop: false,
        onPopInvoked: (didPop) {
          if (didPop) {
            return;
          }
          ChatRepo().dispose();
          Navigator.pushReplacementNamed(context, "home");
        },
        child: Container(
          width: size.width,
          height: size.height,
          padding: const EdgeInsets.only(bottom: 10),
          decoration: const BoxDecoration(
            image: DecorationImage(
              image: AssetImage("assets/home_bg.jpg"),
              fit: BoxFit.cover,
            ),
          ),
          child: const _BlocBuilder(),
        ),
      ),
    );
  }
}

class _BlocBuilder extends StatefulWidget {
  const _BlocBuilder();

  @override
  State<_BlocBuilder> createState() => __BlocBuilderState();
}

class __BlocBuilderState extends State<_BlocBuilder> {
  List<UserAndLastMessage> usersAndLastMsg = [];

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
    context.read<ChatBloc>().add(EventSearchUser(searchPattern: ""));
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<ChatBloc, ChatState>(
      buildWhen: (previous, current) {
        // Only rebuild if the current state is not StateNothing and is different from the previous state.
        return current.runtimeType != StateNothing && previous != current;
        // return current is! StateNothing && previous != current;
      },
      builder: (context, state) {
        final currentState = state;
        if (currentState is StateChatLoading) {
          return const ChatHomeSkeletonPage();
        }

        if (currentState is StateChatError) {
          handleErrorsBlocBuilder(context, currentState.errorMessage);
        }

        if (currentState is StateUserLoaded) {
          usersAndLastMsg = currentState.usersAndLastMsg;
        }

        if (currentState is StateNewMsgReceived) {
          if (currentState.message.messageType == "Message") {
            for (int i = 0; i < usersAndLastMsg.length; ++i) {
              if (usersAndLastMsg[i].username == currentState.message.sender) {
                usersAndLastMsg[i].lastMessage = Message(
                  sender: currentState.message.sender,
                  receiver: currentState.message.receiver,
                  messageContent: currentState.message.messageContent,
                  dateTime: currentState.message.dateTime,
                  status: currentState.message.status,
                );
                break;
              }
            }
            showNotificationDialog(
              context,
              currentState.message.sender,
              currentState.message.messageContent,
              () => navigateToChatPage(currentState.message.sender),
            );
          } else {
            for (int i = 0; i < usersAndLastMsg.length; ++i) {
              if (usersAndLastMsg[i].username == currentState.message.sender) {
                usersAndLastMsg[i].lastMessage.status =
                    currentState.message.messageType;
                break;
              }
            }
          }
        }

        return ListView(
          children: List.generate(
            usersAndLastMsg.length,
            (index) => UserMessageCard(
              usersAndLastMessage: usersAndLastMsg[index],
              navigateToChatPage: navigateToChatPage,
            ),
          ),
        );
      },
    );
  }
}

class _BuildAppBar extends StatefulWidget implements PreferredSizeWidget {
  const _BuildAppBar();

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  State<_BuildAppBar> createState() => _BuildAppBarState();
}

class _BuildAppBarState extends State<_BuildAppBar> {
  bool isSearchBarClosed = true;
  final controllerSearch = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return AppBar(
      title: isSearchBarClosed
          ? const Text(
              "Nemu Chat",
              style: TextStyle(fontWeight: FontWeight.w700, fontSize: 24),
            )
          : _SearchTextWidget(controllerSearch: controllerSearch),
      // Back Button
      leading: IconButton(
        onPressed: () {
          if (isSearchBarClosed) {
            ChatRepo().dispose();
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
    );
  }
}

class _SearchTextWidget extends StatelessWidget {
  const _SearchTextWidget({
    required this.controllerSearch,
  });

  final TextEditingController controllerSearch;

  @override
  Widget build(BuildContext context) {
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
        counterText: "",
        hintText: "Search Users",
        border: InputBorder.none,
      ),
      maxLength: 20,
      controller: controllerSearch,
      onChanged: (val) {
        context.read<ChatBloc>().add(EventSearchUser(searchPattern: val));
      },
    );
  }
}
