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
