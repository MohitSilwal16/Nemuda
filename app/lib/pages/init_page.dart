import 'package:flutter/material.dart';
import 'package:hive/hive.dart';
import 'dart:async';

import 'package:app/services/auth.dart';
import 'package:app/pages/home.dart';
import 'package:app/pages/register_login.dart';
import 'package:app/utils/components/error.dart';
import 'package:app/pages/static/splash_screen.dart';

class InitPage extends StatelessWidget {
  const InitPage({super.key});

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          final isUserValidated = snapshot.data;
          if (isUserValidated == null || isUserValidated == false) {
            return const RegisterLoginPage();
          }
          return const HomePage();
        }
        if (snapshot.hasError) {
          handleErrorsFutureBuilder(context, snapshot.error.toString());
        }
        return const SplashScreen();
      },
      future: Future(
        () async {
          final sessionToken = Hive.box("session").get("sessionToken");
          if (sessionToken == null) {
            return false;
          }
          final res = await validateSessionToken();
          if (res.isUserVerified) {
            return true;
          }
          return false;
        },
      ),
    );
  }
}
