import 'regex.dart';

class Validators {
  static String? validateUsername(String? value) {
    if (value == null || value.isEmpty) {
      return 'This field is mandatory';
    }

    if (!RegexConstants.usernameRegex.hasMatch(value)) {
      return 'Alphanumeric & b\'twin 5-20 chars';
    }

    return null; // Return null if input is valid
  }

  static String? validatePassword(String? value) {
    if (value == null || value.isEmpty) {
      return 'This field is mandatory';
    }

    if (!RegexConstants.passwordRegex.hasMatch(value)) {
      return '8+ chars, lower & upper case, digit, symbol';
    }

    return null; // Return null if input is valid
  }
}
