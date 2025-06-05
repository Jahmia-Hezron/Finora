import 'dart:convert';

import 'package:http/http.dart' as http;

import '../base_service.dart' show BaseService;

class UserService extends BaseService {
  Future<Map<String, dynamic>?> registerUser(Map<String, dynamic> data) async {
    try {
      final token = await getAuthToken();
      final response = await http.post(
        url('user/register'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      final body = jsonDecode(response.body);
      return response.statusCode == 200
          ? body
          : {
            'error': body['message'] ?? 'Failed to register user',
            'statusCode': response.statusCode,
          };
    } on Exception catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<Map<String, dynamic>?> loginUser(Map<String, dynamic> data) async {
    try {
      final response = await http.post(
        url('user/login'),
        headers: headersWithToken(null),
        body: jsonEncode(data),
      );

      final body = jsonDecode(response.body);
      if (response.statusCode == 200) {
        await saveAuthToken(body['token']);
        return body;
      } else {
        return {
          'error': body['message'] ?? 'Failed to login user',
          'statusCode': response.statusCode,
        };
      }
    } on Exception catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<bool> logoutUser() async {
    try {
      final token = await getAuthToken();
      final response = await http.post(
        url('user/logout'),
        headers: headersWithToken(token),
      );
      if (response.statusCode == 200) {
        await clearAuthToken();
        return true;
      } else {
        return false;
      }
    } on Exception catch (_) {
      return false;
    }
  }

  Future<List<Map<String, dynamic>>> getUserByID(int userID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['user']);
      } else {
        throw Exception('Failed to fetch income by source');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> updateUser(
    int userID,
    Map<String, dynamic> userData,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.put(
        url('user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        return jsonDecode(response.body)['user'];
      } else {
        return {
          'error': 'Failed to udpate user',
          'statusCode': response.statusCode,
        };
      }
    } on Exception catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<Map<String, dynamic>?> resetPassword(
    int userID,
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.put(
        url('user/$userID/reset-password'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      if (response.statusCode == 200) {
        return jsonDecode(response.body)['user'];
      } else {
        return {
          'error': 'Failed to reset password',
          'statusCode': response.statusCode,
        };
      }
    } on Exception catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<Map<String, dynamic>?> deleteUser(int userID) async {
    try {
      final token = await getAuthToken();
      final response = await http.delete(
        url('user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        return jsonDecode(response.body)['user'];
      } else {
        return {
          'error': 'Failed to delete user',
          'statusCode': response.statusCode,
        };
      }
    } on Exception catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }
}
