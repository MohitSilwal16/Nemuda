import 'package:app/pb/user.pb.dart';
import 'package:app/utils/colors.dart';
import 'package:flutter/material.dart';

class MessageCard extends StatelessWidget {
  final Message message;
  final String user;

  const MessageCard({
    super.key,
    required this.user,
    required this.message,
  });

  @override
  Widget build(BuildContext context) {
    bool isSelfMessage = message.sender == message.receiver;
    bool isMessageSentByUser = message.sender == user;

    return Container(
      margin: const EdgeInsets.symmetric(vertical: 10, horizontal: 5),
      alignment:
          isMessageSentByUser && !isSelfMessage ? Alignment.centerLeft : Alignment.centerRight,
      child: Container(
        constraints:
            BoxConstraints(maxWidth: MediaQuery.of(context).size.width * 0.5),
        padding: const EdgeInsets.all(8),
        decoration: BoxDecoration(
          color: MyColors.primaryColor,
          borderRadius: BorderRadius.circular(8),
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              message.dateTime,
              style: const TextStyle(
                color: Colors.white70,
                fontSize: 12,
              ),
            ),
            const SizedBox(height: 4),
            Text(
              message.messageContent,
              style: const TextStyle(
                color: Colors.white,
                fontWeight: FontWeight.bold,
                fontSize: 16,
              ),
            ),
            if (isSelfMessage || !isMessageSentByUser)
              Align(
                alignment: Alignment.bottomRight,
                child: Text(
                  message.status,
                  style: const TextStyle(
                    color: Colors.white70,
                    fontSize: 12,
                  ),
                ),
              ),
          ],
        ),
      ),
    );
  }
}
