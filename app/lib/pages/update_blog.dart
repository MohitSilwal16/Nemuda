import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:async';
import 'dart:io';

import 'package:app/main.dart';
import 'package:app/services/blog.dart';
import 'package:app/pb/blogs.pb.dart';
import 'package:app/utils/components/loading.dart';
import 'package:app/utils/components/alert_dialogue.dart';
import 'package:app/utils/components/error.dart';
import 'package:app/utils/components/title_textfield.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/components/textfield.dart';
import 'package:app/utils/utils.dart';
import 'package:app/utils/validator.dart';

class UpdateBlogPage extends StatefulWidget {
  const UpdateBlogPage({
    super.key,
    required this.blog,
  });

  final Blog blog;

  @override
  State<UpdateBlogPage> createState() => _UpdateBlogPageState();
}

class _UpdateBlogPageState extends State<UpdateBlogPage> {
  final controllerTitle = TextEditingController();
  final controllerDescription = TextEditingController();
  final tags = tagsListPostUpdateBlog;
  final formKey = GlobalKey<FormState>();

  String selectedTag = "Political";
  File? selectedFile;

  Future pickImageFromGallery() async {
    try {
      const maxFileSizeInBytes = 2 * 1024 * 1024; // 2MB
      final image = await ImagePicker().pickImage(source: ImageSource.gallery);
      var imagePath = await image!.readAsBytes();

      var fileSize = imagePath.length;
      if (fileSize > maxFileSizeInBytes) {
        if (mounted) {
          showErrorDialog(context, "Image should not exceed 2 MB");
        }
        return;
      }

      setState(() {
        selectedFile = File(image.path);
      });
    } catch (e) {
      // No need to catch
    }
  }

  onSubmit() {
    if (!formKey.currentState!.validate()) {
      return;
    }
    if (selectedFile == null) {
      showErrorDialog(context, "Please Select an Image");
      return;
    }
    showDialog(
      context: context,
      builder: (context) => const CustomCircularProgressIndicator(),
    );

    fileToBytes(selectedFile!).then((bytesImage) {
      updateBlog(
        widget.blog.title,
        controllerTitle.text,
        controllerDescription.text,
        selectedTag,
        bytesImage,
      ).then((res) {
        ScaffoldMessenger.of(context)
            .showSnackBar(returnSnackbar("Blog Updated"));

        Navigator.of(context).pushNamedAndRemoveUntil(
          "home",
          (Route<dynamic> route) => false, // Predicate to remove all routes
        );
      }).catchError((err) {
        final trimmedGRPCError = trimGrpcErrorMessage(err.toString());

        if (trimmedGRPCError == "IMAGE SIZE EXCEEDS 2 MB") {
          showErrorDialog(context, "Image Size Exceeds 2 MB");
        } else if (trimmedGRPCError == "USER CANNOT UPDATE THIS BLOG") {
          showErrorDialog(context, "User Cannot Update this Blog");
        } else {
          handleErrors(context, err);
        }
      });
    }).catchError((err) {
      showErrorDialog(context, "Failed to Convert Images into Bytes");
    });
  }

  @override
  void initState() {
    controllerTitle.text = widget.blog.title;
    controllerDescription.text = widget.blog.description;
    selectedTag = widget.blog.tag;

    super.initState();
  }

  @override
  void dispose() {
    controllerTitle.dispose();
    controllerDescription.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Container(
          width: size.width,
          height: size.height,
          decoration: const BoxDecoration(
            image: DecorationImage(
              image: AssetImage("assets/background.jpg"),
              fit: BoxFit.cover,
            ),
          ),
          child: CustomScrollView(
            slivers: [
              SliverFillRemaining(
                hasScrollBody: false,
                child: Form(
                  key: formKey,
                  autovalidateMode: AutovalidateMode.onUserInteraction,
                  child: Column(
                    children: [
                      SizedBox(height: size.height * .03),

                      // Back Button & Update Blog Title
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 20),
                        child: Row(
                          children: [
                            // Back Button
                            IconButton(
                              onPressed: () => Navigator.pop(context),
                              icon: const Icon(
                                Icons.arrow_back_ios_new,
                                weight: 100,
                                size: 30,
                              ),
                            ),

                            SizedBox(width: size.width * .14),
                            const Text(
                              "Update Blog",
                              style: TextStyle(
                                fontSize: 30,
                                fontWeight: FontWeight.w800,
                              ),
                            ),
                          ],
                        ),
                      ),

                      SizedBox(height: size.height * .05),

                      // Title
                      MyBlogTitleTextField(
                        controller: controllerTitle,
                        errorText: "Title is already used",
                      ),

                      // Description
                      MyTextField(
                        hintText: "Description",
                        obscureText: false,
                        validator: Validators.validateDescription,
                        controller: controllerDescription,
                        keyboardType: TextInputType.multiline,
                        maxLength: 50,
                        maxLines: 2,
                      ),

                      // Tag
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 35),
                        child: DropdownButtonFormField<String>(
                          decoration: InputDecoration(
                            filled: true,
                            fillColor: Colors.black, // Background color
                            border: OutlineInputBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          style: const TextStyle(
                            color: Colors.white, // Text color
                            fontWeight: FontWeight.w700,
                            fontSize: 20,
                          ),
                          value: selectedTag,
                          items: List.generate(
                            tags.length,
                            (index) => DropdownMenuItem<String>(
                              value: tags[index],
                              child: Text(tags[index]),
                            ),
                          ),
                          onChanged: (val) {
                            setState(() {
                              selectedTag = val!;
                            });
                          },
                        ),
                      ),

                      const SizedBox(height: 40),

                      Row(
                        mainAxisAlignment: MainAxisAlignment.spaceAround,
                        children: [
                          // Image
                          ClipRRect(
                            borderRadius: const BorderRadius.all(
                              Radius.circular(20),
                            ),
                            child: Image.network(
                              widget.blog.imagePath,
                              width: size.width * .45,
                            ),
                          ),

                          // Image Picker
                          selectedFile == null
                              ? GestureDetector(
                                  onTap: pickImageFromGallery,
                                  child: Icon(
                                    Icons.photo,
                                    size: size.width * .45,
                                  ),
                                )
                              : GestureDetector(
                                  onTap: pickImageFromGallery,
                                  child: ClipRRect(
                                    borderRadius: const BorderRadius.all(
                                      Radius.circular(20),
                                    ),
                                    child: Image.file(
                                      selectedFile!,
                                      width: size.width * .45,
                                    ),
                                  ),
                                ),
                        ],
                      ),

                      const Spacer(),

                      // Post Blog Button
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 25),
                        child: MyButton(
                          size: size,
                          text: "Update",
                          onPressed: onSubmit,
                          heightWRTScreen: .07,
                          widthWRTScreen: .9,
                          fontSize: 22,
                        ),
                      ),

                      const SizedBox(height: 20),

                      // END
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
