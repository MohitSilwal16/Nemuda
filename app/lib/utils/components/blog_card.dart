import 'package:app/pages/view_blog.dart';
import 'package:app/pb/blogs.pb.dart';
import 'package:flutter/material.dart';

class BlogCard extends StatelessWidget {
  const BlogCard({
    super.key,
    required this.blog,
    required this.size,
  });

  final Size size;
  final Blog blog;
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () => Navigator.push(
        context,
        MaterialPageRoute(
          builder: (context) => ViewBlogPage(title: blog.title),
        ),
      ),
      child: Container(
        decoration: BoxDecoration(
          border: Border.all(color: Colors.white, width: 1),
          borderRadius: const BorderRadius.all(Radius.circular(10)),
        ),
        margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 10),
        padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment
              .start, // Aligns the text properly with the image
          children: [
            Image(
              image: NetworkImage(blog.imagePath),
              height: size.height * .13,
              width: size.width * .33,
              fit: BoxFit.fill,
            ),
            const SizedBox(width: 20),
            Expanded(
              // Allows the text column to take up remaining space
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      const Text(
                        "Username: ",
                        style: TextStyle(
                            fontWeight: FontWeight.w600, fontSize: 15),
                      ),
                      Text(
                        blog.username,
                        style: const TextStyle(fontSize: 15),
                      ),
                    ],
                  ),
                  Row(
                    children: [
                      const Text(
                        "Title: ",
                        style: TextStyle(
                            fontWeight: FontWeight.w600, fontSize: 15),
                      ),
                      // Expanded widget to handle overflow in description
                      Expanded(
                        child: Text(
                          blog.title,
                          overflow: TextOverflow.ellipsis,
                          maxLines: 1,
                          style: const TextStyle(fontSize: 15),
                        ),
                      ),
                    ],
                  ),
                  Row(
                    children: [
                      const Text(
                        "Tag: ",
                        style: TextStyle(
                            fontWeight: FontWeight.w600, fontSize: 15),
                      ),
                      Text(
                        blog.tag,
                        style: const TextStyle(fontSize: 15),
                      ),
                    ],
                  ),
                  Row(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const Text(
                        "Desc: ",
                        style: TextStyle(
                            fontWeight: FontWeight.w600, fontSize: 15),
                      ),
                      // Expanded widget to handle overflow in description
                      Expanded(
                        child: Text(
                          blog.description,
                          overflow: TextOverflow.ellipsis,
                          maxLines:
                              1, // Adjust this if you want more than one line before ellipses
                          style: const TextStyle(fontSize: 15),
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),

            // End
          ],
        ),
      ),
    );
  }
}
