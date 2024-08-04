import 'package:flutter/material.dart';

class MyMessageTextField extends StatefulWidget {
  const MyMessageTextField({
    super.key,
    required this.controller,
  });

  final TextEditingController controller;

  @override
  State<MyMessageTextField> createState() => _MyMessageTextFieldState();
}

class _MyMessageTextFieldState extends State<MyMessageTextField> {
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 10),
      child: TextFormField(
        controller: widget.controller,
        buildCounter: (context,
            {required currentLength, required isFocused, maxLength}) {
          return Container(
            transform: Matrix4.translationValues(0, -110, 0),
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
            onPressed: () {},
            icon: const Icon(Icons.send),
          ),
        ),
        keyboardType: TextInputType.multiline,
        maxLength: 50,
        maxLines: 2,
      ),
    );
  }
}
