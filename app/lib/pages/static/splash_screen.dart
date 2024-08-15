import 'package:flutter/material.dart';

import 'package:app/utils/colors.dart';
import 'package:app/utils/components/loading.dart';

class SplashScreen extends StatelessWidget {
  const SplashScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: MyColors.primaryColor,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Image.asset(
              'assets/icon.png',
              width: 150,
            ),
            const SizedBox(height: 20),
            const Text(
              'Welcome to Nemuda',
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
                color: Colors.white,
              ),
            ),
            const SizedBox(height: 30),
            const CustomCircularProgressIndicator(),
          ],
        ),
      ),
    );
  }
}
