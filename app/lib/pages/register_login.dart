import 'package:app/pages/login.dart';
import 'package:app/pages/register.dart';
import 'package:flutter/material.dart';

class RegisterLoginPage extends StatefulWidget {
  const RegisterLoginPage({super.key});

  @override
  State<RegisterLoginPage> createState() => _RegisterLoginPageState();
}

class _RegisterLoginPageState extends State<RegisterLoginPage> {
  bool isLoginPageOpen = true;

  togglePage() {
    setState(() {
      isLoginPageOpen = !isLoginPageOpen;
    });
  }

  @override
  Widget build(BuildContext context) {
    return isLoginPageOpen
        ? LoginPage(togglePage: togglePage)
        : RegisterPage(togglePage: togglePage);
  }
}
