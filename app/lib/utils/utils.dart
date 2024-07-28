String trimGrpcErrorMessage(String errorMessage) {
  // Define the pattern to match the message part
  final RegExp pattern = RegExp(r'message: ([^,]+)');

  // Find the first match
  final Match? match = pattern.firstMatch(errorMessage);

  // If a match is found, return the captured group; otherwise, return the original message
  return match != null ? match.group(1) ?? errorMessage : errorMessage;
}