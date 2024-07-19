import 'package:flutter/material.dart';

class RegisterLoginTextButton extends StatelessWidget {
  const RegisterLoginTextButton({
    super.key,
    required this.onTap,
    required this.text,
    required this.buttonText,
  });

  final Function()? onTap;
  final String text;
  final String buttonText;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 5),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            text,
            style: const TextStyle(
              fontSize: 20,
              fontWeight: FontWeight.w600,
            ),
          ),
          const SizedBox(width: 20),
          GestureDetector(
            onTap: onTap,
            child: Text(
              buttonText,
              style: const TextStyle(
                decoration: TextDecoration.underline,
                fontSize: 20,
                fontWeight: FontWeight.w600,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
