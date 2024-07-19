class RegexConstants {
  static final RegExp usernameRegex = RegExp(r'^[a-zA-Z0-9]{5,20}$');

  static final RegExp passwordRegex =
      RegExp(r'^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[!@#\$&*~]).{8,}$');
}
