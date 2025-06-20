class Chatbot {
  final String response;
  final String? conversationId;

  Chatbot({required this.response, required this.conversationId});

  factory Chatbot.fromJson(Map<String, dynamic> json) {
    return Chatbot(
      response: json['response'] ?? '',
      conversationId: json['conversation_id'],
    );
  }
}

class ChatbotResponse {
  final bool status;
  final String message;
  final String timestamp;
  final Chatbot data;

  ChatbotResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory ChatbotResponse.fromJson(Map<String, dynamic> json) {
    return ChatbotResponse(
      status: json['status'] ?? false,
      message: json['message'] ?? '',
      timestamp: json['timestamp'] ?? '',
      data:
          json['data'] != null
              ? Chatbot.fromJson(json['data'])
              : Chatbot(response: '', conversationId: null),
    );
  }
}
