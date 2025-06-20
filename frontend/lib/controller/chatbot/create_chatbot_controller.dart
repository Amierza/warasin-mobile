import 'package:frontend/model/chatbot.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/service/chatbot_service.dart';
import 'package:get/get.dart';

class CreateChatbotController extends GetxController {
  var message = ''.obs;
  var messageResult = Rxn<Chatbot>();
  var isLoading = false.obs;

  Future<void> createMessage() async {
    if (message.value.isEmpty) return;

    try {
      isLoading.value = true;
      final result = await ChatbotService.createChatbotService(message.value);

      if (result is ChatbotResponse) {
        messageResult.value = result.data;
      }

      if (result is ErrorResponse) {
        Get.snackbar(result.error, result.message);
      }
    } catch (error) {
      print('Unexpected error or null result, $error');
    } finally {
      isLoading.value = false;
    }
  }
}
