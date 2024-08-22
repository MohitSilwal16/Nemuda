import 'package:flutter/material.dart';
import 'package:audioplayers/audioplayers.dart';

_notiSnackBar(
    {required String title,
    required String body,
    required Function onOpen,
    required BuildContext context}) {
  return SnackBar(
    showCloseIcon: true,
    closeIconColor: Colors.white,
    action: SnackBarAction(
      label: "Open",
      textColor: Colors.white,
      onPressed: () => onOpen(),
    ),
    content: Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          title,
          style: const TextStyle(
            fontWeight: FontWeight.w800,
            fontSize: 18,
            color: Colors.white,
          ),
        ),
        Text(
          body,
          maxLines: 1,
          overflow: TextOverflow.ellipsis,
          style: const TextStyle(
            fontWeight: FontWeight.w600,
            fontSize: 16,
            color: Colors.white,
          ),
        ),
      ],
    ),
    dismissDirection: DismissDirection.up,
    animation: kAlwaysCompleteAnimation,
    duration: const Duration(seconds: 2),
    elevation: 10,
    backgroundColor: Colors.blue,
    behavior: SnackBarBehavior.floating,
    margin: EdgeInsets.only(
      top: 10,
      left: 10,
      right: 10,
      bottom: MediaQuery.of(context).size.height * .8,
    ),
  );
}

void showNotificationDialog(
    BuildContext context, String title, String body, Function onTap) {
  AudioPlayer().play(AssetSource("notification.mp3"));
  WidgetsBinding.instance.addPostFrameCallback((_) {
    final snackBar = _notiSnackBar(
        title: title, body: body, onOpen: () => onTap(), context: context);

    final scaffoldMessenger = ScaffoldMessenger.of(context);
    scaffoldMessenger.hideCurrentSnackBar(); // Hide any existing SnackBar
    scaffoldMessenger.showSnackBar(snackBar);
  });
}
