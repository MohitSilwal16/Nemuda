import 'package:flutter/material.dart';

import 'package:app/services/blog.dart';
import 'package:app/utils/validator.dart';

class MyBlogTitleTextField extends StatefulWidget {
  const MyBlogTitleTextField({
    super.key,
    required this.controller,
    this.errorText,
  });

  final TextEditingController controller;
  final String? errorText;

  @override
  State<MyBlogTitleTextField> createState() => _MyBlogTitleTextFieldState();
}

class _MyBlogTitleTextFieldState extends State<MyBlogTitleTextField> {
  String? errorText;

  @override
  void initState() {
    errorText = widget.errorText;
    super.initState();
  }

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
          errorText: errorText,
          hintText: "Title",
          hintStyle: const TextStyle(color: Colors.white),
          counterStyle: const TextStyle(color: Colors.white),
          filled: true,
          fillColor: Colors.black,
        ),
        keyboardType: TextInputType.name,
        maxLength: 20,
        validator: (val) => Validators.validateTitle(val, errorText),
        onChanged: (value) async {
          final res = await searchBlog(value);
          setState(() {
            if (res.doesBlogExists) {
              errorText = "Title is already used";
            } else {
              errorText = "";
            }
          });
        },
      ),
    );
  }
}
