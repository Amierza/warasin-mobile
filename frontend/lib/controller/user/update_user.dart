import 'dart:convert';
import 'dart:io';

import 'package:frontend/model/user_detail_response.dart';
import 'package:frontend/service/api_service.dart';
import 'package:get/get.dart';
import 'package:image_picker/image_picker.dart';

class UpdateUser extends GetxController {
  var isLoading = false.obs;
  var name = ''.obs;
  var email = ''.obs;
  var phoneNumber = ''.obs;
  var birthDate = ''.obs;
  var cityId = ''.obs;
  var gender = RxnBool(null);
  var province = ''.obs;

  var imageFile = Rx<File?>(null);
  var imageBase64 = ''.obs;

  @override
  void onInit() {
    fetchUserDetail();
    super.onInit();
  }

  Future<void> pickImage() async {
    final picked = await ImagePicker().pickImage(source: ImageSource.gallery);
    if (picked != null) {
      imageFile.value = File(picked.path);
      final bytes = await picked.readAsBytes();
      imageBase64.value = "data:image/png;base64,${base64Encode(bytes)}";
    }
  }

  Future<void> updateProfile() async {
    isLoading.value = true;

    final Response = await ApiService.updateUserService(
      name.value.isNotEmpty ? name.value : null,
      email.value.isNotEmpty ? email.value : null,
      imageBase64.value.isNotEmpty ? imageBase64.value : null,
      gender.value,
      birthDate.value.isNotEmpty ? birthDate.value : null,
      phoneNumber.value.isNotEmpty ? phoneNumber.value : null,
      cityId.value.isNotEmpty ? cityId.value : null,
    );

    isLoading.value = false;

    if (Response != null) {
      Get.snackbar("Berhasil", "Profil berhasil diperbarui");
    } else {
      Get.snackbar("Gagal", "Terjadi kesalahan saat mengupdate profil");
    }
  }

  Future<void> fetchUserDetail() async {
    isLoading.value = true;

    final result = await ApiService.getUserDetailService();

    if (result != null && result is UserDetailResponse) {
      name.value = result.data.userName;
      email.value = result.data.userEmail;
      phoneNumber.value = result.data.userPhoneNumber;
      birthDate.value = result.data.userBirthDate;
      cityId.value = result.data.city.cityId.toString();
      province.value = result.data.city.province?.provinceId ?? '';
      gender.value = result.data.userGender;
      // imageBase64 tidak perlu di-set di sini kecuali kamu ingin tampilkan preview dari server
    }

    isLoading.value = false;
  }
}
