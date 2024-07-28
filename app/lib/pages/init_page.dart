import 'package:flutter/material.dart';
import 'package:hive/hive.dart';

import 'package:app/services/auth.dart';
import 'package:app/pages/home.dart';
import 'package:app/pages/login.dart';
import 'package:app/utils/components/loading.dart';

class InitPage extends StatelessWidget {
  const InitPage({super.key});

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const CustomCircularProgressIndicator();
        }
        final data = snapshot.data;
        if (data == null) {
          return LoginPage();
        }
        if (data) {
          return const HomePage();
        }
        return LoginPage();
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
