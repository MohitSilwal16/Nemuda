import 'package:flutter/material.dart';
import 'package:double_back_to_close_app/double_back_to_close_app.dart';
import 'package:visibility_detector/visibility_detector.dart';

import 'package:app/pb/blogs.pb.dart';
import 'package:app/utils/colors.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/size.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/services/auth.dart';
import 'package:app/services/blog.dart';
import 'package:app/services/service_init.dart';
import 'package:app/utils/components/blog_card.dart';
import 'package:app/utils/utils.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  late List<Blog> blogs = [];
  late int offset = 0;
  var title = "All";

  final tags = [
    "All",
    "Political",
    "Technical",
    "Educational",
    "Geographical",
    "Programming",
    "Other"
  ];

  @override
  void initState() {
    getBlogsWithPagination("All", 0).then((res) {
      setState(() {
        blogs = res.blogs;
        offset = res.nextOffset;
      });
    }).catchError((err) {
      final trimmedGrpcError = trimGrpcErrorMessage(err.toString());

      ScaffoldMessenger.of(context)
          .showSnackBar(returnSnackbar(trimmedGrpcError));

      if (trimmedGrpcError == "INVALID SESSION TOKEN") {
        Navigator.pushReplacementNamed(context, "login");
      }
    });

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final size = returnSize(context);
    return Scaffold(
      appBar: AppBar(
        centerTitle: true,
        title: Text(title),
        leading: Builder(builder: (context) {
          return IconButton(
            onPressed: () {
              Scaffold.of(context).openDrawer();
            },
            icon: const Icon(Icons.menu),
          );
        }),
        actions: [
          // Log out Button
          MyButton(
            size: size,
            text: "Logout",
            onPressed: () {
              logout().then((res) {
                Clients().hiveBox.delete("sessionToken");
                Navigator.pushReplacementNamed(context, "login");
              }).catchError((err) {
                final trimmedGrpcError = trimGrpcErrorMessage(err.toString());

                ScaffoldMessenger.of(context)
                    .showSnackBar(returnSnackbar(trimmedGrpcError));
              });
            },
            widthWRTScreen: .26,
            heightWRTScreen: .05,
            fontSize: 16,
          ),

          const SizedBox(width: 10),
        ],
      ),
      floatingActionButton: CircleAvatar(
        backgroundColor: MyColors.primaryColor,
        radius: 35,
        child: IconButton(
          onPressed: () {},
          icon: const Icon(
            Icons.post_add_rounded,
            size: 45,
            color: Colors.white,
          ),
        ),
      ),
      drawer: appDrawer(context),
      body: DoubleBackToCloseApp(
        snackBar: returnSnackbar("Tag Again to Exit"),
        child: Container(
          width: size.width,
          height: size.height,
          decoration: const BoxDecoration(
            image: DecorationImage(
              image: AssetImage("assets/background-home.jpg"),
              fit: BoxFit.cover,
            ),
          ),
          child: SingleChildScrollView(
            physics: const BouncingScrollPhysics(),
            child: blogs.isNotEmpty
                ? Column(
                    children: List.generate(
                      blogs.length + 1,
                      (index) {
                        if (index == blogs.length) {
                          return VisibilityDetector(
                            key: const Key("load-more"),
                            child: noMoreBlogsContainer(),
                            onVisibilityChanged: (VisibilityInfo info) {
                              if (info.visibleFraction > 0) {
                                if (offset == -1) {
                                  return;
                                }
                                getBlogsWithPagination(title, offset)
                                    .then((res) {
                                  setState(() {
                                    blogs += res.blogs;
                                    offset = res.nextOffset;
                                  });
                                }).catchError((err) {
                                  final trimmedGrpcError =
                                      trimGrpcErrorMessage(err.toString());

                                  ScaffoldMessenger.of(context).showSnackBar(
                                      returnSnackbar(trimmedGrpcError));

                                  if (trimmedGrpcError ==
                                      "INVALID SESSION TOKEN") {
                                    Navigator.pushReplacementNamed(
                                        context, "login");
                                  }
                                });
                              }
                            },
                          );
                        }

                        return BlogCard(
                          size: size,
                          imagePath: blogs[index].imagePath,
                          description: blogs[index].description,
                          tag: blogs[index].tag,
                          title: blogs[index].title,
                          username: blogs[index].username,
                        );
                      },
                    ),
                  )
                : Padding(
                    padding: const EdgeInsets.symmetric(vertical: 40),
                    child: Center(
                      child: Text(
                        "No Blogs for $title tag 😢",
                        style: const TextStyle(
                            fontSize: 20, fontWeight: FontWeight.w600),
                      ),
                    ),
                  ),
          ),
        ),
      ),
    );
  }

  getBlogsByTagForNavbar(String tag) {
    getBlogsWithPagination(tag, 0).then((res) {
      setState(() {
        blogs = res.blogs;
        offset = res.nextOffset;
        title = tag;
      });

      Navigator.pop(context);
    }).catchError((err) {
      final trimmedGrpcError = trimGrpcErrorMessage(err.toString());

      ScaffoldMessenger.of(context)
          .showSnackBar(returnSnackbar(trimmedGrpcError));

      if (trimmedGrpcError == "INVALID SESSION TOKEN") {
        Navigator.pushReplacementNamed(context, "login");
      }
    });
  }

  Drawer appDrawer(BuildContext context) {
    return Drawer(
      child: Container(
        decoration: const BoxDecoration(
          image: DecorationImage(
            image: AssetImage("assets/background-home.jpg"),
            fit: BoxFit.fill,
          ),
        ),
        child: ListView(
          padding: EdgeInsets.zero,
          children: [
            DrawerHeader(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  IconButton(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    icon: const Icon(Icons.arrow_back_ios),
                  ),

                  // Message Icon
                  IconButton(
                    onPressed: () {},
                    icon: const Icon(Icons.message_outlined),
                  ),
                ],
              ),
            ),

            // All
            ListTile(
              title: Text(
                "All",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 20,
                  decoration: title == "All"
                      ? TextDecoration.underline
                      : TextDecoration.none,
                ),
              ),
              onTap: () => getBlogsByTagForNavbar("All"),
            ),

            // Political
            ListTile(
              title: Text(
                "Political",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 20,
                  decoration: title == "Political"
                      ? TextDecoration.underline
                      : TextDecoration.none,
                ),
              ),
              onTap: () => getBlogsByTagForNavbar("Political"),
            ),

            // Technical
            ListTile(
              title: Text(
                "Technical",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 20,
                  decoration: title == "Technical"
                      ? TextDecoration.underline
                      : TextDecoration.none,
                ),
              ),
              onTap: () => getBlogsByTagForNavbar("Technical"),
            ),

            // Educational
            ListTile(
              title: Text(
                "Educational",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 20,
                  decoration: title == "Educational"
                      ? TextDecoration.underline
                      : TextDecoration.none,
                ),
              ),
              onTap: () => getBlogsByTagForNavbar("Educational"),
            ),

            // Geographical
            ListTile(
              title: Text(
                "Geographical",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 20,
                  decoration: title == "Geographical"
                      ? TextDecoration.underline
                      : TextDecoration.none,
                ),
              ),
              onTap: () => getBlogsByTagForNavbar("Geographical"),
            ),

            // Programming
            ListTile(
              title: Text(
                "Programming",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 20,
                  decoration: title == "Programming"
                      ? TextDecoration.underline
                      : TextDecoration.none,
                ),
              ),
              onTap: () => getBlogsByTagForNavbar("Programming"),
            ),

            // Other
            ListTile(
              title: Text(
                "Other",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 20,
                  decoration: title == "Other"
                      ? TextDecoration.underline
                      : TextDecoration.none,
                ),
              ),
              onTap: () => getBlogsByTagForNavbar("Other"),
            ),

            // End
          ],
        ),
      ),
    );
  }

  Padding noMoreBlogsContainer() {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Center(
        child: Container(
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 12),
          decoration: BoxDecoration(
            color: Colors.blue.shade600,
            borderRadius: BorderRadius.circular(10),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.2),
                spreadRadius: 1,
                blurRadius: 4,
              ),
            ],
          ),
          child: Text(
            'No more Blogs',
            style: TextStyle(
              fontSize: 15,
              fontWeight: FontWeight.bold,
              color: Colors.grey.shade100,
            ),
          ),
        ),
      ),
    );
  }
}