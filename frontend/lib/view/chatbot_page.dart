import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:frontend/controller/chatbot/create_chatbot_controller.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/navigation_bar.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class ChatbotPage extends StatefulWidget {
  @override
  _ChatScreenState createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatbotPage> {
  final TextEditingController _messageController = TextEditingController();
  final ScrollController _scrollController = ScrollController();
  final List<Message> _messages = [];
  final CreateChatbotController _chatbotController = Get.put(
    CreateChatbotController(),
  );

  @override
  void initState() {
    super.initState();
    _messages.add(
      Message(
        text: "Halo! Saya Warabot. Ada yang bisa saya bantu?",
        isBot: true,
        timestamp: DateTime.now(),
      ),
    );

    _chatbotController.messageResult.listen((chatbot) {
      if (chatbot != null) {
        _addBotMessage(chatbot.response);
      }
    });
  }

  void _addBotMessage(String text) {
    setState(() {
      _messages.add(
        Message(text: text, isBot: true, timestamp: DateTime.now()),
      );
    });
    _scrollToBottom();
  }

  void _sendMessage() async {
    if (_messageController.text.trim().isEmpty) return;

    final userMessage = _messageController.text.trim();

    setState(() {
      _messages.add(
        Message(text: userMessage, isBot: false, timestamp: DateTime.now()),
      );
    });

    _messageController.clear();
    _scrollToBottom();
    _chatbotController.message.value = userMessage;
    await _chatbotController.createMessage();
  }

  void _scrollToBottom() {
    Future.delayed(Duration(milliseconds: 100), () {
      if (_scrollController.hasClients) {
        _scrollController.animateTo(
          _scrollController.position.maxScrollExtent,
          duration: Duration(milliseconds: 300),
          curve: Curves.easeOut,
        );
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(Icons.smart_toy, color: Colors.white, size: 28),
            SizedBox(width: 8),
            Text(
              'Warabot',
              style: TextStyle(
                color: Colors.white,
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
            ),
          ],
        ),
        backgroundColor: primaryColor,
        centerTitle: true,
        automaticallyImplyLeading: false,
      ),
      body: Column(
        children: [
          Expanded(
            child: Container(
              color: Colors.grey[50],
              child: Obx(() {
                // Show loading indicator when waiting for response
                if (_chatbotController.isLoading.value) {
                  return Center(child: CircularProgressIndicator());
                }
                return ListView.builder(
                  controller: _scrollController,
                  padding: EdgeInsets.all(16),
                  itemCount: _messages.length,
                  itemBuilder: (context, index) {
                    final message = _messages[index];
                    return _buildMessageBubble(message);
                  },
                );
              }),
            ),
          ),
          _buildMessageInput(),
        ],
      ),
      bottomNavigationBar: const CustomNavigationBar(currentIndex: 1),
    );
  }

  Widget _buildMessageBubble(Message message) {
    return Padding(
      padding: EdgeInsets.symmetric(vertical: 4),
      child: Row(
        mainAxisAlignment:
            message.isBot ? MainAxisAlignment.start : MainAxisAlignment.end,
        children: [
          if (message.isBot) ...[
            CircleAvatar(
              radius: 16,
              backgroundColor: primaryColor,
              child: Icon(Icons.smart_toy, size: 18, color: Colors.white),
            ),
            SizedBox(width: 8),
          ],
          Flexible(
            child: Container(
              padding: EdgeInsets.symmetric(horizontal: 16, vertical: 12),
              decoration: BoxDecoration(
                color: message.isBot ? Colors.white : primaryColor,
                borderRadius: BorderRadius.circular(20),
                boxShadow: [
                  BoxShadow(
                    color: Colors.black.withOpacity(0.1),
                    blurRadius: 4,
                    offset: Offset(0, 2),
                  ),
                ],
              ),
              child: MarkdownBody(
                data: message.text,
                styleSheet: MarkdownStyleSheet.fromTheme(
                  Theme.of(context),
                ).copyWith(
                  p:GoogleFonts.poppins(
                    color: message.isBot ? primaryTextColor : Colors.white,
                    fontSize: 16,
                  ),
                ),
              ),
            ),
          ),
          if (!message.isBot) ...[
            SizedBox(width: 8),
            CircleAvatar(
              radius: 16,
              backgroundColor: Colors.grey[400],
              child: Icon(Icons.person, size: 18, color: Colors.white),
            ),
          ],
        ],
      ),
    );
  }

  Widget _buildMessageInput() {
    return Container(
      padding: EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.white,
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.1),
            blurRadius: 4,
            offset: Offset(0, -2),
          ),
        ],
      ),
      child: Row(
        children: [
          Expanded(
            child: Container(
              decoration: BoxDecoration(
                color: Colors.grey[100],
                borderRadius: BorderRadius.circular(25),
                border: Border.all(color: Colors.grey[300]!, width: 1),
              ),
              child: TextField(
                controller: _messageController,
                decoration: InputDecoration(
                  hintText: 'Ketik pesan Anda...',
                  hintStyle: TextStyle(color: Colors.grey[600]),
                  border: InputBorder.none,
                  contentPadding: EdgeInsets.symmetric(
                    horizontal: 20,
                    vertical: 12,
                  ),
                ),
                maxLines: null,
                onSubmitted: (_) => _sendMessage(),
              ),
            ),
          ),
          SizedBox(width: 12),
          Container(
            decoration: BoxDecoration(
              color: primaryColor,
              shape: BoxShape.circle,
            ),
            child: IconButton(
              icon: Icon(Icons.send, color: Colors.white),
              onPressed: _sendMessage,
            ),
          ),
        ],
      ),
    );
  }
}

class Message {
  final String text;
  final bool isBot;
  final DateTime timestamp;

  Message({required this.text, required this.isBot, required this.timestamp});
}
