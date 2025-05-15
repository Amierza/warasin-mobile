class TokenPayload {
  final String userId;
  final String roleId;
  final List<String> endpoints;
  final String iss;
  final int exp;
  final int iat;

  TokenPayload({
    required this.userId,
    required this.roleId,
    required this.endpoints,
    required this.iss,
    required this.exp,
    required this.iat,
  });

  factory TokenPayload.fromJson(Map<String, dynamic> json) {
    return TokenPayload(
      endpoints: List<String>.from(json['endpoints']),
      userId: json['user_id'],
      roleId: json['role_id'],
      iss: json['iss'],
      exp: json['exp'],
      iat: json['iat'],
    );
  }
}
