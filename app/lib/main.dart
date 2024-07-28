import 'package:flutter/material.dart';
import 'package:hive/hive.dart';
import 'package:flutter/services.dart';
import 'package:path_provider/path_provider.dart';

import 'package:app/pages/init_page.dart';
import 'package:app/services/service_init.dart';
import 'package:app/pages/home.dart';
import 'package:app/pages/login.dart';
import 'package:app/pages/register.dart';

const serviceURL = "nemuda.hopto.org";
const servicePort = 8080;

void main() async {
  // Init Hive
  WidgetsFlutterBinding.ensureInitialized();
  final documentsDir = await getApplicationDocumentsDirectory();
  Hive.init(documentsDir.path);
  await Hive.openBox("session");
  // .then((box) {
  //   // box.delete("sessionToken");
  //   print(box.get("sessionToken"));
  // });

  // Init GRPC Clients & Hivebox
  await Clients().init();

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
      routes: {
        "home": (context) => const HomePage(),
        "login": (context) => LoginPage(),
        "register": (context) => RegisterPage(),
      },
      home: const InitPage(),
    );
  }
}
