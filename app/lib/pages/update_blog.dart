import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

import 'package:app/services/blog.dart';
import 'package:app/pb/blogs.pb.dart';
import 'package:app/utils/components/title_textfield.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/components/textfield.dart';
import 'package:app/utils/utils.dart';
import 'package:app/utils/validator.dart';
import 'package:app/utils/size.dart';

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
    final image = await ImagePicker().pickImage(source: ImageSource.gallery);
    setState(() {
      selectedFile = File(image!.path);
    });
  }

  onSubmit() {
    if (!formKey.currentState!.validate()) {
      return;
    }
    if (selectedFile == null) {
      ScaffoldMessenger.of(context)
          .showSnackBar(returnSnackbar("Please Select an Image"));
      return;
    }

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

        Navigator.of(context)
          ..pop()
          ..pop();
      }).catchError((err) {
        final trimmedGrpcError = trimGrpcErrorMessage(err.toString());
        ScaffoldMessenger.of(context)
            .showSnackBar(returnSnackbar(trimmedGrpcError));

        if (trimmedGrpcError == "INVALID SESSION TOKEN") {
          Navigator.pushReplacementNamed(context, "login");
        }
      });
    }).catchError((err) {
      ScaffoldMessenger.of(context)
          .showSnackBar(returnSnackbar("Failed to Convert Images into Bytes"));
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
  Widget build(BuildContext context) {
    final size = returnSize(context);

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
          child: Form(
            key: formKey,
            autovalidateMode: AutovalidateMode.onUserInteraction,
            child: SingleChildScrollView(
              child: Column(
                children: [
                  SizedBox(height: size.height * .02),

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

                  const SizedBox(height: 30),

                  SizedBox(
                    height: size.height * .06,
                  ),

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

                  // END
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
