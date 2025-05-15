class ErrorResponse {
  final bool status;
  final String message;
  final String error;
  final String timestamp;

  ErrorResponse({
    required this.status,
    required this.message,
    required this.error,
    required this.timestamp,
  });

  factory ErrorResponse.fromJson(Map<String, dynamic> json) {
    return ErrorResponse(
      status: json['status'],
      message: json['message'],
      error: json['error'],
      timestamp: json['timestamp'],
    );
  }
}
