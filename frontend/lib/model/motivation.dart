class MotivationCategory {
  final String categoryId;
  final String categoryName;

  MotivationCategory({required this.categoryId, required this.categoryName});

  factory MotivationCategory.fromJson(Map<String, dynamic> json) {
    return MotivationCategory(
      categoryId: json['motivation_category_id'],
      categoryName: json['motivation_category_name'],
    );
  }
}

class Motivation {
  final String motivationId;
  final String motivationAuthor;
  final String motivationContent;
  final MotivationCategory motivationCategory;

  Motivation({
    required this.motivationId,
    required this.motivationAuthor,
    required this.motivationContent,
    required this.motivationCategory,
  });

  factory Motivation.fromJson(Map<String, dynamic> json) {
    return Motivation(
      motivationId: json['motivation_id'],
      motivationAuthor: json['motivation_author'],
      motivationContent: json['motivation_content'],
      motivationCategory: MotivationCategory.fromJson(
        json['motivation_category'],
      ),
    );
  }
}

class AllMotivationResponse {
  final bool status;
  final String message;
  final String timestamp;
  final List<Motivation> data;

  AllMotivationResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory AllMotivationResponse.fromJson(Map<String, dynamic> json) {
    return AllMotivationResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data:
          (json['data'] as List)
              .map((motivation) => Motivation.fromJson(motivation))
              .toList(),
    );
  }
}

class UserMotivation {
  final String userMotivationId;
  final String userMotivationDate;
  final int userMotivationReaction;
  final Motivation motivation;

  UserMotivation({
    required this.userMotivationId,
    required this.userMotivationDate,
    required this.userMotivationReaction,
    required this.motivation,
  });

  factory UserMotivation.fromJson(Map<String, dynamic> json) {
    return UserMotivation(
      userMotivationId: json['user_mot_id'],
      userMotivationDate: json['user_mot_date'],
      userMotivationReaction: json['user_mot_reaction'],
      motivation: Motivation.fromJson(json['motivation']),
    );
  }
}

class CreateUserMotivationResponse {
  final bool status;
  final String message;
  final String timestamp;
  final UserMotivation data;

  CreateUserMotivationResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory CreateUserMotivationResponse.fromJson(Map<String, dynamic> json) {
    return CreateUserMotivationResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data: UserMotivation.fromJson(json['data']),
    );
  }
}

class AllUserMotivationResponse {
  final bool status;
  final String message;
  final String timestamp;
  final List<UserMotivation> data;

  AllUserMotivationResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory AllUserMotivationResponse.fromJson(Map<String, dynamic> json) {
    return AllUserMotivationResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data:
          (json['data'] as List)
              .map((motivation) => UserMotivation.fromJson(motivation))
              .toList(),
    );
  }
}
