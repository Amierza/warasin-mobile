import 'package:frontend/model/consultation.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class GetAllPractice extends GetxController {
  var practiceList = <Practice>[].obs;
  var isLoading = false.obs;
  var allPractice = Rxn<PracticeResponse>();

  Future<void> fetchAllPractice(String psyId) async {
    isLoading.value = true;

    try {
      final response = await ConsultationService.getAllPractice(psyId);
      if (response is PracticeResponse) {
        allPractice.value = response;
        practiceList.value = response.data;
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
