import 'package:flutter/material.dart';

class NewsDetail extends StatelessWidget {
  final String image;
  final String title;
  final String author;
  final String date;
  final String desc;

  const NewsDetail({
    Key? key,
    required this.image,
    required this.title,
    required this.author,
    required this.date,
    required this.desc,
  }): super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}