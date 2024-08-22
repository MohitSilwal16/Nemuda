import 'package:flutter/material.dart';

import 'package:app/utils/utils.dart';
import 'package:app/utils/components/alert_dialogue.dart';

// Called when there is Error
navigateAndPOPbgPages(
    String routeName, String errorMessage, BuildContext context) {
  Navigator.of(context).pushNamedAndRemoveUntil(
    routeName,
    (Route<dynamic> route) => false, // Predicate to remove all routes
  );
  showErrorDialog(context, errorMessage);
}

handleErrors(BuildContext context, String err) {
  if (err.startsWith("TimeoutException")) {
    navigateAndPOPbgPages("server_busy", "Request Timed Out", context);
    return;
  }

  final trimmedGrpcError = trimGrpcErrorMessage(err.toString());

  // Common Errors
  if (trimmedGrpcError == "INVALID SESSION TOKEN") {
    navigateAndPOPbgPages("login", "Session Timed Out", context);
  }
  // Blog Errors
  else if (trimmedGrpcError == "BLOG NOT FOUND" ||
      trimmedGrpcError == "USER CANNOT UPDATE THIS BLOG" ||
      trimmedGrpcError == "USER CANNOT DELETE THIS BLOG") {
    navigateAndPOPbgPages("home", trimmedGrpcError, context);
  } else {
    showErrorDialog(context, trimmedGrpcError);
  }
}

// Called when there is Error in Future Builders
navigateAndPOPbgPagesFutureBuilder(
    String routeName, String errorMessage, BuildContext context) {
  WidgetsBinding.instance.addPostFrameCallback((_) {
    Navigator.of(context).pushNamedAndRemoveUntil(
      routeName,
      (Route<dynamic> route) => false, // Predicate to remove all routes
    );
    showErrorDialog(context, errorMessage);
  });
}

handleErrorsFutureBuilder(BuildContext context, String err) {
  if (err.startsWith("TimeoutException")) {
    navigateAndPOPbgPagesFutureBuilder(
        "server_busy", "Request Timed Out", context);
    return;
  }

  final trimmedGrpcError = trimGrpcErrorMessage(err.toString());

  // Common Errors
  if (trimmedGrpcError == "INVALID SESSION TOKEN") {
    navigateAndPOPbgPagesFutureBuilder("login", "Session Timed Out", context);
  }
  // Blog Errors
  else if (trimmedGrpcError == "BLOG NOT FOUND") {
    navigateAndPOPbgPagesFutureBuilder("home", "Blog Not Found", context);
  } else {
    showErrorDialog(context, trimmedGrpcError);
  }
}

// Called when there is Error in Bloc Builders
navigateAndPOPbgPagesBlocBuilder(
    String routeName, String errorMessage, BuildContext context) {
  WidgetsBinding.instance.addPostFrameCallback((_) {
    Navigator.pushReplacementNamed(context, routeName);
    showErrorDialog(context, errorMessage);
  });
}

handleErrorsBlocBuilder(BuildContext context, String err) {
  if (err.startsWith("TimeoutException")) {
    navigateAndPOPbgPagesFutureBuilder(
        "server_busy", "Request Timed Out", context);
  } else if (err == "INVALID SESSION TOKEN") {
    navigateAndPOPbgPagesFutureBuilder("login", "Session Timed Out", context);
  } else {
    showErrorDialog(context, err.toString());
  }
}
