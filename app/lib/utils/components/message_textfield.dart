import 'package:flutter/material.dart';

class MyMessageTextField extends StatefulWidget {
  const MyMessageTextField({
    super.key,
    required this.controller,
    required this.sendMessage,
    required this.focusNode,
  });

  final TextEditingController controller;
  final Function sendMessage;
  final FocusNode focusNode;

  @override
  State<MyMessageTextField> createState() => _MyMessageTextFieldState();
}

class _MyMessageTextFieldState extends State<MyMessageTextField> {
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 10),
      child: TextField(
        focusNode: widget.focusNode,
        controller: widget.controller,
        buildCounter: (context,
            {required currentLength, required isFocused, maxLength}) {
          return Container(
            transform: Matrix4.translationValues(
                0, -110, 0), // Shift Counter Text to Top
            child: Text("$currentLength/$maxLength"),
          );
        },
        decoration: InputDecoration(
          border: const OutlineInputBorder(
            borderRadius: BorderRadius.all(Radius.circular(10)),
          ),
          hintText: "Enter Message ...",
          hintStyle: const TextStyle(color: Colors.white),
          counterStyle: const TextStyle(color: Colors.white),
          filled: true,
          fillColor: Colors.black,
          suffixIcon: IconButton(
            onPressed: () => widget.sendMessage(widget.controller.text),
            icon: const Icon(Icons.send),
          ),
        ),
        keyboardType: TextInputType.text,
        maxLength: 50,
        maxLines: 2,
        textInputAction: TextInputAction.none,
      ),
    );
  }
}
