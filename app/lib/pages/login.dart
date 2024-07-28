import 'package:flutter/material.dart';
import 'package:hive/hive.dart';
import 'package:double_back_to_close_app/double_back_to_close_app.dart';

import 'package:app/services/auth.dart';
import 'package:app/utils/utils.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/utils/components/welcome_to_nemuda.dart';
import 'package:app/utils/components/register_login_text_button.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/components/textfield.dart';
import 'package:app/utils/validator.dart';
import 'package:app/utils/size.dart';

class LoginPage extends StatelessWidget {
  LoginPage({super.key});

  final formKey = GlobalKey<FormState>();
  final controllerUsername = TextEditingController();
  final controllerPassword = TextEditingController();

  onSubmit(BuildContext context) {
    if (!formKey.currentState!.validate()) {
      return;
    }

    final response = login(controllerUsername.text, controllerPassword.text);

    response.then((responseData) {
      Hive.box("session").put("sessionToken", responseData.sessionToken);

      Navigator.pushReplacementNamed(context, "home");
    }).catchError((error) {
      ScaffoldMessenger.of(context)
          .showSnackBar(returnSnackbar(trimGrpcErrorMessage(error.toString())));
    });
  }

  redirectToRegisterPage(BuildContext context) {
    Navigator.pushReplacementNamed(context, "register");
  }

  @override
  Widget build(BuildContext context) {
    final size = returnSize(context);

    return Scaffold(
      body: DoubleBackToCloseApp(
        snackBar: returnSnackbar("Tag Again to Exit"),
        child: SafeArea(
          child: Container(
            width: size.width,
            height: size.height,
            decoration: const BoxDecoration(
              image: DecorationImage(
                image: AssetImage("assets/background.jpg"),
                fit: BoxFit.cover,
              ),
            ),
            child: SingleChildScrollView(
              child: Form(
                autovalidateMode: AutovalidateMode.onUserInteraction,
                key: formKey,
                child: Column(
                  children: [
                    SizedBox(
                      height: size.height * .1,
                    ),

                    // Welcome to Nemu 2.0
                    const WelcomeToNemudaText(),

                    SizedBox(
                      height: size.height * .05,
                    ),

                    // Login Text
                    const Text(
                      "Login",
                      style: TextStyle(
                        fontSize: 25,
                        fontWeight: FontWeight.w600,
                      ),
                    ),

                    SizedBox(
                      height: size.height * .05,
                    ),

                    // Username textfield
                    MyTextField(
                      controller: controllerUsername,
                      hintText: "Username",
                      obscureText: false,
                      validator: (val) =>
                          Validators.validateUsername(val, null),
                      keyboardType: TextInputType.name,
                      suffixIconData: Icons.person,
                    ),

                    // // Password Textfield
                    MyTextField(
                      hintText: "Password",
                      obscureText: true,
                      validator: Validators.validatePassword,
                      controller: controllerPassword,
                      keyboardType: TextInputType.visiblePassword,
                      suffixIconData: Icons.visibility,
                    ),

                    SizedBox(height: size.height * .05),

                    // Login Button
                    MyButton(
                      text: "Login",
                      onPressed: () => onSubmit(context),
                      size: size,
                      heightWRTScreen: .07,
                      widthWRTScreen: .85,
                      fontSize: 18,
                    ),

                    SizedBox(height: size.height * .05),

                    // Redirect to Register Page
                    RegisterLoginTextButton(
                      text: "New to Nemuda ?",
                      buttonText: "Register",
                      onTap: () => redirectToRegisterPage(context),
                    ),

                    // End
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
