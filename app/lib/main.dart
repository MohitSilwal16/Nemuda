import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:app/pages/login.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    SystemChrome.setSystemUIOverlayStyle(const SystemUiOverlayStyle(
      statusBarColor: Colors.black, // Change this to your desired color
      statusBarIconBrightness: Brightness.light, // This makes the icons white
    ));

    return MaterialApp(
      theme: ThemeData(
        brightness: Brightness.dark,
        fontFamily: 'Roboto',
        textTheme: const TextTheme(
          bodyLarge: TextStyle(color: Colors.white), // Normal text style
          bodyMedium: TextStyle(color: Colors.white), // Normal text style
          titleLarge: TextStyle(color: Colors.white), // Text style for headers
        ),
      ),
      debugShowCheckedModeBanner: false,
      home: LoginPage(),
    );
  }
}
