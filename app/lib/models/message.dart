class ModelMessage {
  String sender;
  String receiver;
  String messageContent;
  String status;
  String dateTime;
  bool selfMessage;
  String error;
  String messageType;

  ModelMessage({
    required this.sender,
    required this.receiver,
    required this.messageContent,
    required this.status,
    required this.dateTime,
    required this.selfMessage,
    required this.error,
    required this.messageType,
  });

  // Factory constructor to create a Message object from JSON
  factory ModelMessage.fromJson(Map<String, dynamic> json) {
    return ModelMessage(
      sender: json['sender'],
      receiver: json['receiver'],
      messageContent: json['messageContent'],
      status: json['status'],
      dateTime: json['dateTime'],
      selfMessage: json['selfMessage'],
      error: json['error'],
      messageType: json['messageType'],
    );
  }

  // Method to convert a Message object to JSON
  Map<String, dynamic> toJson() {
    return {
      'sender': sender,
      'receiver': receiver,
      'messageContent': messageContent,
      'status': status,
      'dateTime': dateTime,
      'selfMessage': selfMessage,
      'error': error,
      'messageType': messageType,
    };
  }
}

class ModelWSMessage {
  String message;
  String receiver;
  String sessionToken;
  String messageType;

  ModelWSMessage({
    required this.message,
    required this.receiver,
    required this.sessionToken,
    required this.messageType,
  });

  // Factory constructor to create a WSMessage object from JSON
  factory ModelWSMessage.fromJson(Map<String, dynamic> json) {
    return ModelWSMessage(
      message: json['message'],
      receiver: json['receiver'],
      sessionToken: json['sessionToken'],
      messageType: json['messageType'],
    );
  }

  // Method to convert a WSMessage object to JSON
  Map<String, dynamic> toJson() {
    return {
      'message': message,
      'receiver': receiver,
      'sessionToken': sessionToken,
      'messageType': messageType,
    };
  }
}
