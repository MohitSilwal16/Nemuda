import 'package:flutter/material.dart';
import 'dart:async';

class WelcomeToNemudaText extends StatefulWidget {
  const WelcomeToNemudaText({
    super.key,
  });

  @override
  State<WelcomeToNemudaText> createState() => _WelcomeToNemudaTextState();
}

class _WelcomeToNemudaTextState extends State<WelcomeToNemudaText> {
  String displayText = "";
  String fullText = "Welcome to Nemu 2.0";
  String replacementText = "Nemuda";
  int index = 0;
  Timer? timer;
  bool isFullTextDisplayed = false;
  bool isRemoving = false;

  @override
  void initState() {
    super.initState();
    startTimer();
  }

  void startTimer() {
    timer = Timer.periodic(const Duration(milliseconds: 100), (Timer t) {
      if (!isFullTextDisplayed) {
        if (index < fullText.length) {
          setState(() {
            displayText += fullText[index];
            index++;
          });
        } else {
          isFullTextDisplayed = true;
          index = fullText.length;
        }
      } else if (isFullTextDisplayed && !isRemoving) {
        if (index > 11) {
          setState(() {
            displayText = displayText.substring(0, index - 1);
            index--;
          });
        } else {
          isRemoving = true;
          index = 11;
        }
      } else if (isRemoving) {
        if (index - 11 < replacementText.length) {
          setState(() {
            displayText = displayText.substring(0, 11) +
                replacementText.substring(0, index - 11 + 1);
            index++;
          });
        } else {
          timer?.cancel();
        }
      }
    });
  }

  @override
  void dispose() {
    timer?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Text(
      displayText,
      style: const TextStyle(
        fontSize: 30,
        fontWeight: FontWeight.w800,
      ),
    );
  }
}
