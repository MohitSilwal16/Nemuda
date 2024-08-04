import 'package:flutter/material.dart';

import 'package:app/pb/user.pb.dart';
import 'package:app/utils/colors.dart';

class UserMessageCard extends StatelessWidget {
  final UserAndLastMessage usersAndLastMessage;
  final Function navigateToChatPage;
  const UserMessageCard({
    super.key,
    required this.usersAndLastMessage,
    required this.navigateToChatPage,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => navigateToChatPage(usersAndLastMessage.username),
      child: Container(
        padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 10),
        margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
        decoration: BoxDecoration(
          color: MyColors.primaryColor,
          borderRadius: BorderRadius.circular(10),
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Username
            Padding(
              padding: const EdgeInsets.only(bottom: 5),
              child: Text(
                usersAndLastMessage.username,
                style: const TextStyle(
                  fontWeight: FontWeight.bold,
                  fontSize: 20,
                ),
              ),
            ),
            // Message Content and Status
            usersAndLastMessage.lastMessage.messageContent != ""
                ? Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      // Message Content
                      Expanded(
                        child: Padding(
                          padding: const EdgeInsets.only(right: 20),
                          child: Text(
                            usersAndLastMessage.lastMessage.messageContent,
                            style: const TextStyle(
                              fontSize: 18,
                              overflow: TextOverflow.ellipsis,
                            ),
                          ),
                        ),
                      ),
                      // Message Status
                      Text(
                        usersAndLastMessage.lastMessage.receiver ==
                                usersAndLastMessage
                                    .username // I sent the Message so shomme Message Status
                            ? usersAndLastMessage.lastMessage.status
                            : usersAndLastMessage.lastMessage.status != "Read"
                                ? "New Message" // He/She sent the message but I didn't read
                                : "", // He/She sent message & I've read
                        style: TextStyle(
                          fontSize: 14,
                          color: usersAndLastMessage.lastMessage.receiver ==
                                  usersAndLastMessage.username
                              ? Colors.white
                              : Colors.green[400],
                          fontWeight:
                              usersAndLastMessage.lastMessage.receiver ==
                                      usersAndLastMessage.username
                                  ? FontWeight.normal
                                  : FontWeight.bold,
                          height: 1.5,
                        ),
                      ),
                    ],
                  )
                : Text(
                    "No Messages",
                    style: TextStyle(
                      fontSize: 19,
                      color: Colors.red[400],
                      fontWeight: FontWeight.bold,
                    ),
                  ),
          ],
        ),
      ),
    );
  }
}
