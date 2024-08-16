import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:hive/hive.dart';
import 'package:flutter/services.dart';
import 'package:path_provider/path_provider.dart';

import 'package:app/bloc/chat_bloc.dart';
import 'package:app/bloc/chat_repo.dart';
import 'package:app/services/service_init.dart';
import 'package:app/pages/register_login.dart';
import 'package:app/pages/chat_home.dart';
import 'package:app/pages/static/server_error.dart';
import 'package:app/pages/static/server_busy_page.dart';
import 'package:app/pages/init_page.dart';
import 'package:app/pages/post_blog.dart';
import 'package:app/pages/home.dart';

// const serviceURL = "nemuda.hopto.org";
const serviceURL = "13.60.113.45";
const servicePort = 8080;
const wsPort = 3000;
late Size size;

void main() async {
  // Init Hive
  WidgetsFlutterBinding.ensureInitialized();
  final documentsDir = await getApplicationDocumentsDirectory();
  Hive.init(documentsDir.path);
  await Hive.openBox("session");

  // Init GRPC Clients & Hivebox
  await ServiceManager().init();

  runApp(
    MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (context) => ChatBloc(repo: ChatRepo()),
        )
      ],
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    size = MediaQuery.of(context).size;
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
        "login": (context) => const RegisterLoginPage(),
        "post_blog": (context) => const PostBlogPage(),
        "chat_home": (context) => const ChatHomePage(),
        "server_error": (context) => const ServerErrorPage(),
        "server_busy": (context) => const ServerBusyPage(),
      },
      home: const InitPage(),
    );
  }
}
