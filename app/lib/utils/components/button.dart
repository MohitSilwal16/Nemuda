import 'package:flutter/material.dart';

import 'package:app/utils/colors.dart';

class MyButton extends StatelessWidget {
  const MyButton({
    super.key,
    required this.size,
    required this.text,
    required this.onPressed,
    required this.heightWRTScreen,
    required this.widthWRTScreen,
    required this.fontSize,
  });

  final Size size;
  final String text;
  final void Function()? onPressed;
  final double widthWRTScreen;
  final double heightWRTScreen;
  final double fontSize;

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      style: ButtonStyle(
        fixedSize: WidgetStateProperty.all(
          Size(size.width * widthWRTScreen, size.height * heightWRTScreen),
        ),
        backgroundColor: WidgetStatePropertyAll(MyColors.primaryColor),
        shape: WidgetStateProperty.all<RoundedRectangleBorder>(
          const RoundedRectangleBorder(
            borderRadius: BorderRadius.all(
              Radius.circular(10),
            ),
          ),
        ),
      ),
      onPressed: onPressed,
      child: Text(
        text,
        style: TextStyle(
          color: Colors.white,
          fontWeight: FontWeight.w600,
          fontSize: fontSize,
        ),
      ),
    );
  }
}
