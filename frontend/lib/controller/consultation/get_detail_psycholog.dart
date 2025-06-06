import 'package:frontend/model/error.dart';
import 'package:frontend/model/psycholog.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class GetDetailPsycholog extends GetxController {
  var isLoading = false.obs;
  var detailPsycholog = Rxn<PsychologResponse>();

  Future<void> fetchDetailPsycholog(String psyId) async {
    isLoading.value = true;

    try {
      final response = await ConsultationService.getDetailPsycholog(psyId);

      if (response is PsychologResponse) {
        detailPsycholog.value = response;
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
