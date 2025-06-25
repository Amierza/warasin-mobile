import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:frontend/controller/location_controller.dart';
import 'package:frontend/controller/user/update_user.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';

class EditProfilePage extends StatelessWidget {
  EditProfilePage({super.key});

  final UpdateUser controller = Get.put(UpdateUser());
  final LocationController locationController = Get.put(LocationController());
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          "Edit Profile",
          style: GoogleFonts.poppins(
            fontSize: 20,
            fontWeight: semiBold,
            color: Colors.white,
          ),
        ),
        backgroundColor: primaryColor,
        elevation: 0,
        iconTheme: const IconThemeData(color: Colors.white),
      ),
      backgroundColor: backgroundColor,
      body: Obx(() {
        if (controller.isLoading.value && controller.name.value.isEmpty) {
          return const Center(child: CircularProgressIndicator());
        }
        return SingleChildScrollView(
          padding: const EdgeInsets.all(24.0),
          child: Form(
            key: _formKey,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                _buildProfileImageSection(),
                const SizedBox(height: 32),
                _buildFormFields(),
                const SizedBox(height: 40),
                _buildSaveButton(),
                const SizedBox(height: 20),
              ],
            ),
          ),
        );
      }),
    );
  }

  Widget _buildProfileImageSection() {
    return GestureDetector(
      onTap: () => controller.pickImage(),
      child: Stack(
        alignment: Alignment.center,
        children: [
          Container(
            width: 120,
            height: 120,
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              color: Colors.grey[200],
            ),
            child:
                controller.selectedPhoto != null
                    ? ClipOval(
                      child: Image.file(
                        controller.selectedPhoto!,
                        width: 120,
                        height: 120,
                        fit: BoxFit.cover,
                      ),
                    )
                    : controller.user.value?.userImage != null
                    ? ClipOval(
                      child: Image.network(
                        controller.user.value!.userImage,
                        width: 120,
                        height: 120,
                        fit: BoxFit.cover,
                        errorBuilder:
                            (context, error, stackTrace) =>
                                const Icon(Icons.person, size: 60),
                      ),
                    )
                    : const Icon(Icons.person, size: 60),
          ),
          Positioned(
            bottom: 0,
            right: 0,
            child: Container(
              padding: const EdgeInsets.all(4),
              decoration: const BoxDecoration(
                shape: BoxShape.circle,
                color: Colors.green,
              ),
              child: const Icon(
                Icons.camera_alt,
                color: Colors.white,
                size: 20,
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildFormFields() {
    return Column(
      children: [
        _buildTextFormField(
          label: "Nama Lengkap",
          icon: Icons.person_outline,
          controller: controller.nameController,
          onChanged: (val) => controller.name.value = val,
          validator:
              (val) => val?.isEmpty ?? true ? "Nama tidak boleh kosong" : null,
        ),
        const SizedBox(height: 20),
        _buildTextFormField(
          label: "Email",
          icon: Icons.email_outlined,
          controller: controller.emailController,
          onChanged: (val) => controller.email.value = val,
          keyboardType: TextInputType.emailAddress,
          validator: (val) {
            if (val?.isEmpty ?? true) return "Email tidak boleh kosong";
            if (!GetUtils.isEmail(val!)) return "Format email tidak valid";
            return null;
          },
        ),
        const SizedBox(height: 20),
        _buildTextFormField(
          label: "Nomor Telepon",
          icon: Icons.phone_outlined,
          controller: controller.phoneController,
          onChanged: (val) => controller.phoneNumber.value = val,
          keyboardType: TextInputType.phone,
          inputFormatters: [FilteringTextInputFormatter.digitsOnly],
          validator: (val) {
            if (val?.isEmpty ?? true) return "Nomor telepon tidak boleh kosong";
            if (val!.length < 10) return "Nomor telepon minimal 10 digit";
            return null;
          },
        ),
        const SizedBox(height: 20),
        _buildDateField(),
        const SizedBox(height: 20),
        _buildGenderField(),
        const SizedBox(height: 20),
        _buildProvinceDropdown(),
        const SizedBox(height: 20),
        _buildCityDropdown(),
      ],
    );
  }

  Widget _buildTextFormField({
    required String label,
    required IconData icon,
    required TextEditingController controller,
    required Function(String) onChanged,
    String? Function(String?)? validator,
    TextInputType? keyboardType,
    List<TextInputFormatter>? inputFormatters,
  }) {
    return TextFormField(
      controller: controller,
      onChanged: onChanged,
      validator: validator,
      keyboardType: keyboardType,
      inputFormatters: inputFormatters,
      decoration: InputDecoration(
        labelText: label,
        prefixIcon: Icon(icon, color: primaryColor),
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: Colors.grey.shade300),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: Colors.grey.shade300),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide(color: primaryColor, width: 2),
        ),
        filled: true,
        fillColor: Colors.grey.shade50,
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 16,
        ),
      ),
    );
  }

  Widget _buildDateField() {
    return TextFormField(
      controller: controller.birthDateController,
      readOnly: true,
      onTap: () async {
        final DateTime? picked = await showDatePicker(
          context: Get.context!,
          initialDate:
              controller.birthDate.value.isNotEmpty
                  ? DateFormat('yyyy-MM-dd').parse(controller.birthDate.value)
                  : DateTime.now().subtract(
                    const Duration(days: 6570),
                  ), // 18 tahun
          firstDate: DateTime(1900),
          lastDate: DateTime.now(),
          builder:
              (context, child) => Theme(
                data: Theme.of(context).copyWith(
                  colorScheme: ColorScheme.light(
                    primary: primaryColor,
                    onPrimary: Colors.white,
                    onSurface: Colors.black,
                  ),
                ),
                child: child!,
              ),
        );

        if (picked != null) {
          final formattedDate = DateFormat('yyyy-MM-dd').format(picked);
          controller.birthDate.value = formattedDate;
          controller.birthDateController.text =
              formattedDate; // <- Tambahkan ini
        }
      },
      validator:
          (val) =>
              val?.isEmpty ?? true ? "Tanggal lahir tidak boleh kosong" : null,
      decoration: InputDecoration(
        labelText: "Tanggal Lahir",
        prefixIcon: Icon(Icons.calendar_today_outlined, color: primaryColor),
        suffixIcon: Icon(Icons.arrow_drop_down, color: primaryColor),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
        filled: true,
        fillColor: Colors.grey.shade50,
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 16,
        ),
      ),
    );
  }

  Widget _buildGenderField() {
    return DropdownButtonFormField<bool>(
      value: controller.gender.value,
      onChanged: (val) => controller.gender.value = val,
      decoration: InputDecoration(
        labelText: "Jenis Kelamin",
        prefixIcon: Icon(Icons.person_outline, color: primaryColor),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
        filled: true,
        fillColor: Colors.grey.shade50,
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 16,
        ),
      ),
      items: const [
        DropdownMenuItem(value: true, child: Text("Perempuan")),
        DropdownMenuItem(value: false, child: Text("Laki-laki")),
      ],
      validator: (val) => val == null ? "Pilih jenis kelamin" : null,
    );
  }

  Widget _buildProvinceDropdown() {
    return Obx(() {
      if (locationController.isLoadingProvince.value) {
        return _buildLoadingInput("Memuat provinsi...");
      }

      if (!controller.isInitialized.value &&
          controller.province.value.isNotEmpty &&
          locationController.selectedProvinceId.value.isEmpty) {
        WidgetsBinding.instance.addPostFrameCallback((_) {
          locationController.selectProvince(controller.province.value);
          controller.isInitialized.value = true;
        });
      }

      return DropdownButtonFormField<String>(
        value:
            locationController.selectedProvinceId.value.isNotEmpty
                ? locationController.selectedProvinceId.value
                : null,
        onChanged: (String? val) {
          if (val != null) {
            locationController.selectProvince(val); // sudah termasuk fetch city
          }
        },
        decoration: InputDecoration(
          labelText: "Provinsi",
          prefixIcon: Icon(Icons.location_on, color: primaryColor),
          border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
          filled: true,
          fillColor: Colors.grey.shade50,
          contentPadding: const EdgeInsets.symmetric(
            horizontal: 16,
            vertical: 16,
          ),
        ),
        items:
            locationController.provinces.map((province) {
              return DropdownMenuItem<String>(
                value: province.provinceId,
                child: Text(province.provinceName),
              );
            }).toList(),
        validator: (val) => val?.isEmpty ?? true ? "Pilih provinsi" : null,
      );
    });
  }

  Widget _buildCityDropdown() {
    return Obx(() {
      if (locationController.isLoadingCity.value) {
        return _buildLoadingInput("Memuat kota...");
      }

      // Set initial value if user already has city
      if (controller.cityId.value.isNotEmpty &&
          locationController.selectedCityId.value.isEmpty) {
        locationController.selectedCityId.value = controller.cityId.value;
      }

      // Disable if no province selected
      if (locationController.selectedProvinceId.value.isEmpty) {
        return _buildDisabledInput("Pilih provinsi terlebih dahulu");
      }

      final cities = locationController.cities;
      final selectedCityId = locationController.selectedCityId.value;

      // Cek apakah selectedCityId ada di daftar cities
      final isSelectedCityValid = cities.any(
        (city) => city.cityId == selectedCityId,
      );

      return DropdownButtonFormField<String>(
        value: isSelectedCityValid ? selectedCityId : null,
        onChanged: (String? val) {
          if (val != null) {
            locationController.selectedCityId.value = val;
            controller.cityId.value = val;
          }
        },
        decoration: InputDecoration(
          labelText: "Kota",
          prefixIcon: Icon(Icons.location_city, color: primaryColor),
          border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
          filled: true,
          fillColor: Colors.grey.shade50,
          contentPadding: const EdgeInsets.symmetric(
            horizontal: 16,
            vertical: 16,
          ),
        ),
        items:
            cities.map((city) {
              return DropdownMenuItem<String>(
                value: city.cityId,
                child: Text(city.cityName),
              );
            }).toList(),
        validator: (val) => val?.isEmpty ?? true ? "Pilih kota" : null,
      );
    });
  }

  Widget _buildLoadingInput(String message) {
    return InputDecorator(
      decoration: InputDecoration(
        labelText: message,
        prefixIcon: Icon(Icons.hourglass_empty, color: Colors.grey),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
        filled: true,
        fillColor: Colors.grey.shade50,
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 16,
        ),
      ),
      child: Row(
        children: [
          Text(message, style: TextStyle(color: Colors.grey)),
          const Spacer(),
          const SizedBox(
            width: 20,
            height: 20,
            child: CircularProgressIndicator(strokeWidth: 2),
          ),
        ],
      ),
    );
  }

  Widget _buildDisabledInput(String message) {
    return InputDecorator(
      decoration: InputDecoration(
        labelText: message,
        prefixIcon: Icon(Icons.info_outline, color: Colors.grey),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
        filled: true,
        fillColor: Colors.grey.shade100,
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 16,
          vertical: 16,
        ),
      ),
      child: Text(message, style: TextStyle(color: Colors.grey)),
    );
  }

  Widget _buildSaveButton() {
    return Obx(() {
      return SizedBox(
        width: double.infinity,
        height: 56,
        child: ElevatedButton(
          onPressed: controller.isLoading.value ? null : _submitForm,
          style: ElevatedButton.styleFrom(
            backgroundColor: primaryColor,
            foregroundColor: Colors.white,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(16),
            ),
            elevation: 2,
            shadowColor: primaryColor.withOpacity(0.3),
          ),
          child:
              controller.isLoading.value
                  ? const SizedBox(
                    width: 24,
                    height: 24,
                    child: CircularProgressIndicator(
                      color: Colors.white,
                      strokeWidth: 2,
                    ),
                  )
                  : Text(
                    "Simpan Perubahan",
                    style: GoogleFonts.poppins(
                      fontSize: 16,
                      fontWeight: FontWeight.w600,
                    ),
                  ),
        ),
      );
    });
  }

  void _submitForm() {
    if (_formKey.currentState!.validate()) {
      // Hanya update jika user memilih kota baru
      final selectedCity = locationController.selectedCityId.value;
      if (selectedCity.isNotEmpty && selectedCity != controller.cityId.value) {
        controller.cityId.value = selectedCity;
      }
      controller.updateProfile();
    }
  }
}
