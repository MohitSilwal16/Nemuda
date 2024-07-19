import 'package:flutter/material.dart';
import 'package:double_back_to_close_app/double_back_to_close_app.dart';

import 'package:app/utils/components/snackbar.dart';
import 'package:app/pages/login.dart';
import 'package:app/utils/components/welcome_to_nemuda.dart';
import 'package:app/utils/components/register_login_text_button.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/components/textfield.dart';
import 'package:app/utils/validator.dart';
import 'package:app/utils/size.dart';

class RegisterPage extends StatelessWidget {
  RegisterPage({super.key});

  final formKey = GlobalKey<FormState>();
  final controllerUsername = TextEditingController();
  final controllerPassword = TextEditingController();

  register() {
    // TODO: Add Registration Logic

    if (formKey.currentState!.validate()) {
      print(controllerUsername.text);
      print(controllerPassword.text);
      print("Form Validated");
    }
  }

  redirectToLoginPage(BuildContext context) {
    Navigator.pushReplacement(
      context,
      MaterialPageRoute(
        builder: (context) => LoginPage(),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    final size = returnSize(context);

    return Scaffold(
      body: DoubleBackToCloseApp(
        snackBar: returnSnackbar("Tap Again to Exit"),
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
        
                    // Register Text
                    const Text(
                      "Register",
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
                      validator: Validators.validateUsername,
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
        
                    // Register Button
                    MyButton(
                      text: "Register",
                      onPressed: register,
                      size: size,
                    ),
        
                    SizedBox(height: size.height * .05),
        
                    // Redirect to Register Page
                    RegisterLoginTextButton(
                      text: "Have an Account ?",
                      buttonText: "Login",
                      onTap:()=> redirectToLoginPage(context),
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
