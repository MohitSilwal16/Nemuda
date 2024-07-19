import 'package:flutter/material.dart';

SnackBar returnSnackbar(String message) {
  return SnackBar(
    content: Text(
      message,
      style: const TextStyle(
        fontWeight: FontWeight.w600,
        fontSize: 16,
        color: Colors.white,
      ),
    ),
    dismissDirection: DismissDirection.down,
    animation: kAlwaysCompleteAnimation,
    duration: const Duration(
      seconds: 2,
    ),
    elevation: 10,
    backgroundColor: Colors.blue,
    behavior: SnackBarBehavior.floating,
  );
}
