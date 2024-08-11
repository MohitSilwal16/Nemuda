import 'package:flutter/material.dart';

class NotificationDialog extends StatelessWidget {
  final String title;
  final String body;
  final Function onOpen;

  const NotificationDialog({
    super.key,
    required this.title,
    required this.body,
    required this.onOpen,
  });

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      backgroundColor: Colors.blueAccent,
      title: Text(
        title,
        style: const TextStyle(
          color: Colors.white,
          fontWeight: FontWeight.bold,
          fontSize: 18,
        ),
      ),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Text(
            body,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 16,
            ),
          ),
          const SizedBox(height: 20),
          const Icon(
            Icons.notifications,
            color: Colors.white,
            size: 48,
          ),
        ],
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: const Text(
            'Close',
            style: TextStyle(
              color: Colors.white,
            ),
          ),
        ),
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
            onOpen();
          },
          child: const Text(
            'Open',
            style: TextStyle(
              color: Colors.white,
            ),
          ),
        ),
      ],
    );
  }
}

void showNotificationDialog(
    BuildContext context, String title, String body, Function onTap) {
  WidgetsBinding.instance.addPostFrameCallback((_) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return NotificationDialog(
          title: title,
          body: body,
          onOpen: onTap,
        );
      },
    );
  });
}
