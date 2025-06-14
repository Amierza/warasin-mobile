import 'package:frontend/model/error.dart';
import 'package:frontend/model/motivation.dart';
import 'package:frontend/service/motivation_service.dart';
import 'package:get/get.dart';

class GetAllMotivationController extends GetxController {
  var motivationList = <Motivation>[].obs;
  var isLoading = false.obs;
  var allMotivation = Rxn<AllMotivationResponse>();

  Future<void> fetchAllMotivatio() async {
    isLoading.value = true;

    try {
      final response = await MotivationService.getAllMotivationService();

      if (response is AllMotivationResponse) {
        allMotivation.value = response;
        motivationList.value = response.data;
      }

      if (response is ErrorResponse) {
        Get.snackbar(response.error, response.message);
      }
    } catch (error) {
      Get.snackbar('Something went wrong', error.toString());
    } finally {
      isLoading.value = false;
    }
  }
}
