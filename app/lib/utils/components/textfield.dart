import 'package:flutter/material.dart';

class MyTextField extends StatefulWidget {
  const MyTextField({
    super.key,
    required this.hintText,
    required this.obscureText,
    this.validator,
    required this.controller,
    this.keyboardType,
    this.suffixIconData,
    this.maxLength = 20,
    this.maxLines = 1,
  });

  final TextEditingController controller;
  final String hintText;
  final bool obscureText;
  final String? Function(String?)? validator;
  final TextInputType? keyboardType;
  final IconData? suffixIconData;
  final int maxLength;
  final int maxLines;

  @override
  State<MyTextField> createState() => _MyTextFieldState();
}

class _MyTextFieldState extends State<MyTextField> {
  bool isHidden = true;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 5),
      child: TextFormField(
        controller: widget.controller,
        decoration: InputDecoration(
          border: const OutlineInputBorder(
            borderRadius: BorderRadius.all(Radius.circular(10)),
          ),
          hintText: widget.hintText,
          hintStyle: const TextStyle(color: Colors.white),
          counterStyle: const TextStyle(color: Colors.white),
          filled: true,
          fillColor: Colors.black,
          suffixIcon: IconButton(
            onPressed: () {
              setState(() {
                isHidden = !isHidden;
              });
            },
            icon: Icon(
              widget.obscureText && isHidden
                  ? Icons.visibility
                  : widget.obscureText && !isHidden
                      ? Icons.visibility_off
                      : widget.suffixIconData,
            ),
          ),
        ),
        keyboardType: widget.keyboardType,
        maxLength: widget.maxLength,
        maxLines: widget.maxLines,
        validator: widget.validator,
        obscureText: widget.obscureText && isHidden,
      ),
    );
  }
}
