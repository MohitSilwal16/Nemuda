import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

import 'package:app/services/blog.dart';
import 'package:app/utils/components/error.dart';
import 'package:app/utils/components/alert_dialogue.dart';
import 'package:app/utils/components/title_textfield.dart';
import 'package:app/utils/components/snackbar.dart';
import 'package:app/utils/components/button.dart';
import 'package:app/utils/components/textfield.dart';
import 'package:app/utils/utils.dart';
import 'package:app/utils/validator.dart';
import 'package:app/utils/size.dart';

class PostBlogPage extends StatefulWidget {
  const PostBlogPage({super.key});

  @override
  State<PostBlogPage> createState() => _PostBlogPageState();
}

class _PostBlogPageState extends State<PostBlogPage> {
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
        showErrorDialog(context, "Image should not exceed 2 MB");
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

    fileToBytes(selectedFile!).then((bytesImage) {
      postBlog(
        controllerTitle.text,
        controllerDescription.text,
        selectedTag,
        bytesImage,
      ).then((res) {
        ScaffoldMessenger.of(context)
            .showSnackBar(returnSnackbar("Blog Added"));
        Navigator.pop(context);
      }).catchError((err) {
        handleErrors(context, err);
      });
    }).catchError((err) {
      showErrorDialog(context, "Failed to Convert Images into Bytes");
    });
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
                mainAxisAlignment: MainAxisAlignment.center,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  SizedBox(height: size.height * .02),

                  // Back Button & Post Blog Title
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

                        SizedBox(width: size.width * .18),
                        const Text(
                          "Post Blog",
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

                  // Image Picker
                  SizedBox(
                    height: size.height * .2,
                    child: Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 30),
                      child: selectedFile == null
                          ? GestureDetector(
                              onTap: pickImageFromGallery,
                              child: SizedBox(
                                height: size.height * .2,
                                child: Icon(
                                  Icons.photo,
                                  size: size.height * .2,
                                ),
                              ),
                            )
                          : GestureDetector(
                              onTap: pickImageFromGallery,
                              child: ClipRRect(
                                borderRadius:
                                    const BorderRadius.all(Radius.circular(20)),
                                child: Image.file(
                                  selectedFile!,
                                  height: size.height * .2,
                                ),
                              ),
                            ),
                    ),
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
                      text: "Post",
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