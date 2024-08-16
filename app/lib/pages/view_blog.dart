import 'package:flutter/material.dart';
import 'dart:async';

import 'package:app/main.dart';
import 'package:app/pb/blogs.pb.dart';
import 'package:app/pages/update_blog.dart';
import 'package:app/services/blog.dart';
import 'package:app/utils/colors.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/utils/components/error.dart';
import 'package:app/utils/components/alert_dialogue.dart';
import 'package:app/pages/static/view_blog_skeleton.dart';

class ViewBlogPage extends StatefulWidget {
  const ViewBlogPage({
    super.key,
    required this.blog,
  });

  final Blog blog;

  @override
  State<ViewBlogPage> createState() => _ViewBlogPageState();
}

class _ViewBlogPageState extends State<ViewBlogPage>
    with WidgetsBindingObserver {
  late Blog blog;
  late bool isBlogLiked = false;
  late bool isBlogUpdatableDeletable = false;
  late Future<void> _futureBlog;

  final verticalScrollBar = ScrollController();
  final horizontalScrollBar = ScrollController();
  final controllerComment = TextEditingController();

  bool _isKeyboardVisible = false;

  onDeleteBlog() {
    deleteBlog(blog.title).then((res) {
      ScaffoldMessenger.of(context)
          .showSnackBar(returnSnackbar("Blog Deleted"));
      Navigator.pop(context);
    }).catchError((err) {
      handleErrors(context, err);
    });
  }

  likeDislikeBlog() {
    if (isBlogLiked) {
      dislikeBlog(blog.title).then((_) {
        getBlogByTitle(blog.title).then((res) {
          setState(() {
            blog = res.blog;
            isBlogLiked = res.isBlogLiked;
          });
        }).catchError((err) {
          handleErrors(context, err);
        });
      }).catchError((err) {
        handleErrors(context, err);
      });
      return;
    }

    likeBlog(blog.title).then((_) {
      getBlogByTitle(blog.title).then((res) {
        setState(() {
          blog = res.blog;
          isBlogLiked = res.isBlogLiked;
        });
      }).catchError((err) {
        handleErrors(context, err);
      });
    }).catchError((err) {
      handleErrors(context, err);
    });
  }

  onCommentSubmit() {
    if (controllerComment.text == "") {
      showErrorDialog(context, "Comment is Empty");
      return;
    }
    addComment(blog.title, controllerComment.text).then((_) {
      getBlogByTitle(blog.title).then((res) {
        setState(() {
          blog = res.blog;
          isBlogLiked = res.isBlogLiked;
        });

        controllerComment.text = "";

        ScaffoldMessenger.of(context)
            .showSnackBar(returnSnackbar("Comment Added"));
      }).catchError((err) {
        handleErrors(context, err);
      });
    }).catchError((err) {
      handleErrors(context, err);
    });
    return;
  }

  navigateToUpdateBlogPage() {
    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => UpdateBlogPage(blog: blog),
      ),
    );
  }

  futureFunction() async {
    try {
      final res = await getBlogByTitle(widget.blog.title);
      blog = res.blog;
      isBlogLiked = res.isBlogLiked;
      isBlogUpdatableDeletable = res.isBlogUpdatableDeletable;
    } catch (err) {
      if (mounted) {
        handleErrors(context, err);
      }
    }
  }

  @override
  void initState() {
    blog = widget.blog;
    _futureBlog = futureFunction();
    super.initState();
    WidgetsBinding.instance.addObserver(this);
  }

  @override
  void dispose() {
    WidgetsBinding.instance.removeObserver(this);
    controllerComment.dispose();

    super.dispose();
  }

  @override
  void didChangeMetrics() {
    super.didChangeMetrics();
    final bottomInset = MediaQuery.of(context).viewInsets.bottom;

    final newValue = bottomInset > 0.0;
    if (_isKeyboardVisible != newValue) {
      setState(() {
        _isKeyboardVisible = newValue;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: _futureBlog,
      builder: (context, snapshot) {
        if (snapshot.hasError) {
          handleErrorsFutureBuilder(context, snapshot.error!);
        }
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const ViewBlogSkeletonPage();
        }
        if (snapshot.connectionState == ConnectionState.done) {
          return viewBlogPage(context, size);
        }

        return const ViewBlogSkeletonPage();
      },
    );
  }

  Scaffold viewBlogPage(BuildContext context, Size size) {
    return Scaffold(
      // Like Button
      floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
      floatingActionButton: !_isKeyboardVisible
          ? CircleAvatar(
              backgroundColor: MyColors.primaryColor,
              radius: 30,
              child: IconButton(
                onPressed: likeDislikeBlog,
                icon: Icon(
                  isBlogLiked ? Icons.heart_broken : Icons.favorite,
                  color: Colors.pink,
                  size: 40,
                ),
              ),
            )
          : const SizedBox(),
      body: SafeArea(
        child: Container(
          height: size.height,
          width: size.width,
          decoration: const BoxDecoration(
            image: DecorationImage(
              fit: BoxFit.fill,
              image: AssetImage("assets/view_blog_bg.jpg"),
            ),
          ),
          child: Stack(
            children: [
              SizedBox(
                height: size.height * .85,
                child: ListView(
                  children: [
                    Image(
                      image: NetworkImage(blog.imagePath),
                      fit: BoxFit.fill,
                    ),
                    const SizedBox(height: 10),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 20),
                      child: Text(
                        blog.username,
                        style: const TextStyle(fontSize: 16),
                      ),
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(
                          vertical: 5, horizontal: 20),
                      child: Text(
                        blog.title,
                        style: const TextStyle(
                            fontSize: 27, fontWeight: FontWeight.w800),
                      ),
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(
                          vertical: 5, horizontal: 20),
                      child: Text(
                        overflow: TextOverflow.ellipsis,
                        maxLines: 3,
                        blog.description,
                        style: const TextStyle(fontSize: 18),
                      ),
                    ),
                    const SizedBox(height: 20),

                    // Liked Usernames Section
                    const Padding(
                      padding: EdgeInsets.symmetric(horizontal: 20),
                      child: Text(
                        "Liked by",
                        style: TextStyle(
                            fontSize: 22, fontWeight: FontWeight.bold),
                      ),
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 20),
                      child: Scrollbar(
                        controller: horizontalScrollBar,
                        thumbVisibility: true,
                        child: SingleChildScrollView(
                          controller: horizontalScrollBar,
                          scrollDirection: Axis.horizontal,
                          child: Row(
                            children: blog.likedUsername
                                .map(
                                  (username) => Padding(
                                    padding: const EdgeInsets.symmetric(
                                        horizontal: 7, vertical: 5),
                                    child: Chip(
                                      label: Text(username),
                                    ),
                                  ),
                                )
                                .toList(),
                          ),
                        ),
                      ),
                    ),
                    const SizedBox(height: 20),

                    // Comments Section
                    const Padding(
                      padding: EdgeInsets.symmetric(horizontal: 20),
                      child: Text(
                        "Comments",
                        style: TextStyle(
                            fontSize: 22, fontWeight: FontWeight.bold),
                      ),
                    ),

                    blog.comments.isNotEmpty
                        ? Container(
                            height: size.height * .15,
                            padding: const EdgeInsets.symmetric(horizontal: 20),
                            child: Scrollbar(
                              controller: verticalScrollBar,
                              thumbVisibility: true,
                              child: ListView.builder(
                                controller: verticalScrollBar,
                                itemCount: blog.comments.length,
                                itemBuilder: (context, index) {
                                  return Card(
                                    margin:
                                        const EdgeInsets.symmetric(vertical: 5),
                                    child: Padding(
                                      padding: const EdgeInsets.all(10.0),
                                      child: Column(
                                        crossAxisAlignment:
                                            CrossAxisAlignment.start,
                                        children: [
                                          Text(
                                            blog.comments[index].username,
                                            style: const TextStyle(
                                              fontWeight: FontWeight.bold,
                                              fontSize: 16,
                                            ),
                                          ),
                                          const SizedBox(height: 5),
                                          Text(
                                            blog.comments[index].description,
                                            style:
                                                const TextStyle(fontSize: 16),
                                          ),
                                        ],
                                      ),
                                    ),
                                  );
                                },
                              ),
                            ),
                          )
                        : const SizedBox(),

                    const SizedBox(height: 10),

                    // Add Comment Section
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 20),
                      child: Row(
                        children: [
                          Expanded(
                            child: TextField(
                              controller: controllerComment,
                              decoration: const InputDecoration(
                                hintText: 'Add a comment ...',
                                border: OutlineInputBorder(),
                              ),
                            ),
                          ),
                          IconButton(
                            icon: const Icon(Icons.send),
                            onPressed: onCommentSubmit,
                          ),
                        ],
                      ),
                    ),
                    const SizedBox(height: 20),
                  ],
                ),
              ),

              // Back button
              Positioned(
                top: 17,
                left: 10,
                child: IconButton(
                  onPressed: () => Navigator.pop(context),
                  icon: const Icon(
                    Icons.arrow_back_ios,
                    size: 30,
                    color: Colors.black,
                  ),
                ),
              ),

              // Delete Blog
              isBlogUpdatableDeletable && !_isKeyboardVisible
                  ? Positioned(
                      bottom: 17,
                      right: 20,
                      child: CircleAvatar(
                        backgroundColor: MyColors.primaryColor,
                        radius: 30,
                        child: IconButton(
                          onPressed: onDeleteBlog,
                          icon: const Icon(
                            Icons.delete_forever,
                            size: 40,
                            color: Colors.white,
                          ),
                        ),
                      ),
                    )
                  : const SizedBox(),

              // Update Blog
              isBlogUpdatableDeletable && !_isKeyboardVisible
                  ? Positioned(
                      bottom: 17,
                      left: 20,
                      child: CircleAvatar(
                        backgroundColor: MyColors.primaryColor,
                        radius: 30,
                        child: IconButton(
                          onPressed: navigateToUpdateBlogPage,
                          icon: const Icon(
                            Icons.update,
                            size: 40,
                            color: Colors.white,
                          ),
                        ),
                      ),
                    )
                  : const SizedBox(),
            ],
          ),
        ),
      ),
    );
  }
}
