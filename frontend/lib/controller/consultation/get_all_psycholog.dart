import 'package:frontend/model/error.dart';
import 'package:frontend/model/psycholog.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class GetAllPsycholog extends GetxController {
  var psychologList = <Psycholog>[].obs;
  var isLoading = false.obs;
  var allPsycholog = Rxn<AllPsychologResponse>();

  @override
  void onInit() {
    fetchAllPsycholog();
    super.onInit();
  }

  Future<void> fetchAllPsycholog() async {
    isLoading.value = true;

    try {
      final response = await ConsultationService.getAllPsycholog();
      if (response is AllPsychologResponse) {
        allPsycholog.value = response;
        psychologList.value = response.data;
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
