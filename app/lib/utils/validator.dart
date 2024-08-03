import 'regex.dart';

class Validators {
  static String? validateUsername(String? value, String? errorMessage) {
    if (value == null || value.isEmpty) {
      return 'This field is mandatory';
    }

    if (!RegexConstants.usernameRegex.hasMatch(value)) {
      return 'Alphanumeric & b\'twin 5-20 chars';
    }

    if (errorMessage != null ){
      return errorMessage;
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

  static String? validateTitle(String? value) {
    if (value == null || value.isEmpty) {
      return 'This field is mandatory';
    }

    if (!RegexConstants.titleRegex.hasMatch(value)) {
      return 'Min 5 & Max 20 chars';
    }

    return null; // Return null if input is valid
  }

  static String? validateDescription(String? value) {
    if (value == null || value.isEmpty) {
      return 'This field is mandatory';
    }

    if (value.length < 5 || value.length > 50) {
      return 'Min 5 & Max 50 chars';
    }

    return null; // Return null if input is valid
  }
}
