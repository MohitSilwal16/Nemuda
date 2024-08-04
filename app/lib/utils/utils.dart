import 'dart:io';
import 'dart:typed_data';

final tagsListHomePage = [
  "All",
  "Political",
  "Technical",
  "Educational",
  "Geographical",
  "Programming",
  "Other"
];

final tagsListPostUpdateBlog = [
  "Political",
  "Technical",
  "Educational",
  "Geographical",
  "Programming",
  "Other"
];

String trimGrpcErrorMessage(String errorMessage) {
  // Define the pattern to match the message part
  final RegExp pattern = RegExp(r'message: ([^,]+)');

  // Find the first match
  final Match? match = pattern.firstMatch(errorMessage);

  // If a match is found, return the captured group; otherwise, return the original message
  return match != null ? match.group(1) ?? errorMessage : errorMessage;
}

// Function to convert File to bytes
Future<Uint8List> fileToBytes(File file) async {
  // Read the file as bytes
  Uint8List bytes = await file.readAsBytes();
  return bytes;
}
