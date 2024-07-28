import 'package:app/services/user.dart';
import 'package:app/utils/validator.dart';
import 'package:flutter/material.dart';

class MyUsernameTextField extends StatefulWidget {
  const MyUsernameTextField({
    super.key,
    required this.controller,
  });

  final TextEditingController controller;

  @override
  State<MyUsernameTextField> createState() => _MyUsernameTextFieldState();
}

class _MyUsernameTextFieldState extends State<MyUsernameTextField> {
  String? errorText;

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
          hintText: "Username",
          hintStyle: const TextStyle(color: Colors.white),
          counterStyle: const TextStyle(color: Colors.white),
          filled: true,
          fillColor: Colors.black,
          suffixIcon: const Icon(Icons.person),
        ),
        keyboardType: TextInputType.name,
        maxLength: 20,
        validator: (val) => Validators.validateUsername(val, errorText),
        onChanged: (value) async {
          final res = await doesUserExists(value);
          setState(() {
            if (res.doesUserExists) {
              errorText = "Username is already used";
            } else {
              errorText = "";
            }
          });
        },
      ),
    );
  }
}
