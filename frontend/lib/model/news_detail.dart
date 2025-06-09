import 'package:frontend/model/news.dart';
import 'package:frontend/model/user_detail_response.dart';

class NewsDetail {
  final String newsDetailId;
  final String newsDetailDate;
  final User user;
  final News news;

  NewsDetail({
    required this.newsDetailId,
    required this.newsDetailDate,
    required this.user,
    required this.news,
  });

  factory NewsDetail.fromJson(Map<String, dynamic> json) {
    return NewsDetail(
      newsDetailId: json['news_detail_id'],
      newsDetailDate: json['news_detail_date'],
      user: User.fromJson(json['user']),
      news: News.fromJson(json['news']),
    );
  }
}

class CreateNewsDetail {
  final bool status;
  final String message;
  final String timestamp;
  final NewsDetail data;

  CreateNewsDetail({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory CreateNewsDetail.fromJson(Map<String, dynamic> json) {
    return CreateNewsDetail(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data: NewsDetail.fromJson(json['data']),
    );
  }
}
