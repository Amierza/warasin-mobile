class News {
  final String newsId;
  final String newsImage;
  final String newsTitle;
  final String newsBody;
  final String newsDate;

  News({
    required this.newsId,
    required this.newsImage,
    required this.newsTitle,
    required this.newsBody,
    required this.newsDate,
  });

  @override
  String toString() {
    return 'News(newsId: $newsId, title: $newsTitle, date: $newsDate)';
  }

  factory News.fromJson(Map<String, dynamic> json) {
    return News(
      newsId: json['news_id'],
      newsImage: json['news_image'],
      newsTitle: json['news_title'],
      newsBody: json['news_body'],
      newsDate: json['news_date'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'news_id': newsId,
      'news_image': newsImage,
      'news_title': newsTitle,
      'news_body': newsBody,
      'news_date': newsDate,
    };
  }
}

class GetAllNewsResponse {
  final bool status;
  final String message;
  final List<News> data;
  final String timestamp;

  GetAllNewsResponse({
    required this.status,
    required this.message,
    required this.data,
    required this.timestamp,
  });

  @override
  String toString() {
    return 'GetAllNewsResponse(status: $status, message: $message, data: $data, timestamp: $timestamp)';
  }

  factory GetAllNewsResponse.fromJson(Map<String, dynamic> json) {
    return GetAllNewsResponse(
      status: json['status'],
      message: json['message'],
      data: List<News>.from(json['data'].map((item) => News.fromJson(item))),
      timestamp: json['timestamp'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'status': status,
      'message': message,
      'data': data.map((news) => news.toJson()).toList(),
      'timestamp': timestamp,
    };
  }
}
